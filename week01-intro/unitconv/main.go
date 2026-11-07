package main

import "fmt"

const ratio float64 = 9.0 / 5.0
const banner = "=== Unit Converter ==="

func main() {

	fmt.Println(banner)
	fmt.Println()
	celsius1 := 0.0
	celsius2 := 100.00
	celsius3 := 37.0

	//converting from Celsius to fahrenheit
	fahrenheit1 := (celsius1 * ratio) + 32
	fahrenheit2 := (celsius2 * ratio) + 32
	fahrenheit3 := (celsius3 * ratio) + 32

	// displaying results
	fmt.Printf("%.2f°C = %.2f°F\n", celsius1, fahrenheit1)
	fmt.Printf("%.2f°C = %.2f°F\n", celsius2, fahrenheit2)
	fmt.Printf("%.2f°C = %.2f°F\n", celsius3, fahrenheit3)

}
