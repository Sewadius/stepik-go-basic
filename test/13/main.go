package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	WriteNumber(a)
	WriteNumber(b)
	WriteNumber(c)
}

func WriteNumber(number int) {
	var name string = "Ноль"

	if number == 1 {
		name = "Один"
	} else if number == 2 {
		name = "Два"
	} else if number == 3 {
		name = "Три"
	} else if number == 4 {
		name = "Четыре"
	} else if number == 5 {
		name = "Пять"
	} else if number == 6 {
		name = "Шесть"
	} else if number == 7 {
		name = "Семь"
	} else if number == 8 {
		name = "Восемь"
	} else if number == 9 {
		name = "Девять"
	} else if number == 10 {
		name = "Десять"
	}

	fmt.Println(name)
}
