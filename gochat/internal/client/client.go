package client

import "gochat/internal/protocol"

type Client interface {
	Dial(address string) error
	Start()
	Close()
	Send(message any) error
	SetName(newName string) error
	SendMessage(message string) error
	Incoming() chan protocol.NotifyMessage
}
