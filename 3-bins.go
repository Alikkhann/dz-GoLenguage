package main

import (
	"fmt"
	"myproject/bins"
	"myproject/files"
	"myproject/storage"
)



func main() {
	binList := bins.BinList{}
	var n int
	fmt.Println("Функция создания файлов типа `ключ` - `значение`.")
	fmt.Println("Сколько списков вы хотите создать?")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
	creatBin := bins.CreateBin()
	binList.Bins = append(binList.Bins, creatBin)
	}
	data, err := storage.ToBytes(binList)

	files.WriteFile(data, "file.json")
	_, err = files.ReadAnyFile("file.json")
	if err != nil {
		return
	}
	binList.PrintBinList()
}
