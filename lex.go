package main

import (
	"fmt"
	"text/scanner"
	"strconv"
)

var (
	singleQuotedSpecials = []rune{-1, '\\', '\''}
	doubleQuotedSpecials = []rune{-1, '\\', '"', '$'}
	specials = []rune{-1, '\t', '\n', '$', '?', '*', '~', '#', '(', ')', '{', '}', '[', ']', '<', '>', '^', '&', ';', '\'', '"', '\\', ' '}
)

type ExpectIdentError error
type ExpectVarNameError error

var keywords = map[int]string{
	IF: "if",
	ELSE: "else",
	BEGIN: "begin",
	END: "end",
	FUNCTION: "function",
}

type Scanner struct {
	scanner.Scanner
}

func (s *Scanner) scanSingleQuoted() (StrsExpr, error) {
	strs := StrsExpr{}
	s.Next()  // The starting '\''

	for {
		fmt.Printf("THE NEXT CHAR: %c\n", byte(s.Peek()))
		switch {
		case s.Peek() == scanner.EOF:
			return StrsExpr{}, fmt.Errorf("unexpected EOF")

		case s.isSingleQuotedIdentChar(s.Peek()):
			ident, err := s.scanSingleQuotedIdent()
			if err != nil {
				return StrsExpr{}, err
			}

			strs = append(strs, ident)

		case s.Peek() == '\'':
			s.Next()
			return strs, nil

		default:
			ident, err := s.scanSingleQuotedIdent()
			if err != nil {
				return StrsExpr{}, err
			}

			strs = append(strs, ident)
		}
	}
}

func (s *Scanner) scanDoubleQuoted() (StrsExpr, error) {
	strs := StrsExpr{}
	s.Next()  // The starting '"'

	for {
		fmt.Printf("THE NEXT CHAR: %c\n", byte(s.Peek()))
		switch {
		case s.Peek() == scanner.EOF:
			return StrsExpr{}, fmt.Errorf("unexpected EOF")

		case s.isDoubleQuotedIdentChar(s.Peek()):
			ident, err := s.scanDoubleQuotedIdent()
			if err != nil {
				return StrsExpr{}, err
			}

			strs = append(strs, ident)

		case s.Peek() == '$':
			vare, err := s.scanVar()
			if err != nil {
				return StrsExpr{}, err
			}

			strs = append(strs, vare)

		case s.Peek() == '"':
			s.Next()
			return strs, nil

		default:
			ident, err := s.scanDoubleQuotedIdent()
			if err != nil {
				return StrsExpr{}, err
			}

			strs = append(strs, ident)
		}
	}
}

func (s *Scanner) isSpecialChar(c rune) bool {
	for _, sc := range specials {
		if sc == c {
			return true
		}
	}
	return false
}

func (s *Scanner) isSingleQuotedIdentChar(c rune) bool {
	return !s.isSingleQuotedSpecialChar(c)
}

func (s *Scanner) isDoubleQuotedIdentChar(c rune) bool {
	return !s.isDoubleQuotedSpecialChar(c)
}

func (s *Scanner) isDoubleQuotedSpecialChar(c rune) bool {
	for _, sc := range doubleQuotedSpecials {
		if sc == c {
			return true
		}
	}
	return false
}

func (s *Scanner) isSingleQuotedSpecialChar(c rune) bool {
	for _, sc := range singleQuotedSpecials {
		if sc == c {
			return true
		}
	}
	return false
}

func (s *Scanner) scanDoubleQuotedIdent() (Ident, error) {
	var ret []rune
	for s.isDoubleQuotedIdentChar(s.Peek()) {
		ret = append(ret, s.Peek())
		s.Next()
	}
	if len(ret) == 0 {
		return Ident{}, ExpectIdentError(fmt.Errorf("expected IDENT"))
	}
	return Ident{Name: string(ret)}, nil
}

func (s *Scanner) scanSingleQuotedIdent() (Ident, error) {
	var ret []rune
	for s.isSingleQuotedIdentChar(s.Peek()) {
		ret = append(ret, s.Peek())
		s.Next()
	}
	if len(ret) == 0 {
		return Ident{}, ExpectIdentError(fmt.Errorf("expected IDENT"))
	}
	return Ident{Name: string(ret)}, nil
}

func (s *Scanner) isNumber(c rune) bool {
	return '0' <= c && c <= '9'
}

func (s *Scanner) isIdentChar(c rune) bool {
	return !s.isSpecialChar(c)
}

func (s *Scanner) isLetter(c rune) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z'
}

func (s *Scanner) isBrank(c rune) bool {
	return c == ' ' || c == '\t'
}

func (s *Scanner) scanVar() (VarExpr, error) {
	s.Next() // The '$'
	var ret []rune
	for s.isVarChar(s.Peek()) {
		ret = append(ret, s.Peek())
		s.Next()
	}
	if len(ret) == 0 {
		return VarExpr{}, ExpectVarNameError(fmt.Errorf("expected variable name after $"))
	}

	name := string(ret)
	return VarExpr{Name: name}, nil
}

func (s *Scanner) isVarChar(c rune) bool {
	return s.isLetter(c) || s.isNumber(c) || c == '_'
}

func (s *Scanner) scanIdent() (Ident, error) {
	var ret []rune
	for s.isIdentChar(s.Peek()) {
		fmt.Printf("RETERET: %#v\n", ret)
		ret = append(ret, s.Peek())
		s.Next()
	}
	if len(ret) == 0 {
		return Ident{}, ExpectIdentError(fmt.Errorf("expected IDENT"))
	}
	return Ident{Name: string(ret)}, nil
}

func (s *Scanner) isStrChar(c rune) bool {
	return true
}

func (s *Scanner) scanStrs() (int, StrsExpr, error) {
	withRightParen := false
	strs := StrsExpr{}
	for {
		fmt.Printf("NENENENEXET: %c\n", byte(s.Peek()))
		switch {
		case s.Peek() == '$':
			vare, err := s.scanVar()
			if err != nil {
				return 0, nil, err
			}
			strs = append(strs, vare)

		case s.Peek() == '"':
			quotedStrs, err := s.scanDoubleQuoted()
			if err != nil {
				return 0, nil, err
			}

			strs = append(strs, quotedStrs...)

		case s.Peek() == '\'':
			quotedStrs, err := s.scanSingleQuoted()
			if err != nil {
				return 0, nil, err
			}

			strs = append(strs, quotedStrs...)

		case s.Peek() == ')':
			if len(strs) > 0 {
				return STRS, strs, nil
			}
			if withRightParen {
				if len(strs) == 0 {
					return ')', strs, nil
				}
				return RIGHT_PAREN_AND_STRS, strs, nil
			}
			s.Next()
			withRightParen = true

		case s.Peek() == '(':
			s.Next()
			if len(strs) == 0 && !withRightParen {
				return '(', strs, nil
			}
			if withRightParen {
				return RIGHT_PAREN_AND_STRS_AND_LEFT_PAREN, strs, nil
			}
			return STRS_AND_LEFT_PAREN, strs, nil

		case s.isSpecialChar(s.Peek()):
			if withRightParen {
				return RIGHT_PAREN_AND_STRS, strs, nil
			}
			return STRS, strs, nil

		default:
			ident, err := s.scanIdent()
			if err != nil {
				return 0, nil, err
			}
			strs = append(strs, ident)
			fmt.Printf("IDENT: %#v\n", ident)
		}
	}
}

func (s *Scanner) skipBrank() {
	for s.isBrank(s.Peek()) {
		s.Next()
	}
}

func (s *Scanner) skipComment() {
	for s.Peek() != '\n' {
		s.Next()
	}
}

func (s *Scanner) scanFD() (FD, error) {
	s.Next() // '&'
	if s.Peek() == '-' {
		return FD{N: -1}, nil
	}

	ret := []rune{}
	for s.isNumber(s.Peek()) {
		ret = append(ret, s.Peek())
		s.Next()
	}
	n, err := strconv.Atoi(string(ret))
	if err != nil {
		return FD{}, err
	}
	return FD{N: n}, nil
}

func (*Scanner) tokenOf(v StrsExpr) (int, bool) {
	if len(v) != 1 {
		return 0, false
	}

	ident, ok := v[0].(Ident)
	if !ok {
		return 0, false
	}

	lit := ident.Name

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
	s      *Scanner
	result []Stmt
}

func (l *Lexer) mainLex(lval *yySymType) (int, error) {
	s := l.s

	RETRY:
	s.skipBrank()
	c := s.Peek()
	fmt.Printf("NEXT CHAR: %c\n", byte(c))

	switch c {
	case scanner.EOF:
		return int(c), nil

	case '#':
		s.skipComment()
		goto RETRY

	// ">", ">&N", ">>"
	case '>':
		s.Next()
		switch s.Peek() {
		case '&':
			fd, err := s.scanFD()
			if err != nil {
				return 0, err
			}
			lval.expr = fd
			return REDIRECT_TO_FD, nil
		case '>':
			return APPEND_REDIRECT, nil
		}
		return '>', nil

	// "^", "^&N", "^^"
	case '^':
		s.Next()
		switch s.Peek() {
		case '&':
			fd, err := s.scanFD()
			if err != nil {
				return 0, err
			}
			lval.expr = fd
			return ERR_REDIRECT_TO_FD, nil

		case '^':
			return APPEND_ERR_REDIRECT, nil
		}
		return '^', nil

	case ';', '\n', '|', '<':
		s.Next()
		return int(c), nil

	default:
		tok, v, err := s.scanStrs()
		if err != nil {
			return 0, err
		}
		if tok != STRS {
			lval.strs = v
			return tok, nil
		}
		if tok, ok := s.tokenOf(v); ok {
			fmt.Printf("TOKEN: %#v\n", tok)
			return tok, nil
		}
		lval.strs = v

		fmt.Printf("STRS: %#v\n", v)
		return STRS, nil
	}
}

func (l *Lexer) Lex(lval *yySymType) int {
	tok, err := l.mainLex(lval)
	if err != nil {
		panic(err)
	}

	fmt.Printf("POSITION: %s\n", l.s.Pos())
	return tok
}

func (l *Lexer) Error(e string) {
	fmt.Printf("%#+v\n", yyErrorMessages)
	panic(e)
}
