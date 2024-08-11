package main

import "fmt"

var laptop string = "Asus"

func main() {
	var country = "Singapore"
	fmt.Println(country)
	var country2 string
	country2 = "South Korea"
	fmt.Println(country2)

	var temperature = 25.5
	fmt.Println(temperature)

	var temperature2 int
	fmt.Println(temperature2)
	temperature2 = 36
	fmt.Println("The temperature today is ", temperature)
	fmt.Printf("The temperature2 today is %d celsius\n", temperature2)

	// MULTIPLE VARIABLE DECLARATION
	var num1, num2 int = 71, 72
	fmt.Println("num1: ", num1, " num2: ", num2)

	// Calculate area of rectangle
	var width, height float64 = 81.5, 91.0
	fmt.Println("Area of Rectangle is: ", width*height)

	// Zero values examples
	var truth bool
	fmt.Println("Default value for Truth is: ", !truth)

	var score float64
	var name string
	var amount int

	fmt.Println("------------------------")
	fmt.Println("Default values are: ")
	fmt.Printf("score: %f, name: %s, amount: %d, truth: %v\n", score, name, amount, truth)
	fmt.Println("------------------------")

	// Short hand variable declaration
	stockValue := 83.3
	fmt.Println("Stock Value is ", stockValue)
	fmt.Printf("Stock Value Type is %T\n", stockValue)

	codeWord := "Never Give Up"
	fmt.Println("Code Word is: ", codeWord)
	fmt.Printf("codeWord Type is %T\n", codeWord)
	fmt.Println("----------------------------------")

	// Global variable example
	fmt.Println("Laptop Brand is ", laptop)
	mobile := "Samsung"
	fmt.Println("Mobile Brand is ", mobile)
	// Always use Short hand declaration inside function

	// TYPE conversion example
	var i int = 42
	i = 24
	var j float64 = float64(i)
	fmt.Printf("i: %T, j: %T\n", i, j)
	fmt.Printf("i: %d, j: %f\n", i, j)

	// CONSTANTS
	const PI = 3.14159
	fmt.Println("Constant PI is: ", PI)
}
