package main

import "fmt"

func main() {
	var targetNumber int

	fmt.Print("Enter the target number >> ")
	fmt.Scanf("%d\n", &targetNumber)

	runningSum := 0
	inputCount := 0

	for runningSum < targetNumber {
		var inputNumber int
		inputCount++

		fmt.Printf("Enter integer #%d to add to the sum >> ", inputCount)
		fmt.Scanf("%d\n", &inputNumber)

		runningSum += inputNumber

	}

	fmt.Println("--- Threshold Reached! ---")
	fmt.Printf("Final Sum: %d (Target was %d)\n", runningSum, targetNumber)
	fmt.Printf("It took %d total entries to reach the target.\n", inputCount)
}
