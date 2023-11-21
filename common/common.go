package common

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

const (
	PrivateKey      = "./keys/private_key.pem"
	PublicKey       = "./keys/public_key.pem"
	BinaryToSign    = "./heartbeat/bin/heartbeat"
	BinarySignature = "./signatures/heartbeat.sig"
)

// LoadPrivateKey loads a private key from a given file.
func LoadPrivateKey(path string) (*rsa.PrivateKey, error) {
	privateKeyData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(privateKeyData)
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing the key")
	}

	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	privKey, _ := (priv).(*rsa.PrivateKey)
	return privKey, nil
}

// LoadPublicKey loads a public key from a given file.
func LoadPublicKey(path string) (*rsa.PublicKey, error) {
	publicKeyData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(publicKeyData)
	if block == nil {
		return nil, err
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		return nil, err
	}
}

// Sign signs the binary data using a private key and returns the signature.
func Sign(binaryPath string, privateKey *rsa.PrivateKey) ([]byte, error) {
	binaryData, err := os.ReadFile(binaryPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read binary data: %v", err)
	}
	hashed := sha256.Sum256(binaryData)

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return nil, fmt.Errorf("ailed to sign data: %v", err)
	}
	return signature, nil
}

// VerifySignature verifies the binary's signature.
func VerifySignature(binaryPath, signaturePath string, publicKey *rsa.PublicKey) error {
	binaryData, err := os.ReadFile(binaryPath)
	if err != nil {
		return fmt.Errorf("failed to read binary data: %v", err)
	}
	signature, err := os.ReadFile(signaturePath)
	if err != nil {
		return fmt.Errorf("failed to read signature data: %v", err)
	}
	hashed := sha256.Sum256(binaryData)

	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return nil
}
