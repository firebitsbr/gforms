package gforms

import (
	"reflect"
)

type Data map[string]*V

type V struct {
	RawValues []string
	// not ptr-value
	Value  interface{}
	IsNill bool
	// value's kind
	Kind reflect.Kind
}

func newV(values []string, kind reflect.Kind) *V {
	v := new(V)
	v.RawValues = values
	v.Kind = kind
	v.IsNill = true
	return v
}

func nilV() *V {
	v := new(V)
	v.IsNill = true
	return v
}