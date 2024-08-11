package main

import "fmt"

func main() {
	// String Concatenation
	fmt.Println("Hello " + "World")

	fmt.Println("I am" + " Iron " + "Man")

	name := "Captain America"
	greeting := "Welcome back, " + name
	fmt.Println(greeting)

	message := fmt.Sprintf("Captain America is %d years Old", 200)
	fmt.Println(message)

	// Arthimetic Examples
	fmt.Println("8 + 8 = ", 8+8)
	fmt.Println("8 - 4 = ", 8-4)
	fmt.Println("8 * 8 = ", 8*8)
	fmt.Println("10 / 4 = ", 10/4)
	fmt.Println("20 % 4 = ", 20%4)
	fmt.Println("20 % 3 = ", 20%3)

	// Floating point Arthimetic
	fmt.Println("3.14159 * 2 = ", 3.14159*2)
	fmt.Println("Exponential: ", 1.5e3)
	fmt.Println("9.0 / 2.0 = ", 9.0/2.0)
	fmt.Printf("%.4f\n", 10.0/3.0)

	// Boolean operations
	fmt.Println("true && true: ", true && true)
	fmt.Println("true && false: ", true && false)
	fmt.Println("false || true: ", false || true)
	fmt.Println("false || false: ", false || false)
	fmt.Println("!false: ", !false)

	// Comparison
	x := 5  // declaring variables
	y := 10 // declaring variables
	fmt.Println("x < y: ", x < y)
	fmt.Println("x > y: ", x > y)

}
