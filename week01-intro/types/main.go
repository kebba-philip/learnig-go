package main

import "fmt"

func main() {
	var name string = "Kebba"
	var age int = 20
	var userID int64 = 1001
	balance := 42.5
	isActive := true
	var firstByte uint8 = 65

	const appName string = "GoAcademy"
	const portNumber int = 8080

	fmt.Printf("Name = %s (type %T)\n", name, name)
	fmt.Printf("age = %d (type %T)\n", age, age)
	fmt.Printf("UserID = %d (type %T)\n", userID, userID)
	fmt.Printf("balance = %v (type %T)\n", balance, balance)
	fmt.Printf("IsActive = %v (type %T)\n", isActive, isActive)
	fmt.Printf("firstByte = %d (type %T)\n", firstByte, firstByte)

	ageAsFloat := float64(age)
	fmt.Printf("Age = %.2f (type %T)\n", ageAsFloat, ageAsFloat)

	fmt.Printf("App:%s\n", appName)
	fmt.Printf("Port:%v\n", portNumber)

}
