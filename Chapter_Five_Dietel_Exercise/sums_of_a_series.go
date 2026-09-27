package main

import "fmt"

func main() {
	totalSum := 0
	for count := 1; count <= 100; count++ {
		totalSum += count
		fmt.Printf("%d. total sum = %d\n", count, totalSum)
	}
}
