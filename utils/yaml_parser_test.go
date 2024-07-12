package utils

import (
	"testing"
)

func TestParseYAML(t *testing.T) {
	contract, err := ParseYAML("../testdata/test_contract.yaml")
	if err != nil {
		t.Fatalf("Failed to parse YAML: %v", err)
	}

	if contract.Name != "Customer Data Contract" {
		t.Errorf("Expected name to be 'Customer Data Contract', got %s", contract.Name)
	}
	// Additional assertions for other fields can be added here
}
