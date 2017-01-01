#! bin/bash

ssh wimill@182.92.82.68

cd $GOPATH/src/github.com/showntop/journey

godep go build 
go install