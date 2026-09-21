package main

import "fmt"

func main() {
	var unitsSold int
	var personThatSoldLargest int = 0
	var largestUnit int = 0
	var count int

	for count = 1; count <= 10; count++ {

		fmt.Printf("----Salesman #%d----\n", count)
		fmt.Print("Enter unitsSold >> ")
		fmt.Scanf("%d\n", &unitsSold)

		if unitsSold > largestUnit {
			largestUnit = unitsSold
			personThatSoldLargest = count
		}

	}

	fmt.Printf("Salesman #%d With %d units sold wins!\n", personThatSoldLargest, largestUnit)

}
