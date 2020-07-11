package ahapattern

import (
	"reflect"
)

func Match(e interface{}) *pattern {
	return &pattern{
		target:       reflect.ValueOf(e),
		typeOfTarget: reflect.TypeOf(e),
	}
}

type pattern struct {
	target       reflect.Value
	typeOfTarget reflect.Type
	// pattern should be something like
	//   func(i int) int {
	//   	return i
	//   }
	ofPatterns  []interface{}
	elsePattern interface{}
}

func (p *pattern) Of(t interface{}, f interface{}) {
	// NOTE: t should be `int`, `struct xxx`
	// TODO: check f via t
}

func (p *pattern) Else(f interface{}) interface{} {
	rf := reflect.ValueOf(f)
	// must be a function with one parameter and one output
	// TODO: maybe handle function with error output?
	mustNumber(rf.Type().NumIn(), 1)
	mustNumber(rf.Type().NumOut(), 1)
	// first parameter
	mustTrue(rf.Type().In(0).AssignableTo(p.typeOfTarget))
	return rf.Call([]reflect.Value{p.target})[0].Interface()
}

func mustNumber(actual int, expected int) {
	if actual != expected {
		panic("")
	}
}

func mustTrue(b bool) {
	if !b {
		panic("")
	}
}
