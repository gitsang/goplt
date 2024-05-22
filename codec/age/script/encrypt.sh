#!/bin/bash

encrypt_file_path=${1}
if [ -z "${encrypt_file_path}" ]; then
    read -p "Enter the path to the file you want to encrypt: " file_path
fi
encrypted_file_path=${2:-$(basename ${encrypt_file_path}).age}
public_key=${3:-$(sed -n 's/.*public key: \(.*\)/\1/p' ./key.age | tr -d '\n')}

echo "Encrypting ${encrypt_file_path} to ${encrypted_file_path} with public key ${public_key}"

pv < ${encrypt_file_path} | \
    age -r "${public_key}" | \
    dd of=${encrypted_file_path}
