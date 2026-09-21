package main

import "fmt"

func main() {

	fmt.Println("<---The Tax Calculator--->")

	var totalTax float64

	var amountToDeduct float64 = 0.0

	var earnings float64
	fmt.Print("Enter earnings >> ")
	fmt.Scanf("%f\n", &earnings)

	if earnings <= 30000 {
		amountToDeduct = earnings * 0.15
		totalTax = earnings - amountToDeduct
	} else {
		amountToDeduct = earnings * 0.20
		totalTax = earnings - amountToDeduct

	}

	fmt.Printf("Pay you tax of >> %.2f\n", totalTax)

}
