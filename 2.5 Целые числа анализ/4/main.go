package main

import "fmt"

func main() {
	var number int

	fmt.Scan(&number)

	first := number / 100
	second := number / 10 % 10
	third := number % 10

	result := third*100 + second*10 + first
	fmt.Println(result)
}
