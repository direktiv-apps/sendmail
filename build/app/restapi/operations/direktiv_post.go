package operations

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/direktiv/apps/go/pkg/apps"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	"app/models"
)

const (
	successKey = "success"
	resultKey  = "result"

	// http related
	statusKey  = "status"
	codeKey    = "code"
	headersKey = "headers"
)

var sm sync.Map

const (
	cmdErr = "io.direktiv.command.error"
	outErr = "io.direktiv.output.error"
	riErr  = "io.direktiv.ri.error"
)

type accParams struct {
	PostParams
	Commands    []interface{}
	DirektivDir string
}

type accParamsTemplate struct {
	models.PostParamsBody
	Commands    []interface{}
	DirektivDir string
}

type ctxInfo struct {
	cf        context.CancelFunc
	cancelled bool
}

func PostDirektivHandle(params PostParams) middleware.Responder {
	resp := &models.PostOKBody{}

	var (
		err  error
		ret  interface{}
		cont bool
	)

	ri, err := apps.RequestinfoFromRequest(params.HTTPRequest)
	if err != nil {
		return generateError(riErr, err)
	}

	ctx, cancel := context.WithCancel(params.HTTPRequest.Context())

	sm.Store(*params.DirektivActionID, &ctxInfo{
		cancel,
		false,
	})

	defer sm.Delete(*params.DirektivActionID)

	var responses []interface{}

	var paramsCollector []interface{}
	accParams := accParams{
		params,
		nil,
		ri.Dir(),
	}

	ret, err = runCommand0(ctx, accParams, ri)

	responses = append(responses, ret)

	// if foreach returns an error there is no continue
	//
	// default we do not continue
	cont = convertTemplateToBool("<no value>", accParams, false)
	// cont = convertTemplateToBool("<no value>", accParams, true)
	//

	if err != nil && !cont {

		errName := cmdErr

		// if the delete function added the cancel tag
		ci, ok := sm.Load(*params.DirektivActionID)
		if ok {
			cinfo, ok := ci.(*ctxInfo)
			if ok && cinfo.cancelled {
				errName = "direktiv.actionCancelled"
				err = fmt.Errorf("action got cancel request")
			}
		}

		return generateError(errName, err)
	}

	paramsCollector = append(paramsCollector, ret)
	accParams.Commands = paramsCollector

	ret, err = runCommand1(ctx, accParams, ri)

	responses = append(responses, ret)

	// if foreach returns an error there is no continue
	//
	// cont = false
	//

	if err != nil && !cont {

		errName := cmdErr

		// if the delete function added the cancel tag
		ci, ok := sm.Load(*params.DirektivActionID)
		if ok {
			cinfo, ok := ci.(*ctxInfo)
			if ok && cinfo.cancelled {
				errName = "direktiv.actionCancelled"
				err = fmt.Errorf("action got cancel request")
			}
		}

		return generateError(errName, err)
	}

	paramsCollector = append(paramsCollector, ret)
	accParams.Commands = paramsCollector

	s, err := templateString(`{
  "sendmail": {{ index . 1 | toJson }}
}
`, responses)
	if err != nil {
		return generateError(outErr, err)
	}

	responseBytes := []byte(s)

	// validate
	resp.UnmarshalBinary(responseBytes)
	err = resp.Validate(strfmt.Default)

	if err != nil {
		return generateError(outErr, err)
	}

	return NewPostOK().WithPayload(resp)
}

// exec
func runCommand0(ctx context.Context,
	params accParams, ri *apps.RequestInfo) (map[string]interface{}, error) {

	ir := make(map[string]interface{})
	ir[successKey] = false

	at := accParamsTemplate{
		*params.Body,
		params.Commands,
		params.DirektivDir,
	}

	cmd, err := templateString(`/replace.sh '{{ .SMTP.Server }}' '{{ .SMTP.User }}' '{{ .SMTP.Password }}'`, at)
	if err != nil {
		ri.Logger().Infof("error executing command: %v", err)
		ir[resultKey] = err.Error()
		return ir, err
	}
	cmd = strings.Replace(cmd, "\n", "", -1)

	silent := convertTemplateToBool("true", at, false)
	print := convertTemplateToBool("false", at, true)
	output := ""

	envs := []string{}

	return runCmd(ctx, cmd, envs, output, silent, print, ri)

}

// end commands

// foreach command
type LoopStruct1 struct {
	accParams
	Item        interface{}
	DirektivDir string
}

func runCommand1(ctx context.Context,
	params accParams, ri *apps.RequestInfo) ([]map[string]interface{}, error) {

	var cmds []map[string]interface{}

	if params.Body == nil {
		return cmds, nil
	}

	for a := range params.Body.Emails {

		ls := &LoopStruct1{
			params,
			params.Body.Emails[a],
			params.DirektivDir,
		}

		cmd, err := templateString(`bash -c 'cat message | s-nail {{- if .Item.Verbose }} -vv {{- end }} -r "{{ .Item.From }}" -A mail {{- if .Item.Subject }} -s "{{ .Item.Subject }}" {{- end }} 
{{- range $i, $a := .Item.Bcc }} -b {{ $a }} {{- end }}
{{- range $i, $a := .Item.Cc }} -c {{ $a }} {{- end }}
{{- range $i, $a := .Item.Attachments }} -a {{ $a }} {{- end }}
{{- range $i, $a := .Item.To }} {{ $a }} {{- end }}'`, ls)
		if err != nil {
			ir := make(map[string]interface{})
			ir[successKey] = false
			ir[resultKey] = err.Error()
			cmds = append(cmds, ir)
			continue
		}

		silent := convertTemplateToBool("false", ls, false)
		print := convertTemplateToBool("true", ls, true)
		cont := convertTemplateToBool("<no value>", ls, false)
		output := ""

		envs := []string{}
		env0, _ := templateString(`MAILRC=account.config`, ls)
		envs = append(envs, env0)

		r, err := runCmd(ctx, cmd, envs, output, silent, print, ri)
		if err != nil {
			ir := make(map[string]interface{})
			ir[successKey] = false
			ir[resultKey] = err.Error()
			cmds = append(cmds, ir)

			if cont {
				continue
			}

			return cmds, err

		}
		cmds = append(cmds, r)

	}

	return cmds, nil

}

// end commands

func generateError(code string, err error) *PostDefault {

	d := NewPostDefault(0).WithDirektivErrorCode(code).
		WithDirektivErrorMessage(err.Error())

	errString := err.Error()

	errResp := models.Error{
		ErrorCode:    &code,
		ErrorMessage: &errString,
	}

	d.SetPayload(&errResp)

	return d
}

func HandleShutdown() {
	// nothing for generated functions
}
