package storage

import (
	"encoding/json"
	"fmt"
	"myproject/bins"
	"myproject/files"
	"os"
	"github.com/fatih/color"
)

func SaveBinlistToFile(bin bins.BinList, filename string) {
	file, err := ToBytes(bin) 
	if err != nil {
		color.Yellow("Не удалось преобразовать")
		return
	}
	writeFile(file, filename)
}

func ReadJsonFile(filename string) []byte{
	data, err := files.ReadAnyFile(string(filename))
	if err != nil {
		color.Red("Ошибка")
		return data
	}
	var bins bins.BinList
  err = json.Unmarshal(data, &bins)
	if err != nil {
		color.Red("Не удалось разобрать файл data.json")
		}
	return data
	}


func 	ToBytes(bin bins.BinList) ([]byte, error) {
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