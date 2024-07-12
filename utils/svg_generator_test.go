package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"my_project/models"
	"os"
	"testing"
)

func TestGenerateSVG(t *testing.T) {
	data, err := ioutil.ReadFile("../testdata/test_contract.yaml")
	if err != nil {
		t.Fatalf("Failed to read test data file: %v", err)
	}

	var contract models.DataContract
	err = yaml.Unmarshal(data, &contract)
	if err != nil {
		t.Fatalf("Failed to unmarshal YAML: %v", err)
	}

	svgData, err := GenerateSVG(&contract)
	if err != nil {
		t.Fatalf("Failed to generate SVG: %v", err)
	}

	tempFile, err := ioutil.TempFile("", "test*.svg")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	if err := SaveSVG(svgData, tempFile.Name()); err != nil {
		t.Fatalf("Failed to save SVG: %v", err)
	}

	// Further tests could include checking the content of the SVG file
}
