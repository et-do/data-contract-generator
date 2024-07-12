package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ReadYAMLFiles reads all YAML files in the specified directory
func ReadYAMLFiles(directory string) ([]string, error) {
	var yamlFiles []string
	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, fmt.Errorf("error reading directory: %v", err)
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".yaml") {
			yamlFiles = append(yamlFiles, file.Name())
		}
	}
	return yamlFiles, nil
}

// ProcessYAMLFile processes a single YAML file, generating an SVG
func ProcessYAMLFile(yamlFile, directory string) error {
	yamlPath := filepath.Join(directory, yamlFile)
	contract, err := ParseYAML(yamlPath)
	if err != nil {
		return fmt.Errorf("error parsing YAML file %s: %v", yamlFile, err)
	}

	svgData, err := GenerateSVG(contract)
	if err != nil {
		return fmt.Errorf("error generating SVG for %s: %v", yamlFile, err)
	}

	outputFile := strings.TrimSuffix(yamlPath, ".yaml") + ".svg"
	if err := SaveSVG(svgData, outputFile); err != nil {
		return fmt.Errorf("error saving SVG to %s: %v", outputFile, err)
	}

	fmt.Printf("SVG file generated: %s\n", outputFile)
	return nil
}

// ProcessDirectory processes all YAML files in the specified directory
func ProcessDirectory(directory string) error {
	yamlFiles, err := ReadYAMLFiles(directory)
	if err != nil {
		return err
	}

	for _, yamlFile := range yamlFiles {
		if err := ProcessYAMLFile(yamlFile, directory); err != nil {
			fmt.Printf("Error processing file %s: %v\n", yamlFile, err)
		}
	}
	return nil
}
