package main

import (
	"data-contract-generator/utils"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: yaml_to_svg_converter <directory>")
		return
	}

	directory := os.Args[1]
	if err := utils.ProcessDirectory(directory); err != nil {
		fmt.Printf("Error processing directory: %v\n", err)
	}
}
