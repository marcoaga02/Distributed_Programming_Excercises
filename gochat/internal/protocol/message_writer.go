package protocol

import (
	"fmt"
	"io"
)

// MessageWriter allows to serialize a Message to a stream
type MessageWriter struct {
	writer io.Writer
}

// NewMessageWriter returns a new MessageWriter that uses writer to serialize a Message
func NewMessageWriter(writer io.Writer) *MessageWriter {
	return &MessageWriter{writer}
}

// writeString write a message to the underlying writer.
// It is an auxiliary function for Write.
func (w *MessageWriter) writeString(message string) (n int, err error) {
	return w.writer.Write([]byte(message))
}

// Write serializes a Message into a string and writes it to the underlying writer.
func (w *MessageWriter) Write(message Message) (n int, err error) {
	switch v := message.(type) {
	case SendMessage:
		return w.writeString(fmt.Sprintf("SEND %s\n", v.Content))
	case ChangeNameMessage:
		return w.writeString(fmt.Sprintf("NAME %s\n", v.NewName))
	case NotifyMessage:
		return w.writeString(fmt.Sprintf("MESSAGE %s %s\n", v.Author, v.Content))
	default:
		return 0, UnknownMessage
	}
}
