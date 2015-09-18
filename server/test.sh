#!/usr/bin/env bash

go test --cover --race $(go list ./... | grep -v /vendor/)
