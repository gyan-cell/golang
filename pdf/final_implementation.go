package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"html/template"
	"os"
	"path/filepath"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

type InvoiceData struct {
	Title    string
	Customer string
	Amount   int
	Date     string
	Logo     string
}

func RenderHtml(templatePath string, data any) (string, error) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, data)
	return buf.String(), err
}

func PrintPDF(html string, outputDir, fileName string) error {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	tmpDir := os.TempDir()
	htmlPath := filepath.Join(tmpDir, "invoice_render.html")

	if err := os.WriteFile(htmlPath, []byte(html), 0644); err != nil {
		return err
	}
	defer os.Remove(htmlPath)

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	var pdf []byte

	err := chromedp.Run(ctx,
		chromedp.Navigate("file://"+htmlPath),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			pdf, _, err = page.PrintToPDF().
				WithPrintBackground(true).
				WithPreferCSSPageSize(true).
				Do(ctx)
			return err
		}),
	)
	if err != nil {
		return err
	}

	pdfPath := filepath.Join(outputDir, fileName)
	return os.WriteFile(pdfPath, pdf, 0644)
}

func main() {

	logoBytes, err := os.ReadFile("assets/random.jpg")
	logoBase64 := base64.StdEncoding.EncodeToString(logoBytes)

	html, err := RenderHtml(
		"templates/invoice.html",
		InvoiceData{
			Title:    "Invoice",
			Customer: "Naresh Kumar",
			Amount:   500,
			Date:     "04 Feb 2026",
			Logo:     logoBase64,
		},
	)
	if err != nil {
		panic(err)
	}

	err = PrintPDF(html, "output", "invoice.pdf")
	if err != nil {
		panic(err)
	}
}
