package models

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)

func TestUnmarshalDataContract(t *testing.T) {
	data, err := ioutil.ReadFile("../testdata/test_contract.yaml")
	if err != nil {
		t.Fatalf("Failed to read test data file: %v", err)
	}

	var contract DataContract
	err = yaml.Unmarshal(data, &contract)
	if err != nil {
		t.Fatalf("Failed to unmarshal YAML: %v", err)
	}

	if contract.Name != "Customer Data Contract" {
		t.Errorf("Expected name to be 'Customer Data Contract', got %s", contract.Name)
	}
	// Additional assertions for other fields can be added here
}
