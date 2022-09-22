
# sendmail 1.0

Run sendmail in Direktiv

---
- #### Categories: unknown
- #### Image: gcr.io/direktiv/functions/sendmail 
- #### License: [Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)
- #### Issue Tracking: https://github.com/direktiv-apps/sendmail/issues
- #### URL: https://github.com/direktiv-apps/sendmail
- #### Maintainer: [direktiv.io](https://www.direktiv.io) 
---

## About sendmail

Run sendmail in Direktiv as a function

### Example(s)
  #### Function Configuration
```yaml
functions:
- id: sendmail
  image: gcr.io/direktiv/functions/sendmail:1.0
  type: knative-workflow
```
   #### Basic
```yaml
- id: sendmail
  type: action
  action:
    function: sendmail
    input: 
      commands:
      - command: Example of running sendmail
```
   #### Advanced
```yaml
- id: sendmail
  type: action
  action:
    function: sendmail
    input: 
      files:
      - name: hello.txt
        data: Hello World
        mode: '0755'
      commands:
      - command: Example of running sendmail
```

   ### Secrets


- **smtppwd**: SMTP password






### Request



#### Request Attributes
[PostParamsBody](#post-params-body)

### Response
  List of executed commands.
#### Reponse Types
    
  

[PostOKBody](#post-o-k-body)
#### Example Reponses
    
```json
[
  {
    "result": null,
    "success": true
  },
  {
    "result": null,
    "success": true
  }
]
```

### Errors
| Type | Description
|------|---------|
| io.direktiv.command.error | Command execution failed |
| io.direktiv.output.error | Template error for output generation of the service |
| io.direktiv.ri.error | Can not create information object from request |


### Types
#### <span id="post-o-k-body"></span> postOKBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| sendmail | [][PostOKBodySendmailItems](#post-o-k-body-sendmail-items)| `[]*PostOKBodySendmailItems` |  | |  |  |


#### <span id="post-o-k-body-sendmail-items"></span> postOKBodySendmailItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| result | [interface{}](#interface)| `interface{}` | ✓ | |  |  |
| success | boolean| `bool` | ✓ | |  |  |


#### <span id="post-params-body"></span> postParamsBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| emails | [][PostParamsBodyEmailsItems](#post-params-body-emails-items)| `[]*PostParamsBodyEmailsItems` |  | | List of emails to send. |  |
| smtp | [PostParamsBodySMTP](#post-params-body-smtp)| `PostParamsBodySMTP` |  | |  |  |


#### <span id="post-params-body-emails-items"></span> postParamsBodyEmailsItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| attachments | []string| `[]string` |  | | Files to attach to the email. Can be provided with Direktiv action `files` |  |
| bcc | []string| `[]string` |  | | Email addresses to send email to (blind copy) |  |
| cc | []string| `[]string` |  | | Email addresses to send email to (carbon copy) |  |
| from | string| `string` |  | | Name used as `from` value, e.g. "My Name\<myname@direktiv.io\>" | `alue, e.g. \"My Name\\\u003cmyname@direktiv.io\\\u003e\` |
| message | string| `string` |  | | Email message. JQ can be used in the text. |  |
| subject | string| `string` |  | | Subject of the email |  |
| to | []string| `[]string` |  | | Email addresses to send email to |  |


#### <span id="post-params-body-smtp"></span> postParamsBodySmtp

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| password | string| `string` |  | |  |  |
| server | string| `string` |  | | SMTP server address | `smtp.sendgrid.net:587` |
| user | string| `string` |  | |  |  |

 
