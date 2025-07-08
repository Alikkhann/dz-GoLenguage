package main

import "fmt"
	const USDToEUR = 0.852 
	const USDToRUB = 78.72
	const EURToRUB = USDToRUB/USDToEUR

func main() {
	usrInp := userInput()
	fmt.Println(usrInp)
}

func userInput() int64{
	var usrInput int64
	fmt.Scan(&usrInput)
	return usrInput
}

func converter(int64, string, string) int64 {
//return 
}