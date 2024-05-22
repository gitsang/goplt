#!/bin/bash

encrypted_file_path=${1}
if [ -z "${encrypted_file_path}" ]; then
    read -p "Enter the path to the file you want to decrypt: " encrypted_file_path
fi
encrypted_passphrase_file=${encrypted_file_path%.enc}.passphrase.enc
decrypted_file_path=${2}
if [ -z "${decrypted_file_path}" ]; then
    decrypted_file_path=${encrypted_file_path%.enc}
fi
private_key=${3:-private.pem}

echo "Decrypting ${encrypted_file_path} to ${decrypted_file_path} with passphrase ${encrypted_passphrase_file} and private key ${private_key}"

passphrase=$(openssl rsautl -decrypt -inkey ${private_key} -in ${encrypted_passphrase_file})
pv < ${encrypted_file_path} | \
    openssl enc -d -a -salt -d -k ${passphrase} | \
    dd of=${decrypted_file_path}
