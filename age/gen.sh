#!/bin/bash

key_file=${1:-./key.age}

rm -f ${key_file}
age-keygen -o ${key_file} &> /dev/null

public_key=$(sed -n 's/.*public key: \(.*\)/\1/p' ${key_file} | tr -d '\n')
secret_key=$(sed -n 's/.*AGE-SECRET-KEY-\(.*\)/\1/p' ${key_file} | tr -d '\n')
echo "public key: ${public_key}"
echo "secret key: ${secret_key}"
