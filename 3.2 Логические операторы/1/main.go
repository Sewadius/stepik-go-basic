package main

import "fmt"

func main() {
	var result string = "NO"
	var number int
	fmt.Scan(&number)

	if number > 99 && number < 1000 {
		result = "YES"
	}

	fmt.Println(result)
}
