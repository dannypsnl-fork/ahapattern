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
	mustNumber(rf.Type().NumIn(), "should only one parameter", 1)
	mustNumber(rf.Type().NumOut(), "should only one output", 1)
	// first parameter
	mustTrue(rf.Type().In(0).AssignableTo(p.typeOfTarget),
		"parameter and target have different type")
	mustTrue(re.Type().AssignableTo(p.typeOfTarget),
		"pattern and target have different type")
	p.ofExpressions = append(p.ofExpressions, re)
	p.ofPatterns = append(p.ofPatterns, rf)
	return p
}

func (p *pattern) Else(f interface{}) interface{} {
	rf := reflect.ValueOf(f)
	// TODO: maybe handle function with error output?
	mustNumber(rf.Type().NumIn(), "should only one parameter", 1)
	mustNumber(rf.Type().NumOut(), "should only one output", 1)
	// first parameter
	mustTrue(rf.Type().In(0).AssignableTo(p.typeOfTarget),
		"parameter and target have different type")
	for i, ofPattern := range p.ofPatterns {
		if p.target.Interface() == p.ofExpressions[i].Interface() {
			return ofPattern.Call([]reflect.Value{p.target})[0].Interface()
		}
	}
	return rf.Call([]reflect.Value{p.target})[0].Interface()
}

func mustNumber(actual int, message string, expected int) {
	if actual != expected {
		panic(message)
	}
}

func mustTrue(b bool, message string) {
	if !b {
		panic(message)
	}
}
