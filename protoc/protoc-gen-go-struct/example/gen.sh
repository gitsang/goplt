#!/bin/bash

pushd ..
    bash ./build.sh
popd

protoc \
    -I . \
    -I ../.. \
    --plugin=protoc-gen-go-struct=../protoc-gen-go-struct \
    --go-struct_out=. \
    hello.proto
