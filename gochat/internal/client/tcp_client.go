package client

import (
	"gochat/internal/protocol"
	"io"
	"log"
	"net"
)

type TcpClient struct {
	conn     net.Conn
	writer   *protocol.MessageWriter
	reader   *protocol.MessageReader
	name     string
	incoming chan protocol.NotifyMessage
}

func NewTcpClient() *TcpClient {
	return &TcpClient{incoming: make(chan protocol.NotifyMessage)}
}

func (c *TcpClient) Close() {
	c.conn.Close()
}

func (c *TcpClient) Send(message any) error {
	_, err := c.writer.Write(message)
	return err
}

func (c *TcpClient) SendMessage(message string) error {
	return c.Send(protocol.SendMessage{Content: message})
}

func (c *TcpClient) SetName(newName string) error {
	c.name = newName
	return c.Send(protocol.ChangeNameMessage{NewName: newName})
}

func (c *TcpClient) Incoming() chan protocol.NotifyMessage {
	return c.incoming
}

func (c *TcpClient) Dial(address string) error {
	conn, err := net.Dial("tcp", address)

	if err == nil {
		c.conn = conn
	}

	c.reader = protocol.NewMessageReader(conn)
	c.writer = protocol.NewMessageWriter(conn)

	return err
}

func (c *TcpClient) Start() {
	for {
		msg, err := c.reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Printf("Error read: %v\n", msg)
			continue
		}

		switch v := msg.(type) {
		case protocol.NotifyMessage:
			c.incoming <- v
		default:
			log.Printf("Error unknown message: %v", v)
		}
	}
}
