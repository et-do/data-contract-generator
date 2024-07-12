package utils

import (
	"bytes"
	"fmt"
	"github.com/ajstarks/svgo"
	"io/ioutil"
	"data-contract-generator/models"
)

func GenerateSVG(contract *models.DataContract) ([]byte, error) {
	buffer := new(bytes.Buffer)
	canvas := svg.New(buffer)

	width := 800
	height := 600
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height, "fill:white")

	y := 30
	canvas.Text(20, y, fmt.Sprintf("Name: %s", contract.Name), "font-size:20px; fill:black")
	y += 30
	canvas.Text(20, y, fmt.Sprintf("Date: %s", contract.Date), "font-size:16px; fill:black")
	y += 30
	canvas.Text(20, y, fmt.Sprintf("Data Source: %s", contract.DataSource), "font-size:16px; fill:black")
	y += 30
	canvas.Text(20, y, fmt.Sprintf("Table Name: %s", contract.TableName), "font-size:16px; fill:black")
	y += 30
	canvas.Text(20, y, "Expected Schema:", "font-size:16px; fill:black")
	y += 30
	for _, field := range contract.ExpectedSchema {
		canvas.Text(40, y, fmt.Sprintf("%s: %s", field.Name, field.Type), "font-size:16px; fill:black")
		y += 20
	}
	y += 10
	canvas.Text(20, y, fmt.Sprintf("Frequency: %s", contract.Frequency), "font-size:16px; fill:black")
	y += 30
	canvas.Text(20, y, fmt.Sprintf("PII Sensitive: %t", contract.PIISensitive), "font-size:16px; fill:black")
	y += 30
	canvas.Text(20, y, fmt.Sprintf("Description: %s", contract.Description), "font-size:16px; fill:black")
	y += 30
	canvas.Text(20, y, "Contact Person:", "font-size:16px; fill:black")
	y += 30
	canvas.Text(40, y, fmt.Sprintf("Name: %s", contract.ContactPerson.Name), "font-size:16px; fill:black")
	y += 20
	canvas.Text(40, y, fmt.Sprintf("Email: %s", contract.ContactPerson.Email), "font-size:16px; fill:black")
	y += 20
	canvas.Text(40, y, fmt.Sprintf("Phone: %s", contract.ContactPerson.Phone), "font-size:16px; fill:black")
	y += 30
	canvas.Text(20, y, fmt.Sprintf("Data Retention Policy: %s", contract.DataRetentionPolicy), "font-size:16px; fill:black")
	y += 30
	canvas.Text(20, y, fmt.Sprintf("Security Classification: %s", contract.SecurityClassification), "font-size:16px; fill:black")
	y += 30
	canvas.Text(20, y, fmt.Sprintf("Version: %s", contract.Version), "font-size:16px; fill:black")

	canvas.End()
	return buffer.Bytes(), nil
}

func SaveSVG(svgData []byte, outputFilename string) error {
	return ioutil.WriteFile(outputFilename, svgData, 0644)
}
