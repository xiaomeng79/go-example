#!/bin/bash

set -e

#build
build() {
    #判断bin是否存在
    if [ ! -d bin ];then
    mkdir bin
    fi
    #build

    dirname=./cmd/$1
    if [ -d $dirname ];then
		for f in $dirname/$2.go; do \
		    if [ -f $f ];then \
		        CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -i -o bin/$1_$2 -tags $1_$2 .
                echo build over: $1_$2; \
            fi \
		done \
	fi
}

#全部build
allbuild() {
    build user api
    build user srv
    build account api
    build account srv
    build auth srv
}
#判断如何build
case $1 in
    allbuild) echo "全部build"
    allbuild
    ;;
    build) echo "build:"$2,$3
    if [ -z $2 -o -z $3 ];then
    echo "参数错误"
    exit 2
    fi
    build $2 $3
    ;;
    *)
    echo "build error"
    exit 2
    ;;
esac
