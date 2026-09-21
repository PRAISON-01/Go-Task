package main

import "fmt"

func main() {

	fmt.Println("N\tN2\tN3\tN4")

	var count int = 1

	for count = 1; count <= 5; count++ {

		number2 := count * count
		number3 := count * count * count
		number4 := count * count * count * count

		fmt.Printf("%d\t%d\t%d\t%d\n", count, number2, number3, number4)
	}
}
