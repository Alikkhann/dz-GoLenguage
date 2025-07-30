package main

import (
	"fmt"
)

func main() {
	convMap := map[string]map[string]float64 {
		"USD": {"EUR": 0.853, "RUB": 78.17},
		"RUB": {"USD": 0.012, "EUR": 0.10},
		"EUR": {"USD": 1.17, "RUB": 91.64},
	}
	valut, res, targVal, sum := userInput(convMap) 
	if valut == targVal {
		fmt.Println(sum)
	}else {
	convertResult(valut, res, targVal, sum)
	}
  // usrInp1, usrInp2, usrInp3 := userInput()
	// result := converter(usrInp1, usrInp2, usrInp3)
}

func userInput(convMap map[string]map[string]float64) (string, float64, string, float64) {
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
	res := sum * convMap[valut][targValut]
	//delete(convMap, "USD")
	return valut, res, targValut, sum
	
}


func convertResult(valut string, result float64, targVal string, sum float64){
	switch {
		case valut == "USD" && targVal == "EUR":
			fmt.Printf("%.2f доллар(ов) США = %.2f евро\n", sum, result)
		case valut == "USD" && targVal == "RUB":
			fmt.Printf("%.2f доллар(ов) США = %.2f рублей\n", sum, result)

		case valut == "EUR" && targVal == "USD":
			fmt.Printf("%.2f евро  = %.2f долларов\n", sum, result)
		case valut == "EUR" && targVal == "RUB":
			fmt.Printf("%.2f евро = %.2f рублей\n", sum, result)

		case valut == "RUB" && targVal == "USD":
			fmt.Printf("%.2f рубль(ей) = %.2f долларов\n", sum, result)
		case valut == "RUB" && targVal == "EUR":
			fmt.Printf("%.2f рубль(ей) = %.2f евро\n", sum, result)
	}
}


// func converter(valut string, sum float64, targValut string) float64 {
//   res := 0.0
// 	m := map[string]map[string]float64 {
// 		"USD": {"EUR": 0.853, "RUB": 78.17},
// 		"RUB": {"USD": 0.012, "EUR": 0.10},
// 		"EUR": {"USD": 1.17, "RUB": 91.64},
// 	}
// 	res = sum * m[valut][targValut]
// return res
// }

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

// func converterr() {
// 	m := map[string]float64 {
// 		"RUB": 1,
// 		"USD": 78.17,
// 		"EUR": 91.64,
// 	} 
// 	fmt.Print(m["USD"])
// }