package main

import "fmt"

func main() {
	var answer string = "Не принадлежит"
	var x int
	fmt.Scan(&x)

	if x <= -3 || x >= 7 {
		answer = "Принадлежит"
	}

	fmt.Println(answer)
}
