package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMinMax(t *testing.T) {
	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 5, 8}
	min, max := GetMinMax(number)

	assert.Equal(t, 1, *min, "Minimum value should be 1")
	assert.Equal(t, 10, *max, "Maximum value should be 10")
}
