package parser

import (
	"io"

	"github.com/acomagu/parsefish/internal/yyparse"
	"github.com/acomagu/parsefish/ast"
	"strings"
)

func Parse(input io.Reader) ast.Node {
	return yyparse.Parse(input)
}

// Node might be Expr.
func ParseExpr(x string) ast.Node {
	return Parse(strings.NewReader(x))
}

type inspector func(ast.Node) bool
type Visitor func(ast.Node) Visitor

func (f inspector) visit(node ast.Node) Visitor {
	if f(node) {
		return f.visit
	}
	return nil
}

func Inspect(node ast.Node, f func(ast.Node) bool) {
	Walk(inspector(f).visit, node)
}

func Walk(v Visitor, node ast.Node) {
	if node == nil {
		return
	}

	f := v(node)
	switch n := node.(type) {
	case ast.Exprs:
		for _, e := range n {
			Walk(f, e)
		}
	case ast.Stmts:
		for _, s := range n {
			Walk(f, s)
		}
	case ast.CmdStmt:
		Walk(f, n.Cmd)
		walkExprs(f, ast.Exprs(n.Args))
	case ast.BeginStmt:
		Walk(f, ast.Stmts(n.Body))
	case ast.IfStmt:
		Walk(f, n.Cond)
		walkStmts(f, ast.Stmts(n.Body))
		walkStmts(f, ast.Stmts(n.Else))
	case ast.FunctionStmt:
		walkExprs(f, ast.Exprs(n.Args))
		walkStmts(f, ast.Stmts(n.Body))
	case ast.PipeStmt:
		Walk(f, n.Lhs)
		Walk(f, n.Rhs)
	case ast.RedirectStmt:
		Walk(f, n.Lhs)
		Walk(f, n.Rhs)
	}
	v(nil)
}

func walkStmts(f Visitor, stmts ast.Stmts) {
	if stmts != nil && len(stmts) > 0 {
		Walk(f, stmts)
	}
}

func walkExprs(f Visitor, exprs ast.Exprs) {
	if exprs != nil && len(exprs) > 0 {
		Walk(f, exprs)
	}
}
