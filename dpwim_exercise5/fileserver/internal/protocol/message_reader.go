package protocol

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

// MessageReader allows to deserialize a Message from a stream
type MessageReader struct {
	reader bufio.Reader
}

// NewMessageReader returns a new MessageReader that gets data from an underlying reader.
func NewMessageReader(reader io.Reader) *MessageReader {
	return &MessageReader{*bufio.NewReader(reader)}
}

// Read deserializes a Message from a stream.
func (r *MessageReader) Read() (Command, error) {
	str, err := r.reader.ReadString('\n')

	str = strings.TrimSpace(str)
	log.Printf("Command: %v\n", str)

	if err != nil {
		return nil, err
	}

	parts := strings.Split(str, " ")
	operationType := parts[0]

	switch operationType {
	case "ls":
		return NewLsCommand(), nil
	case "cat":
		if len(parts) >= 2 { // the file name is present
			return NewCatCommand(parts[1]), nil
		}
		return nil, fmt.Errorf("too few arguments. You should pass the file name with the %s command", operationType)
	case "rm":
		if len(parts) >= 2 { // the file name is present
			return NewRmCommand(parts[1]), nil
		}
		return nil, fmt.Errorf("too few arguments. You should pass the file name with the %s command", operationType)

	case "get":
		if len(parts) >= 2 { // the file name is present
			return NewGetCommand(parts[1]), nil
		}
		return nil, fmt.Errorf("too few arguments. You should pass the file name with the %s command", operationType)
	case "info":
		if len(parts) >= 2 { // the file name is present
			return NewInfoCommand(parts[1]), nil
		}
		return nil, fmt.Errorf("too few arguments. You should pass the file name with the %s command", operationType)
	default:
		return nil, ErrUnknownCommand
	}
}
