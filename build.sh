#!/bin/bash

GOPROXY=https://goproxy.io
go build -i -v
cd console && npm install && npm run build

echo 'build vega done'
