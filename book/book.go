package book

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func Getbook(resp *http.Response) {
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("div.card-deck.mb-3text-center").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title := s.Find("div").Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})
}
