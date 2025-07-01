#!/bin/bash
goos=$(GOOS)
goarch=${GOARCH}
APP_NAME=DataReportUpdate

default: build

build:
	CGO_ENABLED=0 GOOS=$(goos) GOARCH=$(goarch) go build -o ${APP_NAME} report_update.go

run:
	./${APP_NAME} -env=local -conf=./config

push-dev:
	rsync -av -overwrite -e "ssh -i ~/bt.pem" ./${APP_NAME} user@233.233.233.233:/home/centos/bt-coordinator/

run-dev:
	ssh -i "~/bt.pem" user@233.233.233.233 ./bt-coordinator/run.sh restart

clean:
	rm ./${APP_NAME}
	go clean -i

help:
	@echo "make: make build"

	@echo "make build: build, support OS and ARCH, like: make build OS=darwin ARCH=amd64"

	@echo "make clean: remove object files and cached files"

	@echo "make push-dev: send then built 'bt-coordinator' file to dev remote"
