package main

type Node interface{}

type Stmt interface {
	Node
	stmt()
}

type stmtImpl struct{}

func (stmtImpl) stmt() {}

type CmdStmt struct {
	stmtImpl
	Cmd  Expr
	Args []Expr
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

type RedirectStmt struct {
	stmtImpl
	Lhs Stmt
	Rhs Expr
	Err bool
	Append bool
}

type VarExpr struct {
	strExprImpl
	Name string
}

type Expr interface {
	expr()
}

type exprImpl struct{}

func (exprImpl) expr() {}

type StrsExpr []StrExpr

func (StrsExpr) expr() {}

func (StrsExpr) cmdExpr() {}

type StrExpr interface{
	strExpr()
}

type strExprImpl struct{}

func (strExprImpl) strExpr() {}

type CmdSubExpr struct {
	exprImpl
	strExprImpl
	Body []Stmt
}

type Ident struct {
	strExprImpl
	Name string
}

type FD struct {
	exprImpl
	N int
}

type Token int
