/*
 * This example shows how to make a simple HTTP POST request using the net/http
 * package. We will send a POST request along with
 * some data in the request body.
 */

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	url := "https://httpbin.org/post"

	// create post body
	body := strings.NewReader("This is the request body.")

	resp, err := http.Post(url, "text/plain", body)
	if err != nil {
		// we will get an error at this stage if the request fails, such as if the
		// requested URL is not found, or if the server is not reachable.
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// print the status code
	fmt.Println("Status:", resp.Status)

	// Leggi il corpo della risposta
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Stampa il corpo della risposta
	fmt.Println("Response body:", string(responseBody))
}
