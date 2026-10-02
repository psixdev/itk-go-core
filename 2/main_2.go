package main

import (
	"fmt"
	"math/rand/v2"
)

var getRandomInt = rand.IntN

func createRandomSlice(length int, max int) []int {
	slice := make([]int, length)

	for i := 0; i < length; i++ {
		slice[i] = getRandomInt(max)
	}

	return slice
}

func sliceExample(slice []int) []int {
	var filtered []int

	for _, el := range slice {
		if el%2 == 0 {
			filtered = append(filtered, el)
		}
	}

	return filtered
}

func addElement(slice []int, n int) []int {
	result := make([]int, len(slice)+1)

	copy(result, slice)
	result[len(slice)] = n

	return result
}

func copySlice(slice []int) []int {
	result := make([]int, len(slice))

	copy(result, slice)

	return result
}

func removeElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return append([]int(nil), slice...)
	}

	result := make([]int, 0, len(slice)-1)

	result = append(result, slice[:index]...)
	result = append(result, slice[index+1:]...)

	return result
}

func main() {
	originalSlice := createRandomSlice(10, 100)
	fmt.Printf("%-20s %v\n", "original:", originalSlice)

	filteredSlice := sliceExample(originalSlice)
	fmt.Printf("%-20s %v\n", "filtered:", filteredSlice)

	newElSlice := addElement(originalSlice, 2)
	fmt.Printf("%-20s %v\n", "with new el:", newElSlice)

	copiedSlice := copySlice(originalSlice)
	copiedSlice[0] = -1
	fmt.Printf("%-20s %v\n", "copied:", copiedSlice)

	withoutElSlice := removeElement(originalSlice, 2)
	withoutElSlice = removeElement(withoutElSlice, -1)
	withoutElSlice = removeElement(withoutElSlice, 10)
	fmt.Printf("%-20s %v\n", "without el:", withoutElSlice)

	fmt.Printf("%-20s %v\n", "original:", originalSlice)
}
