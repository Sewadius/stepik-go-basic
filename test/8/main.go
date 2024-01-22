package main

import "fmt"

func main() {
	var a, b int
	var s string
	fmt.Scan(&a, &b, &s)
	fmt.Println(oneOrTwo(a, b, s))
}

func oneOrTwo(a, b int, s string) (int, string) {
	if s == "one" {
		return a, s
	} else if s == "two" {
		return b, s
	}
	return 0, s
}
