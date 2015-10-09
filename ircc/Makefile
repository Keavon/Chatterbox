.PHONY: compile-cli start-ircd stop-ircd

include ../includes.mk

compile-cli:
	cli/compile.sh

deps:
	docker run --rm \
	-v $(CBX_ROOT):$(CONTIANER_CBX_ROOT) \
	-w "$(CONTIANER_CBX_ROOT)/ircc" -e "DEBIAN_FRONTEND=noninteractive" -e "GO15VENDOREXPERIMENT=1" \
	$(GO_ENV_CONTIANER) \
	glide install

test:
	./test.sh

start-ircd:
	docker run --name=cbx-ircd -d -p 6667:6667 xena/elemental-ircd
	@echo Development IRCD running at localhost:6667
	@echo To stop the ircd, run make stop-ircd

stop-ircd:
	docker rm -f cbx-ircd
