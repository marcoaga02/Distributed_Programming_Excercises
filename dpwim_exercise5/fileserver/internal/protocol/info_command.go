package protocol

type InfoCommand struct {
	FileName string
}

func NewInfoCommand(fileName string) Command {
	return &InfoCommand{fileName}
}

func (cmd *InfoCommand) Handle() ([]byte, error) {
	return nil, nil
}
