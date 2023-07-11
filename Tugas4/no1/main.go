package main

import "fmt"

type Car struct {
	Type   string
	FuelIn float64
}

func (c Car) EstimateDistance() float64 {
	const fuelEfficieny = 1.5
	distance := c.FuelIn * fuelEfficieny
	return distance
}

func main() {
	car := Car{
		Type:   "SUV",
		FuelIn: 100,
	}

	estimateDistance := car.EstimateDistance()
	fmt.Printf("Perkiraan Jarak Tempuh : %.2F", estimateDistance)
}
