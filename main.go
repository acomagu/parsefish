package main

import (
	"strings"
	"fmt"
	"os"
)

func main() {
    l := new(Lexer)
    l.Init(strings.NewReader(os.Args[1]))
    yyParse(l)
    fmt.Printf("%#v\n", l.result)
}
