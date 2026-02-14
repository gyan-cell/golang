package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/jung-kurt/gofpdf"
)

/*
=============================
        CONSTANTS
=============================
*/

const (
	MaxXMLFileSize     = 10 << 20 // 10MB
	MaxBase64PhotoSize = 5 << 20  // 5MB
	MinBase64PhotoSize = 1000
	MaxFieldLength     = 500
	UIDLength          = 12

	// Magic numbers
	GzipMagic1    = 0x1f
	GzipMagic2    = 0x8b
	JPEGSOIMagic1 = 0xFF
	JPEGSOIMagic2 = 0xD8
	JPEGEOIMagic1 = 0xFF
	JPEGEOIMagic2 = 0xD9
)

/*
=============================
        LAYOUT CONSTANTS
=============================
*/

const (
	CardX      = 20.0
	CardY      = 30.0
	CardWidth  = 170.0
	CardHeight = 110.0
	CardRadius = 4.0

	PhotoX      = 25.0
	PhotoY      = 58.0
	PhotoWidth  = 30.0
	PhotoHeight = 38.0

	TextStartX    = 60.0
	TextLabelX    = 85.0
	TextStartY    = 60.0
	TextRowHeight = 8.0

	AddressLineY  = 100.0
	AddressTextY  = 104.0
	AddressStartY = 110.0

	UIDTextY = 130.0
)

var uidRegex = regexp.MustCompile(`^\d{12}$`)

/*
=============================
        ERROR TYPE
=============================
*/

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error [%s]: %s", e.Field, e.Message)
}

/*
=============================
        XML PARSING
=============================
*/

func ParseXML(path string) (*models.UidData, string, error) {

	info, err := os.Stat(path)
	if err != nil {
		return nil, "", fmt.Errorf("failed to stat file: %w", err)
	}

	if info.Size() > MaxXMLFileSize {
		return nil, "", &ValidationError{
			Field:   "file_size",
			Message: "XML file too large",
		}
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, "", fmt.Errorf("failed to read XML: %w", err)
	}

	return parseXMLBytes(data)
}

func parseXMLBytes(data []byte) (*models.UidData, string, error) {

	var cert models.Certificate

	if err := xml.Unmarshal(data, &cert); err != nil {
		return nil, "", fmt.Errorf("failed to unmarshal XML: %w", err)
	}

	u := &cert.CertificateData.KycRes.UidData

	if !validateUID(u.UID) {
		return nil, "", &ValidationError{
			Field:   "uid",
			Message: "invalid UID format",
		}
	}

	if err := validateUidData(u); err != nil {
		return nil, "", err
	}

	issueDate, err := extractDate(cert.CertificateData.KycRes.Ts)
	if err != nil {
		return nil, "", err
	}

	return u, issueDate, nil
}

/*
=============================
        VALIDATION
=============================
*/

func validateUID(uid string) bool {
	uid = strings.ReplaceAll(uid, " ", "")
	return uidRegex.MatchString(uid)
}

func validateUidData(u *models.UidData) error {

	if strings.TrimSpace(u.Poi.Name) == "" {
		return &ValidationError{Field: "name", Message: "name is required"}
	}

	if len(u.Poi.Name) > MaxFieldLength {
		return &ValidationError{Field: "name", Message: "name too long"}
	}

	if u.Poi.Dob == "" {
		return &ValidationError{Field: "dob", Message: "DOB required"}
	}

	return nil
}

/*
=============================
        PHOTO DECODING
=============================
*/

func DecodePhoto(pht string) ([]byte, error) {

	clean := strings.Join(strings.Fields(pht), "")

	if len(clean) < MinBase64PhotoSize {
		return nil, &ValidationError{Field: "photo", Message: "photo too small"}
	}

	if len(clean) > MaxBase64PhotoSize {
		return nil, &ValidationError{Field: "photo", Message: "photo too large"}
	}

	raw, err := base64.StdEncoding.DecodeString(clean)
	if err != nil {
		return nil, fmt.Errorf("base64 decode failed: %w", err)
	}

	// GZIP handling
	if len(raw) > 2 && raw[0] == GzipMagic1 && raw[1] == GzipMagic2 {
		r, err := gzip.NewReader(bytes.NewReader(raw))
		if err != nil {
			return nil, err
		}
		defer r.Close()

		raw, err = io.ReadAll(r)
		if err != nil {
			return nil, err
		}
	}

	if !isValidJPEG(raw) {
		return nil, &ValidationError{Field: "photo", Message: "invalid JPEG"}
	}

	return raw, nil
}

func isValidJPEG(data []byte) bool {
	if len(data) < 4 {
		return false
	}

	hasSOI := data[0] == JPEGSOIMagic1 && data[1] == JPEGSOIMagic2
	hasEOI := data[len(data)-2] == JPEGEOIMagic1 && data[len(data)-1] == JPEGEOIMagic2

	return hasSOI && hasEOI
}

/*
=============================
        PDF GENERATION
=============================
*/

func GenerateUIDAIPDF(
	u *models.UidData,
	issueDate string,
	photo []byte,
	hasPhoto bool,
	outputPath string,
) error {

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	drawCardBackground(pdf)
	drawHeader(pdf)

	if hasPhoto && len(photo) > 0 {
		if err := pdf.RegisterImageOptionsReader(
			"photo",
			gofpdf.ImageOptions{ImageType: "JPG"},
			bytes.NewReader(photo),
		); err != nil {
			return err
		}

		pdf.Image("photo", PhotoX, PhotoY, PhotoWidth, PhotoHeight, false, "", 0, "")
	}

	pdf.Rect(PhotoX, PhotoY, PhotoWidth, PhotoHeight, "D")

	drawPersonalDetails(pdf, u, issueDate)
	drawAddress(pdf, u)
	drawUID(pdf, u)

	return pdf.OutputFileAndClose(outputPath)
}

/*
=============================
        DRAW HELPERS
=============================
*/

func drawCardBackground(pdf *gofpdf.Fpdf) {
	pdf.SetFillColor(245, 245, 245)
	pdf.RoundedRect(CardX, CardY, CardWidth, CardHeight, CardRadius, "1234", "F")
	pdf.SetDrawColor(180, 180, 180)
	pdf.RoundedRect(CardX, CardY, CardWidth, CardHeight, CardRadius, "1234", "D")
}

func drawHeader(pdf *gofpdf.Fpdf) {

	pdf.SetFillColor(255, 153, 51)
	pdf.Rect(CardX, CardY, CardWidth, 6, "F")

	pdf.SetFillColor(255, 255, 255)
	pdf.Rect(CardX, CardY+6, CardWidth, 10, "F")

	pdf.SetFillColor(19, 136, 8)
	pdf.Rect(CardX, CardY+16, CardWidth, 4, "F")

	pdf.SetFont("Arial", "B", 15)
	pdf.SetXY(CardX+5, CardY+9)
	pdf.Cell(0, 6, "Government of India")

	pdf.SetFont("Arial", "", 10)
	pdf.SetXY(CardX+5, CardY+15)
	pdf.Cell(0, 5, "Unique Identification Authority of India")
}

func drawPersonalDetails(pdf *gofpdf.Fpdf, u *models.UidData, issueDate string) {

	y := TextStartY

	drawRow(pdf, TextStartX, TextLabelX, y, "Name", u.Poi.Name)
	y += TextRowHeight

	drawRow(pdf, TextStartX, TextLabelX, y, "DOB", u.Poi.Dob)
	y += TextRowHeight

	drawRow(pdf, TextStartX, TextLabelX, y, "Gender", formatGender(u.Poi.Gender))
	y += TextRowHeight

	drawRow(pdf, TextStartX, TextLabelX, y, "Issue Date", issueDate)
}

func drawAddress(pdf *gofpdf.Fpdf, u *models.UidData) {

	pdf.Line(PhotoX, AddressLineY, CardX+CardWidth-5, AddressLineY)

	pdf.SetFont("Arial", "B", 11)
	pdf.SetXY(PhotoX, AddressTextY)
	pdf.Cell(0, 6, "Address")

	addr := formatAddress(u.Poa)

	pdf.SetFont("Arial", "", 10)
	pdf.SetXY(PhotoX, AddressStartY)
	pdf.MultiCell(CardWidth-10, 5, addr, "", "L", false)
}

func drawUID(pdf *gofpdf.Fpdf, u *models.UidData) {

	formatted := formatUID(u.UID)

	pdf.SetFont("Arial", "B", 18)
	pdf.SetXY(PhotoX, UIDTextY)
	pdf.CellFormat(CardWidth-10, 8, formatted, "", 0, "C", false, 0, "")
}

func drawRow(pdf *gofpdf.Fpdf, lx, vx, y float64, label, value string) {
	pdf.SetFont("Arial", "B", 11)
	pdf.SetXY(lx, y)
	pdf.Cell(22, 6, label)

	pdf.SetFont("Arial", "", 11)
	pdf.SetXY(vx, y)
	pdf.Cell(0, 6, ": "+value)
}

/*
=============================
        FORMAT HELPERS
=============================
*/

func formatUID(uid string) string {
	uid = strings.ReplaceAll(uid, " ", "")
	if len(uid) != UIDLength {
		return uid
	}
	return uid[:4] + " " + uid[4:8] + " " + uid[8:]
}

func formatGender(g string) string {
	switch strings.ToUpper(strings.TrimSpace(g)) {
	case "M":
		return "Male"
	case "F":
		return "Female"
	case "T":
		return "Transgender"
	default:
		return "Other"
	}
}

func formatAddress(p models.Poa) string {

	parts := []string{}

	add := func(s string) {
		if strings.TrimSpace(s) != "" {
			parts = append(parts, strings.TrimSpace(s))
		}
	}

	add(p.Co)
	add(p.House)
	add(p.Street)
	add(p.Vtc)

	addr := strings.Join(parts, ", ")

	if p.Pc != "" {
		addr += " - " + strings.TrimSpace(p.Pc)
	}

	return addr
}

func extractDate(ts string) (string, error) {

	if ts == "" {
		return "", &ValidationError{Field: "timestamp", Message: "empty timestamp"}
	}

	if len(ts) < 10 {
		return "", &ValidationError{Field: "timestamp", Message: "invalid timestamp"}
	}

	dateStr := ts[:10]

	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return dateStr, nil
	}

	return t.Format("02 January 2006"), nil
}

func main() {

	// Make sure XML file path is passed
	if len(os.Args) < 2 {
		fmt.Print("Usage: go run main.go <aadhaar.xml>")
	}

	xmlPath := os.Args[1]

	// Step 1: Parse XML
	uidData, issueDate, err := ParseXML(xmlPath)
	if err != nil {
		fmt.Printf("ParseXML failed: %v", err)
	}

	// Step 2: Decode Photo (optional)
	var photo []byte
	var hasPhoto bool

	if uidData.Pht != "" {
		photo, err = DecodePhoto(uidData.Pht)
		if err != nil {
			fmt.Printf("Photo decode failed (continuing without photo): %v", err)
			hasPhoto = false
		} else {
			hasPhoto = true
		}
	}

	// Step 3: Generate PDF
	outputFile := "aadhaar_output.pdf"

	err = GenerateUIDAIPDF(
		uidData,
		issueDate,
		photo,
		hasPhoto,
		outputFile,
	)
	if err != nil {
		fmt.Printf("PDF generation failed: %v", err)
	}

	fmt.Println("✅ PDF generated successfully:", outputFile)
}
