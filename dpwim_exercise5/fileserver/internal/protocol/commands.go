package protocol

import "errors"

type Command interface {
	Handle() ([]byte, error)
}

var ErrUnknownCommand = errors.New("unknown command")
