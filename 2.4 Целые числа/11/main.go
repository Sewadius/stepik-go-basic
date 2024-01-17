package main

import "fmt"

func main() {
	var number int

	fmt.Scan(&number)

	next := number + 1
	prev := number - 1

	fmt.Println("Следующее за числом", number, "число:", next)
	fmt.Println("Для числа", number, "предыдущее число:", prev)
}
