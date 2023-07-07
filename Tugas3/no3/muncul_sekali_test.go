package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMunculSekali(t *testing.T) {
	number := "1 2 3 2 1 4 5 4 6"

	res := munculSekali(number)

	expected := []int{3, 5, 6}
	assert.ElementsMatch(t, expected, res)
}
