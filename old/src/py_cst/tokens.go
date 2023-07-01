package py_cst

import (
	"regexp"
)

var (
	symbols = map[string]int{
		"+":   ADD,
		"+=":  ADD_ASSIGN,
		"&=":  AND_ASSIGN,
		"&":   AND_OP,
		"=":   ASSIGN,
		"@":   AT,
		"`":   BACK_QUOTE,
		":":   COLON,
		",":   COMMA,
		"/":   DIV,
		"/=":  DIV_ASSIGN,
		".":   DOT,
		"==":  EQUALS,
		">":   GREATER_THAN,
		">=":  GT_EQ,
		"//":  IDIV,
		"//=": IDIV_ASSIGN,
		"{":   LEFT_BRACE,
		"[":   LEFT_BRACKET,
		"(":   LEFT_PARENTHESIS,
		"<<":  LEFT_SHIFT,
		"<<=": LEFT_SHIFT_ASSIGN,
		"<":   LESS_THAN,
		"<=":  LT_EQ,
		"-":   MINUS,
		"%":   MOD,
		"%=":  MOD_ASSIGN,
		"*=":  MULT_ASSIGN,
		"<>":  NOT_EQUAL,
		"!=":  NOT_EQUAL,
		"~":   NOT_OP,
		"|=":  OR_ASSIGN,
		"|":   OR_OP,
		"**":  STAR_STAR,
		"**=": POWER_ASSIGN,
		"}":   RIGHT_BRACE,
		"]":   RIGHT_BRACKET,
		")":   RIGHT_PARENTHESIS,
		">>":  RIGHT_SHIFT,
		">>=": RIGHT_SHIFT_ASSIGN,
		";":   SEMI_COLON,
		"*":   STAR,
		"-=":  SUB_ASSIGN,
		"^":   XOR,
		"^=":  XOR_ASSIGN,

		"\\\n":   LINE_CONTINUATION,
		"\\\r\n": LINE_CONTINUATION,
		"\n":     NEWLINE,
		"\r\n":   NEWLINE,
	}

	keywords = map[string]int{
		"and":      AND,
		"as":       AS,
		"assert":   ASSERT,
		"break":    BREAK,
		"class":    CLASS,
		"continue": CONTINUE,
		"def":      DEF,
		"del":      DEL,
		"elif":     ELIF,
		"else":     ELSE,
		"except":   EXCEPT,
		"exec":     EXEC,
		"finally":  FINALLY,
		"for":      FOR,
		"from":     FROM,
		"global":   GLOBAL,
		"if":       IF,
		"import":   IMPORT,
		"in":       IN,
		"is":       IS,
		"lambda":   LAMBDA,
		"None":     NONE,
		"not":      NOT,
		"or":       OR,
		"pass":     PASS,
		"print":    PRINT,
		"raise":    RAISE,
		"return":   RETURN,
		"try":      TRY,
		"while":    WHILE,
		"with":     WITH,
		"yield":    YIELD,
	}

	spaceRe = regexp.MustCompile(`^([ \t]+)`)
	nameRe  = regexp.MustCompile(`^([a-zA-Z_]\w*)`)

	// NOTE: must match longStrRe before shortStrRe
	shortStrRe = regexp.MustCompile(
		`^([bBuU]?[rR]?(?:` + // unicode + raw
			`(?:"(?:(?:\\.)|[^"])*")|` + // "string"
			`(?:'(?:(?:\\.)|[^'])*')))`)
	longStrRe = regexp.MustCompile(
		`(?msU)^([bBuU]?[rR]?(?:` +
			`(?:'''(?:(?:\\.)|.)*''')|` +
			`(?:"""(?:(?:\\.)|.)*""")))`)

	// NOTE: order of | matters. 0 must be last.
	intRe = regexp.MustCompile(
		"^((?:(?:0[xX][0-9a-fA-F]+)|" +
			"(?:0[oO]?[0-7]+)|" +
			"(?:0[bB][0-1]+)|" +
			"(?:[1-9]\\d*)|0" +
			")[lL]?)")

	floatRe = regexp.MustCompile(
		"^((?:\\d+[eE][+-]?\\d+)|" +
			"(?:\\d+\\.\\d*(?:[eE][+-]?\\d+)?)|" +
			"(\\.\\d+(?:[eE][+-]?\\d+)?))")
)
