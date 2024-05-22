#!/bin/bash

private_key_file=${1:-./private.pem}
public_key_file=${2:-./public.pem}
key_size=${3:-4096}

openssl genrsa -out ${private_key_file} ${key_size}
openssl rsa -in ${private_key_file} -pubout -out ${public_key_file}
