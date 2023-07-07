package main

import (
	"fmt"
	"strconv"
	"strings"
)

func munculSekali(number string) []int {
	numbers := make(map[int]int)
	result := []int{}

	strArr := strings.Fields(number)

	for _, numStr := range strArr {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			continue
		}
		numbers[num]++
	}

	for num, count := range numbers {
		if count == 1 {
			result = append(result, num)
		}
	}
	return result
}

func main() {
	str := "1 2 3 1 2 4 5 6 7 8 4 5 6 7 8 10 10 3"
	res := munculSekali(str)
	fmt.Printf("Hasil : %v", res)
}
