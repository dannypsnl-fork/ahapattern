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
	p.check(rf)
	re := reflect.ValueOf(e)
	mustTrue(re.Type().AssignableTo(p.typeOfTarget),
		"pattern and target have different type")
	p.ofExpressions = append(p.ofExpressions, re)
	p.ofPatterns = append(p.ofPatterns, rf)
	return p
}

func (p *pattern) Else(f interface{}) interface{} {
	rf := reflect.ValueOf(f)
	p.check(rf)
	for i, ofPattern := range p.ofPatterns {
		if p.target.Interface() == p.ofExpressions[i].Interface() {
			return ofPattern.Call([]reflect.Value{p.target})[0].Interface()
		}
	}
	return rf.Call([]reflect.Value{p.target})[0].Interface()
}

func (p *pattern) check(handler reflect.Value) {
	// TODO: maybe handle function with error output?
	mustNumber(handler.Type().NumIn(), "should only one parameter", 1)
	mustNumber(handler.Type().NumOut(), "should only one output", 1)
	// first parameter
	mustTrue(handler.Type().In(0).AssignableTo(p.typeOfTarget),
		"parameter and target have different type")
}

func mustNumber(actual int, message string, expected int) {
	mustTrue(actual != expected, message)
}

func mustTrue(b bool, message string) {
	if !b {
		panic(message)
	}
}
