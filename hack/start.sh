#!/bin/bash

server="./sample_apiserver"
let item=0
item=`ps -ef | grep $server | grep -v grep | wc -l`

if [ $item -eq 1 ]; then
	echo "The sample_apiserver is running, shut it down..."
	pid=`ps -ef | grep $server | grep -v grep | awk '{print $2}'`
	kill -9 $pid
fi

echo "Start sample_apiserver now ..."
make build
./build/pkg/cmd/sample_apiserver/sample_apiserver  >> sample_apiserver.log 2>&1 &
