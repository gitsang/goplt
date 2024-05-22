#!/bin/bash

encrypted_file_path=${1}
if [ -z "${encrypted_file_path}" ]; then
    read -p "Enter the path to the file you want to decrypt: " encrypted_file_path
fi

decrypted_file_path=${2}
if [ -z "${decrypted_file_path}" ]; then
    decrypted_file_path=${encrypted_file_path%.age}
fi

private_key=${3:-./key.age}

echo "Decrypting ${encrypted_file_path} to ${decrypted_file_path} with private key ${private_key}"

pv < ${encrypted_file_path} | \
    age -d -i ${private_key} | \
    dd of=${decrypted_file_path}
