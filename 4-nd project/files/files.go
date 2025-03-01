package files

import (
	"Golang/password/output"
	"os"
	"github.com/fatih/color"
)

type JsonDb struct {
	filename string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		filename: name,
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.filename)
	if err != nil {
		return nil, err
	}
	return data, nil

}

func (db *JsonDb) Write(content []byte) {
	file, err := os.Create(db.filename)
	if err != nil {
		output.PrintError(err)
	}
	_, err = file.Write(content)
	defer file.Close()
	if err != nil {
		output.PrintError(err)
	}
	color.Black("Запись успешна")
}
