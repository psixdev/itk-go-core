package main

import "fmt"

func difference(s1, s2 []string) []string {
	var result []string

	hash := make(map[string]struct{}, len(s2))

	for _, value := range s2 {
		hash[value] = struct{}{}
	}

	for _, value := range s1 {
		if _, ok := hash[value]; !ok {
			result = append(result, value)
		}
	}

	return result
}

func main() {
	slice := difference(
		[]string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
		[]string{"banana", "date", "fig"},
	)

	fmt.Println(slice)
}
