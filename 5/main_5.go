package main

import "fmt"

func intersection(s1, s2 []int) (bool, []int) {
	var result []int

	hash := make(map[int]struct{}, len(s2))

	for _, value := range s2 {
		hash[value] = struct{}{}
	}

	for _, value := range s1 {
		if _, ok := hash[value]; ok {
			result = append(result, value)
		}
	}

	return len(result) > 0, result
}

func main() {
	exists, slice := intersection(
		[]int{65, 3, 58, 678, 64},
		[]int{64, 2, 3, 43},
	)

	fmt.Println(exists, slice)
}
