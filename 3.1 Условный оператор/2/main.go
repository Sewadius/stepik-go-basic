package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var password, confirm string
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	password = scanner.Text()
	scanner.Scan()
	confirm = scanner.Text()

	if password == confirm {
		fmt.Println("Пароль принят")
	} else {
		fmt.Println("Пароль не принят")
	}
}
