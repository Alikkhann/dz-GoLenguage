package main

import (
	"fmt"
	"myProject/bins"
)



func main() {
	binList := bins.BinList{}
	var n int
	fmt.Println("Функция создания файлов типа `ключ` - `значение`.")
	fmt.Println("Сколько списков вы хотите создать?")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
	creatBin := bins.CreatBin()
	binList.Bins = append(binList.Bins, creatBin)
	}
	binList.PrintBinList()
}
