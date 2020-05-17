package main

import "testing"

func checkTokTypeID(token Token, expectedTypeID uint, t *testing.T) {
	if token.TypeID != expectedTypeID {
		t.Errorf("token.TypeID is %d but should be %d\n", token.TypeID, expectedTypeID)
	}
}

func TestLexString(t *testing.T) {
	mustCompileRegexes()
	source := []byte(`"Hello, world!"`)
	tokens := mustLex(source)
	token := tokens[0]
	checkTokTypeID(token, TOK_STRING, t)
}

func TestLexInt(t *testing.T) {
	mustCompileRegexes()
	source := []byte("42")
	tokens := mustLex(source)
	token := tokens[0]
	checkTokTypeID(token, TOK_INT, t)
}

func TestLexNegInt(t *testing.T) {
	mustCompileRegexes()
	source := []byte("-42")
	tokens := mustLex(source)
	token := tokens[0]
	checkTokTypeID(token, TOK_INT, t)
}

func TestLexReal(t *testing.T) {
	mustCompileRegexes()
	source := []byte("3.14159")
	tokens := mustLex(source)
	token := tokens[0]
	checkTokTypeID(token, TOK_REAL, t)
}

func TestLexNegReal(t *testing.T) {
	mustCompileRegexes()
	source := []byte("-3.14159")
	tokens := mustLex(source)
	token := tokens[0]
	checkTokTypeID(token, TOK_REAL, t)
}

func TestLexSymbol(t *testing.T) {
	mustCompileRegexes()
	source := []byte("foo")
	tokens := mustLex(source)
	token := tokens[0]
	checkTokTypeID(token, TOK_SYMBOL, t)
}

func TestLexKeyword(t *testing.T) {
	mustCompileRegexes()
	source := []byte(":foo")
	tokens := mustLex(source)
	token := tokens[0]
	checkTokTypeID(token, TOK_KEYWORD, t)
}

func TestLexEmptyList(t *testing.T) {
	mustCompileRegexes()
	source := []byte("()")
	tokens := mustLex(source)
	token := tokens[0]
	checkTokTypeID(token, TOK_OPAREN, t)
	token = tokens[1]
	checkTokTypeID(token, TOK_CPAREN, t)
}

func TestLexEmptyVector(t *testing.T) {
	mustCompileRegexes()
	source := []byte("[]")
	tokens := mustLex(source)
	token := tokens[0]
	checkTokTypeID(token, TOK_OBRACK, t)
	token = tokens[1]
	checkTokTypeID(token, TOK_CBRACK, t)
}

func TestLexEmptyMap(t *testing.T) {
	mustCompileRegexes()
	source := []byte("{}")
	tokens := mustLex(source)
	token := tokens[0]
	checkTokTypeID(token, TOK_OBRACE, t)
	token = tokens[1]
	checkTokTypeID(token, TOK_CBRACE, t)
}

func TestLexEmptySet(t *testing.T) {
	mustCompileRegexes()
	source := []byte("#{}")
	tokens := mustLex(source)
	token := tokens[0]
	checkTokTypeID(token, TOK_OHASHBRACE, t)
	token = tokens[1]
	checkTokTypeID(token, TOK_CBRACE, t)
}
