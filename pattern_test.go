package ahapattern

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchSimpleInteger(t *testing.T) {
	a := assert.New(t)

	result := Match(1).
		Of(1, func(i int) int { return 2 }).
		Else(func(i int) int { return i })

	a.Equal(2, result)
}

func TestMatchSimpleString(t *testing.T) {
	a := assert.New(t)

	result := Match("abc").
		Of("abc", func(s string) int { return 2 }).
		Else(func(s string) string { return s })

	a.Equal(2, result)
}

func TestMatchStruct(t *testing.T) {
	a := assert.New(t)

	type Foo struct {
		A int
		B int
	}
	result := Match(Foo{A: 1, B: 2}).
		Of(Foo{A: 1, B: 2}, func(f Foo) int { return f.B }).
		Else(func(f Foo) int { return f.A })

	a.Equal(2, result)
}
