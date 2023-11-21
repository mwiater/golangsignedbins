package common

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

var privateKeyPath = GetAbsPath(PrivateKey)
var publicKeyPath = GetAbsPath(PublicKey)
var binaryToSignPath = GetAbsPath(BinaryToSign)
var binarySignaturePath = GetAbsPath(BinarySignature)

func GetBasePath() string {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	return filepath.Join(basepath, "..")
}

func GetAbsPath(relativePath string) string {
	return filepath.Join(GetBasePath(), relativePath)
}

func TestLoadPrivateKey(t *testing.T) {
	// Test with a valid private key file
	privateKey, err := LoadPrivateKey(privateKeyPath)
	if err != nil {
		t.Errorf("Failed to load private key: %v", err)
	}
	if privateKey == nil {
		t.Errorf("Loaded key is nil")
	}

	// Test with an invalid file path
	_, err = LoadPrivateKey("../keys/non_existent_key.pem")
	if err == nil {
		t.Errorf("Expected an error when loading from a non-existent file path, but got none")
	}
}

func TestSign(t *testing.T) {
	// This test assumes that LoadPrivateKey is working correctly.
	privateKey, _ := LoadPrivateKey(privateKeyPath)

	// Test signing a valid binary
	_, err := Sign(binaryToSignPath, privateKey)
	if err != nil {
		t.Errorf("Failed to sign binary: %v", err)
	}
}

func TestLoadPublicKey(t *testing.T) {
	// Test with a valid public key file
	publicKey, err := LoadPublicKey(publicKeyPath)
	if err != nil {
		t.Errorf("Failed to load public key: %v", err)
	}
	if publicKey == nil {
		t.Errorf("Loaded key is nil")
	}

	// Test with an invalid file path
	_, err = LoadPublicKey("../keys/non_existent_key.pem")
	if err == nil {
		t.Errorf("Expected an error when loading from a non-existent file path, but got none")
	}
}

func TestVerifySignature(t *testing.T) {
	// This test depends on both LoadPublicKey and Sign working correctly.
	// You might need to generate a signature file before this test.
	publicKey, _ := LoadPublicKey(publicKeyPath)

	err := VerifySignature(binaryToSignPath, binarySignaturePath, publicKey)
	if err != nil {
		t.Errorf("Failed to verify signature: %v", err)
	}
}

func TestVerifyTamperedBinary(t *testing.T) {

	// Backup the original binary
	originalBinary, err := os.ReadFile(binaryToSignPath)
	if err != nil {
		t.Fatalf("Failed to backup the original binary: %v", err)
	}

	// Load the public key
	publicKey, err := LoadPublicKey(publicKeyPath)
	if err != nil {
		t.Fatalf("Failed to load public key: %v", err)
	}

	// Create a valid signature for the binary
	privateKey, _ := LoadPrivateKey(privateKeyPath)
	Sign(binaryToSignPath, privateKey)

	// Tamper with the binary
	f, err := os.OpenFile(binaryToSignPath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		t.Fatalf("Failed to open binary for tampering: %v", err)
	}
	_, err = f.Write([]byte("extra data"))
	f.Close()
	if err != nil {
		t.Fatalf("Failed to tamper with binary: %v", err)
	}

	// Verify the signature of the tampered binary
	err = VerifySignature(binaryToSignPath, binarySignaturePath, publicKey)
	if err == nil {
		t.Errorf("Verification passed, but expected it to fail for a tampered binary")
	}

	// Restore the original binary
	err = os.WriteFile(binaryToSignPath, originalBinary, 0644)
	if err != nil {
		t.Fatalf("Failed to restore the original binary: %v", err)
	}
}
