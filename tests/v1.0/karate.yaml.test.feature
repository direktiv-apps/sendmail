
Feature: Basic

# The secrects can be used in the payload with the following syntax #(mysecretname)
Background:
* def smtppwd = karate.properties['smtppwd']

Scenario: sendmail

	Given url karate.properties['testURL']

	And path '/'
	And header Direktiv-ActionID = 'development'
	And header Direktiv-TempDir = '/tmp'
	And request
	"""
	{
		"smtp": {
			"server": "smtp.gmail.com:587",
			"user": "jens.gerke@direktiv.io",
			"password": "#(smtppwd)"
		},
		"emails": [
			{
				"from": "Jens Gerke<jens.gerke@direktiv.io>",
				"to": [
					"jens.gerke@direktiv.io"
				],
				"bcc": [
					"jens.gerke@direktiv.io"
				],
				"cc": [
					"jens.gerke@direktiv.io"
				],
				"attachments": [
					"message"
				],
				"contentType": "text/html",
				"verbose": true,
				"subject": "This Is A Message",
				"message": {
					"name": "message",
					"data": "Hello World!!!<br/>\nJENS\n\t\t\tTEST ME"
				}
			}
		]
	}
	"""
	When method POST
	Then status 200
	And match $ ==
	"""
	{
	"sendmail": [
	{
		"result": "#notnull",
		"success": true
	}
	]
	}
	"""
	