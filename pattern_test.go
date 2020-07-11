package ahapattern

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatch(t *testing.T) {
	a := assert.New(t)

	result := Match(1).
		Of(1, func(i int) int { return 2 }).
		Else(func(i int) int { return i })

	a.Equal(2, result)
}
