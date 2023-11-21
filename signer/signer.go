package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/mwiater/golangsignedbins/common"
)

func main() {
	// Load private key (for signing)
	privateKey, err := common.LoadPrivateKey(common.PrivateKey)
	if err != nil {
		color.Red("Error:" + err.Error())
		os.Exit(1)
	}

	// Sign the binary
	signature, err := common.Sign(common.BinaryToSign, privateKey)
	if err != nil {
		color.Red("Failed to sign binary: " + err.Error())
		os.Exit(1)
	}

	// Save the signature to a file
	err = os.WriteFile(common.BinarySignature, signature, 0644)
	if err != nil {
		color.Red("Failed to write signature file: " + err.Error())
		os.Exit(1)
	}

	fmt.Println("")
	color.Green("Heartbeat binary signed: " + common.BinaryToSign + " -> " + common.BinarySignature)
	fmt.Println("")
}
