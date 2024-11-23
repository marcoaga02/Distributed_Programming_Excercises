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

// Write serializes a Message into a string and writes it to the underlying writer.
func (w *MessageWriter) Write(data []byte) (n int, err error) {
	return fmt.Fprintf(w.writer, "%s\n>>>", data)
}
