package main

import (
	"fmt"
	"io/ioutil"
	"data-contract-generator/models"
	"data-contract-generator/utils"
	"os"
	"path/filepath"
	"strings"
)

func processYAMLFile(file string, directory string) {
	yamlFile := filepath.Join(directory, file)
	contract, err := utils.ParseYAML(yamlFile)
	if err != nil {
		fmt.Printf("Error parsing YAML file %s: %v\n", file, err)
		return
	}

	svgData, err := utils.GenerateSVG(contract)
	if err != nil {
		fmt.Printf("Error generating SVG for %s: %v\n", file, err)
		return
	}

	outputFile := strings.TrimSuffix(yamlFile, ".yaml") + ".svg"
	err = utils.SaveSVG(svgData, outputFile)
	if err != nil {
		fmt.Printf("Error saving SVG to %s: %v\n", outputFile, err)
		return
	}

	fmt.Printf("SVG file generated: %s\n", outputFile)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <directory>")
		return
	}

	directory := os.Args[1]
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".yaml") {
			processYAMLFile(file.Name(), directory)
		}
	}
}
