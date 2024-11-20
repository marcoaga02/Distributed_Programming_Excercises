package main

import (
	"fmt"
	"sync"
)

// Fetcher interface come prima
type Fetcher interface {
	// Fetch restituisce il corpo della pagina e una lista di URL trovati su quella pagina
	Fetch(url string) (body string, urls []string, err error)
}

// La funzione Crawl mantiene la firma invariata
func Crawl(url string, depth int, fetcher Fetcher) {
	// Mappa per tenere traccia degli URL già visitati
	visited := make(map[string]bool)

	// Mutex per sincronizzare l'accesso alla mappa visited
	var mu sync.Mutex

	// WaitGroup per aspettare che tutte le goroutine siano completate
	var wg sync.WaitGroup

	// Funzione per fare il crawl in parallelo
	var crawlHelper func(url string, depth int)

	crawlHelper = func(url string, depth int) {
		defer wg.Done() // Decrementa il contatore del WaitGroup quando la goroutine termina

		// Se il depth è 0, fermati
		if depth <= 0 {
			return
		}

		// Proteggi l'accesso alla mappa visited con il mutex
		mu.Lock()
		if visited[url] {
			mu.Unlock()
			return
		}
		visited[url] = true
		mu.Unlock()

		// Chiamata a fetcher.Fetch
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Stampa l'URL trovato e il corpo della pagina
		fmt.Printf("found: %s %q\n", url, body)

		// Avvia una goroutine per ciascun URL trovato
		for _, u := range urls {
			wg.Add(1) // Aggiungi un lavoro al WaitGroup
			go crawlHelper(u, depth-1)
		}
	}

	// Avvia la prima goroutine con l'URL iniziale
	wg.Add(1)
	go crawlHelper(url, depth)

	// Aspetta che tutte le goroutine siano completate
	wg.Wait()
}

func main() {
	// Test della funzione Crawl
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher come prima, senza cambiare nulla
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// Popolamento del fakeFetcher
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
