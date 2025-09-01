package files

import (
	"os"
	"path/filepath"
	"fmt"
	"github.com/fatih/color"
)

func ReadAnyFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil{
		return nil, err
	}
	return data, nil
}

func ExamFile(name string) {
	if filepath.Ext(name) == ".json" {
 	 color.Red("Расширение переданного файла .JSON")
	}else{
		color.Red("Расширение переданного файла не .JSON")
	}
}

func WriteFile(myString []byte, name string) {
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

