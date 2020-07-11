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
	ofExpressions []reflect.Value
	ofPatterns    []reflect.Value
}

func (p *pattern) Of(e interface{}, f interface{}) *pattern {
	rf := reflect.ValueOf(f)
	re := reflect.ValueOf(e)
	mustNumber(rf.Type().NumIn(), 1)
	mustNumber(rf.Type().NumOut(), 1)
	// first parameter
	mustTrue(rf.Type().In(0).AssignableTo(p.typeOfTarget))
	mustTrue(re.Type().AssignableTo(p.typeOfTarget))
	p.ofExpressions = append(p.ofExpressions, re)
	p.ofPatterns = append(p.ofPatterns, rf)
	return p
}

func (p *pattern) Else(f interface{}) interface{} {
	rf := reflect.ValueOf(f)
	// TODO: maybe handle function with error output?
	mustNumber(rf.Type().NumIn(), 1)
	mustNumber(rf.Type().NumOut(), 1)
	// first parameter
	mustTrue(rf.Type().In(0).AssignableTo(p.typeOfTarget))
	for i, ofPattern := range p.ofPatterns {
		if p.target.Interface() == p.ofExpressions[i].Interface() {
			return ofPattern.Call([]reflect.Value{p.target})[0].Interface()
		}
	}
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
