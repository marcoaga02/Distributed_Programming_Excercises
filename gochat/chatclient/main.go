package main

import (
	"bufio"
	"flag"
	"fmt"
	"gochat/internal/client"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
)

const DEFAULT_ADDRESS = "localhost:5050"

func checkerr(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func main() {
	address := flag.String("address", DEFAULT_ADDRESS, "The address of the chat server in the format address:port")
	name := flag.String("name", fmt.Sprintf("Client%d", rand.Intn(30)), "The name to use in chat")
	flag.Parse()

	client := client.NewTcpClient()
	defer client.Close()

	checkerr(client.Dial(*address))
	checkerr(client.SetName(*name))

	go client.Start()
	go func() {
		for msg := range client.Incoming() {
			fmt.Printf("[%s]:%s\n", msg.Author, msg.Content)
		}
	}()

	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "Error read: %v", err)
			continue
		}

		client.SendMessage(strings.TrimSpace(line))
	}

}
