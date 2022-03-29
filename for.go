package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
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
