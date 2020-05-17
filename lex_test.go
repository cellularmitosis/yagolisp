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

func TestLexEmptyList(t *testing.T) {
	mustCompileRegexes()
	source := []byte("()")
	tokens := mustLex(source)
	token := tokens[0]
	if token.TypeID != TOK_OPAREN {
		t.Errorf("ast.TypeID is %d but should be %d\n", token.TypeID, TOK_OPAREN)
	}
	token = tokens[1]
	if token.TypeID != TOK_CPAREN {
		t.Errorf("ast.TypeID is %d but should be %d\n", token.TypeID, TOK_CPAREN)
	}
}

func TestLexEmptyVector(t *testing.T) {
	mustCompileRegexes()
	source := []byte("[]")
	tokens := mustLex(source)
	token := tokens[0]
	if token.TypeID != TOK_OBRACK {
		t.Errorf("ast.TypeID is %d but should be %d\n", token.TypeID, TOK_OBRACK)
	}
	token = tokens[1]
	if token.TypeID != TOK_CBRACK {
		t.Errorf("ast.TypeID is %d but should be %d\n", token.TypeID, TOK_CBRACK)
	}
}

func TestLexEmptyMap(t *testing.T) {
	mustCompileRegexes()
	source := []byte("{}")
	tokens := mustLex(source)
	token := tokens[0]
	if token.TypeID != TOK_OBRACE {
		t.Errorf("ast.TypeID is %d but should be %d\n", token.TypeID, TOK_OBRACE)
	}
	token = tokens[1]
	if token.TypeID != TOK_CBRACE {
		t.Errorf("ast.TypeID is %d but should be %d\n", token.TypeID, TOK_CBRACE)
	}
}

func TestLexEmptySet(t *testing.T) {
	mustCompileRegexes()
	source := []byte("#{}")
	tokens := mustLex(source)
	token := tokens[0]
	if token.TypeID != TOK_OHASHBRACE {
		t.Errorf("ast.TypeID is %d but should be %d\n", token.TypeID, TOK_OHASHBRACE)
	}
	token = tokens[1]
	if token.TypeID != TOK_CBRACE {
		t.Errorf("ast.TypeID is %d but should be %d\n", token.TypeID, TOK_CBRACE)
	}
}
