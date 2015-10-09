#!/usr/bin/env bash

# TODO: add -race support
go test --cover $(go list ./... | grep -v /vendor/)
