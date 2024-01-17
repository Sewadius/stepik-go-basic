package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)

	result := "NO"

	if number%2 == 0 {
		result = "YES"
	}

	fmt.Println(result)
}
