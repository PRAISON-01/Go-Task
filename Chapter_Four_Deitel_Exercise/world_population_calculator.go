package main

import "fmt"

func main() {

	const initialPopulation float64 = 8316053434.0
	const growthRate float64 = 0.0084

	currentPopulation := initialPopulation
	doublingYear := -1

	fmt.Printf("Initial Baseline Population: %.0f\n", initialPopulation)
	fmt.Printf("Assumed Constant Growth Rate: %.2f%%\n\n", growthRate*100)

	fmt.Println("Year\tExpected Population\tNumerical Increase")
	fmt.Println("------------------------------------------------------------")

	for year := 1; year <= 75; year++ {

		numericalIncrease := currentPopulation * growthRate

		currentPopulation += numericalIncrease

		fmt.Printf("%d\t%.1f\t\t\t%.1f\n", year, currentPopulation, numericalIncrease)

		if doublingYear == -1 && currentPopulation >= (initialPopulation*2) {
			doublingYear = year
		}
	}

	if doublingYear != -1 {
		fmt.Printf("\n\nThe world population will double %d.\n", doublingYear)
	} else {
		fmt.Println("\n\n The population does not double ")

		fmt.Printf("Growth Rate >>  %.2f%% .\n", growthRate*100)
	}
}
