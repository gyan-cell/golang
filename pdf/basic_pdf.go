package main

import (
	"fmt"
	"log"

	"github.com/jung-kurt/gofpdf"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

type TextBlock struct {
	ID      string
	Text    string
	X, Y    float64
	Font    string
	Size    float64
	R, G, B int
	Bold    bool
	Italic  bool
}

func CreatePDF(blocks []TextBlock, output string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	for _, b := range blocks {
		drawText(pdf, b)
	}

	return pdf.OutputFileAndClose(output)
}

func ReadBlocks(blocks []TextBlock) {
	fmt.Println("Reading logical PDF model:")
	for _, b := range blocks {
		fmt.Printf("ID=%s TEXT=%q @ (%.1f, %.1f)\n", b.ID, b.Text, b.X, b.Y)
	}
}

func UpdateBlock(blocks []TextBlock, id string, newText string) []TextBlock {
	for i := range blocks {
		if blocks[i].ID == id {
			blocks[i].Text = newText
		}
	}
	return blocks
}

func DeleteBlock(blocks []TextBlock, id string) []TextBlock {
	out := []TextBlock{}
	for _, b := range blocks {
		if b.ID != id {
			out = append(out, b)
		}
	}
	return out
}

func drawText(pdf *gofpdf.Fpdf, b TextBlock) {
	style := ""
	if b.Bold {
		style += "B"
	}
	if b.Italic {
		style += "I"
	}

	pdf.SetFont(b.Font, style, b.Size)
	pdf.SetTextColor(b.R, b.G, b.B)
	pdf.Text(b.X, b.Y, b.Text)
}

func DeletePages(input, output string, pages []string) error {
	return api.RemovePagesFile(input, output, pages, nil)
}

func main() {
	blocks := []TextBlock{
		{
			ID:   "title",
			Text: "PDF CRUD DEMO",
			X:    20, Y: 30,
			Font: "Helvetica", Size: 24,
			R: 200, G: 50, B: 50,
			Bold: true,
		},
		{
			ID:   "body",
			Text: "This text is placed using coordinates.",
			X:    20, Y: 50,
			Font: "Helvetica", Size: 14,
			R: 30, G: 30, B: 30,
		},
		{
			ID:   "italic",
			Text: "Styled & italic text block",
			X:    20, Y: 65,
			Font: "Times", Size: 14,
			R: 50, G: 100, B: 200,
			Italic: true,
		},
	}

	ReadBlocks(blocks)

	if err := CreatePDF(blocks, "output_created.pdf"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created: output_created.pdf")

	blocks = UpdateBlock(blocks, "body", " This text was UPDATED")
	if err := CreatePDF(blocks, "output_updated.pdf"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated: output_updated.pdf")

	blocks = DeleteBlock(blocks, "italic")
	if err := CreatePDF(blocks, "output_deleted_block.pdf"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted block: output_deleted_block.pdf")

	if err := DeletePages(
		"output_created.pdf",
		"output_pages_deleted.pdf",
		[]string{"2"}, // safe even if page doesnt exist
	); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Page deletion via pdfcpu done: output_pages_deleted.pdf")
}
