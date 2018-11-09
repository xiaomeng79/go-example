#!/bin/bash

set -e

go-bindata -o data/bindata.go -pkg data data/*.json
