package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func normalizeFieldName(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

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
			field := strings.TrimSpace(strings.TrimPrefix(line, "FieldName:"))
			fields = append(fields, field)
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

func buildFDF(data map[string]string) string {
	var sb strings.Builder
	sb.WriteString("%FDF-1.2\n")
	sb.WriteString("1 0 obj\n<<\n/FDF << /Fields [\n")

	for k, v := range data {
		sb.WriteString(fmt.Sprintf("<< /T (%s) /V (%s) >>\n", k, v))
	}

	sb.WriteString("] >>\n>>\nendobj\ntrailer\n<<\n/Root 1 0 R\n>>\n%%EOF")
	return sb.String()
}

func loadFormData(jsonPath string, inline map[string]string) (map[string]string, error) {
	if jsonPath != "" {
		b, err := os.ReadFile(jsonPath)
		if err != nil {
			return nil, fmt.Errorf("reading json: %w", err)
		}
		var data map[string]string
		if err := json.Unmarshal(b, &data); err != nil {
			return nil, fmt.Errorf("parsing json: %w", err)
		}
		return data, nil
	}

	if len(inline) > 0 {
		return inline, nil
	}

	return nil, fmt.Errorf("no form data provided")
}

func fillPDF(inputPDF, outputPDF string, data map[string]string) error {
	normData := make(map[string]string)
	for k, v := range data {
		normData[normalizeFieldName(k)] = v
	}

	pdfFields, err := getPDFFields(inputPDF)
	if err != nil {
		return err
	}

	for _, f := range pdfFields {
		if _, ok := normData[normalizeFieldName(f)]; !ok {
			return fmt.Errorf("missing value for field: %q", f)
		}
	}

	fdf := buildFDF(normData)
	tmp := "temp.fdf"
	if err := os.WriteFile(tmp, []byte(fdf), 0644); err != nil {
		return err
	}
	defer os.Remove(tmp)

	cmd := exec.Command("pdftk", inputPDF, "fill_form", tmp, "output", outputPDF, "flatten")
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("pdftk error: %v\n%s", err, out)
	}

	return nil
}

func main() {
	inputPDF := "foreSome.pdf"
	outputPDF := "form_filled.pdf"

	// set to "" to disable had to be done
	jsonFile := ""

	// 🔹 Option 2: inline fields (used if jsonFile == "")
	inlineFields := map[string]string{
		" Address 1 Text Box":        "MG Road",
		" Address 2 Text Box":        "Near City Mall",
		" City Text Box":             "Bengaluru",
		" Country Combo Box":         "India",
		" Given Name Text Box":       "Gyanranjan",
		" Family Name Text Box":      "Jha",
		" Driving License Check Box": "Yes",
		" Gender List Box":           "Man",
		" Height Formatted Field":    "1i72",
		" Postcode Text Box":         "560001",
		" House nr Text Box":         "42",
		" Favourite Colour List Box": "Violet",
		" Language 1 Check Box":      "Off",
		" Language 2 Check Box":      "Yes",
		" Language 3 Check Box":      "Off",
		" Language 4 Check Box":      "Off",
		" Language 5 Check Box":      "Off",
	}

	data, err := loadFormData(jsonFile, inlineFields)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err := fillPDF(inputPDF, outputPDF, data); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("PDF filled successfully:", outputPDF)
}
