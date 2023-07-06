package main

import (
	"fmt"
	luas2 "tugas2/luas"
	"tugas2/multiple7"
	"tugas2/prima"
)

func main() {
	number := 19
	base1 := 2.5
	base2 := 3.5
	height := 4.5

	isPrime := prima.IsPrime(number)
	multiple7 := multiple7.IsMultipleOfSeven(number)
	luas := luas2.CalculateTrapezoidArea(base1, base2, height)

	if isPrime {
		fmt.Printf("%d adalah bilangan prima!\n", number)
	} else {
		fmt.Printf("%d bukan bilangan prima!\n", number)
	}

	if multiple7 {
		fmt.Printf("%d adalah kelipatan 7!\n", number)
	} else {
		fmt.Printf("%d bukan kelipatan 7!\n", number)
	}

	fmt.Printf("Luas Dari Trapesium : %v\n", luas)
}
gi