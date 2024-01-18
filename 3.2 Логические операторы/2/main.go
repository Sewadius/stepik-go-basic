package main

import "fmt"

func main() {
	var answer string = "Не принадлежит"
	var x int
	fmt.Scan(&x)

	if x > -1 && x < 17 {
		answer = "Принадлежит"
	}

	fmt.Println(answer)
}
