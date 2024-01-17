package main

import "fmt"

func main() {
	var number int

	fmt.Scan(&number)

	first := number % 10
	second := number / 10 % 10
	third := number / 100 % 10

	fmt.Println(first + second + third)
}
