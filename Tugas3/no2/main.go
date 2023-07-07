package main

import "fmt"

func Mapping(slice []string) map[string]int {
	res := make(map[string]int)

	for _, val := range slice {
		res[val]++
	}
	return res
}

func main() {
	fmt.Println(Mapping([]string{"abc", "abc", "acd"}))
}
