package main

import (
	"math/rand"
	"testing"
	"time"
)

var benchTmp interface{}

// ---

const (
	TYPE_INT    = 1
	TYPE_BOOL   = 2
	TYPE_STRING = 3
)

type IntTypedValue struct {
	Type  int
	Value interface{}
}

func makeRandomIntTypedValue() IntTypedValue {
	t := rand.Intn(3) + 1
	switch t {
	case TYPE_INT:
		return IntTypedValue{
			Type:  TYPE_INT,
			Value: 42,
		}
	case TYPE_BOOL:
		return IntTypedValue{
			Type:  TYPE_BOOL,
			Value: true,
		}
	case TYPE_STRING:
		return IntTypedValue{
			Type:  TYPE_STRING,
			Value: "hello",
		}
	default:
		panic("makeRandomIntTypedValue hit default case")
	}
}

func BenchmarkTypeDispatchInt(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	value := makeRandomIntTypedValue()
	for i := 0; i < b.N; i++ {
		switch value.Type {
		case TYPE_INT:
			benchTmp = value.Value.(int)
		case TYPE_STRING:
			benchTmp = value.Value.(string)
		case TYPE_BOOL:
			benchTmp = value.Value.(bool)
		}
	}
}

// ---

func makeRandomValue() interface{} {
	t := rand.Intn(3) + 1
	switch t {
	case TYPE_INT:
		return int(42)
	case TYPE_BOOL:
		return true
	case TYPE_STRING:
		return "hello"
	default:
		panic("makeRandomValue hit default case")
	}
}

func BenchmarkTypeDispatchSwitch(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	value := makeRandomValue()
	for i := 0; i < b.N; i++ {
		switch value.(type) {
		case int:
			benchTmp = value.(int) + 1
		case bool:
			benchTmp = value.(bool)
		case string:
			benchTmp = value.(string)
		}
	}
}
