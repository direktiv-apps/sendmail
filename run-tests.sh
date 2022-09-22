#!/bin/bash

if [[ -z "${DIREKTIV_TEST_URL}" ]]; then
	echo "Test URL is not set, setting it to http://localhost:9191"
	DIREKTIV_TEST_URL="http://localhost:9191"
fi

if [[ -z "${DIREKTIV_SECRET_sendmailSecret}" ]]; then
	echo "Secret sendmailSecret is required, set it with DIREKTIV_SECRET_sendmailSecret"
	exit 1
fi

docker run --network=host -v `pwd`/tests/:/tests direktiv/karate java -DtestURL=${DIREKTIV_TEST_URL} -Dlogback.configurationFile=/logging.xml -DsendmailSecret="${DIREKTIV_SECRET_sendmailSecret}"  -jar /karate.jar /tests/v1.0/karate.yaml.test.feature ${*:1}