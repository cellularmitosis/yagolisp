package main

import "testing"

func TestLexString(t *testing.T) {
	mustCompileRegexes()
	source := []byte(`"Hello, world!"`)
	tokens := mustLex(source)
	token := tokens[0]
	if token.TypeID != TOK_STRING {
		t.Errorf("ast.TypeID is %d but should be %d\n", token.TypeID, TOK_STRING)
	}
}

func TestLexInt(t *testing.T) {
	mustCompileRegexes()
	source := []byte("42")
	tokens := mustLex(source)
	token := tokens[0]
	if token.TypeID != TOK_INT {
		t.Errorf("ast.TypeID is %d but should be %d\n", token.TypeID, TOK_INT)
	}
}
