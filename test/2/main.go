package main

import "fmt"

func main() {
	var number int
	var result = 1

	fmt.Scan(&number)

	for i := result; i <= number; i++ {
		result *= i
	}

	fmt.Println(result)
}
