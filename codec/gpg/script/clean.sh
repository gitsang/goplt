#!/bin/bash

export GPG_TTY=$(tty)
gpg --list-secret-keys 2>&1 | grep -A 1 "sec" | grep -vE "sec|-" | awk '{print $1}' | xargs gpg --delete-secret-keys
gpg --list-public-keys 2>&1 | grep -A 1 "pub" | grep -vE "pub|-" | awk '{print $1}' | xargs gpg --delete-keys
gpg --list-secret-keys
gpg --list-public-keys

rm -fr *.gpg
