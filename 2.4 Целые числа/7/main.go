package main

import "fmt"

func main() {
	var number int

	fmt.Scan(&number)
	number--
	for i := 0; i < 3; i++ {
		number++
		fmt.Println(number)
	}
}
