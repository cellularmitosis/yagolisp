package main

func eval(ast *ASTNode) LispValue {
	switch ast.TypeID {
	case AST_STRING:
		return string(ast.Bytes)
	case AST_PROGRAM:
		var value LispValue
		for _, subnode := range ast.Subnodes {
			value = eval(subnode)
		}
		return value
	default:
		panic("don't know how to eval that")
	}
}
