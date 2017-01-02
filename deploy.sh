#!/bin/bash

#ssh wimill@182.92.82.68

#工作目录
cd $GOPATH/src/github.com/showntop/journey

mkdir -p /home/wimill/runtime
mkdir -p /home/wimill/runtime/logs

cp ./supervisord.conf /home/wimill/runtime

godep go build 
go install

/bin/ps auxf | /bin/grep "journey" | /bin/grep -v grep | /usr/bin/awk '{print $2}' | /usr/bin/xargs -n1 /bin/kill -2

nohup supervisord -c /home/wimill/runtime/supervisord.conf > /home/wimill/runtime/supervisord.log &