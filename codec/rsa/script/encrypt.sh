#!/bin/bash

encrypt_file_path=${1}
if [ -z "${encrypt_file_path}" ]; then
    read -p "Enter the path to the file you want to encrypt: " file_path
fi
encrypted_file_path=${2:-$(basename ${encrypt_file_path}).enc}
encrypted_passphrase_file=${encrypted_file_path%.enc}.passphrase.enc
public_key=${3:-public.pem}

echo "Encrypting ${encrypt_file_path} to ${encrypted_file_path} with public key ${public_key}"

passphrase=$(openssl rand -base64 32)
pv < ${encrypt_file_path} | \
    openssl enc -e -a -salt -k ${passphrase} | \
    dd of=${encrypted_file_path}
echo ${passphrase} | \
    openssl rsautl -encrypt -pubin -pubin -inkey ${public_key} \
    -out ${encrypted_passphrase_file}

