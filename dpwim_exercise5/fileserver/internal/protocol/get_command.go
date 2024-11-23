package protocol

type GetCommand struct {
	FileName string
}

func NewGetCommand(fileName string) Command {
	return &GetCommand{fileName}
}

func (msg *GetCommand) Handle() ([]byte, error) {
	return nil, nil
}
