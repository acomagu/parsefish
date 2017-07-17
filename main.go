package main

import (
	"fmt"
	"os"
)

func main() {
	yyDebug = 100
	yyErrorVerbose = true
	l := new(Lexer)

	l.s.Init(os.Stdin)
	yyParse(l)
	fmt.Printf("%#v\n", l.result)
}
