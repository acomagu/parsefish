package token

import (
	"github.com/acomagu/parsefish/internal/yyparse"
)

const Str = yyparse.STR
const StrAndLeftParen = yyparse.STR_AND_LEFT_PAREN
const RightParenAndStr = yyparse.RIGHT_PAREN_AND_STR
const RightParenAndStrAndLeftParen = yyparse.RIGHT_PAREN_AND_STR_AND_LEFT_PAREN
const RedirectToFD = yyparse.REDIRECT_TO_FD
const ErrRedirectToFD = yyparse.ERR_REDIRECT_TO_FD
const CmdSub = yyparse.CMD_SUB
const Var = yyparse.VAR
const AppendRedirect = yyparse.APPEND_REDIRECT
const AppendErrRedirect = yyparse.APPEND_ERR_REDIRECT
const NextRightParen = yyparse.NEXT_RIGHT_PAREN
const NextLeftParen = yyparse.NEXT_LEFT_PAREN
const If = yyparse.IF
const Else = yyparse.ELSE
const Begin = yyparse.BEGIN
const End = yyparse.END
const Function = yyparse.FUNCTION
