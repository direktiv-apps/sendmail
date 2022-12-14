// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Sending emails via SMTP",
    "title": "sendmail",
    "version": "1.0",
    "x-direktiv-meta": {
      "categories": [
        "social"
      ],
      "container": "gcr.io/direktiv/functions/sendmail",
      "issues": "https://github.com/direktiv-apps/sendmail/issues",
      "license": "[Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)",
      "long-description": "This function uses s-nail to send emails via SMTP. It supports CC, BCC and attachments. \n\nIt is required to name the message ` + "`" + `message` + "`" + ` (see examples). Alternatively the message can be a Direktiv attribute stored as ` + "`" + `message` + "`" + `.\n\nFor Gmail an ` + "`" + `App password` + "`" + ` has to be created under the account using this function: [Google Account Admin](https://myaccount.google.com/security).",
      "maintainer": "[direktiv.io](https://www.direktiv.io) ",
      "url": "https://github.com/direktiv-apps/sendmail"
    }
  },
  "paths": {
    "/": {
      "post": {
        "parameters": [
          {
            "type": "string",
            "default": "development",
            "description": "direktiv action id is an UUID. \nFor development it can be set to 'development'\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          },
          {
            "type": "string",
            "default": "/tmp",
            "description": "direktiv temp dir is the working directory for that request\nFor development it can be set to e.g. '/tmp'\n",
            "name": "Direktiv-TempDir",
            "in": "header"
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "required": [
                "smtp",
                "emails"
              ],
              "properties": {
                "emails": {
                  "description": "List of emails to send.",
                  "type": "array",
                  "items": {
                    "type": "object",
                    "required": [
                      "from",
                      "to"
                    ],
                    "properties": {
                      "attachments": {
                        "description": "Files to attach to the email. Can be provided with Direktiv action ` + "`" + `files` + "`" + `",
                        "type": "array",
                        "items": {
                          "type": "string"
                        }
                      },
                      "bcc": {
                        "description": "Email addresses to send email to (blind copy)",
                        "type": "array",
                        "items": {
                          "type": "string"
                        }
                      },
                      "cc": {
                        "description": "Email addresses to send email to (carbon copy)",
                        "type": "array",
                        "items": {
                          "type": "string"
                        }
                      },
                      "contentType": {
                        "description": "Content-Type of the message, e.g. text/html",
                        "type": "string",
                        "default": "text/plain"
                      },
                      "from": {
                        "description": "Name used as ` + "`" + `from` + "`" + ` value, e.g. \"My Name\\\u003cmyname@direktiv.io\\\u003e\"",
                        "type": "string"
                      },
                      "message": {
                        "$ref": "#/definitions/direktivFile"
                      },
                      "subject": {
                        "description": "Subject of the email",
                        "type": "string"
                      },
                      "to": {
                        "description": "Email addresses to send email to",
                        "type": "array",
                        "items": {
                          "type": "string"
                        }
                      },
                      "verbose": {
                        "description": "Enable debug output",
                        "type": "boolean"
                      }
                    }
                  }
                },
                "smtp": {
                  "type": "object",
                  "required": [
                    "server",
                    "user",
                    "password"
                  ],
                  "properties": {
                    "password": {
                      "description": "Password for the SMTP server",
                      "type": "string"
                    },
                    "server": {
                      "description": "SMTP server address",
                      "type": "string",
                      "example": "smtp.sendgrid.net:587"
                    },
                    "user": {
                      "description": "User name for the SMTP server",
                      "type": "string"
                    }
                  }
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Output of executed email commands.",
            "schema": {
              "type": "object",
              "properties": {
                "sendmail": {
                  "type": "array",
                  "items": {
                    "type": "object",
                    "required": [
                      "success",
                      "result"
                    ],
                    "properties": {
                      "result": {
                        "additionalProperties": false
                      },
                      "success": {
                        "type": "boolean"
                      }
                    }
                  }
                }
              }
            },
            "examples": {
              "sendmail": [
                {
                  "result": null,
                  "success": true
                }
              ]
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            },
            "headers": {
              "Direktiv-ErrorCode": {
                "type": "string"
              },
              "Direktiv-ErrorMessage": {
                "type": "string"
              }
            }
          }
        },
        "x-direktiv": {
          "cmds": [
            {
              "action": "exec",
              "exec": "/replace.sh '{{ .SMTP.Server }}' '{{ .SMTP.User }}' '{{ .SMTP.Password }}'",
              "print": false,
              "silent": true
            },
            {
              "action": "foreach",
              "env": [
                "MAILRC=account.config"
              ],
              "exec": "bash -c 'cat message | s-nail {{- if .Item.Verbose }} -vv {{- end }} -r \"{{ .Item.From }}\" -A mail {{- if .Item.ContentType }} -M \"{{ .Item.ContentType }}\" {{- end }} {{- if .Item.Subject }} -s \"{{ .Item.Subject }}\" {{- end }} \n{{- range $i, $a := .Item.Bcc }} -b {{ $a }} {{- end }}\n{{- range $i, $a := .Item.Cc }} -c {{ $a }} {{- end }}\n{{- range $i, $a := .Item.Attachments }} -a {{ $a }} {{- end }}\n{{- range $i, $a := .Item.To }} {{ $a }} {{- end }}'",
              "loop": ".Emails",
              "print": true,
              "silent": false
            }
          ],
          "output": "{\n  \"sendmail\": {{ index . 1 | toJson }}\n}\n"
        },
        "x-direktiv-errors": {
          "io.direktiv.command.error": "Command execution failed",
          "io.direktiv.output.error": "Template error for output generation of the service",
          "io.direktiv.ri.error": "Can not create information object from request"
        },
        "x-direktiv-examples": [
          {
            "content": "- id: values\n  type: noop\n  transform:\n    name: Direktiv\n  transition: sendmail\n- id: sendmail\n  type: action\n  action:\n    function: sendmail\n    secrets: [\"smtppwd\"]\n    input: \n      smtp:\n        server: smtp.server.com:587\n        user: myuser@myemail.com\n        password: jq(.secrets.smtppwd)\n      emails:\n      - from: My User\u003cmyuser@myemail.com\u003e\n        to:\n        - user1@emailtest.com\n        cc:\n        - user2@emailtest.com\n        subject: This Is A Message\n        message:\n          name: message\n          data: |-\n            Hello jq(.name),\n\n            this is an email.\n\n            Good Bye\n  catch:\n  - error: \"*\"",
            "title": "Basic Email"
          },
          {
            "content": "- id: sendmail\n  type: action\n  action:\n    function: sendmail\n    secrets: [\"smtppwd\"]\n    input: \n      smtp:\n        server: smtp.server.com:587\n        user: myuser@myemail.com\n        password: jq(.secrets.smtppwd)\n      emails:\n      - from: My User\u003cmyuser@myemail.com\u003e\n        to: \n        - user1@emailtest.com\n        attachments:\n        - message\n        subject: This Is A Message\n        message:\n          name: message\n          data: |-\n            This is the text and the attachment\n  catch:\n  - error: \"*\"",
            "title": "Emails with attachment"
          }
        ],
        "x-direktiv-function": "functions:\n- id: sendmail\n  image: gcr.io/direktiv/functions/sendmail:1.0\n  type: knative-workflow",
        "x-direktiv-secrets": [
          {
            "description": "SMTP password",
            "name": "smtppwd"
          }
        ]
      },
      "delete": {
        "parameters": [
          {
            "type": "string",
            "description": "On cancel Direktiv sends a DELETE request to\nthe action with id in the header\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          }
        ],
        "responses": {
          "200": {
            "description": ""
          }
        },
        "x-direktiv": {
          "cancel": "echo 'cancel {{ .DirektivActionID }}'"
        }
      }
    }
  },
  "definitions": {
    "direktivFile": {
      "type": "object",
      "x-go-type": {
        "import": {
          "package": "github.com/direktiv/apps/go/pkg/apps"
        },
        "type": "DirektivFile"
      }
    },
    "error": {
      "type": "object",
      "required": [
        "errorCode",
        "errorMessage"
      ],
      "properties": {
        "errorCode": {
          "type": "string"
        },
        "errorMessage": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Sending emails via SMTP",
    "title": "sendmail",
    "version": "1.0",
    "x-direktiv-meta": {
      "categories": [
        "social"
      ],
      "container": "gcr.io/direktiv/functions/sendmail",
      "issues": "https://github.com/direktiv-apps/sendmail/issues",
      "license": "[Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)",
      "long-description": "This function uses s-nail to send emails via SMTP. It supports CC, BCC and attachments. \n\nIt is required to name the message ` + "`" + `message` + "`" + ` (see examples). Alternatively the message can be a Direktiv attribute stored as ` + "`" + `message` + "`" + `.\n\nFor Gmail an ` + "`" + `App password` + "`" + ` has to be created under the account using this function: [Google Account Admin](https://myaccount.google.com/security).",
      "maintainer": "[direktiv.io](https://www.direktiv.io) ",
      "url": "https://github.com/direktiv-apps/sendmail"
    }
  },
  "paths": {
    "/": {
      "post": {
        "parameters": [
          {
            "type": "string",
            "default": "development",
            "description": "direktiv action id is an UUID. \nFor development it can be set to 'development'\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          },
          {
            "type": "string",
            "default": "/tmp",
            "description": "direktiv temp dir is the working directory for that request\nFor development it can be set to e.g. '/tmp'\n",
            "name": "Direktiv-TempDir",
            "in": "header"
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/postParamsBody"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Output of executed email commands.",
            "schema": {
              "$ref": "#/definitions/postOKBody"
            },
            "examples": {
              "sendmail": [
                {
                  "result": null,
                  "success": true
                }
              ]
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            },
            "headers": {
              "Direktiv-ErrorCode": {
                "type": "string"
              },
              "Direktiv-ErrorMessage": {
                "type": "string"
              }
            }
          }
        },
        "x-direktiv": {
          "cmds": [
            {
              "action": "exec",
              "exec": "/replace.sh '{{ .SMTP.Server }}' '{{ .SMTP.User }}' '{{ .SMTP.Password }}'",
              "print": false,
              "silent": true
            },
            {
              "action": "foreach",
              "env": [
                "MAILRC=account.config"
              ],
              "exec": "bash -c 'cat message | s-nail {{- if .Item.Verbose }} -vv {{- end }} -r \"{{ .Item.From }}\" -A mail {{- if .Item.ContentType }} -M \"{{ .Item.ContentType }}\" {{- end }} {{- if .Item.Subject }} -s \"{{ .Item.Subject }}\" {{- end }} \n{{- range $i, $a := .Item.Bcc }} -b {{ $a }} {{- end }}\n{{- range $i, $a := .Item.Cc }} -c {{ $a }} {{- end }}\n{{- range $i, $a := .Item.Attachments }} -a {{ $a }} {{- end }}\n{{- range $i, $a := .Item.To }} {{ $a }} {{- end }}'",
              "loop": ".Emails",
              "print": true,
              "silent": false
            }
          ],
          "output": "{\n  \"sendmail\": {{ index . 1 | toJson }}\n}\n"
        },
        "x-direktiv-errors": {
          "io.direktiv.command.error": "Command execution failed",
          "io.direktiv.output.error": "Template error for output generation of the service",
          "io.direktiv.ri.error": "Can not create information object from request"
        },
        "x-direktiv-examples": [
          {
            "content": "- id: values\n  type: noop\n  transform:\n    name: Direktiv\n  transition: sendmail\n- id: sendmail\n  type: action\n  action:\n    function: sendmail\n    secrets: [\"smtppwd\"]\n    input: \n      smtp:\n        server: smtp.server.com:587\n        user: myuser@myemail.com\n        password: jq(.secrets.smtppwd)\n      emails:\n      - from: My User\u003cmyuser@myemail.com\u003e\n        to:\n        - user1@emailtest.com\n        cc:\n        - user2@emailtest.com\n        subject: This Is A Message\n        message:\n          name: message\n          data: |-\n            Hello jq(.name),\n\n            this is an email.\n\n            Good Bye\n  catch:\n  - error: \"*\"",
            "title": "Basic Email"
          },
          {
            "content": "- id: sendmail\n  type: action\n  action:\n    function: sendmail\n    secrets: [\"smtppwd\"]\n    input: \n      smtp:\n        server: smtp.server.com:587\n        user: myuser@myemail.com\n        password: jq(.secrets.smtppwd)\n      emails:\n      - from: My User\u003cmyuser@myemail.com\u003e\n        to: \n        - user1@emailtest.com\n        attachments:\n        - message\n        subject: This Is A Message\n        message:\n          name: message\n          data: |-\n            This is the text and the attachment\n  catch:\n  - error: \"*\"",
            "title": "Emails with attachment"
          }
        ],
        "x-direktiv-function": "functions:\n- id: sendmail\n  image: gcr.io/direktiv/functions/sendmail:1.0\n  type: knative-workflow",
        "x-direktiv-secrets": [
          {
            "description": "SMTP password",
            "name": "smtppwd"
          }
        ]
      },
      "delete": {
        "parameters": [
          {
            "type": "string",
            "description": "On cancel Direktiv sends a DELETE request to\nthe action with id in the header\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          }
        ],
        "responses": {
          "200": {
            "description": ""
          }
        },
        "x-direktiv": {
          "cancel": "echo 'cancel {{ .DirektivActionID }}'"
        }
      }
    }
  },
  "definitions": {
    "direktivFile": {
      "type": "object",
      "x-go-type": {
        "import": {
          "package": "github.com/direktiv/apps/go/pkg/apps"
        },
        "type": "DirektivFile"
      }
    },
    "error": {
      "type": "object",
      "required": [
        "errorCode",
        "errorMessage"
      ],
      "properties": {
        "errorCode": {
          "type": "string"
        },
        "errorMessage": {
          "type": "string"
        }
      }
    },
    "postOKBody": {
      "type": "object",
      "properties": {
        "sendmail": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/postOKBodySendmailItems"
          }
        }
      },
      "x-go-gen-location": "operations"
    },
    "postOKBodySendmailItems": {
      "type": "object",
      "required": [
        "success",
        "result"
      ],
      "properties": {
        "result": {
          "additionalProperties": false
        },
        "success": {
          "type": "boolean"
        }
      },
      "x-go-gen-location": "operations"
    },
    "postParamsBody": {
      "type": "object",
      "required": [
        "smtp",
        "emails"
      ],
      "properties": {
        "emails": {
          "description": "List of emails to send.",
          "type": "array",
          "items": {
            "$ref": "#/definitions/postParamsBodyEmailsItems"
          }
        },
        "smtp": {
          "$ref": "#/definitions/postParamsBodySmtp"
        }
      },
      "x-go-gen-location": "operations"
    },
    "postParamsBodyEmailsItems": {
      "type": "object",
      "required": [
        "from",
        "to"
      ],
      "properties": {
        "attachments": {
          "description": "Files to attach to the email. Can be provided with Direktiv action ` + "`" + `files` + "`" + `",
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "bcc": {
          "description": "Email addresses to send email to (blind copy)",
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "cc": {
          "description": "Email addresses to send email to (carbon copy)",
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "contentType": {
          "description": "Content-Type of the message, e.g. text/html",
          "type": "string",
          "default": "text/plain"
        },
        "from": {
          "description": "Name used as ` + "`" + `from` + "`" + ` value, e.g. \"My Name\\\u003cmyname@direktiv.io\\\u003e\"",
          "type": "string"
        },
        "message": {
          "$ref": "#/definitions/direktivFile"
        },
        "subject": {
          "description": "Subject of the email",
          "type": "string"
        },
        "to": {
          "description": "Email addresses to send email to",
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "verbose": {
          "description": "Enable debug output",
          "type": "boolean"
        }
      },
      "x-go-gen-location": "operations"
    },
    "postParamsBodySmtp": {
      "type": "object",
      "required": [
        "server",
        "user",
        "password"
      ],
      "properties": {
        "password": {
          "description": "Password for the SMTP server",
          "type": "string"
        },
        "server": {
          "description": "SMTP server address",
          "type": "string",
          "example": "smtp.sendgrid.net:587"
        },
        "user": {
          "description": "User name for the SMTP server",
          "type": "string"
        }
      },
      "x-go-gen-location": "operations"
    }
  }
}`))
}
