// Code generated by goyacc -o parser.go -v parser.output -p demo parser.y. DO NOT EDIT.

//line parser.y:2
package main

import __yyfmt__ "fmt"

//line parser.y:2

//line parser.y:6
type demoSymType struct {
	yys int
	Node
	NodeList []Node
}

const ID = 57346
const PLUS = 57347
const MINUS = 57348
const LBRACE = 57349
const RBRACE = 57350
const BLOCK = 57351
const ERROR = 57352
const EOF = 57353

var demoToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"ID",
	"PLUS",
	"MINUS",
	"LBRACE",
	"RBRACE",
	"BLOCK",
	"ERROR",
	"EOF",
}

var demoStatenames = [...]string{}

const demoEofCode = 1
const demoErrCode = 2
const demoInitialStackSize = 16

//line parser.y:60

func init() {
	demoErrorVerbose = true
}

//line yacctab:1
var demoExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const demoPrivate = 57344

const demoLast = 13

var demoAct = [...]int{
	7, 3, 10, 11, 5, 8, 9, 13, 12, 1,
	4, 6, 2,
}

var demoPact = [...]int{
	-6, -1000, -1000, -1000, -4, -1000, -1000, -3, -1000, -1000,
	4, 3, -1000, -1000,
}

var demoPgo = [...]int{
	0, 12, 11, 10, 9,
}

var demoR1 = [...]int{
	0, 4, 1, 3, 3, 2, 2, 2, 2, 2,
}

var demoR2 = [...]int{
	0, 1, 3, 2, 0, 1, 3, 3, 1, 1,
}

var demoChk = [...]int{
	-1000, -4, -1, 7, -3, 8, -2, 4, 9, 10,
	5, 6, 4, 4,
}

var demoDef = [...]int{
	0, -2, 1, 4, 0, 2, 3, 5, 8, 9,
	0, 0, 6, 7,
}

var demoTok1 = [...]int{
	1,
}

var demoTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
}

var demoTok3 = [...]int{
	0,
}

var demoErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	demoDebug        = 0
	demoErrorVerbose = false
)

type demoLexer interface {
	Lex(lval *demoSymType) int
	Error(s string)
}

type demoParser interface {
	Parse(demoLexer) int
	Lookahead() int
}

type demoParserImpl struct {
	lval  demoSymType
	stack [demoInitialStackSize]demoSymType
	char  int
}

func (p *demoParserImpl) Lookahead() int {
	return p.char
}

func demoNewParser() demoParser {
	return &demoParserImpl{}
}

const demoFlag = -1000

func demoTokname(c int) string {
	if c >= 1 && c-1 < len(demoToknames) {
		if demoToknames[c-1] != "" {
			return demoToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func demoStatname(s int) string {
	if s >= 0 && s < len(demoStatenames) {
		if demoStatenames[s] != "" {
			return demoStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func demoErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !demoErrorVerbose {
		return "syntax error"
	}

	for _, e := range demoErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + demoTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := demoPact[state]
	for tok := TOKSTART; tok-1 < len(demoToknames); tok++ {
		if n := base + tok; n >= 0 && n < demoLast && demoChk[demoAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if demoDef[state] == -2 {
		i := 0
		for demoExca[i] != -1 || demoExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; demoExca[i] >= 0; i += 2 {
			tok := demoExca[i]
			if tok < TOKSTART || demoExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if demoExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += demoTokname(tok)
	}
	return res
}

func demolex1(lex demoLexer, lval *demoSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = demoTok1[0]
		goto out
	}
	if char < len(demoTok1) {
		token = demoTok1[char]
		goto out
	}
	if char >= demoPrivate {
		if char < demoPrivate+len(demoTok2) {
			token = demoTok2[char-demoPrivate]
			goto out
		}
	}
	for i := 0; i < len(demoTok3); i += 2 {
		token = demoTok3[i+0]
		if token == char {
			token = demoTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = demoTok2[1] /* unknown char */
	}
	if demoDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", demoTokname(token), uint(char))
	}
	return char, token
}

func demoParse(demolex demoLexer) int {
	return demoNewParser().Parse(demolex)
}

func (demorcvr *demoParserImpl) Parse(demolex demoLexer) int {
	var demon int
	var demoVAL demoSymType
	var demoDollar []demoSymType
	_ = demoDollar // silence set and not used
	demoS := demorcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	demostate := 0
	demorcvr.char = -1
	demotoken := -1 // demorcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		demostate = -1
		demorcvr.char = -1
		demotoken = -1
	}()
	demop := -1
	goto demostack

ret0:
	return 0

ret1:
	return 1

demostack:
	/* put a state and value onto the stack */
	if demoDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", demoTokname(demotoken), demoStatname(demostate))
	}

	demop++
	if demop >= len(demoS) {
		nyys := make([]demoSymType, len(demoS)*2)
		copy(nyys, demoS)
		demoS = nyys
	}
	demoS[demop] = demoVAL
	demoS[demop].yys = demostate

demonewstate:
	demon = demoPact[demostate]
	if demon <= demoFlag {
		goto demodefault /* simple state */
	}
	if demorcvr.char < 0 {
		demorcvr.char, demotoken = demolex1(demolex, &demorcvr.lval)
	}
	demon += demotoken
	if demon < 0 || demon >= demoLast {
		goto demodefault
	}
	demon = demoAct[demon]
	if demoChk[demon] == demotoken { /* valid shift */
		demorcvr.char = -1
		demotoken = -1
		demoVAL = demorcvr.lval
		demostate = demon
		if Errflag > 0 {
			Errflag--
		}
		goto demostack
	}

demodefault:
	/* default state action */
	demon = demoDef[demostate]
	if demon == -2 {
		if demorcvr.char < 0 {
			demorcvr.char, demotoken = demolex1(demolex, &demorcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if demoExca[xi+0] == -1 && demoExca[xi+1] == demostate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			demon = demoExca[xi+0]
			if demon < 0 || demon == demotoken {
				break
			}
		}
		demon = demoExca[xi+1]
		if demon < 0 {
			goto ret0
		}
	}
	if demon == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			demolex.Error(demoErrorMessage(demostate, demotoken))
			Nerrs++
			if demoDebug >= 1 {
				__yyfmt__.Printf("%s", demoStatname(demostate))
				__yyfmt__.Printf(" saw %s\n", demoTokname(demotoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for demop >= 0 {
				demon = demoPact[demoS[demop].yys] + demoErrCode
				if demon >= 0 && demon < demoLast {
					demostate = demoAct[demon] /* simulate a shift of "error" */
					if demoChk[demostate] == demoErrCode {
						goto demostack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if demoDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", demoS[demop].yys)
				}
				demop--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if demoDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", demoTokname(demotoken))
			}
			if demotoken == demoEofCode {
				goto ret1
			}
			demorcvr.char = -1
			demotoken = -1
			goto demonewstate /* try again in the same state */
		}
	}

	/* reduction by production demon */
	if demoDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", demon, demoStatname(demostate))
	}

	demont := demon
	demopt := demop
	_ = demopt // guard against "declared and not used"

	demop -= demoR2[demon]
	// demop is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if demop+1 >= len(demoS) {
		nyys := make([]demoSymType, len(demoS)*2)
		copy(nyys, demoS)
		demoS = nyys
	}
	demoVAL = demoS[demop+1]

	/* consult goto table to find next state */
	demon = demoR1[demon]
	demog := demoPgo[demon]
	demoj := demog + demoS[demop].yys + 1

	if demoj >= demoLast {
		demostate = demoAct[demog]
	} else {
		demostate = demoAct[demoj]
		if demoChk[demostate] != -demon {
			demostate = demoAct[demog]
		}
	}
	// dummy call; replaced with literal code
	switch demont {

	case 1:
		demoDollar = demoS[demopt-1 : demopt+1]
//line parser.y:22
		{
			demolex.(*parseContext).result = demoDollar[1].Node
		}
	case 2:
		demoDollar = demoS[demopt-3 : demopt+1]
//line parser.y:28
		{
			demoVAL.Node = &Block{demoDollar[1].Node, demoDollar[2].NodeList, demoDollar[3].Node}
		}
	case 3:
		demoDollar = demoS[demopt-2 : demopt+1]
//line parser.y:34
		{
			demoVAL.NodeList = append(demoDollar[1].NodeList, demoDollar[2].Node)
		}
	case 4:
		demoDollar = demoS[demopt-0 : demopt+1]
//line parser.y:37
		{ // empty
			demoVAL.NodeList = nil
		}
	case 5:
		demoDollar = demoS[demopt-1 : demopt+1]
//line parser.y:43
		{
			demoVAL.Node = demoDollar[1].Node
		}
	case 6:
		demoDollar = demoS[demopt-3 : demopt+1]
//line parser.y:46
		{
			demoVAL.Node = &Binary{demoDollar[1].Node, demoDollar[2].Node, demoDollar[3].Node}
		}
	case 7:
		demoDollar = demoS[demopt-3 : demopt+1]
//line parser.y:49
		{
			demoVAL.Node = &Binary{demoDollar[1].Node, demoDollar[2].Node, demoDollar[3].Node}
		}
	case 8:
		demoDollar = demoS[demopt-1 : demopt+1]
//line parser.y:52
		{
			demoVAL.Node = demoDollar[1].Node
		}
	case 9:
		demoDollar = demoS[demopt-1 : demopt+1]
//line parser.y:55
		{
			demoVAL.Node = demoDollar[1].Node
		}
	}
	goto demostack /* stack new state and value */
}
