.PHONY: build compile compile-interactive build-docker deps serve start

GO_ENV_CONTIANER = joshuaanderson/go-build-env:0.0.0
CBX_PACKAGE = github.com/chatterbox-irc/chatterbox
CBX_ROOT = $(GOPATH)/src/$(CBX_PACKAGE)
CONTIANER_CBX_ROOT = /go/src/$(CBX_PACKAGE)

build: compile build-docker

build-docker:
	docker-compose build

compile:
	docker run --rm \
	-v $(CBX_ROOT):$(CONTIANER_CBX_ROOT) \
	-w "$(CONTIANER_CBX_ROOT)/server" $(GO_ENV_CONTIANER) \
	./compile.sh

compile-interactive:
	docker run --rm -it \
	-v $(CBX_ROOT):$(CONTIANER_CBX_ROOT) \
	-w "$(CONTIANER_CBX_ROOT)/server" -e "DEBIAN_FRONTEND=noninteractive" -e "GO15VENDOREXPERIMENT=1" \
	$(GO_ENV_CONTIANER) bash

deps:
	docker run --rm \
	-v $(CBX_ROOT):$(CONTIANER_CBX_ROOT) \
	-w "$(CONTIANER_CBX_ROOT)/server" $(GO_ENV_CONTIANER) \
	glide install

serve: build start

start:
	docker-compose up
