package main

import "fmt"

func main() {
	var answer = "NO"
	var x int
	fmt.Scan(&x)

	check := x >= -3 && x <= 1 || x >= 5 && x <= 9

	if check {
		answer = "YES"
	}

	fmt.Println(answer)
}
