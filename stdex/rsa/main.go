package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"

	log "github.com/gitsang/golog"
	"go.uber.org/zap"
)

func Verify(publicKey *rsa.PublicKey, signature []byte, message string) {
	msg := []byte(message)

	msgHash := sha256.New()
	_, err := msgHash.Write(msg)
	if err != nil {
		panic(err)
	}
	msgHashSum := msgHash.Sum(nil)

	err = rsa.VerifyPSS(publicKey, crypto.SHA256, msgHashSum, signature, nil)
	if err != nil {
		log.Error("could not verify signature: ", zap.Error(err))
		return
	}

	log.Info("signature verified")
}

func Signature(privateKey *rsa.PrivateKey, message string) ([]byte, error) {
	msg := []byte(message)
	msgHash := sha256.New()
	_, err := msgHash.Write(msg)
	if err != nil {
		log.Error("hash write failed", zap.Error(err))
		return nil, err
	}
	msgHashSum := msgHash.Sum(nil)

	return rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, msgHashSum, nil)
}

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	publicKey := privateKey.PublicKey

	log.Info("generate key success",
		zap.Any("privateKey", privateKey.D),
		zap.Any("publicKey", publicKey.E))

	encryptedBytes, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &publicKey,
		[]byte("super secret message"), nil)
	if err != nil {
		panic(err)
	}
	log.Info("encrypt success", zap.ByteString("encrypted_bytes", encryptedBytes))

	decryptedBytes, err := privateKey.Decrypt(nil, encryptedBytes, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		panic(err)
	}
	log.Info("decrypt success", zap.String("decrypted_msg", string(decryptedBytes)))
}
