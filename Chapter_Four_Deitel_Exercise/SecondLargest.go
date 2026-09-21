package main

import "fmt"

func main() {

	var largestNumber int
	var secondLargest int

	for count := 1; count <= 10; count++ {

		fmt.Printf("----Count #%d----\n", count)

		var number int
		fmt.Print("Enter Number >> ")
		fmt.Scanf("%d\n", &number)

		if number > largestNumber {
			secondLargest = largestNumber
			largestNumber = number
		} else if number > secondLargest && number != largestNumber {
			secondLargest = number
		}

	}

	fmt.Printf("Largest number: %d\n", largestNumber)
	fmt.Printf("Second largest number: %d\n", secondLargest)
}
