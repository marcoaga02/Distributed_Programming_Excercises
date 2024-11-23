package protocol

import (
	"encoding/json"
	"fmt"
	"os"
)

type RmCommand struct {
	FileName string
}

type RmFileResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func NewRmCommand(fileName string) Command {
	return &RmCommand{fileName}
}

func (msg *RmCommand) Handle() ([]byte, error) {
	filePath := "dbox/" + msg.FileName

	err := os.Remove(filePath)
	if err != nil {
		response := RmFileResponse{
			Success: false,
			Message: fmt.Sprintf("Failed to remove file '%s': %v", msg.FileName, err),
		}
		return json.Marshal(response)
	}

	response := RmFileResponse{
		Success: true,
		Message: fmt.Sprintf("File '%s' removed successfully", msg.FileName),
	}
	return json.Marshal(response)
}
