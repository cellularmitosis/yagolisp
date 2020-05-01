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
