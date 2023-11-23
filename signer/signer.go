package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/mwiater/golangsignedbins/common"
)

func main() {
	// Load private key for signing
	privateKey, err := common.LoadPrivateKey(common.PrivateKey)
	if err != nil {
		common.PrintError(err)
	}

	// Sign the binary
	signature, err := common.Sign(common.BinaryToSign, privateKey)
	if err != nil {
		common.PrintError(err)
	}

	// Save the signature to a file
	err = os.WriteFile(common.BinarySignature, signature, 0644)
	if err != nil {
		common.PrintError(err)
	}

	fmt.Println("")
	color.Green("Heartbeat binary signed: " + common.BinaryToSign + " -> " + common.BinarySignature)
	fmt.Println("")
}
