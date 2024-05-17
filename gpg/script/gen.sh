#!/bin/bash

gpg_key_id=$(gpg --full-gen-key --batch <<EOF 2>&1 | grep -oE '[0-9A-F]{40}'
Key-Type: 1
Key-Length: 4096
Subkey-Type: 1
Subkey-Length: 4096
Name-Real: Yealinkops
Name-Email: yealinkops@yealink.com
Expire-Date: 0
%no-protection
EOF
)

echo "gpg key id: $gpg_key_id"

rm -rf secret.key.gpg
rm -rf public.key.gpg

gpg --output secret.key.gpg --armor --export-secret-keys $gpg_key_id
gpg --output public.key.gpg --armor --export $gpg_key_id
