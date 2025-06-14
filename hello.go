package main

import "fmt"

func main(){
	/*

	// manifest typing is used in Go, meaning you must declare the type of variable
	// In Go, you can declare a variable using the var keyword followed by the variable name and type
	var firstName string = "Fadhli"


	fmt.Printf("Hello, %s\n", firstName) // Printf is used for formatted output and %s is a placeholder for string
	fmt.Println("Hello", firstName, "!")
	fmt.Println("Hello" + firstName) // plus operator is used for string concatenation

	// Type inference is also available in Go, where the type is inferred from the value assigned
	// but this badpractice is not recommended, because it can lead to confusion
	y := 10 // int32 is the default type for integers in Go

	fmt.Println("The value of y is", y) // Println is used for simple output without formatting

	var thirdNumber int8 = 45
	fmt.Printf("Third number is %d\n", thirdNumber)
	var decimalNumber float32 = 3.12
	fmt.Printf("Decimal number is %.2f\n", decimalNumber) // %.2f is used for formatting float numbers to 2 decimal places

	var isExisting bool = true
	fmt.Printf("Is existing? %t\n", isExisting) // %t is used for boolean values

	var message = `Nama saya "Fadhli njir".
	Salam kenal.
	Mari belajar "Golang".`

	fmt.Println(message) // Backticks are used for raw string literals, allowing multi-line strings and special characters without escaping
	*/

	const speedOfLight int32 = 299792458 // Constants are declared using the const keyword, and their value cannot be changed
	fmt.Printf("The speed of light is %d m/s\n", speedOfLight) // %d is used for integers

	fmt.Println("Hello, world")
	fmt.Println("Hello, world")

	fmt.Print("Hello, world")
	fmt.Print("Hello, world")

	const (
		pi = 3.14 // Constants can also be declared in a block, and they can be of any type
		areYouHappy bool = false
	)

	fmt.Println("Pi is approximately", pi) // Println automatically adds a newline at the end
	fmt.Println("Are you happy?", areYouHappy) // Println can also be used for boolean values

	var value uint8 = (((2 + 6) % 3) * 4 - 2) / 3
	fmt.Println("The value of the expression is", value) // This will print the result of the expression

	var numbers uint8 = 54

	if checkResult := numbers; checkResult <= 60{
		fmt.Println("The number is less than or equal to 60")
	} else {
		fmt.Println("The number is greater than 60")
	}

	for i := 0; i < 5; i++ {
		fmt.Println("Iteration number:", i+1) // Println is used for simple output without formatting
	}

}