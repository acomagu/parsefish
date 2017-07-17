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
	Cmd  Cmd
	Args []Arg
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

type Cmd struct {
	ident Ident
}

type Arg struct {
	ident Ident
}

type Ident string
