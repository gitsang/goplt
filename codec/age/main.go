package main

import (
	"flag"
	"io"
	"os"

	"filippo.io/age"
)

func main() {
	var identityFlag string
	flag.StringVar(&identityFlag, "i", "identity", "identities file `FILE`")
	flag.Parse()

	identityFile, err := os.Open(identityFlag)
	if err != nil {
		panic(err)
	}

	identities, err := age.ParseIdentities(identityFile)
	if err != nil {
		panic(err)
	}

	r, err := age.Decrypt(os.Stdin, identities...)
	if err != nil {
		panic(err)
	}

	if _, err := io.Copy(os.Stdout, r); err != nil {
		panic(err)
	}

	return
}
