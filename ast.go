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
	Cmd  CmdExpr
	Args []ArgExpr
}

type CmdExpr interface {
	cmdExpr()
}

type cmdExprImpl struct{}

func (cmdExprImpl) cmdExpr() {}

type ArgExpr interface {
	argExpr()
}

type argExprImpl struct{}

func (argExprImpl) argExpr() {}

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

type VarExpr struct {
	argExprImpl
	Name string
}

type Expr interface {
	expr()
}

type exprImpl struct{}

func (exprImpl) expr() {}

type Ident struct {
	exprImpl
	cmdExprImpl
	argExprImpl
	Name string
}
