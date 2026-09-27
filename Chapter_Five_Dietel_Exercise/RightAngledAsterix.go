package main

import "fmt"

func main() {

	var number int
	fmt.Print("Enter number: ")
	fmt.Scan(&number)

	fmt.Print("---Rigght Angled Triangle (a) ---\n")
	for count := 1; count <= number; count++ {
		for iterator := 1; iterator <= count; iterator++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Print("--- Upside Down Right Angled Triangle (b) ---\n")

	for count := 1; count <= number; count++ {
		for iterator := 1; iterator <= number-count+1; iterator++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Print("--- inversed Upside Down Right Angled Triangle (c) ---\n")

	for count := 1; count <= number; count++ {
		for iterator := 1; iterator <= count; iterator++ {
			fmt.Print(" ")
		}

		for iterator := 1; iterator <= number-count+1; iterator++ {
			fmt.Print("*")
		}

		fmt.Println()
	}

	fmt.Print("--- inversed Upside Down Right Angled Triangle (c) ---\n")

	for count := 1; count <= number; count++ {
		for iterator := 1; iterator <= number-count+1; iterator++ {
			fmt.Print(" ")
		}

		for iterator := 1; iterator <= count+1; iterator++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}
