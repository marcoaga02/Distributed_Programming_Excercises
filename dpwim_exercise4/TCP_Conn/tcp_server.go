package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() { // run the server with args 8080
	PORT := ":9999"
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	netData, err := bufio.NewReader(c).ReadString('\n')
	netData = netData[:len(netData)-1] // to remove the end-line character

	if err != nil {
		fmt.Println(err)
		return
	}
	location, err := time.LoadLocation(netData)

	if err != nil {
		panic(err)
	}

	timeUTC := time.Now().In(location)
	fmt.Print("-> ", string(netData))
	c.Write([]byte(timeUTC.String()))
}
