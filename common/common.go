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

	"github.com/fatih/color"
)

const (
	PrivateKey      = "./keys/private_key.pem"
	PublicKey       = "./keys/public_key.pem"
	BinaryToSign    = "./heartbeat/bin/heartbeat"
	BinarySignature = "./signatures/heartbeat.sig"
)

// LoadPrivateKey loads an RSA private key from the specified file path.
// The key should be in PEM format. It returns an rsa.PrivateKey and any error encountered.
func LoadPrivateKey(path string) (*rsa.PrivateKey, error) {
	privateKeyData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("load private key: %v", err)
	}

	block, _ := pem.Decode(privateKeyData)
	if block == nil {
		return nil, fmt.Errorf("parse PEM block: key not found")
	}

	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("parse private key: %v", err)
	}

	privKey, ok := priv.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("type assertion to rsa.PrivateKey failed")
	}
	return privKey, nil
}

// LoadPublicKey loads an RSA public key from the specified file path.
// The key should be in PEM format. It returns an rsa.PublicKey and any error encountered.
func LoadPublicKey(path string) (*rsa.PublicKey, error) {
	publicKeyData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("load public key: %v", err)
	}

	block, _ := pem.Decode(publicKeyData)
	if block == nil {
		return nil, fmt.Errorf("parse PEM block: key not found")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("parse public key: %v", err)
	}

	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("type assertion to rsa.PublicKey failed")
	}
	return rsaPub, nil
}

// Sign creates a signature for a binary file located at binaryPath using the provided RSA private key.
// It returns the signature as a byte slice and any error encountered.
func Sign(binaryPath string, privateKey *rsa.PrivateKey) ([]byte, error) {
	binaryData, err := os.ReadFile(binaryPath)
	if err != nil {
		return nil, fmt.Errorf("read binary data: %v", err)
	}

	hashed := sha256.Sum256(binaryData)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return nil, fmt.Errorf("sign data: %v", err)
	}
	return signature, nil
}

// VerifySignature checks the signature of a binary file located at binaryPath against a signature file at signaturePath.
// It uses the provided RSA public key for verification. It returns any error encountered in the verification process.
func VerifySignature(binaryPath, signaturePath string, publicKey *rsa.PublicKey) error {
	binaryData, err := os.ReadFile(binaryPath)
	if err != nil {
		return fmt.Errorf("read binary data: %v", err)
	}

	signature, err := os.ReadFile(signaturePath)
	if err != nil {
		return fmt.Errorf("read signature data: %v", err)
	}

	hashed := sha256.Sum256(binaryData)
	if err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signature); err != nil {
		return fmt.Errorf("verify signature: %v", err)
	}
	return nil
}

func PrintError(err error) {
	fmt.Println()
	fmt.Println(color.HiRedString("[Error]"), err.Error())
	fmt.Println()
	os.Exit(1)
}
