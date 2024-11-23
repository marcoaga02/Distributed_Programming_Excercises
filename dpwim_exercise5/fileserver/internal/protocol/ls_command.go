package protocol

import (
	"encoding/json"
	"fmt"
	"os"
)

type LsCommand struct{}

func NewLsCommand() Command {
	return &LsCommand{}
}

func (msg *LsCommand) Handle() ([]byte, error) {
	dirPath := "dbox"

	files, err := os.ReadDir(dirPath)

	if err != nil {
		return nil, fmt.Errorf("error reading the directory '%s': %v", dirPath, err)
	}

	var filesName []string
	for _, file := range files {
		filesName = append(filesName, file.Name())
	}

	return json.Marshal(filesName)
}
