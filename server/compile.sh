#!/usr/bin/env bash
export GO15VENDOREXPERIMENT=1
export DEBIAN_FRONTEND=noninteractive
CGO_ENABLED=0 go build -o cbx .
