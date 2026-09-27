package main

import "testing"

func Test_ArrayOfNumbers_10IsTheLargestNumbers(t *testing.T) {
	var result, index int
	result, index = FindLargestNumber([]int{10, 3, 6, 7, 2, 1})
	var expected int
	expected = 10
	if result != expected {
		t.Error("Expected", expected, " Got ", result)
	}
	var expectedIndex int

	expectedIndex = 0

	if index != expectedIndex {
		t.Error("Expected Index", expectedIndex, " Got ", index)
	}

}

func Test_ArrayOfNumbers_99IsTheLargestNumbers(t *testing.T) {
	var result, index int
	result, index = FindLargestNumber([]int{22, 55, 11, 99, 0, 44})

	var expectedNumber int
	expectedNumber = 99

	var expectedIndex int
	expectedIndex = 3

	if result != expectedNumber {
		t.Error("Expected Number >> ", expectedNumber, "\nGot >>", result)
	}

	if index != expectedIndex {
		t.Error("Expected Index >> ", expectedNumber, "\nGot >>", result)
	}

}
