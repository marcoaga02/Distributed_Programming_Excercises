package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func extractLinks(html string, baseURL string) []string {
	re := regexp.MustCompile(`(?i)<a\s+[^>]*href="([^"]+)"[^>]*>`)
	matches := re.FindAllStringSubmatch(html, -1)
	linkSet := make(map[string]bool)
	var links []string
	for _, match := range matches {
		link := match[1]
		if !strings.HasPrefix(link, "http://") && !strings.HasPrefix(link, "https://") {
			link = baseURL + link
		}
		// to avoid duplicates
		linkSet[link] = true // if the element is present, simply the value is changed to true. If not present, a new value with key link is added to the map
	}

	for link := range linkSet {
		links = append(links, link)
	}

	return links
}

func main() {
	//url := "http://www.example.com"
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <URL>")
	}

	url := os.Args[1]
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	baseURL := url
	if strings.Contains(url, "://") {
		// Find the position of the first `/` after the `://` to get the domain part
		protocolEndIndex := strings.Index(url, "://") + 3 // +3 to skip "://"
		// Now find the first `/` after the protocol
		domainEndIndex := strings.Index(url[protocolEndIndex:], "/") + protocolEndIndex
		if domainEndIndex == protocolEndIndex-1 {
			// If no `/` is found, then the domain is the whole URL after the protocol
			domainEndIndex = len(url)
		}
		baseURL = url[:domainEndIndex]
	}
	fmt.Println(baseURL)

	links := extractLinks(string(data), baseURL)
	fmt.Println("Found links:")
	for _, link := range links {
		fmt.Println(link)
	}

}
