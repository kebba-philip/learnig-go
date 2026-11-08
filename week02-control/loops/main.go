package main

import "fmt"

func main() {
	fmt.Println("=== 1. Count 1-10 ===")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	fmt.Println()
	sum := 0
	fmt.Println("=== Sum 1 to 100 ===")
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
	fmt.Println()
	fmt.Println("=== Range Over a Slice ===")
	nums := []int{5, 10, 15, 20, 25}

	for i, num := range nums {
		fmt.Printf("%v: %v\n", i, num)
	}
	fmt.Println()
	fmt.Println("=== Skipping Evens ===")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println()
	fmt.Println("=== FizzBuzz ===")
	for i := 1; i <= 100; i++ {
		fizzbuzz(i)
	}

}

func fizzbuzz(n int) {
	switch {
	case n%3 == 0 && n%5 == 0:
		fmt.Println("FizzBuzz")
	case n%3 == 0:
		fmt.Println("Fizz")
	case n%5 == 0:
		fmt.Println("Buzz")
	default:
		fmt.Println(n)
	}
}
