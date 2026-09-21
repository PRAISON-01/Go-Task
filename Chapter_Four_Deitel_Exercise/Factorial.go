package main

import "fmt"

func main() {
	var number int

	fmt.Print("Enter a integer >> ")
	fmt.Scanf("%d\n", &number)

	if number < 0 {
		fmt.Println("You no sabi maths ")
		return
	}

	factorial := 1

	for count := 1; count <= number; count++ {
		factorial *= count
	}

	fmt.Printf("%d! = %d\n", number, factorial)
}
