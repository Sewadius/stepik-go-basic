package main

import "fmt"

func main() {
	const ok = "Доступ разрешен"
	const no = "Доступ запрещен"

	var age int
	fmt.Scan(&age)

	if age >= 12 {
		fmt.Println(ok)
	} else {
		fmt.Println(no)
	}
}
