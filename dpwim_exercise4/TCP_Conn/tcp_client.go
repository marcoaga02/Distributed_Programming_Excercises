package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %s\n", err)
		os.Exit(1)
	}
	filename := os.Getenv("IANA_TIME_ZONE")

	connect := "127.0.0.1:9999"
	c, err := net.Dial("tcp", connect)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(c, filename+"\n")

	message, _ := bufio.NewReader(c).ReadString('\n')
	fmt.Println("->: " + message)
}
