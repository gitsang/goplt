#!/bin/bash

for pb in *.proto
do
  protoc --go_out=plugins=grpc:./ "$pb"
done
