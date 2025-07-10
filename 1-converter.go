package main

import (
	"fmt"
)
	const USDToEUR = 0.853
	const USDToRUB = 78.17
	const EURToRUB = USDToRUB/USDToEUR //91.64
	const EURToUSD = 1.17
  const RUBToUSD = 0.012
	const RUBToEUR = 0.10

func main() {
	// usrInp := userInput()
	// fmt.Println(usrInp)
	//error1, error2, error3 := errorExams(usrInp1, usrInp2, usrInp3)
	usrInp1, usrInp2, usrInp3 := userInput()
	switch { 
	case usrInp1 != "USD" && usrInp1 != "EUR" && usrInp1 != "RUB": 
		fmt.Println("Вы ввели неправильную валюту. Валюта для конвертации(USD/RUB/EUR)")
		fmt.Scan(&usrInp1)
	case usrInp2 < 0: 
		fmt.Println("Введите сумму для конвертации")
		fmt.Scan(&usrInp2)
	case usrInp3 != "USD" && usrInp3 != "EUR" && usrInp3 != "RUB": 
		fmt.Println("Вы ввели неправильную валюту. Валюта для конвертации(USD/RUB/EUR)")
		fmt.Scan(&usrInp3)
	}
	result := converter(usrInp1, usrInp2, usrInp3)
	fmt.Print(result)
}

func userInput() (string, float64, string) {
	var valut string
	var sum float64
	var targValut string

for valut != "USD" && valut != "RUB" && valut != "EUR" {
	fmt.Println("Введите валюту для конвертации(USD/RUB/EUR)")
	fmt.Scan(&valut)
	if valut != "USD" && valut != "EUR" && valut != "RUB" {
    fmt.Println("Вы ввели неправильную валюту.")
  //fmt.Scan(&valut)
	}
}
for sum <= 0 {
	fmt.Println("Введите число")
	fmt.Scan(&sum)
	if sum < 0 {
		fmt.Println("Неправильный ввод! Введите сумму для конвертации")
		//fmt.Scan(&sum)
	}
}
for targValut != "USD" && targValut != "RUB" && targValut != "EUR" {
	fmt.Println("Введите таргетную валюту(USD/RUB/EUR)")
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
	}
return res
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