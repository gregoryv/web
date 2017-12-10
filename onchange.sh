#!/bin/bash -e
GOPATH=$HOME
go install github.com/gregoryv/website/...
#go test -cover -coverprofile /tmp/c.out . > /tmp/test.out
