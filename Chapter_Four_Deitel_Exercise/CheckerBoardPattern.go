package main

import "fmt"

func main() {

	for count := 1; count <= 8; count++ {
		if count%2 == 0 {

			fmt.Print(" ")

		}

		for counter := 1; counter <= 8; counter++ {
			fmt.Print("* ")
		}

		fmt.Println("")
	}
}
