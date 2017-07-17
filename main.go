package main

import (
	"os"
	"go/ast"
)

func main() {
	yyDebug = 100
	yyErrorVerbose = true
	l := new(Lexer)

	l.s.Init(os.Stdin)
	yyParse(l)
	ast.Print(nil, l.result)
}
