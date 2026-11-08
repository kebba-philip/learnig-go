package main

import "fmt"

func main() {
	// Task 1: Nested break
	fmt.Println("=== Task 1: Nested break ===")

	for i := 0; i < 3; i++ {
		fmt.Println("Outer loop i =", i)

		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break
			}
			fmt.Println(i, j)
		}
	}
	// Total lines printed by the loops: 10
	// The outer loop continues after the plain break.

	// Task 2: Labeled break
	fmt.Println("\n=== Task 2: Labeled break ===")

outer:
	for i := 0; i < 3; i++ {
		fmt.Println("Outer loop i =", i)

		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break outer
			}
			fmt.Println(i, j)
		}
	}
	// Total lines printed by the loops: 6
	// break outer exits both the inner and outer loops.

	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("\n=== Task 3: Search Matrix ===")

	if row, col, ok := searchMatrix(matrix, 5); ok {
		fmt.Printf("Found 5 at row %d, col %d\n", row, col)
	} else {
		fmt.Println("not found")
	}

}

func searchMatrix(matrix [][]int, target int) (row, col int, found bool) {
	for r, rowVals := range matrix {
		for c, v := range rowVals {
			if v == target {
				return r, c, true
			}
		}
	}

	return 0, 0, false

}
