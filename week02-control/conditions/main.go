package main

import "fmt"

func main() {
	score := 85
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	case score >= 60:
		fmt.Println("D")
	default:
		fmt.Println("F")

	}

	day := "Wed"
	switch day {
	case "Mon", "Tue", "Wed", "Thu", "Fri":
		fmt.Println("weekday")
	case "Sat", "Sun":
		fmt.Println("weekend")
	default:
		fmt.Println("unknown")
	}
	// FizzBuzz
	fizzbuzz(15)
	fizzbuzz(9)
	fizzbuzz(10)
	fizzbuzz(7)

}

func fizzbuzz(n int) {
	if n%3 == 0 && n%5 == 0 {
		fmt.Println("FizzBuzz")

	} else if remainder := n % 3; remainder == 0 {
		fmt.Println("Fizz")

	} else if remainder := n % 5; remainder == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(n)
	}
}
