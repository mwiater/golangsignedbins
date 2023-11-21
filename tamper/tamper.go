package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/mwiater/golangsignedbins/common"
)

func main() {
	// Path to the binary

	// Tamper with the binary
	f, err := os.OpenFile(common.BinaryToSign, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Printf("Failed to open binary for modification: %v", err)
	}
	_, err = f.Write([]byte("extra data"))
	f.Close()
	if err != nil {
		fmt.Printf("Failed to modify binary: %v", err)
	}

	fmt.Println()
	color.Red("File has been modified: " + common.BinaryToSign)
	fmt.Println()
}
