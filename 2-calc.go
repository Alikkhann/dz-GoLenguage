package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	operation, arr := scanCalculation() 
	
	if operation == "AVG" {
		AVGoperation(arr)
	}else if operation == "SUM" {
		SUMoperation(arr)
	}else if operation == "MED" {
		MEDoperation(arr)
	}
}

func scanCalculation() (string, []int){
	var operation string
	var newSlice []int
	var errorCount int = 0

	fmt.Print("Выберите нужную операцию:(AVG/SUM/MED) ")
	fmt.Scan(&operation)
	for operation != "AVG" && operation != "SUM" && operation != "MED" {
		fmt.Print("Вы ввели неправильную операцию.")
		fmt.Scan(&operation)
	}
	fmt.Print("Введите числа через запятую: ")

	for {
		reader := bufio.NewReader(os.Stdin)
		numsStr, _ := reader.ReadString('\n')
		parts := strings.Split(numsStr, ",")
		for _, value := range parts {
			trims := strings.TrimSpace(value)
			elem, err := strconv.Atoi(trims)
			if err != nil {
				errorCount++
				continue
			}
			newSlice = append(newSlice, elem)
		}

	if errorCount != 0 {
		fmt.Print("Ошибка! Введите числа через запятую: ")
		errorCount = 0
		newSlice = nil
		continue
	}else {
		break
	}
	}
	return operation, newSlice
}

func AVGoperation(arr []int) {
	if len(arr) != 0 {
	res := 0.0
	count := 0.0
	for _, value := range arr {
		res += float64(value)
		count++
	}
	fmt.Printf("%.2f\n", res/count)
	}else {
		fmt.Print("Массив не содержит элементов!")
	}
}

func SUMoperation(arr []int) {
	res := 0
	for _, value := range arr {
		res += value
	}
	fmt.Println(res)
}

func MEDoperation(arr []int) {
	slice := arr[:]
	ln := len(slice)
	for i := 0; i < ln; i++ {
		min := i
		for j := i+1; j < ln; j++ {
			if slice[min] > slice[j] {
				min = j
			}
		}
		slice[min], slice[i] = slice[i], slice[min]
	}
	if ln % 2 != 0 {
		fmt.Print(slice[ln/2])
	}else if ln % 2 == 0 {
		res := float64(slice[ln/2-1] + slice[ln/2])/2.0
		fmt.Print(res)
	}	
}