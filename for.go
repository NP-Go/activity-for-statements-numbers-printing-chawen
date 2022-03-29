package main

import "fmt"

func main() {
	var input1, input2 int
	fmt.Println("Enter first scan")
	fmt.Scanln(&input1)
	fmt.Println("Enter second scan")
	fmt.Scanln(&input2)
	for i := input1; i <= input2; i++ {
		// fmt.Println(i)
		if i%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}

		fmt.Println(i)
	}

	fmt.Println("Last number")
}
