# parsefish

Parser for Fish shell script written in Go.

## Description

This is a library to parse Fish shell script. Use like

```Go
tree := parser.ParseExpr("ls -la | cat -\n")
```

and get AST node tree.

It can be printed like this:

```Go
ast.Print(tree)
```

```
     0  ast.Stmts (len = 1) {
     1  .  0: ast.PipeStmt {
     2  .  .  Lhs: ast.CmdStmt {
     3  .  .  .  Cmd: ast.StrExpr (len = 1) {
     4  .  .  .  .  0: ast.Ident {
     5  .  .  .  .  .  Name: "ls"
     6  .  .  .  .  }
     7  .  .  .  }
     8  .  .  .  Args: []ast.Expr (len = 1) {
     9  .  .  .  .  0: ast.StrExpr (len = 1) {
    10  .  .  .  .  .  0: ast.Ident {
    11  .  .  .  .  .  .  Name: "-la"
    12  .  .  .  .  .  }
    13  .  .  .  .  }
    14  .  .  .  }
    15  .  .  }
    16  .  .  Rhs: ast.CmdStmt {
    17  .  .  .  Cmd: ast.StrExpr (len = 1) {
    18  .  .  .  .  0: ast.Ident {
    19  .  .  .  .  .  Name: "cat"
    20  .  .  .  .  }
    21  .  .  .  }
    22  .  .  .  Args: []ast.Expr (len = 1) {
    23  .  .  .  .  0: ast.StrExpr (len = 1) {
    24  .  .  .  .  .  0: ast.Ident {
    25  .  .  .  .  .  .  Name: "-"
    26  .  .  .  .  .  }
    27  .  .  .  .  }
    28  .  .  .  }
    29  .  .  }
    30  .  }
    31  }
```

## Installation

The `go generate` command must be run before using.

```
$ go get github.com/acomagu/parsefish
$ go generate github.com/acomagu/parsefish/internal/yyparse
```

## Project Status

Experimental. The API can be changed roughly!

## Author

[acomagu](https://github.com/acomagu)

## License

MIT
