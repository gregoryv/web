#!/bin/bash -e
GOPATH=$HOME
go install github.com/gregoryv/web/...
go test -cover -coverprofile /tmp/c.out ./site
