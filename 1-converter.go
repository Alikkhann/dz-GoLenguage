package main

import (
	"fmt"
)

const USDToEUR = 0.853
const USDToRUB = 78.17
const EURToRUB = USDToRUB / USDToEUR //91.64
const EURToUSD = 1.17
const RUBToUSD = 0.012
const RUBToEUR = 0.10

func main() {
	usrInp1, usrInp2, usrInp3 := userInput()
	result := converter(usrInp1, usrInp2, usrInp3)
	convertResult(usrInp1, usrInp2, usrInp3, result)
}

func convertResult(usrInp1 string, usrInp2 float64, usrInp3 string, result float64) {
	switch {
	case usrInp1 == "USD" && usrInp3 == "EUR":
		fmt.Printf("%.2f доллар(ов) США = %.2f евро\n", usrInp2, result)
	case usrInp1 == "USD" && usrInp3 == "RUB":
		fmt.Printf("%.2f доллар(ов) США = %.2f рублей\n", usrInp2, result)

	case usrInp1 == "EUR" && usrInp3 == "USD":
		fmt.Printf("%.2f евро  = %.2f долларов\n", usrInp2, result)
	case usrInp1 == "EUR" && usrInp3 == "RUB":
		fmt.Printf("%.2f евро = %.2f рублей\n", usrInp2, result)

	case usrInp1 == "RUB" && usrInp3 == "USD":
		fmt.Printf("%.2f рубль(ей) = %.2f долларов\n", usrInp2, result)
	case usrInp1 == "RUB" && usrInp3 == "EUR":
		fmt.Printf("%.2f рубль(ей) = %.2f евро\n", usrInp2, result)
	}
}

func userInput() (string, float64, string) {
	var valut string
	var targValut string
	var sum float64
	for valut != "USD" && valut != "RUB" && valut != "EUR" {
		fmt.Println("Добрый день! Программа для перевода одной валюты в другую. Введите валюту для конвертации(USD/RUB/EUR)")
		fmt.Scan(&valut)
		if valut != "USD" && valut != "EUR" && valut != "RUB" {
			fmt.Println("Вы ввели неправильную валюту.")
			//fmt.Scan(&valut)
		}
	}
	sum = inputAmount()

	for targValut != "USD" && targValut != "RUB" && targValut != "EUR" {
		fmt.Println("Введите валюту в которую необходимо конвертировать(USD/RUB/EUR)")
		fmt.Scan(&targValut)
		if targValut != "USD" && targValut != "EUR" && targValut != "RUB" {
			fmt.Println("Вы ввели неправильную валюту.")
			//fmt.Scan(&targValut)
		}
	}
	return valut, sum, targValut
}

func converter(valut string, sum float64, targValut string) float64 {
	var res float64
	switch {
	case valut == targValut && valut == "USD":
		res = sum
	case valut == targValut && valut == "EUR":
		res = sum
	case valut == targValut && valut == "RUB":
		res = sum
	case valut == "USD" && targValut == "EUR":
		res = USDToEUR * sum
	case valut == "USD" && targValut == "RUB":
		res = USDToRUB * sum
	case valut == "EUR" && targValut == "USD":
		res = EURToUSD * sum
	case valut == "EUR" && targValut == "RUB":
		res = EURToRUB * sum
	case valut == "RUB" && targValut == "USD":
		res = RUBToUSD * sum
	case valut == "RUB" && targValut == "EUR":
		res = RUBToEUR * sum
	default:
		fmt.Println("Неправильный ввод!")
	}
	return res
}

func inputAmount() float64 {
	var sum float64
	for sum <= 0 {
		fmt.Println("Введите число!")
		fmt.Scan(&sum)
		if sum <= 0 {
			fmt.Println("Неправильный ввод!")
			//fmt.Scan(&sum)
		}
	}
	return sum
}

// if valut == "USD" && targValut == "EUR" {
// 	res = USDToEUR * sum
// }else {
// 	if valut == "USD" && targValut == "RUB" {
// 	res = USDToRUB * sum
// 	}
// }

// if valut == "EUR" && targValut == "RUB" {
// 	res = EURToRUB * sum
// }else {
// 	if valut == "EUR" && targValut == "USD" {
// 	res = EURToUSD * sum
// 	}
// }

// if valut == "RUB" && targValut == "EUR" {
// 	res = RUBToEUR * sum
// }else {
// 	if valut == "RUB" && targValut == "USD" {
// 	res = RUBToUSD * sum
// 	}
// }

// func errorExams(arg1 string, arg2 float64, arg3 string) (err1, err2, err3 error) {
// 	if arg1 != "USD" && arg1 != "EUR" && arg1 != "RUB" {
// 	  err1 = errors.New("Вы ввели неправильную валюту. Валюта для конвертации(USD/RUB/EUR)")
// 		return err1, err2, err3
// 	} else if arg2 < 0 {
// 		err2 = errors.New("Введите сумму для конвертации")
// 		return err1, err2, err3
// 	} else if arg3 != "USD" && arg3 != "EUR" && arg3 != "RUB" {
// 		err3 = errors.New("Вы ввели неправильную валюту. Валюта для конвертации(USD/RUB/EUR)")
// 		return err1, err2, err3
// 	}
// 	return nil, nil, nil
// }
