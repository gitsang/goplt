#!/bin/bash

encrypt_file_path=${1}
if [ -z "${encrypt_file_path}" ]; then
    read -p "Enter the path to the file you want to encrypt: " file_path
fi
encrypted_file_path=${2:-$(basename ${encrypt_file_path}).gpg}
public_key=${3:-"./public.key.gpg"}

echo "Encrypting ${encrypt_file_path} to ${encrypted_file_path} with public key ${public_key}"

fingerprint=$(gpg --import ${public_key} 2>&1 | grep -oP '\b[0-9A-F]{16}\b')
pv < ${encrypt_file_path} | \
    gpg --encrypt --recipient ${fingerprint} --encrypt-files - | \
    dd of=${encrypted_file_path}
