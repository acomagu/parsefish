package main

import (
	"fmt"
	"text/scanner"
)

var keywords = map[int]string{
	IF: "if",
	ELSE: "else",
	BEGIN: "begin",
	END: "end",
}

type Scanner struct {
	scanner.Scanner
}

func (s *Scanner) Scan() (int, string) {
	s.skipBrank()
	c := s.Peek()
	switch {
	case s.isIdentChar(c):
		lit := s.scanIdent()
		if tok, ok := s.tokenOf(lit); ok {
			return tok, ""
		}
		return IDENT, lit
	}

	s.Next()
	s.skipBrank()
	switch c {
	case ';', '\n':
		return int(c), string(c)
	}
	return scanner.EOF, ""
}

func (s *Scanner) isNumber(c rune) bool {
	return '0' <= c && c <= '9'
}

func (s *Scanner) isIdentChar(c rune) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_' || c == '-' || c == '/' || c == '.'
}

func (s *Scanner) isBrank(c rune) bool {
	return c == ' ' || c == '\t'
}

func (s *Scanner) scanIdent() string {
	var ret []rune
	for s.isIdentChar(s.Peek()) || s.isNumber(s.Peek()) {
		ret = append(ret, s.Peek())
		s.Next()
	}
	return string(ret)
}

func (s *Scanner) skipBrank() {
	for s.isBrank(s.Peek()) {
		s.Next()
	}
}

func (*Scanner) tokenOf(lit string) (int, bool) {
	for tok, l := range keywords {
		if l == lit {
			return tok, true
		}
	}
	return 0, false
}

func newScanner() Scanner {
	return Scanner{}
}

type Lexer struct {
	s      Scanner
	result []Stmt
}

func (l *Lexer) Lex(lval *yySymType) int {
	tok, lit := l.s.Scan()
	switch tok {
	case IDENT:
		lval.ident = Ident{Name: lit}
		fmt.Printf("%s: IDENT\n", lit)
	}
	return tok
}

func (l *Lexer) Error(e string) {
	fmt.Printf("%#+v\n", yyErrorMessages)
	panic(e)
}
