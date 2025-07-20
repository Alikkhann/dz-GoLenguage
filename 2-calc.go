package main

import "fmt"

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

func scanCalculation() (string, []float64){
	var operation string
	var arr []float64
	var element float64

	fmt.Print("Выберите нужную операцию:(AVG/SUM/MED) ")
	fmt.Scan(&operation)
	if operation != "AVG" && operation != "SUM" && operation != "MED" {
		fmt.Print("Вы ввели неправильную операцию.")
		scanCalculation()
	}
	fmt.Print("Введите числа: ")
	for {
		fmt.Scan(&element)
		if element == 0 {
			break
		}
		arr = append(arr, element)
	}
	return operation, arr
}

func AVGoperation(arr []float64) {
	res := 0.0
	count := 0.0
	for _, value := range arr {
		res += value
		count++
	}
	fmt.Println(res/count)
}

func SUMoperation(arr []float64) {
	res := 0.0
	for _, value := range arr {
		res += value
	}
	fmt.Println(res)
}

func MEDoperation(arr []float64) {
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



