#!/bin/bash

encrypt_file_path=${1}
if [ -z "${encrypt_file_path}" ]; then
    read -p "Enter the path to the file you want to encrypt: " file_path
fi

public_key=${2:-$(sed -n 's/.*public key: \(.*\)/\1/p' ./key.age | tr -d '\n')}
if [ -z "${public_key}" ]; then
    read -p "Enter the public key: " public_key
fi

pv < ${encrypt_file_path} | \
    age -r "${public_key}" > $(basename ${encrypt_file_path}).age
