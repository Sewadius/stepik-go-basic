package main

import "fmt"

func main() {
	var phone, cover, charger, headphones int

	fmt.Scan(&phone, &cover, &charger, &headphones)
	price := 3 * (phone + cover + charger + headphones)
	fmt.Println(price)
}
