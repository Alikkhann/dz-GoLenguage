package storage

import (
	"encoding/json"
	"fmt"
	// "myproject/files"
	"os"
	"github.com/fatih/color"
	// "myproject/bins"
)

type DbFiles interface {
	ReadAnyFile(string) ([]byte, error)
}

func SaveBinlistToFile(bin any, filename string) {
	file, err := ToBytes(bin) 
	if err != nil {
		color.Yellow("Не удалось преобразовать")
		return
	}
	writeFile(file, filename)
}

func ReadJsonFile(db DbFiles, filename string, v interface{}) (error) {
	data, err := db.ReadAnyFile(filename)
	if err != nil {
		color.Red("Ошибка")
		return err
	}
  err = json.Unmarshal(data, v)
	if err != nil {
		color.Red("Не удалось разобрать файл data.json")
		return err
	}
	return nil
	}


func 	ToBytes(bin any) ([]byte, error) {
	data, err := json.Marshal(bin)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func writeFile(myString []byte, name string) {
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