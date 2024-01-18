package main

import "fmt"

func main() {
	var answer string = "NO"
	var number int
	fmt.Scan(&number)

	first := number % 10
	second := number / 10 % 10
	third := number / 100 % 10

	if first != second && third != second && first != third {
		answer = "YES"
	}

	fmt.Println(answer)
}
