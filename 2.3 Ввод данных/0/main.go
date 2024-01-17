package main

import "fmt"

func main() {
	fmt.Println("Привет! Как тебя зовут?")

	var name string
	fmt.Scan(&name)

	fmt.Println("Привет, " + name)
}
