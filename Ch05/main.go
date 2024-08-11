package main

import "fmt"

func main() {
	/*
				// While loop
				i := 1
				for i <= 3 {
					fmt.Println(i)
					i = i + 1
				}

				// using all the components of for loop
				sum := 0
				for i := 1; i <= 5; i++ {
					fmt.Println(i)
					sum += i
				}
				fmt.Println("Sum is ", sum)
			for j := 10; j >= 0; j-- {
				fmt.Println(j)
			}
		// Range over integers
		for i := range 5 {
			fmt.Println("Range: ", i)
		}
	*/
	// Forever Loop
	/* -> Uncomment on your risk
	for {
		fmt.Println("Forever Dragon")
	}
	for i := range 1000 {
		fmt.Println("Range: ", i)
		if i == 500 {
			break
		}
	}
	*/
	// for {
	// 	fmt.Println("Forever loop")
	// 	break
	// }
	for n := range 20 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
