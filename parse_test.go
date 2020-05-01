package main

import "testing"

func TestParseEmpty(t *testing.T) {
	mustCompileRegexes()
	program := []byte("")
	tokens := mustLex(program)
	ast := mustParse(tokens)
	if ast.TypeID != AST_PROGRAM {
		t.Errorf("ast.TypeID is %d but should be %d\n", ast.TypeID, AST_PROGRAM)
	}
	if len(ast.Subnodes) != 0 {
		t.Errorf("len(ast.Subnodes) is %d but should be %d\n", len(ast.Subnodes), 1)
	}
}

func TestParseSymbol(t *testing.T) {
	mustCompileRegexes()
	program := []byte("foo")
	tokens := mustLex(program)
	ast := mustParse(tokens)
	if ast.TypeID != AST_PROGRAM {
		t.Errorf("ast.TypeID is %d but should be %d\n", ast.TypeID, AST_PROGRAM)
	}
	if len(ast.Subnodes) != 1 {
		t.Errorf("len(ast.Subnodes) is %d but should be %d\n", len(ast.Subnodes), 1)
	}
	ast = ast.Subnodes[0]
	if ast.TypeID != AST_SYMBOL {
		t.Errorf("ast.TypeID is %d but should be %d\n", ast.TypeID, AST_SYMBOL)
	}
}

func TestParseString(t *testing.T) {
	mustCompileRegexes()
	program := []byte(`"foo"`)
	tokens := mustLex(program)
	ast := mustParse(tokens)
	if ast.TypeID != AST_PROGRAM {
		t.Errorf("ast.TypeID is %d but should be %d\n", ast.TypeID, AST_PROGRAM)
	}
	if len(ast.Subnodes) != 1 {
		t.Errorf("len(ast.Subnodes) is %d but should be %d\n", len(ast.Subnodes), 1)
	}
	ast = ast.Subnodes[0]
	if ast.TypeID != AST_STRING {
		t.Errorf("ast.TypeID is %d but should be %d\n", ast.TypeID, AST_STRING)
	}
}

func TestParseListEmpty(t *testing.T) {
	mustCompileRegexes()
	program := []byte("()")
	tokens := mustLex(program)
	ast := mustParse(tokens)
	if ast.TypeID != AST_PROGRAM {
		t.Errorf("ast.TypeID is %d but should be %d\n", ast.TypeID, AST_PROGRAM)
	}
	if len(ast.Subnodes) != 1 {
		t.Errorf("len(ast.Subnodes) is %d but should be %d\n", len(ast.Subnodes), 1)
	}
	ast = ast.Subnodes[0]
	if ast.TypeID != AST_LIST {
		t.Errorf("ast.TypeID is %d but should be %d\n", ast.TypeID, AST_LIST)
	}
	if len(ast.Subnodes) != 0 {
		t.Errorf("len(ast.Subnodes) is %d but should be %d\n", len(ast.Subnodes), 0)
	}
}

func TestParseListSymbol(t *testing.T) {
	mustCompileRegexes()
	program := []byte("(foo)")
	tokens := mustLex(program)
	ast := mustParse(tokens)
	if ast.TypeID != AST_PROGRAM {
		t.Errorf("ast.TypeID is %d but should be %d\n", ast.TypeID, AST_PROGRAM)
	}
	if len(ast.Subnodes) != 1 {
		t.Errorf("len(ast.Subnodes) is %d but should be %d\n", len(ast.Subnodes), 1)
	}
	ast = ast.Subnodes[0]
	if ast.TypeID != AST_LIST {
		t.Errorf("ast.TypeID is %d but should be %d\n", ast.TypeID, AST_LIST)
	}
	if len(ast.Subnodes) != 1 {
		t.Errorf("len(ast.Subnodes) is %d but should be %d\n", len(ast.Subnodes), 1)
	}
	ast = ast.Subnodes[0]
	if ast.TypeID != AST_SYMBOL {
		t.Errorf("ast.TypeID is %d but should be %d\n", ast.TypeID, AST_SYMBOL)
	}
}
