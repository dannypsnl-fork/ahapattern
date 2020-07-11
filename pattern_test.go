package ahapattern

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatch(t *testing.T) {
	a := assert.New(t)

	result := Match(1).Else(func(i int) int { return i })

	a.Equal(1, result)
}
