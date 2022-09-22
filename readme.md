
# sendmail 1.0

Sending emails via SMTP

---
- #### Categories: social
- #### Image: gcr.io/direktiv/functions/sendmail 
- #### License: [Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)
- #### Issue Tracking: https://github.com/direktiv-apps/sendmail/issues
- #### URL: https://github.com/direktiv-apps/sendmail
- #### Maintainer: [direktiv.io](https://www.direktiv.io) 
---

## About sendmail

This function uses s-nail to send emails via SMTP. It supports CC, BCC and attachments. 
It is required to name the message `message`. Alternatively the message can be a Direktiv attribute stored as `message`.
For Gmail an `App password` has to be created under the account using this function: (Google Account Admin)[https://myaccount.google.com/security].

### Example(s)
  #### Function Configuration
```yaml
functions:
- id: sendmail
  image: gcr.io/direktiv/functions/sendmail:1.0
  type: knative-workflow
```
   #### Basic Email
```yaml
- id: values
  type: noop
  transform:
    name: Direktiv
  transition: sendmail
- id: sendmail
  type: action
  action:
    function: sendmail
    secrets: ["smtppwd"]
    input: 
      smtp:
        server: smtp.gmail.com:587
        user: jens.gerke@direktiv.io
        password: jq(.secrets.smtppwd)
      emails:
      - from: Jens Gerke<jens.gerke@direktiv.io>
        to:
        - jens.gerke@direktiv.io
        cc:
        - jens.gerke@direktiv.io
        subject: This Is A Message
        message:
          name: message
          data: |-
            Hello jq(.name),

            this is an email.

            Good Bye
```
   #### Emails with attachment
```yaml
- id: sendmail
  type: action
  action:
    function: sendmail
    secrets: ["smtppwd"]
    input: 
      smtp:
        server: smtp.gmail.com:587
        user: jens.gerke@direktiv.io
        password: jq(.secrets.smtppwd)
      emails:
      - from: Jens Gerke<jens.gerke@direktiv.io>
        to: 
        - jens.gerke@direktiv.io
        attachments:
        - message
        subject: This Is A Message
        message:
          name: message
          data: |-
            This is the text and the attachment
```

   ### Secrets


- **smtppwd**: SMTP password






### Request



#### Request Attributes
[PostParamsBody](#post-params-body)

### Response
  Output of executed email commands.
#### Reponse Types
    
  

[PostOKBody](#post-o-k-body)
#### Example Reponses
    
```json
[
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
| emails | [][PostParamsBodyEmailsItems](#post-params-body-emails-items)| `[]*PostParamsBodyEmailsItems` | ✓ | | List of emails to send. |  |
| smtp | [PostParamsBodySMTP](#post-params-body-smtp)| `PostParamsBodySMTP` | ✓ | |  |  |


#### <span id="post-params-body-emails-items"></span> postParamsBodyEmailsItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| attachments | []string| `[]string` |  | | Files to attach to the email. Can be provided with Direktiv action `files` |  |
| bcc | []string| `[]string` |  | | Email addresses to send email to (blind copy) |  |
| cc | []string| `[]string` |  | | Email addresses to send email to (carbon copy) |  |
| from | string| `string` | ✓ | | Name used as `from` value, e.g. "My Name\<myname@direktiv.io\>" |  |
| message | [DirektivFile](#direktiv-file)| `apps.DirektivFile` |  | |  |  |
| subject | string| `string` |  | | Subject of the email |  |
| to | []string| `[]string` | ✓ | | Email addresses to send email to |  |
| verbose | boolean| `bool` |  | | Enable debug output |  |


#### <span id="post-params-body-smtp"></span> postParamsBodySmtp

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| password | string| `string` | ✓ | | Password for the SMTP server |  |
| server | string| `string` | ✓ | | SMTP server address | `smtp.sendgrid.net:587` |
| user | string| `string` | ✓ | | User name for the SMTP server |  |

 
