package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func mergeFromDir(dir string, output string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	var files []string
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		if strings.HasSuffix(strings.ToLower(e.Name()), ".pdf") {
			files = append(files, filepath.Join(dir, e.Name()))
		}
	}

	sort.Strings(files)

	if len(files) == 0 {
		return fmt.Errorf("no pdf files found in %s", dir)
	}

	conf := model.NewDefaultConfiguration()
	return api.MergeCreateFile(files, output, false, conf)
}

func mergeFromURLs(ctx context.Context, urls []string, output string) error {
	tmpDir, err := os.MkdirTemp("", "pdfmerge-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpDir)

	client := &http.Client{Timeout: 30 * time.Second}
	var files []string

	for i, url := range urls {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return err
		}

		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			return fmt.Errorf("failed to download %s", url)
		}

		outPath := filepath.Join(tmpDir, fmt.Sprintf("%03d.pdf", i))
		out, err := os.Create(outPath)
		if err != nil {
			resp.Body.Close()
			return err
		}

		_, err = io.Copy(out, io.LimitReader(resp.Body, 50<<20))
		out.Close()
		resp.Body.Close()

		if err != nil {
			return err
		}

		files = append(files, outPath)
	}

	conf := model.NewDefaultConfiguration()
	return api.MergeCreateFile(files, output, false, conf)
}

func main() {
	ctx := context.Background()

	err := mergeFromDir("./output", "merged_offline.pdf")
	if err != nil {
		panic(err)
	}

	err = mergeFromURLs(ctx, []string{
		"https://www.princexml.com/samples/invoice-colorful/invoicesample.pdf",
		"https://www.princexml.com/samples/invoice-colorful/invoicesample.pdf",
	}, "merged_online.pdf")
	if err != nil {
		panic(err)
	}
}
