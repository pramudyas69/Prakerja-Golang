package prima

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPrime(t *testing.T) {
	// Test case untuk bilangan prima
	primeNumbers := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	for _, number := range primeNumbers {
		assert.True(t, IsPrime(number), fmt.Sprintf("%d seharusnya bilangan prima", number))
	}

	// Test case untuk bukan bilangan prima
	nonPrimeNumbers := []int{0, 1, 4, 6, 8, 9, 10, 12, 14, 15}
	for _, number := range nonPrimeNumbers {
		assert.False(t, IsPrime(number), fmt.Sprintf("%d seharusnya bukan bilangan prima", number))
	}
}
