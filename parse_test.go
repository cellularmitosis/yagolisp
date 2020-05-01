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

func TestParseInt(t *testing.T) {
	mustCompileRegexes()
	program := []byte("42")
	tokens := mustLex(program)
	ast := mustParse(tokens)
	if ast.TypeID != AST_PROGRAM {
		t.Errorf("ast.TypeID is %d but should be %d\n", ast.TypeID, AST_PROGRAM)
	}
	if len(ast.Subnodes) != 1 {
		t.Errorf("len(ast.Subnodes) is %d but should be %d\n", len(ast.Subnodes), 1)
	}
	ast = ast.Subnodes[0]
	if ast.TypeID != AST_INT {
		t.Errorf("ast.TypeID is %d but should be %d\n", ast.TypeID, AST_INT)
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
	program := []byte(`(foo "bar" 42)`)
	tokens := mustLex(program)
	prog := mustParse(tokens)
	if prog.TypeID != AST_PROGRAM {
		t.Errorf("ast.TypeID is %d but should be %d\n", prog.TypeID, AST_PROGRAM)
	}
	if len(prog.Subnodes) != 1 {
		t.Errorf("len(ast.Subnodes) is %d but should be %d\n", len(prog.Subnodes), 1)
	}
	list := prog.Subnodes[0]
	if list.TypeID != AST_LIST {
		t.Errorf("ast.TypeID is %d but should be %d\n", list.TypeID, AST_LIST)
	}
	if len(list.Subnodes) != 3 {
		t.Errorf("len(ast.Subnodes) is %d but should be %d\n", len(list.Subnodes), 3)
	}
	sym := list.Subnodes[0]
	if sym.TypeID != AST_SYMBOL {
		t.Errorf("ast.TypeID is %d but should be %d\n", sym.TypeID, AST_SYMBOL)
	}
	str := list.Subnodes[1]
	if str.TypeID != AST_STRING {
		t.Errorf("ast.TypeID is %d but should be %d\n", str.TypeID, AST_STRING)
	}
	i := list.Subnodes[2]
	if i.TypeID != AST_INT {
		t.Errorf("ast.TypeID is %d but should be %d\n", i.TypeID, AST_INT)
	}
}
