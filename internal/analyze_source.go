package internal

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func GoFindEndOfFunction(srcLines []string, lineWithFunctionStart int) int {
	if lineWithFunctionStart >= 0 {
		stack := make([]rune, 0)
		for i := lineWithFunctionStart; i < len(srcLines); i++ {
			line := srcLines[i]
			for _, char := range line {
				switch char {
				case '{':
					stack = append(stack, char)
				case '}':
					if len(stack) == 0 {
						// Error: closing brace without matching opening brace
						return -1
					}
					stack = stack[:len(stack)-1]
				}
			}
			if len(stack) == 0 {
				// Found the end of the function block
				return i
			}
		}
	}
	// Error: end of file reached without finding matching closing brace
	return -1
}

func GoLastWriteToVar(s string, varName string) int {
	src := "package main\n" + s
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "example.go", src, 0)
	if err != nil {
		return -1
	}
	lastLine := -1
	ast.Inspect(node, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.FuncDecl:
			for _, field := range n.Type.Params.List {
				for _, n := range field.Names {
					if n.Name == varName {
						pos := fset.Position(n.Pos())
						lastLine = pos.Line - 1
					}
				}
			}
		case *ast.AssignStmt:
			for i, lhs := range n.Lhs {
				ident, ok := lhs.(*ast.Ident)
				if !ok {
					continue
				}
				if ident.Name == varName {
					pos := fset.Position(n.Lhs[i].Pos())
					lastLine = pos.Line - 1
				}
			}
		}
		return true
	})
	return lastLine
}
