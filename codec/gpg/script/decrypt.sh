#!/bin/bash

encrypted_file_path=${1}
if [ -z "${encrypted_file_path}" ]; then
    read -p "Enter the path to the file you want to decrypt: " encrypted_file_path
fi
decrypted_file_path=${2}
if [ -z "${decrypted_file_path}" ]; then
    decrypted_file_path=${encrypted_file_path%.gpg}
fi
private_key=${3:-./secret.key.gpg}

echo "Decrypting ${encrypted_file_path} to ${decrypted_file_path} with private key ${private_key}"

fingerprint=$(gpg --import ${private_key} 2>&1 | grep -oP '\b[0-9A-F]{16}\b')
pv < ${encrypted_file_path} | \
    gpg --decrypt --decrypt-files - |
    dd of=${decrypted_file_path}
