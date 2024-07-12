package utils

import (
	"data-contract-generator/models"
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestGenerateSVG(t *testing.T) {
	// Read the test YAML file
	data, err := os.ReadFile("../testdata/test_contract.yaml")
	if err != nil {
		t.Fatalf("Failed to read test data file: %v", err)
	}

	// Unmarshal the YAML data into the DataContract struct
	var contract models.DataContract
	err = yaml.Unmarshal(data, &contract)
	if err != nil {
		t.Fatalf("Failed to unmarshal YAML: %v", err)
	}

	// Generate the SVG data from the DataContract struct
	svgData, err := GenerateSVG(&contract)
	if err != nil {
		t.Fatalf("Failed to generate SVG: %v", err)
	}

	// Define the output file path
	outputFilePath := filepath.Join("../testdata", "test_contract.svg")

	// Write the SVG data to the output file
	if err := os.WriteFile(outputFilePath, svgData, 0644); err != nil {
		t.Fatalf("Failed to write SVG data to file: %v", err)
	}

	t.Logf("SVG data written to file: %s", outputFilePath)

	// Verify that the SVG file was created and has content
	fileInfo, err := os.Stat(outputFilePath)
	if err != nil {
		t.Fatalf("Failed to get file info: %v", err)
	}

	if fileInfo.Size() == 0 {
		t.Fatalf("SVG file is empty")
	}

	t.Logf("SVG file size: %d bytes", fileInfo.Size())

	// Optionally, read and print the file content for debugging
	svgContent, err := os.ReadFile(outputFilePath)
	if err != nil {
		t.Fatalf("Failed to read generated SVG file: %v", err)
	}
	t.Logf("Generated SVG content:\n%s", svgContent)
}
