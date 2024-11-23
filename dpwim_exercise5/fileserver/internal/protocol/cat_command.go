package protocol

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type CatCommand struct {
	FileName string
}

type CatFileResponse struct {
	FileName string `json:"fileName"`
	Content  string `json:"content"`
}

func NewCatCommand(fileName string) Command {
	return &CatCommand{fileName}
}

func (msg *CatCommand) Handle() ([]byte, error) {
	filePath := "dbox/" + msg.FileName
	file, err := os.Open(filePath)

	if err != nil {
		return nil, fmt.Errorf("error opening the file '%s': %v", msg.FileName, err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)

	if err != nil {
		return nil, fmt.Errorf("error reading the file '%s': %v", msg.FileName, err)
	}

	response := CatFileResponse{
		FileName: msg.FileName,
		Content:  string(content),
	}

	return json.Marshal(response)
}
