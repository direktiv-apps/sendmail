#!/bin/sh

docker build -t sendmail . && docker run -p 9191:8080 sendmail