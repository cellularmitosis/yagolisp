package main

import "testing"

func TestEvalString(t *testing.T) {
	mustCompileRegexes()
	program := []byte(`"Hello, world!"`)
	tokens := mustLex(program)
	ast := mustParse(tokens)
	value := eval(ast)
	switch typ := value.(type) {
	case string:
		break
	default:
		t.Errorf("value.(type) is %d but should be string\n", typ)
	}
}

func TestEvalInt(t *testing.T) {
	mustCompileRegexes()
	program := []byte(`42`)
	tokens := mustLex(program)
	ast := mustParse(tokens)
	value := eval(ast)
	switch typ := value.(type) {
	case int:
		break
	default:
		t.Errorf("value.(type) is %d but should be int\n", typ)
	}
}

func TestEvalReal(t *testing.T) {
	mustCompileRegexes()
	program := []byte(`3.14159`)
	tokens := mustLex(program)
	ast := mustParse(tokens)
	value := eval(ast)
	switch typ := value.(type) {
	case float64:
		break
	default:
		t.Errorf("value.(type) is %d but should be float64\n", typ)
	}
}
