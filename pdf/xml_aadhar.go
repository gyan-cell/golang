package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/jung-kurt/gofpdf"
)

type Certificate struct {
	CertificateData struct {
		KycRes struct {
			Ts      string  `xml:"ts,attr"`
			UidData UidData `xml:"UidData"`
		} `xml:"KycRes"`
	} `xml:"CertificateData"`
}

type UidData struct {
	UID string `xml:"uid,attr"`
	Poi struct {
		Name   string `xml:"name,attr"`
		Dob    string `xml:"dob,attr"`
		Gender string `xml:"gender,attr"`
	} `xml:"Poi"`
	Poa struct {
		Co     string `xml:"co,attr"`
		House  string `xml:"house,attr"`
		Street string `xml:"street,attr"`
		Vtc    string `xml:"vtc,attr"`
		State  string `xml:"state,attr"`
		Pc     string `xml:"pc,attr"`
	} `xml:"Poa"`
	Pht string `xml:"Pht"`
}

func main() {
	uid, issueDate := parseXML("aadhaar.xml")
	photo, hasPhoto := decodePhoto(uid.Pht)
	generateUIDAIPDF(uid, issueDate, photo, hasPhoto)
	fmt.Println("Aadhaar UIDAI-style PDF generated")
}

func parseXML(path string) (*UidData, string) {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var c Certificate
	if err := xml.Unmarshal(data, &c); err != nil {
		panic(err)
	}

	ts := c.CertificateData.KycRes.Ts
	issueDate := extractDate(ts)

	return &c.CertificateData.KycRes.UidData, issueDate
}

func decodePhoto(pht string) ([]byte, bool) {
	clean := strings.Join(strings.Fields(pht), "")
	if len(clean) < 1000 {
		return nil, false
	}

	raw, err := base64.StdEncoding.DecodeString(clean)
	if err != nil {
		return nil, false
	}

	if len(raw) > 2 && raw[0] == 0x1f && raw[1] == 0x8b {
		r, err := gzip.NewReader(bytes.NewReader(raw))
		if err != nil {
			return nil, false
		}
		defer r.Close()

		raw, err = io.ReadAll(r)
		if err != nil {
			return nil, false
		}
	}

	if len(raw) < 4 || raw[0] != 0xFF || raw[1] != 0xD8 {
		return nil, false
	}

	return raw, true
}

func generateUIDAIPDF(u *UidData, issueDate string, photo []byte, hasPhoto bool) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Card background
	pdf.SetFillColor(245, 245, 245)
	pdf.RoundedRect(20, 30, 170, 110, 4, "1234", "F")
	pdf.SetDrawColor(180, 180, 180)
	pdf.RoundedRect(20, 30, 170, 110, 4, "1234", "D")

	// Tricolour header
	pdf.SetFillColor(255, 153, 51)
	pdf.Rect(20, 30, 170, 6, "F")
	pdf.SetFillColor(255, 255, 255)
	pdf.Rect(20, 36, 170, 10, "F")
	pdf.SetFillColor(19, 136, 8)
	pdf.Rect(20, 46, 170, 4, "F")

	pdf.SetFont("Arial", "B", 15)
	pdf.SetXY(25, 39)
	pdf.Cell(0, 6, "Government of India")
	pdf.SetFont("Arial", "", 10)
	pdf.SetXY(25, 45)
	pdf.Cell(0, 5, "Unique Identification Authority of India")

	// Photo (embedded from memory)
	if hasPhoto {
		pdf.RegisterImageOptionsReader(
			"aadhaar_photo",
			gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true},
			bytes.NewReader(photo),
		)
		pdf.Image("aadhaar_photo", 25, 58, 30, 38, false, "", 0, "")
	}
	pdf.Rect(25, 58, 30, 38, "D")

	// Text rows
	drawRow(pdf, 60, 85, 60, "Name", u.Poi.Name)
	drawRow(pdf, 60, 85, 68, "DOB", u.Poi.Dob)
	drawRow(pdf, 60, 85, 76, "Gender", gender(u.Poi.Gender))
	drawRow(pdf, 60, 85, 84, "Issue Date", issueDate)

	// Address
	pdf.Line(25, 100, 185, 100)
	pdf.SetFont("Arial", "B", 11)
	pdf.SetXY(25, 104)
	pdf.Cell(0, 6, "Address")

	addr := fmt.Sprintf("%s, %s, %s, %s - %s",
		u.Poa.Co, u.Poa.House, u.Poa.Street, u.Poa.Vtc, u.Poa.Pc)

	pdf.SetFont("Arial", "", 10)
	pdf.SetXY(25, 110)
	pdf.MultiCell(160, 5, addr, "", "L", false)

	// UID
	pdf.SetFont("Arial", "B", 18)
	pdf.SetXY(25, 130)
	pdf.CellFormat(170, 8, formatUID(u.UID), "", 0, "C", false, 0, "")

	pdf.OutputFileAndClose("aadhaar_card.pdf")
}

func drawRow(pdf *gofpdf.Fpdf, lx, vx, y float64, k, v string) {
	pdf.SetFont("Arial", "B", 11)
	pdf.SetXY(lx, y)
	pdf.Cell(22, 6, k)
	pdf.SetFont("Arial", "", 11)
	pdf.SetXY(vx, y)
	pdf.Cell(0, 6, ": "+v)
}

func formatUID(u string) string {
	u = strings.ReplaceAll(u, " ", "")
	if len(u) != 12 {
		return u
	}
	return u[:4] + " " + u[4:8] + " " + u[8:]
}

func gender(g string) string {
	switch strings.ToUpper(g) {
	case "M":
		return "Male"
	case "F":
		return "Female"
	case "T":
		return "Transgender"
	default:
		return g
	}
}

func extractDate(ts string) string {
	if len(ts) < 10 {
		return ts
	}

	t, err := time.Parse("2006-01-02", ts[:10])
	if err != nil {
		return ts[:10]
	}

	return t.Format("02 January 2006")
}
