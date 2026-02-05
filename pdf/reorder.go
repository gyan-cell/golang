package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func ReorderPDFKeepRest(input string, order []int, output string) error {
	conf := model.NewDefaultConfiguration()

	ctx, err := api.ReadContextFile(input)
	if err != nil {
		return err
	}

	total := ctx.PageCount
	seen := make(map[int]bool)

	finalOrder := []int{}

	for _, p := range order {
		if p < 1 || p > total {
			return fmt.Errorf("page %d out of range (1-%d)", p, total)
		}
		if !seen[p] {
			finalOrder = append(finalOrder, p)
			seen[p] = true
		}
	}

	for i := 1; i <= total; i++ {
		if !seen[i] {
			finalOrder = append(finalOrder, i)
		}
	}

	tmp, err := os.MkdirTemp("", "reorder-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmp)

	var files []string

	for i, p := range finalOrder {
		dir := filepath.Join(tmp, fmt.Sprintf("%03d", i))
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}

		if err := api.ExtractPagesFile(
			input,
			dir,
			[]string{strconv.Itoa(p)},
			conf,
		); err != nil {
			return err
		}

		pdf, err := filepath.Glob(filepath.Join(dir, "*.pdf"))
		if err != nil || len(pdf) != 1 {
			return fmt.Errorf("failed extracting page %d", p)
		}

		files = append(files, pdf[0])
	}

	if err := os.MkdirAll(filepath.Dir(output), 0755); err != nil {
		return err
	}

	return api.MergeCreateFile(files, output, false, conf)
}

func main() {
	err := ReorderPDFKeepRest(
		"dictionary.pdf",
		[]int{3, 1, 5, 2, 28, 32, 4, 12, 4, 5, 6, 7, 343, 45, 5, 6, 7, 88, 299},
		"./output/reordered.pdf",
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("done")
}
