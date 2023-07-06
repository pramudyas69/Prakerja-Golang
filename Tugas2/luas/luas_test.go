package luas

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateTrapezoidArea(t *testing.T) {
	tests := []struct {
		base1, base2, height, expectedArea float64
	}{
		{3, 4, 2, 7},          // Test case 1
		{5, 7, 3, 18},         // Test case 2
		{0, 0, 0, 0},          // Test case 3
		{10, 20, 5, 75},       // Test case 4
		{8, 6, 4, 28},         // Test case 5
		{2, 3, -4, 0},         // Test case 6
		{2.5, 3.5, 4.5, 13.5}, // Test case 7 (floating-point numbers)
	}

	for _, test := range tests {
		area := CalculateTrapezoidArea(test.base1, test.base2, test.height)
		assert.Equal(t, test.expectedArea, area)
	}
}
