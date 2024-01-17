package main

import "fmt"

func main() {
	const Hours = 60
	var minutes int

	fmt.Scan(&minutes)
	hours := minutes / Hours
	mins := minutes - hours*Hours

	fmt.Println(minutes, "мин - это", hours, "час", mins, "минут.")
}
