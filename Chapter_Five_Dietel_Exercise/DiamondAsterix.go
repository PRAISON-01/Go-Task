package main

import "fmt"

func main() {

	var number int
	fmt.Print("Enter your number: ")
	fmt.Scan(&number)

	for count := 1; count <= number; count++ {
		for counter := count; counter < number; counter++ {
			fmt.Print(" ")
		}

		for counter := 1; counter < count; counter++ {
			fmt.Print(" *")
		}

		fmt.Println()

	}
}
