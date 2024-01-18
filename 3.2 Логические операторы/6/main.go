package main

import "fmt"

func main() {
	var answer = "NO"
	var year int
	fmt.Scan(&year)

	if year%400 == 0 || year%4 == 0 && year%100 != 0 {
		answer = "YES"
	}

	fmt.Println(answer)
}
