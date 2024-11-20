package protocol

import "errors"

type Message interface{}

// SendMessage represents a text message sent by a user to others.
type SendMessage struct {
	Content string
}

// ChangeNameMessage is a message sent by a user to the server to change her name in the system.
type ChangeNameMessage struct {
	NewName string
}

// NotifyMessage is a message sent by the server to clients to notify them of a new message.
type NotifyMessage struct {
	Author  string
	Content string
}

var UnknownMessage = errors.New("Unknown message type")
