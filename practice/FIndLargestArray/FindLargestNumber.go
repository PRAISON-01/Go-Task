package main

func FindLargestNumber(array []int) (int, int) {
	var largest int = array[0]
	var largestIndex int = 0

	for index, number := range array {
		if number > largest {
			largest = number
			largestIndex = index

		}
	}

	return largest, largestIndex
}
