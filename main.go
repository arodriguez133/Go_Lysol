package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/microcosm-cc/bluemonday"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter URL: ")
	url, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading URL: ", err)
	}

	url = strings.TrimSpace(url)

	// Make HTTP GET request
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Error fetching URL: ", err)
	}
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body: ", err)
	}

	// Use bluemonday to sanitize the document
	p := bluemonday.StrictPolicy()
	html, err := document.Html()
	if err != nil {
		log.Fatal("Error generating HTML: ", err)
	}
	text := p.Sanitize(html)

	// Print out the sanitized text
	fmt.Println(text)
}
