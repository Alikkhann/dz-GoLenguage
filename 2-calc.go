package main

import (
	"fmt"
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
	var numsStr string
	var newSlice []int
	
	fmt.Print("Выберите нужную операцию:(AVG/SUM/MED) ")
	fmt.Scan(&operation)
	for operation != "AVG" && operation != "SUM" && operation != "MED" {
		fmt.Print("Вы ввели неправильную операцию.")
		fmt.Scan(&operation)
	}
	fmt.Print("Введите числа через запятую: ")
	fmt.Scanln(&numsStr)
	parts := strings.Split(numsStr, ",")
	for _, value := range parts {
		elem, err := strconv.Atoi(value)
		if err != nil {
			fmt.Print("Ошибка")
			continue
		}else{
			newSlice = append(newSlice, elem)
		}
	}
	
	return operation, newSlice
}

func AVGoperation(arr []int) {
	res := 0
	count := 0
	for _, value := range arr {
		res += value
		count++
	}
	fmt.Println(res/count)
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
		res := (slice[ln/2-1] + slice[ln/2])/2
		fmt.Print(res)
	}	
}








// arr[] int = {4, 1, 8, 5, 2}
// min := arr[0]
// if min < value {
// 	temp = min
// 	min = value
// 	value = temp
// }



