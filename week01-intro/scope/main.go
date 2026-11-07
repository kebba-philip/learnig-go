package main

import (
	"fmt"
	"strconv"
)

var pkgLevel = "visible everywhere in this package"

func main() {
	var age int
	var name string
	var isLearning bool

	fmt.Printf("zero int = %d\n", age)
	fmt.Printf("zero string = %q\n", name)
	fmt.Printf("zero bool = %t\n", isLearning)

	f := 3.99
	f2 := int(f)

	fmt.Printf("float %v -> int %d\n", f, f2)

	// The result is being explicitly converted to an int data type
	// not being rounded

	n := 65
	fmt.Printf("string(%d) = %q\n", n, string(n))
	fmt.Printf("strconv.Itoa(%d) = %q\n", n, strconv.Itoa(n))

	// Use strconv.Itoa for display: string(n) treats n as a Unicode code point,
	// so string(65) is "A", not "65".

	fmt.Printf("pkgLevel = %s\n", pkgLevel)
	funcLevel := "visible in main only"
	fmt.Printf("funcLevel = %s\n", funcLevel)

	if true {
		blockLevel := "local to this block"
		fmt.Printf("blockLevel = %s\n", blockLevel)
	}

	// fmt.Print(blockLevel)
	// error undefined: blockLevel
}
