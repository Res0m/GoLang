package files

import (
	"fmt"
	"os"
)

func ReadFile() {
	data, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	_, err = file.Write(content)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("Запись успешна")
}