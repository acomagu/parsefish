%{
package main

type Statement struct {
    command Command
    args []Argument
}
type Command struct {
    ident Ident
}
type Argument struct {
    ident Ident
}
type Ident struct {
    literal string
}

%}

%union{
    stmt  Statement
    command Command
    arg Argument
    args []Argument
    ident Ident
}

%type<stmt> program
%type<stmt> statement
%type<arg> argument
%type<args> arguments
%type<command> command
%token<ident> IDENT

%left '+' '-'

%%

program
    : statement
    {
        $$ = $1
        yylex.(*Lexer).result = $$
    }

statement
    : command
    {
        $$ = Statement{command: $1}
    }
    | command arguments
    {
        $$ = Statement{command: $1, args: $2}
    }

arguments
    : argument
    {
        $$ = []Argument{$1}
    }
    | argument arguments
    {
        $$ = append($2, $1)
    }

command
    : IDENT
    {
        $$ = Command{ident: $1}
    }

argument
    : IDENT
    {
        $$ = Argument{ident: $1}
    }

%%
