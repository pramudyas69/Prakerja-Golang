package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	uniqueValue := make(map[string]bool)
	strArr := []string{}

	for _, val := range arrayA {
		strArr = append(strArr, val)
		uniqueValue[val] = true
	}

	for _, val := range arrayB {
		if !uniqueValue[val] {
			strArr = append(strArr, val)
			uniqueValue[val] = true
		}
	}
	return strArr
}

func main() {
	fmt.Println(ArrayMerge([]string{"Pandu", "Pramudya"}, []string{"Zenitsu", "Pramudya", "Tanjiro"}))
}
