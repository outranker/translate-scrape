package main

import (
	"context"
	"log"
	"strings"

	"github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var res string
	err := chromedp.Run(ctx,chromedp.Navigate(`https://pkg.go.dev/time`), chromedp.Text(`.Documentation-overview`, &res, chromedp.NodeNotVisible))

	if err != nil { 
		log.Fatal(err)
	}
	log.Println(strings.TrimSpace(res))
}