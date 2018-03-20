#!/bin/bash
export GOOS=linux
export CGO_ENABLED=0

# cd healthchecker;go get;go build -o healthchecker-linux-amd64;echo built `pwd`;cd ..
### account service module
cd accountservice;go get;go build -o _build/accountservice-linux-amd64;echo built `pwd`;cd ..

docker rmi gojobs/accountservice registry.ckmro.com:1180/gojobs/accountservice
docker build -t gojobs/accountservice accountservice/
docker tag gojobs/accountservice registry.ckmro.com:1180/gojobs/accountservice
docker push registry.ckmro.com:1180/gojobs/accountservice

### account service module
cd resumeservice;go get;go build -o _build/resumeservice-linux-amd64;echo built `pwd`;cd ..
docker rmi gojobs/resumeservice registry.ckmro.com:1180/gojobs/resumeservice
docker build -t gojobs/resumeservice resumeservice/
docker tag gojobs/resumeservice registry.ckmro.com:1180/gojobs/resumeservice
docker push registry.ckmro.com:1180/gojobs/resumeservice