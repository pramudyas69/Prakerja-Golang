package main

import "fmt"

func GetMinMax(numbers ...*int) (min int, max int) {
	for i, val := range numbers {
		switch {
		case i == 0:
			max, min = *val, *val
		case *val > max:
			max = *val
		case *val < min:
			min = *val
		}
	}
	return
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scanln(&a1)
	fmt.Scanln(&a2)
	fmt.Scanln(&a3)
	fmt.Scanln(&a4)
	fmt.Scanln(&a5)
	fmt.Scanln(&a6)

	min, max = GetMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Printf("Nilai Min : %d\n", min)
	fmt.Printf("Nilai Max : %d", max)
}
