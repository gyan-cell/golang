package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// getPDFFields extracts all field names from the PDF using pdftk
func getPDFFields(pdfPath string) ([]string, error) {
	cmd := exec.Command("pdftk", pdfPath, "dump_data_fields")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	if err := cmd.Start(); err != nil {
		return nil, err
	}

	var fields []string
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "FieldName:") {
			fieldName := strings.TrimPrefix(line, "FieldName:") // **do not trim spaces**
			fields = append(fields, fieldName)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if err := cmd.Wait(); err != nil {
		return nil, err
	}

	return fields, nil
}

// generateJSONTemplate generates a JSON file with all fields exactly as in PDF
func generateJSONTemplate(pdfPath, outputJSON string) error {
	fields, err := getPDFFields(pdfPath)
	if err != nil {
		return fmt.Errorf("failed to get PDF fields: %v", err)
	}

	data := make(map[string]string)
	for _, f := range fields {
		data[f] = "" // empty value, key exactly as PDF
	}

	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	if err := os.WriteFile(outputJSON, jsonBytes, 0644); err != nil {
		return fmt.Errorf("failed to write JSON file: %v", err)
	}

	fmt.Printf("JSON template generated: %s\n", outputJSON)
	return nil
}

func main() {
	pdfPath := "drafter.pdf"          // your PDF
	outputJSON := "pdf_template.json" // generated JSON

	if err := generateJSONTemplate(pdfPath, outputJSON); err != nil {
		fmt.Println("Error:", err)
		return
	}
}
