package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/mwiater/golangsignedbins/common"
)

func main() {
	// Load public key (for verification)
	publicKey, err := common.LoadPublicKey(common.PublicKey)
	if err != nil {
		color.Red("Error: " + err.Error())
		os.Exit(1)
	}

	// Verify the signature
	err = common.VerifySignature(common.BinaryToSign, common.BinarySignature, publicKey)
	if err != nil {
		color.Red("Signature verification failed: " + err.Error())
		os.Exit(1)
	}
	fmt.Println("")
	color.Green("Heartbeat binary verified..." + common.BinarySignature + " -> " + common.BinaryToSign)
	fmt.Println("")

	// Execute the binary
	cmd := exec.Command(common.BinaryToSign)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		color.Red("Failed to execute binary: " + err.Error())
		os.Exit(1)
	}
}
