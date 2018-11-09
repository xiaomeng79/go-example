#!/bin/bash


proto() {
    dirname=./srv/$1/proto
    if [ -d $dirname ];then
		for f in $dirname/*.proto; do \
		    if [ -f $f ];then \
                protoc -I. --micro_out=. --go_out=. $f; \
                echo compiled protoc: $f; \
            fi \
		done \
	fi
}

proto user
proto account
proto auth