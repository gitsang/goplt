#!/bin/bash

public_key_path="./public.key.gpg"

encrypt_file_path=${1}
if [ -z "${encrypt_file_path}" ]; then
    read -p "Enter the path to the file you want to encrypt: " file_path
fi

fingerprint=$(gpg --import ${public_key_path} 2>&1 | grep -oP '\b[0-9A-F]{16}\b')

pv < ${encrypt_file_path} | \
    gpg --encrypt \
        --recipient ${fingerprint} \
        --encrypt-files - > "${encrypt_file_path}".gpg
