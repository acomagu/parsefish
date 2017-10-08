//line yyparse.go.y:2
package parser

import __yyfmt__ "fmt"

//line yyparse.go.y:2
import (
	"github.com/acomagu/parsefish/ast"
)

//line yyparse.go.y:9
type yySymType struct {
	yys   int
	stmts ast.Stmts
	stmt  ast.Stmt
	str   ast.StrExpr
	expr  ast.Expr
	exprs ast.Exprs
}

const STR = 57346
const STR_AND_LEFT_PAREN = 57347
const RIGHT_PAREN_AND_STR = 57348
const RIGHT_PAREN_AND_STR_AND_LEFT_PAREN = 57349
const REDIRECT_TO_FD = 57350
const ERR_REDIRECT_TO_FD = 57351
const CMD_SUB = 57352
const VAR = 57353
const APPEND_REDIRECT = 57354
const APPEND_ERR_REDIRECT = 57355
const NEXT_RIGHT_PAREN = 57356
const NEXT_LEFT_PAREN = 57357
const IF = 57358
const ELSE = 57359
const BEGIN = 57360
const END = 57361
const FUNCTION = 57362

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"STR",
	"STR_AND_LEFT_PAREN",
	"RIGHT_PAREN_AND_STR",
	"RIGHT_PAREN_AND_STR_AND_LEFT_PAREN",
	"REDIRECT_TO_FD",
	"ERR_REDIRECT_TO_FD",
	"CMD_SUB",
	"VAR",
	"APPEND_REDIRECT",
	"APPEND_ERR_REDIRECT",
	"NEXT_RIGHT_PAREN",
	"NEXT_LEFT_PAREN",
	"IF",
	"ELSE",
	"BEGIN",
	"END",
	"FUNCTION",
	"'|'",
	"'>'",
	"'^'",
	"'('",
	"')'",
	"';'",
	"'\\n'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yyparse.go.y:216

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 101

var yyAct = [...]int{

	2, 7, 21, 55, 20, 49, 16, 19, 4, 15,
	32, 34, 40, 16, 19, 36, 73, 39, 13, 30,
	12, 35, 14, 63, 35, 62, 18, 41, 41, 31,
	32, 43, 42, 18, 70, 31, 32, 51, 52, 65,
	54, 53, 13, 35, 59, 50, 60, 61, 35, 48,
	50, 47, 31, 32, 64, 46, 57, 58, 45, 16,
	19, 33, 44, 29, 17, 69, 68, 41, 41, 72,
	71, 66, 67, 25, 28, 56, 38, 24, 27, 18,
	25, 28, 3, 11, 24, 27, 22, 23, 26, 10,
	9, 31, 32, 22, 23, 26, 37, 8, 6, 5,
	1,
}
var yyPact = [...]int{

	2, -1000, -1000, 2, 65, -1000, -1000, -1000, -1000, -1000,
	-1000, 55, 3, 2, 9, -1000, -1000, -1000, 2, 2,
	-1000, -1000, 2, 58, 54, -1000, 51, 47, -1000, -17,
	-22, -1000, -1000, 55, -1000, -1000, 2, 2, 9, 2,
	50, 65, 50, 72, -1000, -1000, -1000, -1000, -22, -1000,
	-1000, 28, 6, 2, 20, -1000, -1000, -1000, 2, 2,
	-1000, -1000, -1000, 26, 15, -1000, 50, -1000, 2, -1000,
	-1000, -1000, -3, -1000,
}
var yyPgo = [...]int{

	0, 100, 0, 12, 82, 8, 99, 1, 98, 97,
	90, 89, 83, 11, 9, 64, 3, 61, 2, 63,
	19,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 4, 3, 3, 5, 5, 5,
	5, 5, 5, 6, 6, 8, 7, 7, 7, 9,
	9, 10, 11, 11, 11, 11, 11, 11, 17, 17,
	12, 13, 14, 14, 15, 15, 16, 16, 16, 19,
	20, 20, 18, 18, 18,
}
var yyR2 = [...]int{

	0, 1, 0, 2, 2, 1, 3, 1, 1, 1,
	1, 1, 1, 1, 2, 4, 4, 7, 5, 5,
	4, 3, 3, 3, 2, 3, 3, 2, 1, 2,
	1, 1, 1, 1, 3, 3, 1, 1, 3, 1,
	1, 2, 1, 1, 2,
}
var yyChk = [...]int{

	-1000, -1, -2, -4, -5, -6, -8, -7, -9, -10,
	-11, -12, 18, 16, 20, -14, 4, -15, 24, 5,
	-2, -18, 21, 22, 12, 8, 23, 13, 9, -19,
	-20, 26, 27, -17, -13, -14, -18, -4, -17, -18,
	-3, -5, -3, -5, 4, 4, 4, 4, -20, 27,
	-13, -2, -2, -18, -2, -16, 25, 6, 7, -18,
	-16, 19, 19, 17, -2, 19, -3, -3, -18, -7,
	19, -16, -2, 19,
}
var yyDef = [...]int{

	2, -2, 1, 2, 0, 7, 8, 9, 10, 11,
	12, 13, 0, 0, 0, 30, 32, 33, 0, 0,
	3, 4, 0, 0, 0, 24, 0, 0, 27, 42,
	43, 39, 40, 14, 28, 31, 2, 2, 0, 2,
	0, 5, 0, 21, 22, 23, 25, 26, 44, 41,
	29, 0, 0, 2, 0, 34, 36, 37, 0, 0,
	35, 15, 16, 0, 0, 20, 0, 6, 2, 18,
	19, 38, 0, 17,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	27, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	24, 25, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 26,
	3, 3, 22, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 23, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 21,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yyparse.go.y:41
		{
			yyVAL.stmts = yyDollar[1].stmts
			yylex.(*Lexer).result = yyVAL.stmts
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line yyparse.go.y:48
		{
			yyVAL.stmts = ast.Stmts{}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line yyparse.go.y:52
		{
			yyVAL.stmts = append(ast.Stmts{yyDollar[1].stmt}, yyDollar[2].stmts...)
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line yyparse.go.y:58
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yyparse.go.y:64
		{
			yyVAL.stmts = ast.Stmts{yyDollar[1].stmt}
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yyparse.go.y:68
		{
			yyVAL.stmts = append(yyDollar[3].stmts, yyDollar[1].stmt)
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yyparse.go.y:82
		{
			yyVAL.stmt = ast.CmdStmt{Cmd: yyDollar[1].expr}
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line yyparse.go.y:86
		{
			yyVAL.stmt = ast.CmdStmt{Cmd: yyDollar[1].expr, Args: yyDollar[2].exprs}
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line yyparse.go.y:92
		{
			yyVAL.stmt = ast.BeginStmt{Body: yyDollar[3].stmts}
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line yyparse.go.y:98
		{
			yyVAL.stmt = ast.IfStmt{Cond: yyDollar[2].stmt, Body: yyDollar[3].stmts}
		}
	case 17:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line yyparse.go.y:102
		{
			yyVAL.stmt = ast.IfStmt{Cond: yyDollar[2].stmt, Body: yyDollar[3].stmts, Else: yyDollar[6].stmts}
		}
	case 18:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line yyparse.go.y:106
		{
			yyVAL.stmt = ast.IfStmt{Cond: yyDollar[2].stmt, Body: yyDollar[3].stmts, Else: ast.Stmts{yyDollar[5].stmt}}
		}
	case 19:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line yyparse.go.y:112
		{
			yyVAL.stmt = ast.FunctionStmt{Args: yyDollar[2].exprs, Body: yyDollar[4].stmts}
		}
	case 20:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line yyparse.go.y:116
		{
			yyVAL.stmt = ast.FunctionStmt{Body: yyDollar[3].stmts}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yyparse.go.y:122
		{
			yyVAL.stmt = ast.PipeStmt{Lhs: yyDollar[1].stmt, Rhs: yyDollar[3].stmt}
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yyparse.go.y:128
		{
			yyVAL.stmt = ast.RedirectStmt{Lhs: yyDollar[1].stmt, Rhs: yyDollar[3].str, Err: false, Append: false}
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yyparse.go.y:132
		{
			yyVAL.stmt = ast.RedirectStmt{Lhs: yyDollar[1].stmt, Rhs: yyDollar[3].str, Err: false, Append: true}
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line yyparse.go.y:136
		{
			yyVAL.stmt = ast.RedirectStmt{Lhs: yyDollar[1].stmt, Rhs: yyDollar[2].expr, Err: false, Append: false}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yyparse.go.y:140
		{
			yyVAL.stmt = ast.RedirectStmt{Lhs: yyDollar[1].stmt, Rhs: yyDollar[3].str, Err: true, Append: false}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yyparse.go.y:144
		{
			yyVAL.stmt = ast.RedirectStmt{Lhs: yyDollar[1].stmt, Rhs: yyDollar[3].str, Err: true, Append: true}
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line yyparse.go.y:148
		{
			yyVAL.stmt = ast.RedirectStmt{Lhs: yyDollar[1].stmt, Rhs: yyDollar[2].expr, Err: true, Append: false}
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yyparse.go.y:154
		{
			yyVAL.exprs = ast.Exprs{yyDollar[1].expr}
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line yyparse.go.y:158
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[2].expr)
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yyparse.go.y:164
		{
			yyVAL.expr = yyDollar[1].str
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yyparse.go.y:170
		{
			yyVAL.expr = yyDollar[1].str
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yyparse.go.y:176
		{
			yyVAL.str = yyDollar[1].str
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yyparse.go.y:183
		{
			yyVAL.str = append(ast.StrExpr{ast.CmdSub{Body: yyDollar[2].stmts}}, yyDollar[3].str...)
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yyparse.go.y:187
		{
			tmp := append(yyDollar[1].str, ast.CmdSub{Body: yyDollar[2].stmts})
			yyVAL.str = append(tmp, yyDollar[3].str...)
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line yyparse.go.y:194
		{
			yyVAL.str = ast.StrExpr{}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yyparse.go.y:199
		{
			tmp := append(yyDollar[1].str, ast.CmdSub{Body: yyDollar[2].stmts})
			yyVAL.str = append(tmp, yyDollar[3].str...)
		}
	}
	goto yystack /* stack new state and value */
}
