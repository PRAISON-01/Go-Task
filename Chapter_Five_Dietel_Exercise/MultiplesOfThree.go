package main

import "fmt"

func main() {
	index := 1

	for count := 1; count <= 30; count++ {
		if count%3 == 0 {
			fmt.Printf("%d.  %d is divisible by 3\n\n", index, count)
			index++
		}
	}
}
