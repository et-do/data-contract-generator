package utils

import (
	"bytes"
	"data-contract-generator/models"
	"fmt"
	"os"
	"strings"

	svg "github.com/ajstarks/svgo"
)

func GenerateSVG(contract *models.DataContract) ([]byte, error) {
	buffer := new(bytes.Buffer)
	canvas := svg.New(buffer)

	width := 900
	height := 800
	padding := 20
	textSpacing := 30
	borderThickness := 2
	lineLength := 60 // Approximate character length per line

	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height, "fill:white")

	// Draw a border around the canvas
	canvas.Rect(padding, padding, width-2*padding, height-2*padding, fmt.Sprintf("stroke:black;stroke-width:%d;fill:none", borderThickness))

	y := padding + textSpacing
	// Section Header
	canvas.Text(padding*2, y, "Data Contract", "font-size:24px; font-weight:bold; fill:black")
	y += textSpacing + 10

	// Contract details
	canvas.Text(padding*2, y, fmt.Sprintf("Name: %s", contract.Name), "font-size:20px; fill:black")
	y += textSpacing
	canvas.Text(padding*2, y, fmt.Sprintf("Date: %s", contract.Date), "font-size:16px; fill:black")
	y += textSpacing
	canvas.Text(padding*2, y, fmt.Sprintf("Data Source: %s", contract.DataSource), "font-size:16px; fill:black")
	y += textSpacing
	canvas.Text(padding*2, y, fmt.Sprintf("Table Name: %s", contract.TableName), "font-size:16px; fill:black")
	y += textSpacing

	// Section Header
	canvas.Text(padding*2, y, "Expected Schema", "font-size:20px; font-weight:bold; fill:black")
	y += textSpacing
	for _, field := range contract.ExpectedSchema {
		canvas.Text(padding*3, y, fmt.Sprintf("- %s: %s", field.Name, field.Type), "font-size:16px; fill:black")
		y += textSpacing - 10
	}
	y += textSpacing - 20

	// Other details
	canvas.Text(padding*2, y, fmt.Sprintf("Frequency: %s", contract.Frequency), "font-size:16px; fill:black")
	y += textSpacing
	canvas.Text(padding*2, y, fmt.Sprintf("PII Sensitive: %t", contract.PIISensitive), "font-size:16px; fill:black")
	y += textSpacing

	// Section Header
	canvas.Text(padding*2, y, "Description", "font-size:20px; font-weight:bold; fill:black")
	y += textSpacing - 10
	descriptionLines := wrapText(contract.Description, lineLength)
	for _, line := range descriptionLines {
		canvas.Text(padding*3, y, line, "font-size:16px; fill:black")
		y += textSpacing - 10
	}
	y += textSpacing - 20

	// Section Header
	canvas.Text(padding*2, y, "Contact Person", "font-size:20px; font-weight:bold; fill:black")
	y += textSpacing
	canvas.Text(padding*3, y, fmt.Sprintf("Name: %s", contract.ContactPerson.Name), "font-size:16px; fill:black")
	y += textSpacing - 10
	canvas.Text(padding*3, y, fmt.Sprintf("Email: %s", contract.ContactPerson.Email), "font-size:16px; fill:black")
	y += textSpacing - 10
	canvas.Text(padding*3, y, fmt.Sprintf("Phone: %s", contract.ContactPerson.Phone), "font-size:16px; fill:black")
	y += textSpacing

	// Other details
	canvas.Text(padding*2, y, fmt.Sprintf("Data Retention Policy: %s", contract.DataRetentionPolicy), "font-size:16px; fill:black")
	y += textSpacing
	canvas.Text(padding*2, y, fmt.Sprintf("Security Classification: %s", contract.SecurityClassification), "font-size:16px; fill:black")
	y += textSpacing
	canvas.Text(padding*2, y, fmt.Sprintf("Version: %s", contract.Version), "font-size:16px; fill:black")
	y += textSpacing

	// Add some space at the bottom
	y += textSpacing

	canvas.End()
	return buffer.Bytes(), nil
}

// Helper function to wrap text
func wrapText(text string, lineLength int) []string {
	words := strings.Fields(text)
	var lines []string
	var currentLine string

	for _, word := range words {
		if len(currentLine)+len(word)+1 > lineLength {
			lines = append(lines, currentLine)
			currentLine = word
		} else {
			if len(currentLine) > 0 {
				currentLine += " "
			}
			currentLine += word
		}
	}
	if len(currentLine) > 0 {
		lines = append(lines, currentLine)
	}

	return lines
}

func SaveSVG(svgData []byte, outputFilename string) error {
	return os.WriteFile(outputFilename, svgData, 0644)
}
