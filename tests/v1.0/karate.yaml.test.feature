
Feature: Basic

# The secrects can be used in the payload with the following syntax #(mysecretname)
Background:
* def smtppwd = karate.properties['smtppwd']

Scenario: get request

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
					"gerke74@gmail.com"
				],
				"bcc": [
					"jens.gerke@direktiv.io"
				],
				"cc": [
					"jgerke74@gmail.com"
				],
				"attachments": [
					"message"
				],
				"verbose": true,
				"subject": "This Is A Message",
				"message": {
					"name": "message",
					"data": "Hello World!!!\nJENS\n\t\t\tTEST ME"
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
	