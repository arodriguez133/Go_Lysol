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

	fmt.Println(`
	 _______  _______          ___      __   __  _______  _______  ___     
	|       ||       |        |   |    |  | |  ||       ||       ||   |    
	|    ___||   _   |        |   |    |  |_|  ||  _____||   _   ||   |    
	|   | __ |  | |  |        |   |    |       || |_____ |  | |  ||   |    
	|   ||  ||  |_|  |        |   |___ |_     _||_____  ||  |_|  ||   |___ 
	|   |_| ||       | _____  |       |  |   |   _____| ||       ||       |
	|_______||_______||_____| |_______|  |___|  |_______||_______||_______|
		   `)

	fmt.Println("Yo yo yo it's ya app Go_Lysol")
	fmt.Println("")
	fmt.Print("Enter URL to GO GO GO: ")
	url, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Aw shit! Error reading URL: ", err)
	}

	url = strings.TrimSpace(url)

	fmt.Print("Got it fam. Did you want that with a side of gangsta? Y/N:")
	shouldGangsta, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Aw shit! something happened")
	}

	// Make HTTP GET request
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Aw shit! Error fetching URL: ", err)
	}
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Aw shit! Error loading HTTP response body: ", err)
	}

	// Use bluemonday to sanitize the document
	p := bluemonday.StrictPolicy()
	html, err := document.Html()
	if err != nil {
		log.Fatal("Aw shit! Error generating HTML: ", err)
	}
	text := p.Sanitize(html)

	// Print out the sanitized text
	// but first check if they want the gangsta
	if shouldGangsta == "Y" || shouldGangsta == "y" {
		text = gangstify(text)
	}
	fmt.Println(text)

	fmt.Println("Operation is complete yo")
}

func gangstify(sentence string) string {

	replacements := map[string]string{
		"you":   "ya",
		"are":   "be",
		"going": "rollin'",
		"to":    "ta",
		"the":   "da",
		"your":  "ya",
		"for":   "fo'",
		"and":   "an'",
		"with":  "wit'",
	}

	for oldWord, newWord := range replacements {
		sentence = strings.Replace(sentence, oldWord, newWord, -1)
		sentence = strings.Replace(sentence, strings.Title(oldWord), strings.Title(newWord), -1)
	}

	return sentence
}
