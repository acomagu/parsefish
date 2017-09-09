%{
package main

%}

%union{
    stmts Stmts
    stmt  Stmt
    str StrExpr
    expr Expr
    exprs Exprs
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
%type<str> str paren_str open_left_paren
%type<exprs> args
%token<str> STR STR_AND_LEFT_PAREN RIGHT_PAREN_AND_STR RIGHT_PAREN_AND_STR_AND_LEFT_PAREN
%token<expr> REDIRECT_TO_FD ERR_REDIRECT_TO_FD CMD_SUB VAR
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
        $$ = Stmts{}
    }
    | line lines
    {
        $$ = append(Stmts{$1}, $2...)
    }

line
    : stmt eos
    {
        $$ = $1
    }

stmts
    : stmt
    {
        $$ = Stmts{$1}
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
        $$ = IfStmt{Cond: $2, Body: $3, Else: Stmts{$5}}
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
    : stmt '>' STR
    {
        $$ = RedirectStmt{Lhs: $1, Rhs: $3, Err: false, Append: false}
    }
    | stmt APPEND_REDIRECT STR
    {
        $$ = RedirectStmt{Lhs: $1, Rhs: $3, Err: false, Append: true}
    }
    | stmt REDIRECT_TO_FD
    {
        $$ = RedirectStmt{Lhs: $1, Rhs: $2, Err: false, Append: false}
    }
    | stmt '^' STR
    {
        $$ = RedirectStmt{Lhs: $1, Rhs: $3, Err: true, Append: false}
    }
    | stmt APPEND_ERR_REDIRECT STR
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
        $$ = Exprs{$1}
    }
    | args arg
    {
        $$ = append($1, $2)
    }

cmd
    : str
    {
        $$ = $1
    }

arg
    : str
    {
        $$ = $1
    }

str
    : STR
    {
        $$ = $1
    }
    | paren_str

paren_str
    : '(' stmts open_left_paren
    {
        $$ = append(StrExpr{CmdSub{Body: $2}}, $3...)
    }
    | STR_AND_LEFT_PAREN stmts open_left_paren
    {
        tmp := append($1, CmdSub{Body: $2})
        $$ = append(tmp, $3...)
    }

open_left_paren
    : ')'
    {
        $$ = StrExpr{}
    }
    | RIGHT_PAREN_AND_STR
    | RIGHT_PAREN_AND_STR_AND_LEFT_PAREN stmts open_left_paren
    {
        tmp := append($1, CmdSub{Body: $2})
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
