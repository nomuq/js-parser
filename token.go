package js-parser

import (
	"fmt"
)

// Token is an enum of all terminal symbols of the js language.
type Token int

// Token values.
const (
	UNAVAILABLE Token = iota - 1
	EOI
	INVALID_TOKEN
	ERROR
	WHITESPACE
	MULTILINECOMMENT
	SINGLELINECOMMENT
	IDENTIFIER
	AWAIT              // await
	BREAK              // break
	CASE               // case
	CATCH              // catch
	CLASS              // class
	CONST              // const
	CONTINUE           // continue
	DEBUGGER           // debugger
	DEFAULT            // default
	DELETE             // delete
	DO                 // do
	ELSE               // else
	EXPORT             // export
	EXTENDS            // extends
	FINALLY            // finally
	FOR                // for
	FUNCTION           // function
	IF                 // if
	IMPORT             // import
	IN                 // in
	INSTANCEOF         // instanceof
	NEW                // new
	RETURN             // return
	SUPER              // super
	SWITCH             // switch
	THIS               // this
	THROW              // throw
	TRY                // try
	TYPEOF             // typeof
	VAR                // var
	VOID               // void
	WHILE              // while
	WITH               // with
	YIELD              // yield
	ENUM               // enum
	NULL               // null
	TRUE               // true
	FALSE              // false
	AS                 // as
	ASYNC              // async
	FROM               // from
	GET                // get
	LET                // let
	OF                 // of
	SET                // set
	STATIC             // static
	TARGET             // target
	IMPLEMENTS         // implements
	INTERFACE          // interface
	PRIVATE            // private
	PROTECTED          // protected
	PUBLIC             // public
	ANY                // any
	UNKNOWN            // unknown
	BOOLEAN            // boolean
	NUMBER             // number
	STRING             // string
	SYMBOL             // symbol
	ABSTRACT           // abstract
	CONSTRUCTOR        // constructor
	DECLARE            // declare
	IS                 // is
	MODULE             // module
	NAMESPACE          // namespace
	REQUIRE            // require
	TYPE               // type
	READONLY           // readonly
	KEYOF              // keyof
	UNIQUE             // unique
	INFER              // infer
	LBRACE             // {
	RBRACE             // }
	LPAREN             // (
	RPAREN             // )
	LBRACK             // [
	RBRACK             // ]
	DOT                // .
	DOTDOTDOT          // ...
	SEMICOLON          // ;
	COMMA              // ,
	LT                 // <
	GT                 // >
	LTASSIGN           // <=
	GTASSIGN           // >=
	ASSIGNASSIGN       // ==
	EXCLASSIGN         // !=
	ASSIGNASSIGNASSIGN // ===
	EXCLASSIGNASSIGN   // !==
	ATSIGN             // @
	PLUS               // +
	MINUS              // -
	MULT               // *
	DIV                // /
	REM                // %
	PLUSPLUS           // ++
	MINUSMINUS         // --
	LTLT               // <<
	GTGT               // >>
	GTGTGT             // >>>
	AND                // &
	OR                 // |
	XOR                // ^
	EXCL               // !
	TILDE              // ~
	ANDAND             // &&
	OROR               // ||
	QUEST              // ?
	QUESTQUEST         // ??
	QUESTDOT           // ?.
	COLON              // :
	ASSIGN             // =
	PLUSASSIGN         // +=
	MINUSASSIGN        // -=
	MULTASSIGN         // *=
	DIVASSIGN          // /=
	REMASSIGN          // %=
	LTLTASSIGN         // <<=
	GTGTASSIGN         // >>=
	GTGTGTASSIGN       // >>>=
	ANDASSIGN          // &=
	ORASSIGN           // |=
	XORASSIGN          // ^=
	ASSIGNGT           // =>
	MULTMULT           // **
	MULTMULTASSIGN     // **=
	NUMERICLITERAL
	STRINGLITERAL
	NOSUBSTITUTIONTEMPLATE
	TEMPLATEHEAD
	TEMPLATEMIDDLE
	TEMPLATETAIL
	REGULAREXPRESSIONLITERAL
	JSXSTRINGLITERAL
	JSXIDENTIFIER
	JSXTEXT
	RESOLVESHIFT

	NumTokens
)

var tokenStr = [...]string{
	"EOI",
	"INVALID_TOKEN",
	"ERROR",
	"WHITESPACE",
	"MULTILINECOMMENT",
	"SINGLELINECOMMENT",
	"IDENTIFIER",
	"await",
	"break",
	"case",
	"catch",
	"class",
	"const",
	"continue",
	"debugger",
	"default",
	"delete",
	"do",
	"else",
	"export",
	"extends",
	"finally",
	"for",
	"function",
	"if",
	"import",
	"in",
	"instanceof",
	"new",
	"return",
	"super",
	"switch",
	"this",
	"throw",
	"try",
	"typeof",
	"var",
	"void",
	"while",
	"with",
	"yield",
	"enum",
	"null",
	"true",
	"false",
	"as",
	"async",
	"from",
	"get",
	"let",
	"of",
	"set",
	"static",
	"target",
	"implements",
	"interface",
	"private",
	"protected",
	"public",
	"any",
	"unknown",
	"boolean",
	"number",
	"string",
	"symbol",
	"abstract",
	"constructor",
	"declare",
	"is",
	"module",
	"namespace",
	"require",
	"type",
	"readonly",
	"keyof",
	"unique",
	"infer",
	"{",
	"}",
	"(",
	")",
	"[",
	"]",
	".",
	"...",
	";",
	",",
	"<",
	">",
	"<=",
	">=",
	"==",
	"!=",
	"===",
	"!==",
	"@",
	"+",
	"-",
	"*",
	"/",
	"%",
	"++",
	"--",
	"<<",
	">>",
	">>>",
	"&",
	"|",
	"^",
	"!",
	"~",
	"&&",
	"||",
	"?",
	"??",
	"?.",
	":",
	"=",
	"+=",
	"-=",
	"*=",
	"/=",
	"%=",
	"<<=",
	">>=",
	">>>=",
	"&=",
	"|=",
	"^=",
	"=>",
	"**",
	"**=",
	"NUMERICLITERAL",
	"STRINGLITERAL",
	"NOSUBSTITUTIONTEMPLATE",
	"TEMPLATEHEAD",
	"TEMPLATEMIDDLE",
	"TEMPLATETAIL",
	"REGULAREXPRESSIONLITERAL",
	"JSXSTRINGLITERAL",
	"JSXIDENTIFIER",
	"JSXTEXT",
	"RESOLVESHIFT",
}

func (tok Token) String() string {
	if tok >= 0 && int(tok) < len(tokenStr) {
		return tokenStr[tok]
	}
	return fmt.Sprintf("token(%d)", tok)
}
