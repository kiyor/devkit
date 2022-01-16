#!/bin/bash

oapp=devkit
napp=$(basename `pwd`)
oorg=kiyor
norg=$1

if [[ ! -z $1 ]]; then
	echo '$1 not exist; $1 = new org name'
fi

list='go.mod main.go routers/routers.go'
for i in ${list}; do
	sed -i "s|${oorg}/${oapp}|${norg}/${napp}|g" ${i}
done

list='docker-compose.yml Dockerfile go.mod main.go routers/routers.go'
for i in ${list}; do
	sed -i "s|${oapp}|${napp}|g" ${i}
done

go mod tidy
go mod vendor

make build
