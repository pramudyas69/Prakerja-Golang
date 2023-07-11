package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCar_EstimateDistance(t *testing.T) {
	car := Car{
		Type:   "SUV",
		FuelIn: 100,
	}

	expectedDistance := 150.0
	actualDistance := car.EstimateDistance()

	assert.Equal(t, expectedDistance, actualDistance, "Perkiraan jarak yang ditempuh tidak sesuai")
}
