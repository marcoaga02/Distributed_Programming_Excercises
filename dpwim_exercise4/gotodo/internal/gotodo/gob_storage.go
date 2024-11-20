package gotodo

import (
	"encoding/gob"
	"os"
)

type GobStorage struct {
	FileName string
}

func NewGobStorage(filename string) *GobStorage {
	return &GobStorage{filename}
}

func (storage *GobStorage) Save(data []Todo) error {
	// 0644 (rw-r--r--)
	fileData, err := os.OpenFile(storage.FileName, os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		return err
	}

	defer fileData.Close()

	enc := gob.NewEncoder(fileData)
	return enc.Encode(data)
}

func (storage *GobStorage) Load(data *[]Todo) error {
	fileData, err := os.Open(storage.FileName)

	if err != nil {
		return err
	}
	defer fileData.Close()

	dec := gob.NewDecoder(fileData)
	return dec.Decode(data)

}
