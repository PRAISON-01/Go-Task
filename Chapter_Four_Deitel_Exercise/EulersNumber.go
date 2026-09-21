package main

import "fmt"

func main() {
	var number int

	fmt.Print("Enter number of terms  >> ")
	fmt.Scanf("%d\n", &number)

	if number < 0 {
		fmt.Println("You no sabi maths ")
		return
	}

	factorial := 1.0
	eulersNumber := 1.0

	for count := 1; count <= number; count++ {
		factorial *= float64(count)
		eulersNumber += 1.0 / factorial
	}

	fmt.Printf("Eulers Number >> %.5f", eulersNumber)
}
