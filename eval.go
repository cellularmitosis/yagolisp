package main

func eval(ast *ASTNode) LispValue {
	switch ast.TypeID {
	case AST_NIL:
		return ast.Value
	case AST_BOOL:
		return ast.Value
	case AST_INT:
		return ast.Value
	case AST_REAL:
		return ast.Value
	case AST_STRING:
		return ast.Value
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
