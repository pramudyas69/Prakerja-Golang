package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapping(t *testing.T) {
	slice := []string{"abc", "abc", "acd", "adb", "cda", "adb", "abc"}

	result := Mapping(slice)

	expected := map[string]int{"abc": 3, "acd": 1, "adb": 2, "cda": 1}
	assert.Equal(t, expected, result)
}
