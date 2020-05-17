package main

import "testing"

func checkASTTypeID(ast *ASTNode, expectedTypeID uint, t *testing.T) {
	if ast.TypeID != expectedTypeID {
		t.Errorf("ast.TypeID is %d but should be %d\n", ast.TypeID, expectedTypeID)
	}
}

func checkSubnodesLen(ast *ASTNode, expectedLen int, t *testing.T) {
	if len(ast.Subnodes) != expectedLen {
		t.Errorf("len(ast.Subnodes) is %d but should be %d\n", len(ast.Subnodes), expectedLen)
	}
}

func TestParseEmpty(t *testing.T) {
	mustCompileRegexes()
	program := []byte("")
	tokens := mustLex(program)
	ast := mustParse(tokens)
	checkASTTypeID(ast, AST_PROGRAM, t)
	checkSubnodesLen(ast, 0, t)
}

func TestParseSymbol(t *testing.T) {
	mustCompileRegexes()
	program := []byte("foo")
	tokens := mustLex(program)
	ast := mustParse(tokens)
	checkASTTypeID(ast, AST_PROGRAM, t)
	checkSubnodesLen(ast, 1, t)
	ast = ast.Subnodes[0]
	checkASTTypeID(ast, AST_SYMBOL, t)
}

func TestParseString(t *testing.T) {
	mustCompileRegexes()
	program := []byte(`"foo"`)
	tokens := mustLex(program)
	ast := mustParse(tokens)
	checkASTTypeID(ast, AST_PROGRAM, t)
	checkSubnodesLen(ast, 1, t)
	ast = ast.Subnodes[0]
	checkASTTypeID(ast, AST_STRING, t)
}

func TestParseInt(t *testing.T) {
	mustCompileRegexes()
	program := []byte("42")
	tokens := mustLex(program)
	ast := mustParse(tokens)
	checkASTTypeID(ast, AST_PROGRAM, t)
	checkSubnodesLen(ast, 1, t)
	ast = ast.Subnodes[0]
	checkASTTypeID(ast, AST_INT, t)
}

func TestParseNegInt(t *testing.T) {
	mustCompileRegexes()
	program := []byte("-42")
	tokens := mustLex(program)
	ast := mustParse(tokens)
	checkASTTypeID(ast, AST_PROGRAM, t)
	checkSubnodesLen(ast, 1, t)
	ast = ast.Subnodes[0]
	checkASTTypeID(ast, AST_INT, t)
}

func TestParseReal(t *testing.T) {
	mustCompileRegexes()
	program := []byte("3.14159")
	tokens := mustLex(program)
	ast := mustParse(tokens)
	checkASTTypeID(ast, AST_PROGRAM, t)
	checkSubnodesLen(ast, 1, t)
	ast = ast.Subnodes[0]
	checkASTTypeID(ast, AST_REAL, t)
}

func TestParseNegReal(t *testing.T) {
	mustCompileRegexes()
	program := []byte("-3.14159")
	tokens := mustLex(program)
	ast := mustParse(tokens)
	checkASTTypeID(ast, AST_PROGRAM, t)
	checkSubnodesLen(ast, 1, t)
	ast = ast.Subnodes[0]
	checkASTTypeID(ast, AST_REAL, t)
}

func TestParseListEmpty(t *testing.T) {
	mustCompileRegexes()
	program := []byte("()")
	tokens := mustLex(program)
	ast := mustParse(tokens)
	checkASTTypeID(ast, AST_PROGRAM, t)
	checkSubnodesLen(ast, 1, t)
	ast = ast.Subnodes[0]
	checkASTTypeID(ast, AST_LIST, t)
	checkSubnodesLen(ast, 0, t)
}

func TestParseListSymbolStringInt(t *testing.T) {
	mustCompileRegexes()
	program := []byte(`(foo "bar" 42)`)
	tokens := mustLex(program)
	prog := mustParse(tokens)
	checkASTTypeID(prog, AST_PROGRAM, t)
	checkSubnodesLen(prog, 1, t)
	list := prog.Subnodes[0]
	checkASTTypeID(list, AST_LIST, t)
	checkSubnodesLen(list, 3, t)
	sym := list.Subnodes[0]
	checkASTTypeID(sym, AST_SYMBOL, t)
	str := list.Subnodes[1]
	checkASTTypeID(str, AST_STRING, t)
	i := list.Subnodes[2]
	checkASTTypeID(i, AST_INT, t)
}

func TestParseVectorEmpty(t *testing.T) {
	mustCompileRegexes()
	program := []byte("[]")
	tokens := mustLex(program)
	ast := mustParse(tokens)
	checkASTTypeID(ast, AST_PROGRAM, t)
	checkSubnodesLen(ast, 1, t)
	ast = ast.Subnodes[0]
	checkASTTypeID(ast, AST_VECTOR, t)
	checkSubnodesLen(ast, 0, t)
}
