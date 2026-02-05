package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func collectPDFs(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, e := range entries {
		if !e.IsDir() && filepath.Ext(e.Name()) == ".pdf" {
			files = append(files, filepath.Join(dir, e.Name()))
		}
	}

	sort.Strings(files)
	return files, nil
}

func SplitPDFAtIndex(inputPDF string, splitAt int, outDir string) error {
	if splitAt <= 0 {
		return fmt.Errorf("split index must be >= 1")
	}

	tmp1, err := os.MkdirTemp("", "split1-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmp1)

	tmp2, err := os.MkdirTemp("", "split2-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmp2)

	conf := model.NewDefaultConfiguration()

	if err := api.ExtractPagesFile(
		inputPDF,
		tmp1,
		[]string{fmt.Sprintf("1-%d", splitAt)},
		conf,
	); err != nil {
		return err
	}

	if err := api.ExtractPagesFile(
		inputPDF,
		tmp2,
		[]string{fmt.Sprintf("%d-", splitAt+1)},
		conf,
	); err != nil {
		return err
	}

	files1, err := collectPDFs(tmp1)
	if err != nil {
		return err
	}

	files2, err := collectPDFs(tmp2)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(outDir, 0755); err != nil {
		return err
	}

	if err := api.MergeCreateFile(
		files1,
		filepath.Join(outDir, "part1.pdf"),
		false,
		conf,
	); err != nil {
		return err
	}

	if err := api.MergeCreateFile(
		files2,
		filepath.Join(outDir, "part2.pdf"),
		false,
		conf,
	); err != nil {
		return err
	}

	return nil
}

func ExtractPDFPage(inputPDF string, page int, outDir string) error {
	if page <= 0 {
		return fmt.Errorf("page must be >= 1")
	}

	if err := os.MkdirAll(outDir, 0755); err != nil {
		return err
	}

	conf := model.NewDefaultConfiguration()

	return api.ExtractPagesFile(
		inputPDF,
		outDir,
		[]string{fmt.Sprintf("%d", page)},
		conf,
	)
}

func main() {
	if err := SplitPDFAtIndex(
		"dictionary.pdf",
		30,
		"./split_output",
	); err != nil {
		panic(err)
	}

	if err := ExtractPDFPage(
		"dictionary.pdf",
		22,
		"./single_page",
	); err != nil {
		panic(err)
	}

	fmt.Println("done")
}
