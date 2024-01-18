package main

import "fmt"

func main() {
	var answer string = "NO"
	var number int
	fmt.Scan(&number)

	a := number % 10
	b := number / 10 % 10
	c := number / 100 % 10
	d := number / 1000 % 10
	e := number / 10000 % 10
	f := number / 100000 % 10

	if a+b+c == d+e+f {
		answer = "YES"
	}

	fmt.Println(answer)
}
