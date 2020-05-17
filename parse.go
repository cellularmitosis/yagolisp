package main

import (
	"fmt"
	"math/bits"
	"os"
	"strconv"
)

// ASTNodeType is a structure describing a particular type of AST node.
type ASTNodeType struct {
	ID   uint
	Name string
}

// The AST node type identifiers.
const (
	AST_PROGRAM = iota

	AST_STRING
	AST_REAL
	AST_INT
	AST_KEYWORD
	AST_NIL
	AST_BOOL
	AST_SYMBOL

	AST_LIST
	AST_VECTOR
	AST_MAP
	AST_SET
)

var astNodeTypes = []ASTNodeType{
	{ID: AST_PROGRAM, Name: "AST_PROGRAM"},

	{ID: AST_STRING, Name: "AST_STRING"},
	{ID: AST_REAL, Name: "AST_REAL"},
	{ID: AST_INT, Name: "AST_INT"},
	{ID: AST_KEYWORD, Name: "AST_KEYWORD"},
	{ID: AST_NIL, Name: "AST_NIL"},
	{ID: AST_BOOL, Name: "AST_BOOL"},
	{ID: AST_SYMBOL, Name: "AST_SYMBOL"},

	{ID: AST_LIST, Name: "AST_LIST"},
	{ID: AST_VECTOR, Name: "AST_VECTOR"},
	{ID: AST_MAP, Name: "AST_MAP"},
	{ID: AST_SET, Name: "AST_SET"},
}

// ASTNode is a structure describing a node in an abstract syntax tree.
type ASTNode struct {
	TypeID   uint
	Bytes    []byte
	Value    interface{}
	Subnodes []*ASTNode
}

// grammar (discarding whitespace):
// AST_PROGRAM = expr+
// expr = atom | container
// atom = AST_STRING | AST_REAL | AST_INT | AST_KEYWORD | AST_NIL | AST_BOOL | AST_SYMBOL
// container = AST_LIST | AST_VECTOR | AST_MAP | AST_SET
// AST_LIST = TOK_OPAREN expr* TOK_CPAREN
// AST_VECTOR = TOK_OBRACK expr* TOK_CBRACK
// AST_MAP = TOK_OBRACE kv_pair* TOK_CBRACE
// kv_pair = expr expr
// AST_SET = TOK_OHASHBRACE expr* TOK_CBRACE

// Tries to parse the next token as a string.
// Returns an AST node and the index of the next token.
func parseString(tokens []Token, index uint) (*ASTNode, uint) {
	token := tokens[index]
	if token.TypeID != TOK_STRING {
		return nil, index
	}

	ast := ASTNode{
		TypeID:   AST_STRING,
		Bytes:    token.Bytes,
		Value:    string(token.Bytes),
		Subnodes: nil,
	}
	return &ast, index + 1
}

// Tries to parse the next token as a real.
// Returns an AST node and the index of the next token.
func parseReal(tokens []Token, index uint) (*ASTNode, uint) {
	token := tokens[index]
	if token.TypeID != TOK_REAL {
		return nil, index
	}

	bitSize := 64
	value, err := strconv.ParseFloat(string(token.Bytes), bitSize)
	if err != nil {
		return nil, index
	}

	ast := ASTNode{
		TypeID:   AST_REAL,
		Bytes:    token.Bytes,
		Value:    value,
		Subnodes: nil,
	}
	return &ast, index + 1
}

// Tries to parse the next token as an int.
// Returns an AST node and the index of the next token.
func parseInt(tokens []Token, index uint) (*ASTNode, uint) {
	token := tokens[index]
	if token.TypeID != TOK_INT {
		return nil, index
	}

	base := 10
	bitSize := bits.UintSize
	value, err := strconv.ParseInt(string(token.Bytes), base, bitSize)
	if err != nil {
		return nil, index
	}

	ast := ASTNode{
		TypeID:   AST_INT,
		Bytes:    token.Bytes,
		Value:    int(value),
		Subnodes: nil,
	}
	return &ast, index + 1
}

// Tries to parse the next token as a keyword.
// Returns an AST node and the index of the next token.
func parseKeyword(tokens []Token, index uint) (*ASTNode, uint) {
	token := tokens[index]
	if token.TypeID != TOK_KEYWORD {
		return nil, index
	}

	ast := ASTNode{
		TypeID:   AST_KEYWORD,
		Bytes:    token.Bytes,
		Subnodes: nil,
	}
	return &ast, index + 1
}

// LispNil is a type representing the Lisp notion of nil.
type lispNil struct{}

// LispNilValue is the global nil value singleton.
var lispNilValue = lispNil{}

// Tries to parse the next token as nil.
// Returns an AST node and the index of the next token.
func parseNil(tokens []Token, index uint) (*ASTNode, uint) {
	token := tokens[index]
	if token.TypeID != TOK_SYMBOL || string(token.Bytes) != "nil" {
		return nil, index
	}

	ast := ASTNode{
		TypeID:   AST_NIL,
		Bytes:    token.Bytes,
		Value:    lispNilValue,
		Subnodes: nil,
	}
	return &ast, index + 1
}

// Tries to parse the next token as a boolean.
// Returns an AST node and the index of the next token.
func parseBool(tokens []Token, index uint) (*ASTNode, uint) {
	token := tokens[index]
	if token.TypeID != TOK_SYMBOL {
		return nil, index
	}

	if string(token.Bytes) == "true" {
		ast := ASTNode{
			TypeID:   AST_BOOL,
			Bytes:    token.Bytes,
			Value:    true,
			Subnodes: nil,
		}
		return &ast, index + 1
	}

	if string(token.Bytes) == "false" {
		ast := ASTNode{
			TypeID:   AST_BOOL,
			Bytes:    token.Bytes,
			Value:    false,
			Subnodes: nil,
		}
		return &ast, index + 1
	}

	return nil, index
}

// Tries to parse the next token as a symbol.
// Returns an AST node and the index of the next token.
func parseSymbol(tokens []Token, index uint) (*ASTNode, uint) {
	token := tokens[index]
	if token.TypeID != TOK_SYMBOL {
		return nil, index
	}

	ast := ASTNode{
		TypeID:   AST_SYMBOL,
		Bytes:    token.Bytes,
		Subnodes: nil,
	}
	return &ast, index + 1
}

// Tries to parse the next tokens as a list.
// Returns an AST node and the index of the next token.
func parseList(tokens []Token, index uint) (*ASTNode, uint) {
	token := tokens[index]

	// a list must start with an OPAREN
	if token.TypeID != TOK_OPAREN {
		return nil, index
	}
	index2 := index + 1

	// loop over the contents of the list
	buff := make([]*ASTNode, 0, 8)
	subnodeCount := 0
	for index2 < uint(len(tokens)) {
		var ast *ASTNode
		token = tokens[index2]
		// stop the loop if we hit a CPAREN
		if token.TypeID == TOK_CPAREN {
			break
		}
		ast, index2 = parseExpr(tokens, index2)
		if ast == nil {
			return nil, index
		}
		buff = append(buff, ast)
		subnodeCount++
	}
	subnodes := make([]*ASTNode, subnodeCount)
	copy(subnodes, buff)

	// a list must end with a CPAREN
	if index2 == uint(len(tokens)) || tokens[index2].TypeID != TOK_CPAREN {
		return nil, index
	}
	index2++

	ast := ASTNode{
		TypeID:   AST_LIST,
		Bytes:    nil,
		Subnodes: subnodes,
	}
	return &ast, index2
}

// Tries to parse the next tokens as a vector.
// Returns an AST node and the index of the next token.
func parseVector(tokens []Token, index uint) (*ASTNode, uint) {
	token := tokens[index]

	// a vector must start with an OBRACK
	if token.TypeID != TOK_OBRACK {
		return nil, index
	}
	index2 := index + 1

	// loop over the contents of the list
	buff := make([]*ASTNode, 0, 8)
	subnodeCount := 0
	for index2 < uint(len(tokens)) {
		var ast *ASTNode
		token = tokens[index2]
		// stop the loop if we hit a CBRACK
		if token.TypeID == TOK_CBRACK {
			break
		}
		ast, index2 = parseExpr(tokens, index2)
		if ast == nil {
			return nil, index
		}
		buff = append(buff, ast)
		subnodeCount++
	}
	subnodes := make([]*ASTNode, subnodeCount)
	copy(subnodes, buff)

	// a list must end with a CBRACK
	if index2 == uint(len(tokens)) || tokens[index2].TypeID != TOK_CBRACK {
		return nil, index
	}
	index2++

	ast := ASTNode{
		TypeID:   AST_VECTOR,
		Bytes:    nil,
		Subnodes: subnodes,
	}
	return &ast, index2
}

// Tries to parse the next token as an atom.
// Returns an AST node and the index of the next token.
func parseAtom(tokens []Token, index uint) (*ASTNode, uint) {
	ast, index2 := parseString(tokens, index)
	if ast != nil {
		return ast, index2
	}
	ast, index2 = parseReal(tokens, index)
	if ast != nil {
		return ast, index2
	}
	ast, index2 = parseInt(tokens, index)
	if ast != nil {
		return ast, index2
	}
	ast, index2 = parseKeyword(tokens, index)
	if ast != nil {
		return ast, index2
	}
	ast, index2 = parseNil(tokens, index)
	if ast != nil {
		return ast, index2
	}
	ast, index2 = parseBool(tokens, index)
	if ast != nil {
		return ast, index2
	}
	ast, index2 = parseSymbol(tokens, index)
	if ast != nil {
		return ast, index2
	}

	return nil, index
}

// Tries to parse the next tokens as an expression.
// Returns an AST node and the index of the next token.
func parseExpr(tokens []Token, index uint) (*ASTNode, uint) {
	ast, index2 := parseAtom(tokens, index)
	if ast != nil {
		return ast, index2
	}
	ast, index2 = parseList(tokens, index)
	if ast != nil {
		return ast, index2
	}
	ast, index2 = parseVector(tokens, index)
	if ast != nil {
		return ast, index2
	}

	return nil, index
}

// Tries to parse a program.
// Returns an AST node and the index of the next token.
func mustParseProgram(tokens []Token, index uint) (*ASTNode, uint) {
	subnodes := make([]*ASTNode, 0)
	index2 := index
	for index2 < uint(len(tokens)) {
		var ast *ASTNode
		ast, index2 = parseExpr(tokens, index2)
		if ast == nil {
			fmt.Fprintf(
				os.Stderr,
				"Error: failed to parse program at token index %d\n",
				index2,
			)
			os.Exit(EXITCODE_PARSE)
		}
		subnodes = append(subnodes, ast)
	}

	ast := ASTNode{
		TypeID:   AST_PROGRAM,
		Bytes:    nil,
		Subnodes: subnodes,
	}
	return &ast, index2
}

// Parses the loaded file's tokens into an AST.
// Aborts on failure.
func mustParse(tokens []Token) *ASTNode {
	ast, _ := mustParseProgram(tokens, 0)
	return ast
}
