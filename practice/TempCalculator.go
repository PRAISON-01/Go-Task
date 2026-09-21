package main

import "fmt"

func main() {
	var input float64
	fmt.Print("Enter temperature in Fahrenheit >> ")
	fmt.Scanf("%f", &input)

	var celsius float64 = (input - 32) * (5.0 / 9.0)
	fmt.Printf("Temperature in celsius >> %f\n", celsius)
}
