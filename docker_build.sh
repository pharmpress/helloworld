#!/bin/bash

cat >.payload <<EOF
go run third_party.go clean -i net
go run third_party.go install -tags netgo std
CGO_ENABLED=0 go run third_party.go build -a -tags netgo --ldflags '-w -extldflags=-static' -o bin/helloworld-linux64-static github.com/pharmpress/helloworld
EOF

echo "building statically-linked helloworld..."
docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:latest bash .payload
rm -f .payload


docker build -t quay.io/pharmpress/helloworld .
