//go:generate goyacc -o yyparse.go yyparse.go.y

package yyparse

import (
	"io"

	"github.com/acomagu/parsefish/ast"
)

func Parse(input io.Reader) ast.Node {
	s := new(Scanner)
	s.Init(input)
	l := new(Lexer)
	l.s = s
	yyParse(l)
	return ast.Stmts(l.result)
}
