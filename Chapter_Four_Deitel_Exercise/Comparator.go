package main

import "fmt"

func comparator(numberOne int, numberTwo int) int {

	if numberOne > numberTwo {
		return 1
	}

	if numberOne < numberTwo {
		return -1
	}
	if numberOne == numberTwo {

		return 1
	}

	return 0
}

func main() {

	var numberOne int
	fmt.Print("Enter First Number >> ")
	fmt.Scanf("%d\n", &numberOne)

	var numberTwo int
	fmt.Print("Enter Second Number >> ")
	fmt.Scanf("%d\n", &numberTwo)

	result := comparator(numberOne, numberTwo)

	fmt.Println(result)

}
