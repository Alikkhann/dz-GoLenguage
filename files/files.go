package files

import (
	"os"
	"path/filepath"
	"fmt"
)

type FilesManager interface {
	ReadAnyFile(string) ([]byte, error)
	WriteFile([]byte, string)
}

type Files struct {}

func (f *Files) ReadAnyFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil{
		return nil, err
	}
	return data, nil
}

func ExamFile(name string) bool{
 	 return filepath.Ext(name) == ".json"
}

func (f *Files) WriteFile(myString []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, err = file.Write(myString)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")
}