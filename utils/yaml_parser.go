package utils

import (
	"io/ioutil"
	"my_project/models"
	"gopkg.in/yaml.v2"
)

func ParseYAML(filename string) (*models.DataContract, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var contract models.DataContract
	err = yaml.Unmarshal(data, &contract)
	if err != nil {
		return nil, err
	}

	return &contract, nil
}
