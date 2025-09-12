package main

import (
	"fmt"
	"encoding/json"
	"myproject/bins"
	"myproject/files"
	"myproject/storage"
)



func main() {
	var filesmanager files.FilesManager = &files.Files{}
	var storfiles storage.StorManager = &storage.StorManager{}
	binList := bins.BinList{}
	var binsmanager bins.BinsManager = binList
	data, err := filesmanager.ReadAnyFile("file.json")
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
	creatBin := binsmanager.CreateBin()
	binList.Bins = append(binList.Bins, creatBin)
}
	data, err = storage.ToBytes(binList)
	if err != nil {
		fmt.Println("Ошибка!")
		return
	}
	filesmanager.WriteFile(data, "file.json")
	
	binsmanager.PrintBinList()
}
