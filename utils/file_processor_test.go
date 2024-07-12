package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func copyTestYAMLFile(t *testing.T, destDir, srcFile string) string {
	t.Helper()
	destFile := filepath.Join(destDir, filepath.Base(srcFile))
	input, err := ioutil.ReadFile(srcFile)
	if err != nil {
		t.Fatalf("Failed to read source file: %v", err)
	}
	if err := ioutil.WriteFile(destFile, input, 0644); err != nil {
		t.Fatalf("Failed to write destination file: %v", err)
	}
	return destFile
}

func TestReadYAMLFiles(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Copy the test YAML file
	copyTestYAMLFile(t, tempDir, "../testdata/test_contract.yaml")

	// Create an additional non-YAML file
	if err := ioutil.WriteFile(filepath.Join(tempDir, "not_yaml.txt"), []byte("This is not a YAML file"), 0644); err != nil {
		t.Fatalf("Failed to create non-YAML file: %v", err)
	}

	yamlFiles, err := ReadYAMLFiles(tempDir)
	if err != nil {
		t.Fatalf("Failed to read YAML files: %v", err)
	}

	expectedFiles := []string{"test_contract.yaml"}
	if len(yamlFiles) != len(expectedFiles) {
		t.Fatalf("Expected %d YAML files, got %d", len(expectedFiles), len(yamlFiles))
	}

	for i, file := range expectedFiles {
		if yamlFiles[i] != file {
			t.Errorf("Expected file %s, got %s", file, yamlFiles[i])
		}
	}
}

func TestProcessYAMLFile(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Copy the test YAML file
	yamlFile := copyTestYAMLFile(t, tempDir, "../testdata/test_contract.yaml")

	// Process the YAML file
	err = ProcessYAMLFile(filepath.Base(yamlFile), tempDir)
	if err != nil {
		t.Fatalf("Failed to process YAML file: %v", err)
	}

	// Check if the SVG file was created
	svgFile := filepath.Join(tempDir, "test_contract.svg")
	if _, err := os.Stat(svgFile); os.IsNotExist(err) {
		t.Fatalf("SVG file was not created")
	}
}
