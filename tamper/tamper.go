package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/mwiater/golangsignedbins/common"
)

func main() {
	// Tamper with the binary
	f, err := os.OpenFile(common.BinaryToSign, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		common.PrintError(err)
	}
	_, err = f.Write([]byte("extra data"))
	f.Close()
	if err != nil {
		common.PrintError(err)
	}

	fmt.Println()
	fmt.Printf("%s %s\n", color.HiYellowString("Warning: File has been modified:"), common.BinaryToSign)
	fmt.Println()
}
