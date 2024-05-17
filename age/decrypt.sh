#!/bin/bash

encrypted_file_path=${1}
if [ -z "${encrypted_file_path}" ]; then
    read -p "Enter the path to the file you want to decrypt: " encrypted_file_path
fi

decrypted_file_path=${2}
if [ -z "${decrypted_file_path}" ]; then
    decrypted_file_path="${encrypted_file_path%.*}.decrypted"
fi

key_file=${3:-./key.age}

pv < ${encrypted_file_path} | \
    age -d -i ${key_file} > ${decrypted_file_path}
