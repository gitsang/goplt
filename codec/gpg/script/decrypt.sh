#!/bin/bash

secret_key_path="./secret.key.gpg"

encrypted_file_path=${1}
if [ -z "${encrypted_file_path}" ]; then
    read -p "Enter the path to the file you want to decrypt: " encrypted_file_path
fi

decrypted_file_path=${2}
if [ -z "${decrypted_file_path}" ]; then
    decrypted_file_path="${encrypted_file_path%.*}.decrypted"
fi

fingerprint=$(gpg --import ${secret_key_path} 2>&1 | grep -oP '\b[0-9A-F]{16}\b')

pv < ${encrypted_file_path} | \
    gpg --decrypt \
        --decrypt-files - > ${decrypted_file_path}
