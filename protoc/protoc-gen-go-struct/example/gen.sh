#!/bin/bash

pushd ..
    bash -x ./build.sh
popd

protoc \
    --plugin=protoc-gen-go-struct=../protoc-gen-go-struct \
    --go-struct_out=. \
    hello.proto
