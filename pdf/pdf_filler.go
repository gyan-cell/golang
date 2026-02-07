package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

/* -------------------- Utils -------------------- */

func normalize(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

/* -------------------- Load JSON -------------------- */

func loadJSONData(path string) (map[string]string, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var data map[string]string
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, err
	}

	return data, nil
}

/* -------------------- PDF Generation -------------------- */

func createPDF(output string, data map[string]string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	pdf.Cell(0, 10, "Filled PDF Data")
	pdf.Ln(12)

	pdf.SetFont("Arial", "", 12)

	// Iterate over JSON data and print key: value pairs
	for key, val := range data {
		line := fmt.Sprintf("%s: %s", key, val)
		pdf.MultiCell(0, 8, line, "", "L", false)
		pdf.Ln(2)
	}

	// Save PDF
	return pdf.OutputFileAndClose(output)
}

/* -------------------- Main -------------------- */

func main() {
	jsonFile := "data.json"
	outputPDF := "filled.pdf"

	data, err := loadJSONData(jsonFile)
	if err != nil {
		panic(err)
	}

	// Optional: normalize keys
	normalized := make(map[string]string)
	for k, v := range data {
		normalized[normalize(k)] = v
	}

	if err := createPDF(outputPDF, normalized); err != nil {
		panic(err)
	}

	fmt.Println("PDF created successfully:", outputPDF)
}
