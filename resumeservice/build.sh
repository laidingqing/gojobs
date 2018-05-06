#!/bin/bash
export GOOS=linux
export CGO_ENABLED=0

go get;go build -o _build/resumeservice-linux-amd64;echo built `pwd`
