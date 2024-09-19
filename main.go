package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://blog.medium.com/"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".js-trackPostPresentation").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".u-letterSpacingTight.u-lineHeightTighter").Text()
		description := s.Find(".u-fontSize18.u-letterSpacingTight.u-lineHeightTight").Text()
		author := s.Find(".ds-link.ds-link--styleSubtle").Text()
		date := s.Find("time").Text()

		imageURL, _ := s.Find(".u-block").Attr("style")
		imageURL = strings.TrimPrefix(strings.TrimSuffix(strings.Split(imageURL, "url(")[1], ");"), "\"")

		fmt.Printf("Title: %s\n", strings.TrimSpace(title))
		fmt.Printf("Description: %s\n", strings.TrimSpace(description))
		fmt.Printf("Author: %s\n", strings.TrimSpace(author))
		fmt.Printf("Date: %s\n", strings.TrimSpace(date))
		fmt.Printf("Image URL: %s\n", imageURL)
		fmt.Println("---")
	})
}
