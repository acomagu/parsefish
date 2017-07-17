%{
package main

%}

%union{
    stmts []Stmt
    stmt  Stmt
    cmd CmdExpr
    arg ArgExpr
    args []ArgExpr
    ident Ident
    tok string
}

%type<stmts> prog
%type<stmts> stmts
%type<stmt> stmt
%type<stmt> if_stmt
%type<stmt> begin_stmt
%type<arg> arg
%type<args> args
%type<cmd> cmd
%token<ident> IDENT
%token<tok> IF ELSE BEGIN END

%%

prog
    : stmts
    {
        $$ = $1
        yylex.(*Lexer).result = $$
    }

stmts
    : stmt
    {
        $$ = []Stmt{$1}
    }
    | stmt stmts
    {
        $$ = append([]Stmt{$1}, $2...)
    }

stmt
    : cmd eos
    {
        $$ = CmdStmt{Cmd: $1}
    }
    | cmd args eos
    {
        $$ = CmdStmt{Cmd: $1, Args: $2}
    }
    | begin_stmt
    | if_stmt

begin_stmt
    : BEGIN eos stmts END eos
    {
        $$ = BeginStmt{Body: $3}
    }

if_stmt
    : IF stmt stmts END eos
    {
        $$ = IfStmt{Cond: $2, Body: $3}
    }
    | IF stmt stmts ELSE eos stmts END eos
    {
        $$ = IfStmt{Cond: $2, Body: $3, Else: $6}
    }
    | IF stmt stmts ELSE if_stmt
    {
        $$ = IfStmt{Cond: $2, Body: $3, Else: []Stmt{$5}}
    }

args
    : arg
    {
        $$ = []ArgExpr{$1}
    }
    | arg args
    {
        $$ = append([]ArgExpr{$1}, $2...)
    }

cmd
    : IDENT
    {
        $$ = $1
    }

arg
    : IDENT
    {
        $$ = $1
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
