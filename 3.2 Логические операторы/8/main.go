package main

import "fmt"

func main() {
	var answer = "NO"
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	if a == c || b == d {
		answer = "YES"
	}

	fmt.Println(answer)
}
