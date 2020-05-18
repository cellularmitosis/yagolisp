package main

import "fmt"

type LispValueStringer interface {
	LispValueString() string
}

type LispValue interface {
	LispValueStringer
}

// LispNil is a type representing the Lisp notion of nil.
type LispNil struct{}

func (LispNil) LispValueString() string {
	return "#<LispNil>"
}

type LispBool bool

func (b LispBool) LispValueString() string {
	return fmt.Sprintf("#<LispBool %t>", b)
}

type LispInt int

func (i LispInt) LispValueString() string {
	return fmt.Sprintf("#<LispInt %d>", i)
}

type LispReal float64

func (f LispReal) LispValueString() string {
	return fmt.Sprintf("#<LispReal %f>", f)
}

type LispString string

func (s LispString) LispValueString() string {
	return fmt.Sprintf(`#<LispString %s>`, s)
}

type LispSymbol string

func (sym LispSymbol) LispValueString() string {
	return fmt.Sprintf("#<LispSymbol %s>", sym)
}

type LispKeyword string

func (k LispKeyword) LispValueString() string {
	return fmt.Sprintf("#<LispKeyword %s>", k)
}

type Bindings map[LispSymbol]LispValue

type EnvStack struct {
	Bindings Bindings
	Next     *EnvStack
}

func (envStack *EnvStack) lookup(sym LispSymbol) LispValue {
	value, present := envStack.Bindings[sym]
	if present {
		return value
	}
	if envStack.Next != nil {
		return envStack.Next.lookup(sym)
	}
	return LispNil{}
}

func eval(ast *ASTNode, env *EnvStack) LispValue {
	switch ast.TypeID {
	case AST_NIL, AST_BOOL, AST_INT, AST_REAL, AST_STRING, AST_KEYWORD:
		return ast.Value
	case AST_SYMBOL:
		sym := ast.Value.(LispSymbol)
		return env.lookup(sym)
	case AST_PROGRAM:
		var value LispValue
		for _, subnode := range ast.Subnodes {
			value = eval(subnode, env)
		}
		return value
	default:
		panic("don't know how to eval that")
	}
}
