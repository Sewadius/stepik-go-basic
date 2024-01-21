package main

import "fmt"

func main() {
	var amount int
	fmt.Scan(&amount)

	if amount > 1000 {
		fmt.Println("Apple")
	} else if amount >= 500 && amount <= 1000 {
		fmt.Println("Samsung")
	} else {
		fmt.Println("Nokia с фонариком")
	}
}
