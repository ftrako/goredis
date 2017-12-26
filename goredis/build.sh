#!/bin/bash 

echo "create preset dirs..."
rm -rf giredis
mkdir -p giredis/bin

echo "build..."
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64
go build -o giredis/bin/giredis

echo "finished!"

#窗口不自动消失
`read`