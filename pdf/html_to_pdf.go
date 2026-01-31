package main

import (
	"context"
	"os"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func main() {
	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<style>
			body { font-family: Arial; padding: 20px; }
			h1 { color: blue; }
		</style>
	</head>
	<body>
		<h1>My Invoice</h1>
		<p>Customer: John Doe</p>
		<p>Amount: $500</p>
	</body>
	</html>
	`

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var pdf []byte

	tasks := chromedp.Tasks{
		chromedp.Navigate("data:text/html," + html),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			pdf, _, err = page.PrintToPDF().
				WithPrintBackground(true).
				WithPreferCSSPageSize(true).
				Do(ctx)
			return err
		}),
	}

	if err := chromedp.Run(ctx, tasks); err != nil {
		panic(err)
	}

	if err := os.WriteFile("invoice.pdf", pdf, 0644); err != nil {
		panic(err)
	}

	println("PDF created successfully!")

}
