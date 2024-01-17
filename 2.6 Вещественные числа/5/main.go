package main

import "fmt"

func main() {
	var number float64

	fmt.Scan(&number)
	var whole_part int = int(number)
	result := number - float64(whole_part)

	fmt.Println(result)
}
