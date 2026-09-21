package main

import "fmt"

func main() {

	var number int

	fmt.Print("Enter a five digit integer >> ")

	fmt.Scanf("%d\n", &number)

	if number > 99999 || number < 10000 {
		fmt.Println("Number is not a five digit number")
		return
	}

	uNumber := number

	var reversedNumber = 0

	for number != 0 {

		digit := number % 10
		reversedNumber = (reversedNumber * 10) + digit
		number /= 10

	}

	if uNumber == reversedNumber {
		fmt.Print("Number is Palindrome")
	} else {
		fmt.Print("Number is 'NOT' Palindrome")
	}

}
