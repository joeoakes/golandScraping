package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
)

func main() {
	url := "https://example.com" // Replace with the URL you want to crawl
	crawl(url)
}

// This function will crawl a Website
func crawl(url string) {

	// Test html load
	file, err := os.Open("example.html")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	// Parse the HTML file
	doc, err := html.Parse(file)
	if err != nil {
		log.Fatalf("Error parsing HTML: %v", err)
	}

	// Go and get the Web resource response
	/*resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to crawl %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to crawl %s: %s\n", url, resp.Status)
		return
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Printf("Failed to parse HTML: %v\n", err)
		return
	}
	*/

	// Find all anchor tags in the HTML
	var visitNode func(*html.Node)
	visitNode = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					fmt.Println(attr.Val)
				}
			}
		}
		// Recursively traverse child nodes
		for child := n.FirstChild; child != nil; child = child.NextSibling {
			visitNode(child)
		}
	}
	visitNode(doc)
}
