package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := 1 + 2
	assert.Equal(t, 3, sum)
}
