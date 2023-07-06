package multiple7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsMultipleOfSeven(t *testing.T) {
	// Test case untuk bilangan kelipatan 7
	assert.True(t, IsMultipleOfSeven(7))
	assert.True(t, IsMultipleOfSeven(14))
	assert.True(t, IsMultipleOfSeven(21))
	assert.True(t, IsMultipleOfSeven(70))

	// Test case untuk bilangan bukan kelipatan 7
	assert.False(t, IsMultipleOfSeven(5))
	assert.False(t, IsMultipleOfSeven(10))
	assert.False(t, IsMultipleOfSeven(15))
	assert.False(t, IsMultipleOfSeven(23))
}
