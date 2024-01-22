package main

import "fmt"

func main() {
	var number int

	fmt.Scan(&number)
	fmt.Println(isEven(number))
}

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}
