package ast

import (
	"text/scanner"
)

type Node interface {
	Pos() scanner.Position
}

type nodeImpl struct {
	pos scanner.Position
}

func (n nodeImpl) Pos() scanner.Position {
	return n.pos
}

func (n *nodeImpl) SetPos(pos scanner.Position) {
	n.pos = pos
}

type Stmts []Stmt

func (s Stmts) Pos() scanner.Position {
	return s[0].Pos()
}

type Stmt interface {
	Node
	stmt()
}

type stmtImpl struct {
	nodeImpl
}

func (stmtImpl) stmt() {}

type CmdStmt struct {
	stmtImpl
	Cmd  Expr
	Args []Expr
}

func (s CmdStmt) Pos() scanner.Position {
	return s.Cmd.Pos()
}

type BeginStmt struct {
	stmtImpl
	Body []Stmt
}

type IfStmt struct {
	stmtImpl
	Cond Stmt
	Body []Stmt
	Else []Stmt
}

type FunctionStmt struct {
	stmtImpl
	Args []Expr
	Body []Stmt
}

type PipeStmt struct {
	stmtImpl
	Lhs Stmt
	Rhs Stmt
}

func (s PipeStmt) Pos() scanner.Position {
	return s.Lhs.Pos()
}

type RedirectStmt struct {
	stmtImpl
	Lhs    Stmt
	Rhs    Expr
	Err    bool
	Append bool
}

func (s RedirectStmt) Pos() scanner.Position {
	return s.Lhs.Pos()
}

type Exprs []Expr

func (e Exprs) Pos() scanner.Position {
	return e[0].Pos()
}

type Expr interface {
	Node
	expr()
}

type exprImpl struct {
	nodeImpl
}

func (exprImpl) expr() {}

type StrExpr []SubStr

func (e StrExpr) Pos() scanner.Position {
	return e[0].Pos()
}

func (StrExpr) expr() {}

type strExprImpl struct {
	exprImpl
}

type SubStr interface {
	Node
	subStr()
}

type subStrImpl struct {
	nodeImpl
}

func (subStrImpl) subStr() {}

type CmdSub struct {
	subStrImpl
	Body []Stmt
}

type Ident struct {
	subStrImpl
	Name string
}

type FD struct {
	subStrImpl
	N int
}

type VarExpr struct {
	subStrImpl
	Name string
}

type Token int
