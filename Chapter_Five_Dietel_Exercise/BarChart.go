package main

import "fmt"

func main() {
	var input int

	var array [5]int
	for count := 1; count <= 5-1; count++ {
		fmt.Print("Enter a number >> ")
		fmt.Scan(&input)

		if input < 1 || input > 30 {
			println("Number must be in range within 1 and 30")
		}

		array[count] = input

	}

	fmt.Println(array)
	fmt.Println("\n\nBar Chart\n\n")

	for iterator := 0; iterator < 5; iterator++ {
		currentValue := array[iterator]

		fmt.Printf("%2d: ", currentValue)

		for counter := 0; counter < currentValue; counter++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}
