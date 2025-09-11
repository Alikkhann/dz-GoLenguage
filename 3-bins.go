package main

import (
	"fmt"
	"encoding/json"
	"myproject/bins"
	"myproject/files"
	"myproject/storage"
)



func main() {
	binList := bins.BinList{}
	filesStruct := files.Files{}
	data, err := filesStruct.ReadAnyFile("file.json")
	if err == nil {
		err = json.Unmarshal(data, &binList)
			if err != nil {
				fmt.Print("Ошибка разбора JSON")
			}
	}
	var n int
	fmt.Println("Функция создания файлов типа `ключ` - `значение`.")
	fmt.Println("Сколько списков вы хотите создать?")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
	creatBin := bins.CreateBin()
	binList.Bins = append(binList.Bins, creatBin)
}
	data, err = storage.ToBytes(binList)
	if err != nil {
		fmt.Println("Ошибка!")
		return
	}
	files.WriteFile(data, "file.json")
	
	binList.PrintBinList()
}
