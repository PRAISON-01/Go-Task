package main

import "fmt"

func main() {

	var totalMiles int = 0
	var totalGallons int = 0

	for {
		var milesDriven int
		fmt.Print("Enter miles driven (enter -1 to quit) >> ")
		fmt.Scanf("%d\n", &milesDriven)

		if milesDriven < 0 {
			break
		}

		var gallonUsed int
		fmt.Print("Enter Gallon Used >> ")
		fmt.Scanf("%d\n", &gallonUsed)

		totalMiles += milesDriven
		totalGallons += gallonUsed

	}

	if totalGallons == 0 {
		fmt.Println("No gallons used. Cannot calculate MPG.")
		return
	}

	milesPerGallon := float64(totalMiles) / float64(totalGallons)
	fmt.Printf("Miles Per Gallon %.2f\n", milesPerGallon)
}
