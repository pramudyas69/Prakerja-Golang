package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayMerge(t *testing.T) {
	sliceA := []string{"apple", "banana", "orange"}
	sliceB := []string{"banana", "grape", "pineapple"}

	res := ArrayMerge(sliceA, sliceB)

	expected := []string{"apple", "banana", "orange", "grape", "pineapple"}
	assert.Equal(t, expected, res)
}
