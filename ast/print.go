package ast

import (
	"go/ast"
)

func Print(x interface{}) error {
	return ast.Print(nil, x)
}
