package main

import "fmt"

func main() {
	// Integer division trap
	fmt.Println("=== Integer Division ===")
	fmt.Println(10 / 3)     // integer division truncates
	fmt.Println(10.0 / 3.0) // float division
	fmt.Println(float64(10) / 3)

	// Modulo
	// returns the remainder
	fmt.Println("=== Modulo Operator ===")
	fmt.Println(10 % 3) // 1
	fmt.Println(20 % 4) // 0
	fmt.Println(17 % 5) // 2
	fmt.Println()

	fmt.Println("=== Comparisons ===")
	// returns a boolean value
	fmt.Println(5 == 5)
	fmt.Println(5 != 5)
	fmt.Println(10 > 3)
	fmt.Println(10 <= 3)
	fmt.Println()

	// Logical operators
	fmt.Println("=== Logical Operators ===")
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println()

	// Short-circuit demo
	var user *string

	if user != nil && *user == "Philip" { // Because Go's && operator uses short-circuit evaluation.
		fmt.Println("matched")
	} else {
		fmt.Println("no user or no match")
	}

	a, b, c := 5, 10, 15
	without := a + b*c  // 155 because multiplication comes before addition in go
	with := a + (b * c) // 155
	fmt.Println(without == with)
}
