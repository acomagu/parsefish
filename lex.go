package main

import (
	"text/scanner"
)

type Lexer struct {
	scanner.Scanner
	result Statement
}

func (l *Lexer) Lex(lval *yySymType) int {
	t := l.Scan();
	if int(t) == scanner.EOF {
		return scanner.EOF
	}
	if int(t) == scanner.Ident {
		lval.ident = Ident{literal: l.TokenText()}
		return IDENT
	}
	return scanner.EOF
}

func (l *Lexer) Error(e string) {
	panic(e)
}
