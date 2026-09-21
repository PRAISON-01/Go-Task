package main

import "fmt"

func main() {

	var totalSales float64 = 0.0

	for {

		var itemValue float64
		fmt.Print("Enter item value >> ")
		fmt.Scanf("%f\n", &itemValue)

		if itemValue == -1 {
			break
		}

		totalSales += itemValue

	}

	baseSalary := 200.0
	commission := totalSales * 0.09
	totalEarnings := baseSalary + commission

	fmt.Printf("\nLast week sales : $%.2f\n", totalSales)
	fmt.Printf("Salesman Profit >> $%.2f\n", totalEarnings)
}
