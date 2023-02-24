package erro

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func FindEndOfFunction(srcLines []string, lineWithFunctionStart int) int {
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

	// Error: end of file reached without finding matching closing brace
	return -1
}

func lastWriteToVar(s string, varName string) int {
	src := "package main\n" + s
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "example.go", src, 0)
	if err != nil {
		return -1
	}
	lastLine := -1
	ast.Inspect(node, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.FuncDecl:
			f := n.(*ast.FuncDecl)
			for _, field := range f.Type.Params.List {
				for _, n := range field.Names {
					if n.Name == varName {
						pos := fset.Position(f.Pos())
						lastLine = pos.Line - 1
					}
				}
			}
		case *ast.AssignStmt:
			for i, lhs := range n.(*ast.AssignStmt).Lhs {
				ident, ok := lhs.(*ast.Ident)
				if !ok {
					continue
				}
				if ident.Name == varName {
					pos := fset.Position(n.(*ast.AssignStmt).Lhs[i].Pos())
					lastLine = pos.Line - 1
				}
			}
		}
		return true
	})
	return lastLine
}
