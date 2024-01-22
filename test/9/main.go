package main

import "fmt"

func main() {
	var age int
	fmt.Scan(&age)
	fmt.Println(marsAge(age))
}

func marsAge(age int) int {
	return (age * 365) / 687
}
