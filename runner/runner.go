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
		common.PrintError(err)
	}

	// Verify the signature
	err = common.VerifySignature(common.BinaryToSign, common.BinarySignature, publicKey)
	if err != nil {
		common.PrintError(err)
	}

	fmt.Println()
	fmt.Printf("%s %s\n", color.HiGreenString("Heartbeat binary verified:"), common.BinaryToSign)
	fmt.Println()
	color.HiGreen("Executing...")
	fmt.Println()

	// Execute the binary
	cmd := exec.Command(common.BinaryToSign)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		common.PrintError(err)
	}
}
