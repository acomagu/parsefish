%{
package main

%}

%union{
    stmts []Stmt
    stmt  Stmt
    strs StrsExpr
    expr Expr
    exprs []Expr
}

%type<stmts> prog
%type<stmts> lines
%type<stmts> stmts
%type<stmt> line
%type<stmt> stmt
%type<stmt> cmd_stmt
%type<stmt> if_stmt
%type<stmt> begin_stmt
%type<stmt> function_stmt
%type<stmt> pipe_stmt
%type<stmt> redirect_stmt
%type<expr> cmd
%type<expr> arg
%type<strs> strs
%type<strs> paren_strs
%type<strs> open_left_paren
%type<exprs> args
%token<expr> VAR
%token<strs> STRS STRS_AND_LEFT_PAREN RIGHT_PAREN_AND_STRS RIGHT_PAREN_AND_STRS_AND_LEFT_PAREN
%token<expr> REDIRECT_TO_FD ERR_REDIRECT_TO_FD CMD_SUB
%token<symbol> APPEND_REDIRECT APPEND_ERR_REDIRECT NEXT_RIGHT_PAREN NEXT_LEFT_PAREN
%token<kwd> IF ELSE BEGIN END FUNCTION

%%

prog
    : lines
    {
        $$ = $1
        yylex.(*Lexer).result = $$
    }

lines
    :
    {
        $$ = []Stmt{}
    }
    | line lines
    {
        $$ = append([]Stmt{$1}, $2...)
    }

line
    : stmt eos
    {
        $$ = $1
    }

stmts
    : stmt
    {
        $$ = []Stmt{$1}
    }
    | stmt eos stmts
    {
        $$ = append($3, $1)
    }

stmt
    : cmd_stmt
    | begin_stmt
    | if_stmt
    | function_stmt
    | pipe_stmt
    | redirect_stmt

cmd_stmt
    : cmd
    {
        $$ = CmdStmt{Cmd: $1}
    }
    | cmd args
    {
        $$ = CmdStmt{Cmd: $1, Args: $2}
    }

begin_stmt
    : BEGIN eos lines END
    {
        $$ = BeginStmt{Body: $3}
    }

if_stmt
    : IF line lines END
    {
        $$ = IfStmt{Cond: $2, Body: $3}
    }
    | IF line lines ELSE eos lines END
    {
        $$ = IfStmt{Cond: $2, Body: $3, Else: $6}
    }
    | IF line lines ELSE if_stmt
    {
        $$ = IfStmt{Cond: $2, Body: $3, Else: []Stmt{$5}}
    }

function_stmt
    : FUNCTION args eos lines END
    {
        $$ = FunctionStmt{Args: $2, Body: $4}
    }
    | FUNCTION eos lines END
    {
        $$ = FunctionStmt{Body: $3}
    }

pipe_stmt
    : stmt '|' stmt
    {
        $$ = PipeStmt{Lhs: $1, Rhs: $3}
    }

redirect_stmt
    : stmt '>' STRS
    {
        $$ = RedirectStmt{Lhs: $1, Rhs: $3, Err: false, Append: false}
    }
    | stmt APPEND_REDIRECT STRS
    {
        $$ = RedirectStmt{Lhs: $1, Rhs: $3, Err: false, Append: true}
    }
    | stmt REDIRECT_TO_FD
    {
        $$ = RedirectStmt{Lhs: $1, Rhs: $2, Err: false, Append: false}
    }
    | stmt '^' STRS
    {
        $$ = RedirectStmt{Lhs: $1, Rhs: $3, Err: true, Append: false}
    }
    | stmt APPEND_ERR_REDIRECT STRS
    {
        $$ = RedirectStmt{Lhs: $1, Rhs: $3, Err: true, Append: true}
    }
    | stmt ERR_REDIRECT_TO_FD
    {
        $$ = RedirectStmt{Lhs: $1, Rhs: $2, Err: true, Append: false}
    }

args
    : arg
    {
        $$ = []Expr{$1}
    }
    | args arg
    {
        $$ = append($1, $2)
    }

cmd
    : strs
    {
        $$ = $1
    }

arg
    : strs
    {
        $$ = $1
    }

strs
    : STRS
    | paren_strs

paren_strs
    : '(' stmts open_left_paren
    {
        $$ = append(StrsExpr{CmdSubExpr{Body: $2}}, $3...)
    }
    | STRS_AND_LEFT_PAREN stmts open_left_paren
    {
        tmp := append($1, CmdSubExpr{Body: $2})
        $$ = append(tmp, $3...)
    }

open_left_paren
    : ')'
    {
        $$ = StrsExpr{}
    }
    | RIGHT_PAREN_AND_STRS
    | RIGHT_PAREN_AND_STRS_AND_LEFT_PAREN stmts open_left_paren
    {
        tmp := append($1, CmdSubExpr{Body: $2})
        $$ = append(tmp, $3...)
    }

semicolon
    : ';'

break
    : '\n'
    | break '\n'

eos
    : semicolon
    | break
    | semicolon break

%%
