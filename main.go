package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
fmt.Println(`hello`)
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.nytimes.com/live/2022/03/02/world/ukraine-russia-war"), 
		chromedp.Text(".site-content", 
		&res, chromedp.NodeVisible))

	if err != nil { 
		log.Fatal(err)
	}
	log.Println(strings.TrimSpace(res))
	os.Exit(0)
}