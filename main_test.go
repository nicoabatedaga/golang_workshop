package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDivision(t *testing.T) {
	result, err := division(10, 2)
	assert.Nil(t, err)
	assert.Equal(t, 5.0, result)

	result, err = division(10, 0)
	assert.NotNil(t, err)
	assert.Equal(t, "cannot divide by zero", err.Error())
}
