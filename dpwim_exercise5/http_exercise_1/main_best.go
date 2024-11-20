package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// Funzione per estrarre i link dall'HTML
func extractLinks2(doc *html.Node, baseURL string) []string {
	var links []string
	linkSet := make(map[string]bool)

	// Naviga attraverso gli elementi del documento HTML
	var crawl func(*html.Node)
	crawl = func(n *html.Node) {
		// Controlla se l'elemento è un <a> con un attributo href
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					link := attr.Val
					// Se il link è relativo, lo completiamo con il baseURL
					if !strings.HasPrefix(link, "http://") && !strings.HasPrefix(link, "https://") {
						link = baseURL + link
					}
					linkSet[link] = true
				}
			}
		}
		// Continua a cercare nei nodi figli
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			crawl(c)
		}
	}

	// Avvia la ricerca dei link partendo dalla radice
	crawl(doc)

	for link := range linkSet {
		links = append(links, link)
	}

	return links
}

func main() {
	// Assicurati che l'URL venga passato come argomento
	if len(os.Args) < 2 {
		log.Fatal("URL missing")
	}
	url := os.Args[1]

	// Fai una richiesta GET per scaricare la pagina
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Parse l'HTML della risposta
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Estrai il baseURL dal link passato (protocollo + dominio)
	baseURL := url
	if strings.Contains(url, "://") {
		// Trova la parte del dominio
		protocolEndIndex := strings.Index(url, "://") + 3 // +3 per saltare "://"
		domainEndIndex := strings.Index(url[protocolEndIndex:], "/") + protocolEndIndex
		if domainEndIndex == protocolEndIndex-1 {
			domainEndIndex = len(url)
		}
		baseURL = url[:domainEndIndex]
	}

	// Estrai i link
	links := extractLinks2(doc, baseURL)

	// Stampa i link trovati
	fmt.Println("Found links:")
	for _, link := range links {
		fmt.Println(link)
	}
}
