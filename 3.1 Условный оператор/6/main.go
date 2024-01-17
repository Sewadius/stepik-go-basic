package main

import "fmt"

func main() {
	var number int
	var result int = -1

	fmt.Scan(&number)

	if number > 0 {
		result = 1
	}

	if number == 0 {
		result = 0
	}

	fmt.Println(result)
}
