//line py.y:2
package py_cst

import __yyfmt__ "fmt"

//line py.y:3
import ()

//line py.y:9
type yySymType struct {
	yys    int
	intVal int

	token  *Token // lexer will only use Token
	tokens []*Token

	expression  Expression
	expressions *ExprList // (i) and (i,) have different meaning ...

	statement  Statement
	statements []Statement

	argument  *Argument
	arguments []*Argument
	argList   *ArgumentList

	iterator  *Iterator
	iterators []*Iterator

	// either []*Argument (Invocation), []*Subscript, or NAME (Field access)
	trailer  interface{}
	trailers []interface{}

	decorator  *Decorator
	decorators []*Decorator

	importClause  *ImportClause
	importClauses []*ImportClause

	conditionClause  *ConditionClause
	conditionClauses []*ConditionClause

	subscript  *Subscript
	subscripts []*Subscript

	withClause  *WithClause
	withClauses []*WithClause
}

const FLOAT = 57346
const INTEGER = 57347
const STRING = 57348
const NAME = 57349
const NEWLINE = 57350
const LINE_CONTINUATION = 57351
const COMMENT_NEWLINE = 57352
const INDENT = 57353
const DEDENT = 57354
const AND = 57355
const AS = 57356
const ASSERT = 57357
const BREAK = 57358
const CLASS = 57359
const CONTINUE = 57360
const DEF = 57361
const DEL = 57362
const ELIF = 57363
const ELSE = 57364
const EXCEPT = 57365
const EXEC = 57366
const FINALLY = 57367
const FOR = 57368
const FROM = 57369
const GLOBAL = 57370
const IF = 57371
const IMPORT = 57372
const IN = 57373
const IS = 57374
const LAMBDA = 57375
const NONE = 57376
const NOT = 57377
const OR = 57378
const PASS = 57379
const PRINT = 57380
const RAISE = 57381
const RETURN = 57382
const TRY = 57383
const WHILE = 57384
const WITH = 57385
const YIELD = 57386
const NOT_IN = 57387
const IS_NOT = 57388
const ADD = 57389
const ADD_ASSIGN = 57390
const AND_ASSIGN = 57391
const AND_OP = 57392
const ASSIGN = 57393
const AT = 57394
const BACK_QUOTE = 57395
const COLON = 57396
const COMMA = 57397
const DIV = 57398
const DIV_ASSIGN = 57399
const DOT = 57400
const EQUALS = 57401
const GREATER_THAN = 57402
const GT_EQ = 57403
const IDIV = 57404
const IDIV_ASSIGN = 57405
const LEFT_BRACE = 57406
const LEFT_BRACKET = 57407
const LEFT_PARENTHESIS = 57408
const LEFT_SHIFT = 57409
const LEFT_SHIFT_ASSIGN = 57410
const LESS_THAN = 57411
const LT_EQ = 57412
const MINUS = 57413
const MOD = 57414
const MOD_ASSIGN = 57415
const MULT_ASSIGN = 57416
const NOT_EQUAL = 57417
const NOT_OP = 57418
const OR_ASSIGN = 57419
const OR_OP = 57420
const POWER_ASSIGN = 57421
const RIGHT_BRACE = 57422
const RIGHT_BRACKET = 57423
const RIGHT_PARENTHESIS = 57424
const RIGHT_SHIFT = 57425
const RIGHT_SHIFT_ASSIGN = 57426
const SEMI_COLON = 57427
const STAR = 57428
const STAR_STAR = 57429
const SUB_ASSIGN = 57430
const XOR = 57431
const XOR_ASSIGN = 57432

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"FLOAT",
	"INTEGER",
	"STRING",
	"NAME",
	"NEWLINE",
	"LINE_CONTINUATION",
	"COMMENT_NEWLINE",
	"INDENT",
	"DEDENT",
	"AND",
	"AS",
	"ASSERT",
	"BREAK",
	"CLASS",
	"CONTINUE",
	"DEF",
	"DEL",
	"ELIF",
	"ELSE",
	"EXCEPT",
	"EXEC",
	"FINALLY",
	"FOR",
	"FROM",
	"GLOBAL",
	"IF",
	"IMPORT",
	"IN",
	"IS",
	"LAMBDA",
	"NONE",
	"NOT",
	"OR",
	"PASS",
	"PRINT",
	"RAISE",
	"RETURN",
	"TRY",
	"WHILE",
	"WITH",
	"YIELD",
	"NOT_IN",
	"IS_NOT",
	"ADD",
	"ADD_ASSIGN",
	"AND_ASSIGN",
	"AND_OP",
	"ASSIGN",
	"AT",
	"BACK_QUOTE",
	"COLON",
	"COMMA",
	"DIV",
	"DIV_ASSIGN",
	"DOT",
	"EQUALS",
	"GREATER_THAN",
	"GT_EQ",
	"IDIV",
	"IDIV_ASSIGN",
	"LEFT_BRACE",
	"LEFT_BRACKET",
	"LEFT_PARENTHESIS",
	"LEFT_SHIFT",
	"LEFT_SHIFT_ASSIGN",
	"LESS_THAN",
	"LT_EQ",
	"MINUS",
	"MOD",
	"MOD_ASSIGN",
	"MULT_ASSIGN",
	"NOT_EQUAL",
	"NOT_OP",
	"OR_ASSIGN",
	"OR_OP",
	"POWER_ASSIGN",
	"RIGHT_BRACE",
	"RIGHT_BRACKET",
	"RIGHT_PARENTHESIS",
	"RIGHT_SHIFT",
	"RIGHT_SHIFT_ASSIGN",
	"SEMI_COLON",
	"STAR",
	"STAR_STAR",
	"SUB_ASSIGN",
	"XOR",
	"XOR_ASSIGN",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line py.y:1845

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 307
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1297

var yyAct = [...]int{

	57, 59, 206, 270, 207, 6, 325, 6, 3, 267,
	352, 90, 5, 144, 4, 2, 406, 282, 274, 150,
	300, 320, 323, 36, 281, 286, 66, 265, 93, 94,
	149, 132, 101, 71, 95, 72, 100, 179, 133, 53,
	63, 216, 61, 165, 466, 391, 70, 167, 436, 127,
	91, 69, 389, 130, 67, 68, 97, 326, 380, 378,
	151, 121, 182, 168, 341, 97, 151, 338, 280, 181,
	180, 279, 347, 123, 126, 129, 151, 284, 187, 191,
	195, 50, 139, 288, 109, 114, 222, 164, 169, 73,
	151, 262, 76, 112, 135, 138, 213, 201, 401, 120,
	151, 186, 190, 196, 117, 349, 163, 447, 233, 113,
	111, 426, 170, 115, 182, 119, 322, 184, 214, 152,
	118, 181, 180, 220, 110, 152, 116, 92, 134, 224,
	215, 218, 218, 164, 227, 152, 321, 143, 490, 145,
	146, 307, 235, 177, 238, 145, 146, 217, 217, 152,
	219, 231, 427, 209, 210, 335, 336, 172, 228, 152,
	328, 229, 198, 174, 176, 491, 234, 175, 474, 145,
	146, 134, 246, 173, 468, 232, 243, 467, 228, 230,
	247, 271, 276, 237, 236, 107, 478, 171, 228, 228,
	292, 441, 293, 398, 385, 383, 382, 348, 342, 84,
	83, 89, 82, 255, 256, 294, 295, 257, 258, 259,
	260, 101, 303, 261, 253, 254, 263, 309, 252, 250,
	290, 251, 137, 314, 315, 339, 316, 308, 296, 81,
	65, 333, 297, 312, 291, 271, 283, 242, 305, 311,
	332, 226, 86, 225, 337, 223, 221, 304, 80, 205,
	248, 128, 475, 327, 471, 457, 455, 353, 317, 79,
	78, 77, 330, 433, 289, 423, 87, 422, 421, 343,
	344, 88, 419, 334, 417, 414, 379, 371, 351, 370,
	366, 350, 306, 239, 203, 202, 98, 283, 166, 33,
	360, 32, 361, 141, 224, 412, 345, 410, 340, 354,
	160, 161, 372, 473, 162, 49, 472, 6, 354, 377,
	97, 357, 346, 141, 97, 376, 140, 374, 356, 367,
	331, 15, 359, 141, 56, 249, 204, 283, 156, 155,
	157, 16, 390, 301, 141, 299, 381, 287, 154, 158,
	104, 479, 459, 271, 159, 384, 395, 399, 386, 276,
	387, 402, 283, 365, 405, 355, 105, 407, 211, 394,
	407, 403, 404, 354, 142, 418, 106, 400, 287, 420,
	302, 358, 439, 424, 425, 411, 368, 301, 6, 369,
	6, 388, 432, 434, 435, 413, 415, 416, 430, 428,
	431, 480, 437, 364, 363, 440, 438, 442, 443, 271,
	326, 393, 392, 319, 318, 134, 313, 278, 444, 241,
	240, 125, 449, 446, 407, 453, 454, 103, 102, 199,
	456, 99, 458, 448, 460, 461, 462, 463, 464, 452,
	273, 354, 272, 6, 298, 362, 465, 90, 324, 136,
	34, 178, 285, 212, 148, 147, 469, 266, 451, 407,
	450, 245, 244, 96, 197, 194, 185, 408, 476, 189,
	477, 60, 193, 64, 470, 75, 35, 85, 271, 271,
	124, 74, 395, 407, 407, 407, 407, 481, 488, 153,
	108, 489, 43, 492, 493, 14, 11, 13, 484, 485,
	486, 487, 271, 494, 41, 395, 496, 84, 83, 89,
	82, 8, 42, 19, 10, 429, 24, 12, 48, 39,
	33, 40, 32, 38, 22, 18, 25, 47, 20, 29,
	55, 46, 27, 54, 17, 7, 62, 81, 65, 26,
	21, 37, 52, 51, 30, 28, 31, 58, 9, 373,
	86, 375, 23, 44, 45, 56, 80, 1, 131, 0,
	0, 0, 0, 0, 0, 0, 0, 79, 78, 77,
	0, 0, 0, 0, 87, 0, 0, 0, 0, 88,
	84, 83, 89, 82, 8, 0, 0, 0, 0, 0,
	0, 48, 39, 33, 40, 32, 38, 0, 0, 0,
	47, 0, 29, 55, 46, 27, 54, 0, 0, 62,
	81, 65, 0, 21, 37, 52, 51, 30, 28, 31,
	58, 0, 0, 86, 0, 0, 0, 0, 56, 80,
	0, 0, 0, 0, 84, 83, 89, 82, 208, 0,
	79, 78, 77, 0, 0, 48, 39, 87, 40, 0,
	38, 0, 88, 0, 47, 0, 0, 55, 46, 0,
	54, 0, 0, 62, 81, 65, 0, 21, 37, 52,
	51, 0, 0, 0, 58, 0, 0, 86, 84, 83,
	89, 82, 0, 80, 0, 0, 84, 83, 89, 82,
	0, 0, 0, 0, 79, 78, 77, 0, 0, 0,
	0, 87, 0, 0, 0, 0, 88, 62, 81, 65,
	0, 0, 0, 0, 0, 62, 81, 65, 0, 0,
	0, 86, 0, 0, 0, 0, 0, 80, 0, 86,
	0, 0, 0, 0, 0, 80, 0, 0, 79, 78,
	77, 0, 0, 0, 0, 87, 79, 78, 77, 0,
	88, 0, 0, 87, 0, 0, 329, 0, 88, 0,
	268, 269, 0, 0, 264, 0, 0, 0, 268, 269,
	84, 83, 89, 82, 200, 0, 0, 0, 0, 0,
	0, 48, 39, 0, 40, 0, 38, 0, 0, 0,
	47, 0, 0, 55, 46, 0, 54, 0, 0, 62,
	81, 65, 0, 21, 37, 52, 51, 0, 0, 0,
	58, 0, 0, 86, 84, 83, 89, 82, 0, 80,
	0, 0, 84, 83, 89, 82, 0, 0, 0, 0,
	79, 78, 77, 0, 0, 0, 0, 87, 0, 0,
	0, 0, 88, 62, 81, 65, 0, 0, 0, 0,
	0, 62, 81, 65, 0, 0, 0, 86, 0, 0,
	0, 0, 0, 80, 0, 86, 0, 0, 0, 0,
	0, 80, 0, 0, 79, 78, 77, 84, 83, 89,
	82, 87, 79, 78, 77, 0, 88, 0, 0, 87,
	0, 0, 0, 0, 88, 0, 396, 397, 0, 0,
	0, 0, 0, 0, 0, 495, 62, 81, 65, 0,
	0, 0, 84, 83, 89, 82, 84, 83, 89, 82,
	86, 0, 0, 0, 0, 0, 80, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 79, 78, 77,
	0, 62, 81, 65, 87, 62, 81, 65, 0, 88,
	0, 84, 83, 89, 82, 86, 0, 0, 0, 86,
	483, 80, 0, 0, 0, 80, 84, 83, 89, 82,
	0, 0, 79, 78, 77, 0, 79, 78, 77, 87,
	62, 81, 65, 87, 88, 0, 0, 0, 88, 0,
	0, 58, 0, 0, 86, 482, 81, 0, 0, 445,
	80, 84, 83, 89, 82, 84, 83, 89, 82, 86,
	0, 79, 78, 77, 0, 80, 0, 0, 87, 0,
	0, 0, 0, 88, 0, 0, 79, 78, 77, 183,
	62, 81, 65, 87, 62, 81, 65, 0, 88, 0,
	0, 0, 0, 0, 86, 0, 0, 0, 86, 0,
	80, 0, 0, 0, 80, 84, 83, 89, 82, 0,
	0, 79, 78, 77, 0, 79, 78, 77, 87, 0,
	0, 0, 87, 88, 0, 0, 0, 88, 0, 0,
	122, 0, 0, 310, 62, 81, 65, 0, 0, 0,
	84, 83, 89, 82, 0, 0, 0, 0, 86, 0,
	0, 0, 0, 0, 80, 0, 0, 0, 0, 0,
	0, 84, 83, 89, 82, 79, 78, 77, 0, 62,
	81, 65, 87, 0, 0, 0, 0, 88, 0, 0,
	0, 0, 188, 86, 0, 0, 0, 0, 0, 80,
	62, 81, 65, 0, 0, 84, 83, 89, 82, 0,
	79, 78, 77, 0, 86, 0, 0, 87, 0, 0,
	80, 277, 88, 0, 0, 275, 192, 0, 0, 0,
	0, 79, 78, 77, 62, 81, 65, 0, 87, 84,
	83, 89, 82, 88, 0, 0, 0, 0, 86, 0,
	0, 0, 0, 0, 80, 353, 0, 0, 0, 0,
	84, 83, 89, 82, 0, 79, 78, 77, 62, 81,
	65, 0, 87, 0, 0, 0, 0, 88, 0, 58,
	0, 0, 86, 0, 0, 0, 0, 0, 80, 62,
	81, 65, 0, 0, 84, 83, 89, 82, 0, 79,
	78, 77, 0, 86, 0, 0, 87, 0, 0, 80,
	0, 88, 0, 0, 0, 0, 0, 0, 0, 0,
	79, 78, 77, 409, 81, 65, 0, 87, 0, 0,
	0, 0, 88, 0, 0, 0, 0, 86, 0, 0,
	0, 0, 0, 80, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 79, 78, 77, 0, 0, 0,
	0, 87, 0, 0, 0, 0, 88,
}
var yyPact = [...]int{

	566, -1000, 566, -1000, -1000, -1000, -1000, -1000, -1000, 42,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 1186, 1186, 952,
	232, 1186, 411, 410, 272, 134, 36, 987, 952, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 404, 952, 1186, -1000,
	196, 1186, 1186, -1000, 398, 164, 398, -1000, 1186, 287,
	-1000, 351, 83, -1000, 269, 195, 9, -46, 238, -20,
	41, 101, -1000, -1000, 952, 56, -1000, 937, 1041, 1076,
	1186, -1000, -1000, -1000, -1000, 413, -1000, -1000, -1000, -1000,
	-1000, -1000, 756, 231, 230, 295, 194, 9, 620, 99,
	-1000, 344, 30, 64, -1000, -1000, -1000, 1165, 1165, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 1186, -1000, 191, -1000, 55, 190, 1186, -1000,
	188, 186, -1000, 120, -1000, 131, 121, -1000, 100, -1000,
	195, 195, 195, 1186, 229, 403, 402, -1000, 182, -1000,
	125, -1000, 93, 952, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 215, 294, -1000, 952, 952, 952, 952, 952, 952,
	952, 952, 952, 952, 952, -1000, -1000, 952, 4, -1000,
	672, 1097, 400, -1000, -11, -14, -1000, 301, -1000, -4,
	-1000, 311, -1000, 3, -1000, 210, -1000, 179, 137, -1000,
	-1000, -1000, 620, 620, 1186, 952, 310, -1000, 359, 620,
	1186, 952, 228, 59, 620, 991, -1000, -1000, -1000, -1000,
	178, 399, 1186, 1186, -1000, 1186, 398, 397, 396, 50,
	-1000, 50, 130, -1000, 664, 298, 351, -1000, -1000, 1186,
	176, -1000, 69, 1186, -15, 170, -1000, 9, -1000, -1000,
	-46, 238, -20, 41, 41, 101, 101, -1000, -1000, -1000,
	-1000, -1000, 952, -1000, -1000, -18, -1000, 143, 1186, 1186,
	-1000, 261, -9, 142, -1000, 47, 227, 1131, -1000, -1000,
	-1000, 326, -1000, 952, -1000, 342, -1000, 952, -1000, 1186,
	326, 1186, -1000, 1186, 372, 331, 226, 9, 354, 225,
	223, 1186, 566, -1000, -1000, 9, 620, -1000, -23, -1000,
	222, -24, 1186, -1000, 141, -1000, 140, -1000, -1000, -1000,
	-1000, -1000, 393, -1000, 139, -1000, 334, -1000, 50, 373,
	-30, 1186, -1000, -42, -1000, 395, 394, -1000, -1000, 93,
	-1000, -1000, 800, 138, -1000, 326, 1186, -1000, 1097, 40,
	1131, 203, -1000, 1186, -1000, 1220, 266, -1000, 1220, 264,
	301, 221, 365, 220, 1186, 218, 620, 214, 213, 211,
	620, 620, 97, 493, -1000, 566, -1000, -1000, -1000, 620,
	209, -1000, 1186, 1186, -34, 393, 389, -1000, -1000, 364,
	-1000, 388, 136, -1000, -1000, -1000, 1186, 1186, 902, -1000,
	-1000, -1000, 203, -1000, -1000, -1000, -1000, 257, -1000, 53,
	195, -1000, 1220, 326, 1186, 1186, 202, 620, 201, 620,
	320, 620, 620, 620, -1000, -1000, 1186, 1186, 566, -1000,
	-1000, -1000, -1000, 620, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -43, 122, -1000, 119, 1186, -1000, 1220, 200, 277,
	274, 113, -1000, -1000, 198, 620, -1000, 620, -1000, 132,
	-1000, 316, -1000, -1000, -1000, -1000, 384, 898, 863, -1000,
	-1000, 1220, 1220, 1220, 1220, 620, -1000, -1000, 620, 84,
	-1000, 110, 1186, 1186, -1000, -1000, -1000, -1000, -1000, -1000,
	620, 808, -1000, -1000, -1000, 1186, -1000,
}
var yyPgo = [...]int{

	0, 548, 547, 544, 543, 542, 541, 539, 4, 538,
	14, 8, 15, 2, 529, 331, 525, 524, 518, 31,
	516, 515, 514, 507, 321, 506, 504, 12, 503, 502,
	494, 487, 486, 485, 482, 480, 479, 92, 38, 471,
	470, 467, 55, 46, 466, 465, 463, 462, 26, 35,
	19, 461, 459, 457, 16, 1, 42, 40, 89, 51,
	33, 0, 456, 54, 39, 41, 455, 454, 34, 453,
	452, 451, 450, 448, 23, 81, 3, 30, 447, 9,
	27, 13, 445, 444, 443, 25, 17, 442, 24, 37,
	441, 305, 440, 439, 6, 21, 22, 438, 20, 435,
	434, 10, 18, 432, 430, 36, 421,
}
var yyR1 = [...]int{

	0, 2, 2, 27, 11, 11, 12, 12, 6, 6,
	7, 7, 7, 91, 91, 91, 92, 92, 17, 17,
	24, 84, 84, 77, 77, 83, 83, 82, 82, 81,
	81, 81, 81, 81, 81, 81, 50, 50, 71, 71,
	70, 70, 10, 10, 8, 8, 8, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 44, 44, 65, 65,
	21, 21, 35, 35, 35, 35, 35, 35, 35, 35,
	35, 35, 35, 35, 28, 28, 28, 28, 18, 22,
	22, 22, 22, 22, 30, 30, 34, 29, 29, 29,
	29, 5, 5, 4, 95, 95, 95, 93, 93, 3,
	3, 3, 94, 94, 19, 19, 97, 97, 96, 96,
	1, 1, 38, 38, 40, 40, 25, 20, 20, 20,
	14, 14, 16, 16, 16, 16, 16, 16, 16, 16,
	99, 99, 26, 26, 26, 26, 32, 32, 23, 23,
	100, 100, 31, 31, 31, 31, 31, 33, 106, 106,
	105, 105, 98, 98, 98, 98, 13, 13, 54, 54,
	53, 53, 61, 61, 61, 55, 55, 56, 56, 57,
	57, 46, 46, 36, 36, 36, 36, 36, 36, 36,
	36, 36, 36, 48, 48, 63, 63, 42, 42, 59,
	59, 59, 43, 43, 43, 60, 60, 60, 60, 60,
	37, 37, 37, 39, 39, 49, 49, 58, 58, 58,
	58, 90, 90, 45, 45, 45, 45, 45, 45, 45,
	45, 45, 45, 45, 45, 45, 41, 41, 52, 52,
	62, 62, 51, 51, 89, 89, 89, 89, 104, 104,
	103, 103, 102, 102, 102, 102, 102, 102, 102, 102,
	102, 102, 101, 101, 73, 73, 72, 72, 69, 69,
	68, 68, 75, 75, 74, 74, 67, 67, 66, 66,
	47, 47, 47, 47, 15, 15, 15, 79, 79, 78,
	78, 80, 80, 80, 80, 80, 80, 80, 80, 80,
	80, 80, 76, 76, 76, 85, 85, 87, 87, 87,
	86, 86, 88, 88, 88, 64, 64,
}
var yyR2 = [...]int{

	0, 0, 1, 1, 1, 1, 1, 2, 1, 2,
	1, 2, 2, 3, 5, 6, 1, 2, 2, 2,
	5, 2, 3, 1, 3, 1, 3, 1, 2, 2,
	2, 5, 1, 4, 4, 7, 1, 3, 1, 3,
	1, 2, 1, 1, 2, 3, 3, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 3, 1, 1,
	1, 3, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 3, 5, 2, 1,
	1, 1, 1, 1, 1, 2, 1, 1, 2, 4,
	6, 1, 1, 2, 1, 3, 1, 1, 2, 4,
	4, 5, 1, 3, 1, 3, 1, 3, 1, 2,
	1, 3, 1, 3, 1, 3, 2, 2, 4, 6,
	2, 4, 1, 1, 1, 1, 1, 1, 1, 1,
	4, 5, 4, 5, 8, 7, 4, 7, 6, 9,
	3, 4, 4, 7, 7, 10, 6, 4, 1, 3,
	1, 3, 1, 2, 4, 4, 1, 4, 1, 1,
	3, 4, 1, 5, 1, 1, 3, 1, 3, 1,
	2, 1, 3, 1, 1, 1, 1, 1, 1, 1,
	1, 2, 2, 1, 3, 1, 3, 1, 3, 1,
	3, 3, 1, 3, 3, 1, 3, 3, 3, 3,
	1, 1, 1, 1, 2, 1, 2, 1, 3, 2,
	4, 1, 2, 2, 3, 3, 2, 3, 2, 3,
	3, 1, 1, 1, 1, 1, 1, 2, 1, 2,
	1, 2, 3, 4, 2, 3, 3, 2, 1, 3,
	1, 2, 3, 1, 1, 2, 2, 3, 2, 3,
	3, 4, 1, 2, 1, 3, 1, 2, 1, 3,
	1, 2, 1, 3, 1, 2, 3, 5, 1, 2,
	1, 4, 1, 2, 4, 6, 7, 1, 3, 1,
	2, 1, 4, 6, 7, 9, 4, 2, 4, 5,
	7, 2, 1, 2, 3, 4, 6, 1, 2, 3,
	4, 6, 1, 2, 3, 1, 2,
}
var yyChk = [...]int{

	-1000, -2, -12, -11, -10, -27, -8, -16, 8, -9,
	-26, -32, -23, -31, -33, -24, -15, -17, -21, -28,
	-18, 37, -22, -5, -25, -20, -14, 29, 42, 26,
	41, 43, 19, 17, -92, -44, -74, 38, 20, 16,
	18, -30, -29, -34, -4, -3, 28, 24, 15, -91,
	-75, 40, 39, -64, 30, 27, 52, -61, 44, -55,
	-51, -56, 33, -57, -46, 35, -48, -63, -42, -59,
	-43, -60, -49, -58, -39, -45, -37, 66, 65, 64,
	53, 34, 7, 5, 4, -41, 47, 71, 76, 6,
	-11, 8, 85, -61, -61, -68, -69, -48, 54, -106,
	-105, -61, 7, 7, -91, -24, -15, 51, -35, 48,
	88, 74, 57, 73, 49, 77, 90, 68, 84, 79,
	63, -74, 83, -68, -40, 7, -48, -61, 55, -74,
	-61, -1, -19, -38, 7, -38, -93, 58, -38, -74,
	29, 36, 13, 54, -81, 86, 87, -82, -83, -77,
	-50, 7, 66, -36, 69, 60, 59, 61, 70, 75,
	31, 32, 35, -57, 78, 89, 50, 67, 83, 47,
	71, 86, 56, 72, 62, -37, -58, 87, -90, -89,
	66, 65, 58, 82, -64, -62, -74, -61, 81, -52,
	-74, -61, 80, -47, -66, -61, -74, -67, -75, 6,
	8, -8, 54, 54, 31, 55, -13, -8, 8, 54,
	55, 14, -84, 66, 54, 66, -65, -64, -74, -65,
	-61, 55, 31, 55, -61, 55, 55, 14, 58, 30,
	58, 30, -38, 8, 66, -55, -56, -57, -61, 54,
	7, 7, 55, 51, -70, -71, -50, -48, 35, 31,
	-63, -42, -59, -43, -43, -60, -60, -49, -49, -49,
	-49, -49, 87, -89, 82, -80, -78, -79, 86, 87,
	-76, -61, -103, -104, -102, 58, -61, 54, 7, 82,
	82, -88, -86, 26, 81, -87, -85, 26, 80, 54,
	-88, 55, 53, 55, -13, -13, -74, -48, -100, 25,
	-98, 23, 11, -13, -105, -48, 54, 82, -81, -13,
	82, -74, 55, 7, -61, -61, -61, -19, 7, 7,
	-95, 86, 66, -96, -97, -94, 7, -95, 30, 82,
	-80, 22, -61, 55, -77, 86, 87, -61, 82, 55,
	-49, 82, 55, -61, -61, -88, 51, 81, 55, 58,
	54, -61, -101, 54, -86, 29, -68, -85, 29, -68,
	-61, -61, -99, 22, 21, 22, 54, -98, 22, 25,
	54, 54, -61, -7, -10, -6, -27, -13, 82, 54,
	82, -74, 55, 55, -96, 55, 14, -95, 8, 82,
	-61, 87, 7, 7, -50, -76, 86, 87, 55, -61,
	-102, 58, -61, -101, -101, -61, -54, -55, -53, 33,
	31, -54, 31, -88, 54, 21, 22, 54, -61, 54,
	-13, 54, 54, 54, -13, -13, 14, 55, -12, 12,
	-27, -10, -13, 54, -61, -61, 82, -94, 7, 8,
	7, 55, -61, -61, -79, 87, -101, 54, -81, -55,
	-72, -73, -54, -61, -61, 54, -13, 54, -13, 22,
	-13, -13, -13, -61, -61, -13, 87, 55, 55, -61,
	-54, 54, 29, 29, 55, 54, -13, -13, 54, 25,
	7, -79, 87, 87, -54, -54, -54, -54, -13, -13,
	54, 55, -61, -61, -13, 87, -61,
}
var yyDef = [...]int{

	1, -2, 2, 6, 4, 5, 42, 43, 3, 0,
	122, 123, 124, 125, 126, 127, 128, 129, 47, 48,
	49, 50, 51, 52, 53, 54, 55, 0, 0, 0,
	0, 0, 0, 0, 0, 60, 56, 74, 0, 79,
	80, 81, 82, 83, 91, 92, 0, 0, 0, 16,
	264, 84, 87, 86, 0, 0, 0, 262, 305, 162,
	164, 165, 0, 167, 169, 0, 171, 183, 185, 187,
	189, 192, 195, 205, 0, 207, 203, 0, 0, 0,
	0, 221, 222, 223, 224, 225, 200, 201, 202, 226,
	7, 44, 0, 0, 0, 0, 260, 258, 0, 0,
	148, 150, 0, 0, 17, 18, 19, 0, 0, 62,
	63, 64, 65, 66, 67, 68, 69, 70, 71, 72,
	73, 75, 0, 78, 116, 114, 117, 120, 265, 85,
	88, 93, 110, 104, 112, 0, 0, 97, 0, 306,
	0, 0, 0, 0, 0, 0, 0, 32, 27, 25,
	23, 36, 0, 0, 173, 174, 175, 176, 177, 178,
	179, 180, 0, 170, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 204, 206, 0, 209, 211,
	0, 0, 0, 213, 0, 0, 230, 262, 216, 0,
	228, 262, 218, 0, 270, 262, 272, 268, 0, 227,
	45, 46, 0, 0, 0, 261, 0, 156, 0, 0,
	0, 0, 0, 0, 0, 0, 57, 58, 59, 61,
	76, 0, 0, 0, 263, 0, 0, 0, 0, 0,
	98, 0, 0, 13, 0, 0, 166, 168, 232, 0,
	29, 30, 28, 0, 0, 40, 38, 172, 182, 181,
	184, 186, 188, 190, 191, 193, 194, 196, 197, 198,
	199, 208, 0, 212, 234, 0, 281, 279, 0, 0,
	277, 292, 0, 240, 238, 0, 243, 244, 237, 214,
	215, 231, 302, 0, 217, 229, 297, 0, 219, 0,
	273, 269, 220, 0, 132, 136, 0, 259, 142, 0,
	0, 152, 0, 147, 149, 151, 0, 21, 0, 274,
	0, 0, 0, 115, 118, 121, 89, 111, 105, 113,
	99, 94, 0, 96, 108, 106, 102, 100, 0, 0,
	0, 0, 233, 0, 26, 0, 0, 24, 37, 41,
	210, 235, 280, 287, 291, 293, 0, 236, 241, 0,
	245, 246, 248, 252, 303, 0, 0, 298, 0, 0,
	266, 0, 133, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 153, 0, 10, 0, 8, 20, 22, 0,
	0, 77, 0, 0, 0, 109, 0, 101, 14, 0,
	163, 0, 33, 34, 39, 278, 0, 0, 0, 294,
	239, 242, 247, 249, 250, 253, 304, 158, 159, 0,
	0, 299, 0, 271, 0, 0, 0, 0, 0, 0,
	138, 0, 0, 0, 146, 140, 0, 0, 12, 157,
	9, 11, 275, 0, 119, 90, 95, 107, 103, 15,
	31, 0, 282, 286, 288, 0, 251, 0, 0, 300,
	295, 256, 254, 267, 0, 0, 135, 0, 137, 0,
	141, 143, 144, 154, 155, 276, 0, 0, 0, 289,
	160, 0, 0, 0, 257, 0, 134, 130, 0, 0,
	35, 283, 0, 0, 161, 301, 296, 255, 131, 139,
	0, 0, 284, 290, 145, 0, 285,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73, 74, 75, 76, 77, 78, 79, 80, 81,
	82, 83, 84, 85, 86, 87, 88, 89, 90,
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
		yyDollar = yyS[yypt-0 : yypt+1]
		//line py.y:279
		{
			// empty file is ok
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:282
		{
			// %parser-param is not supported =...(
			yylex.(*Context).Statements = yyDollar[1].statements
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:289
		{
			yyVAL.statement = NewPassStmt(yyDollar[1].token, true)
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:295
		{
			yyVAL.statements = yyDollar[1].statements
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:298
		{
			yyVAL.statements = []Statement{yyDollar[1].statement}
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:304
		{
			yyVAL.statements = yyDollar[1].statements
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:307
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statements...)
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:312
		{
			yyVAL.statements = []Statement{yyDollar[1].statement}
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:315
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:321
		{
			yyVAL.statements = yyDollar[1].statements
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:324
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statements...)
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:327
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statements...)
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:333
		{
			yyVAL.decorator = NewDecorator(yyDollar[1].token, DottedNameToExpr(yyDollar[2].tokens), yyDollar[3].token)
		}
	case 14:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line py.y:336
		{
			call := NewCallExpr(DottedNameToExpr(yyDollar[2].tokens), NewArgumentList(yyDollar[3].token, nil, yyDollar[4].token))

			yyVAL.decorator = NewDecorator(yyDollar[1].token, call, yyDollar[5].token)
		}
	case 15:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line py.y:341
		{
			call := NewCallExpr(DottedNameToExpr(yyDollar[2].tokens), NewArgumentList(yyDollar[3].token, yyDollar[4].arguments, yyDollar[5].token))

			yyVAL.decorator = NewDecorator(yyDollar[1].token, call, yyDollar[6].token)
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:349
		{
			yyVAL.decorators = []*Decorator{yyDollar[1].decorator}
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:352
		{
			yyVAL.decorators = append(yyDollar[1].decorators, yyDollar[2].decorator)
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:358
		{
			yyDollar[2].statement.(*FuncDef).SetDecorators(yyDollar[1].decorators)
			yyVAL.statement = yyDollar[2].statement
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:362
		{
			yyDollar[2].statement.(*ClassDef).SetDecorators(yyDollar[1].decorators)
			yyVAL.statement = yyDollar[2].statement
		}
	case 20:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line py.y:369
		{
			yyVAL.statement = NewFuncDef(yyDollar[1].token, yyDollar[2].token, yyDollar[3].argList, yyDollar[4].token, yyDollar[5].statements)
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:375
		{
			yyVAL.argList = NewArgumentList(yyDollar[1].token, nil, yyDollar[2].token)
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:378
		{
			yyVAL.argList = NewArgumentList(yyDollar[1].token, yyDollar[2].arguments, yyDollar[3].token)
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:384
		{
			yyVAL.argument = NewArgument(yyDollar[1].expression, nil, nil)
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:387
		{
			yyVAL.argument = NewArgument(yyDollar[1].expression, yyDollar[2].token, yyDollar[3].expression)
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:393
		{
			yyVAL.arguments = []*Argument{yyDollar[1].argument}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:396
		{
			yyDollar[1].arguments[len(yyDollar[1].arguments)-1].MergeFrom(&yyDollar[2].token.Node)
			yyVAL.arguments = append(yyDollar[1].arguments, yyDollar[3].argument)
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:403
		{
			yyVAL.arguments = yyDollar[1].arguments
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:406
		{
			yyVAL.arguments = yyDollar[1].arguments
			yyVAL.arguments[len(yyVAL.arguments)-1].NodeInfo().MergeFrom(&yyDollar[2].token.Node)
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:413
		{
			yyVAL.arguments = []*Argument{NewPositionVarParam(yyDollar[1].token, NewIdentifier(yyDollar[2].token))}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:416
		{
			yyVAL.arguments = []*Argument{NewKeywordVarParam(yyDollar[1].token, NewIdentifier(yyDollar[2].token))}
		}
	case 31:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line py.y:419
		{
			arg := NewPositionVarParam(yyDollar[1].token, NewIdentifier(yyDollar[2].token))
			arg.MergeFrom(&yyDollar[3].token.Node)

			yyVAL.arguments = []*Argument{arg, NewKeywordVarParam(yyDollar[4].token, NewIdentifier(yyDollar[5].token))}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:425
		{
			yyVAL.arguments = yyDollar[1].arguments
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line py.y:428
		{
			yyDollar[1].arguments[len(yyDollar[1].arguments)-1].MergeFrom(&yyDollar[2].token.Node)

			yyVAL.arguments = append(yyDollar[1].arguments, NewPositionVarParam(yyDollar[3].token, NewIdentifier(yyDollar[4].token)))
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line py.y:433
		{
			yyDollar[1].arguments[len(yyDollar[1].arguments)-1].MergeFrom(&yyDollar[2].token.Node)

			yyVAL.arguments = append(yyDollar[1].arguments, NewKeywordVarParam(yyDollar[3].token, NewIdentifier(yyDollar[4].token)))
		}
	case 35:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line py.y:438
		{
			yyDollar[1].arguments[len(yyDollar[1].arguments)-1].MergeFrom(&yyDollar[2].token.Node)

			arg := NewPositionVarParam(yyDollar[3].token, NewIdentifier(yyDollar[4].token))
			arg.MergeFrom(&yyDollar[5].token.Node)

			yyVAL.arguments = append(yyDollar[1].arguments, arg)
			yyVAL.arguments = append(yyDollar[1].arguments, NewKeywordVarParam(yyDollar[6].token, NewIdentifier(yyDollar[7].token)))
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:450
		{
			yyVAL.expression = NewIdentifier(yyDollar[1].token)
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:453
		{
			yyVAL.expression = yyDollar[2].expressions.ConvertToExpr()

			yyDollar[1].token.MergeFrom(yyVAL.expression.NodeInfo())
			yyDollar[1].token.MergeFrom(&yyDollar[3].token.Node)

			*yyVAL.expression.NodeInfo() = yyDollar[1].token.Node
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:464
		{
			yyVAL.expressions = NewExprList([]Expression{yyDollar[1].expression})
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:467
		{
			yyVAL.expressions = yyDollar[1].expressions
			yyVAL.expressions.Expressions[len(yyVAL.expressions.Expressions)-1].NodeInfo().MergeFrom(&yyDollar[2].token.Node)
			yyVAL.expressions.Expressions = append(yyVAL.expressions.Expressions, yyDollar[3].expression)
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:475
		{
			yyVAL.expressions = yyDollar[1].expressions
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:478
		{
			yyVAL.expressions = yyDollar[1].expressions
			yyVAL.expressions.Expressions[len(yyVAL.expressions.Expressions)-1].NodeInfo().MergeFrom(&yyDollar[2].token.Node)
			yyVAL.expressions.ExplicitCollection = true
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:486
		{
			yyVAL.statements = yyDollar[1].statements
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:489
		{
			yyVAL.statements = []Statement{yyDollar[1].statement}
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:495
		{
			yyVAL.statements = yyDollar[1].statements
			yyVAL.statements[len(yyVAL.statements)-1].NodeInfo().MergeTrailingFrom(&yyDollar[2].token.Node)
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:499
		{
			yyVAL.statements = yyDollar[1].statements
			yyVAL.statements[len(yyVAL.statements)-1].NodeInfo().MergeTrailingFrom(&yyDollar[2].token.Node)
			yyVAL.statements[len(yyVAL.statements)-1].NodeInfo().MergeTrailingFrom(&yyDollar[3].token.Node)
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:504
		{
			yyDollar[1].statements[len(yyDollar[1].statements)-1].NodeInfo().MergeTrailingFrom(&yyDollar[2].token.Node)
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[3].statements...)
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:510
		{
			yyVAL.statements = []Statement{yyDollar[1].statement}
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:513
		{
			yyVAL.statements = []Statement{yyDollar[1].statement}
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:516
		{
			yyVAL.statements = []Statement{yyDollar[1].statement}
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:519
		{
			yyVAL.statements = []Statement{NewPassStmt(yyDollar[1].token, false)}
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:522
		{
			yyVAL.statements = []Statement{yyDollar[1].statement}
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:525
		{
			yyVAL.statements = yyDollar[1].statements
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:528
		{
			yyVAL.statements = []Statement{yyDollar[1].statement}
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:531
		{
			yyVAL.statements = []Statement{yyDollar[1].statement}
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:534
		{
			yyVAL.statements = []Statement{yyDollar[1].statement}
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:541
		{
			yyVAL.expression = yyDollar[1].expressions.ConvertToExpr()
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:544
		{
			// Assign is technically an statement
			yyVAL.expression = NewBinaryExpr(yyDollar[1].expression, yyDollar[2].token, yyDollar[3].expression)
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:551
		{
			yyVAL.expression = yyDollar[1].expression
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:554
		{
			yyVAL.expression = yyDollar[1].expressions.ConvertToExpr()
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:560
		{
			yyVAL.statement = &ExprStmt{Expression: yyDollar[1].expression}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:563
		{
			// Augassign is technically an statement
			yyVAL.statement = &ExprStmt{Expression: NewBinaryExpr(yyDollar[1].expressions.ConvertToExpr(), yyDollar[2].token, yyDollar[3].expression)}
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:570
		{
			yyVAL.token = yyDollar[1].token
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:573
		{
			yyVAL.token = yyDollar[1].token
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:576
		{
			yyVAL.token = yyDollar[1].token
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:579
		{
			yyVAL.token = yyDollar[1].token
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:582
		{
			yyVAL.token = yyDollar[1].token
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:585
		{
			yyVAL.token = yyDollar[1].token
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:588
		{
			yyVAL.token = yyDollar[1].token
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:591
		{
			yyVAL.token = yyDollar[1].token
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:594
		{
			yyVAL.token = yyDollar[1].token
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:597
		{
			yyVAL.token = yyDollar[1].token
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:600
		{
			yyVAL.token = yyDollar[1].token
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:603
		{
			yyVAL.token = yyDollar[1].token
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:609
		{
			yyVAL.statement = NewPrintStmt(yyDollar[1].token, nil, nil, nil)
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:612
		{
			yyVAL.statement = NewPrintStmt(yyDollar[1].token, nil, nil, yyDollar[2].expressions)
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:615
		{
			yyVAL.statement = NewPrintStmt(yyDollar[1].token, yyDollar[2].token, yyDollar[3].expression, nil)
		}
	case 77:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line py.y:618
		{
			yyVAL.statement = NewPrintStmt(yyDollar[1].token, yyDollar[2].token, yyDollar[3].expression, yyDollar[5].expressions)
			yyVAL.statement.NodeInfo().MergeFrom(&yyDollar[4].token.Node)
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:625
		{
			yyVAL.statement = NewDelStmt(yyDollar[1].token, yyDollar[2].expressions)
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:631
		{
			yyVAL.statement = NewBreakStmt(yyDollar[1].token)
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:634
		{
			yyVAL.statement = NewContinueStmt(yyDollar[1].token)
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:637
		{
			yyVAL.statement = yyDollar[1].statement
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:640
		{
			yyVAL.statement = yyDollar[1].statement
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:643
		{
			yyVAL.statement = yyDollar[1].statement
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:649
		{
			yyVAL.statement = NewReturnStmt(yyDollar[1].token, nil)
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:652
		{
			yyVAL.statement = NewReturnStmt(yyDollar[1].token, yyDollar[2].expressions.ConvertToExpr())
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:658
		{
			yyVAL.statement = NewExprStmt(yyDollar[1].expression)
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:664
		{
			yyVAL.statement = NewRaiseStmt(yyDollar[1].token, nil, nil, nil, nil, nil)
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:667
		{
			yyVAL.statement = NewRaiseStmt(yyDollar[1].token, yyDollar[2].expression, nil, nil, nil, nil)
		}
	case 89:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line py.y:670
		{
			yyVAL.statement = NewRaiseStmt(yyDollar[1].token, yyDollar[2].expression, yyDollar[3].token, yyDollar[4].expression, nil, nil)
		}
	case 90:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line py.y:673
		{
			yyVAL.statement = NewRaiseStmt(yyDollar[1].token, yyDollar[2].expression, yyDollar[3].token, yyDollar[4].expression, yyDollar[5].token, yyDollar[6].expression)
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:679
		{
			yyVAL.statements = yyDollar[1].statements
		}
	case 92:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:682
		{
			yyVAL.statements = yyDollar[1].statements
		}
	case 93:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:688
		{
			yyDollar[2].statements[0].NodeInfo().MergeLeadingFrom(&yyDollar[1].token.Node)
			yyVAL.statements = yyDollar[2].statements

			//TODO PRESERVE COMMENT FOR BELOW RULES ...

		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:704
		{
			yyVAL.importClauses = []*ImportClause{} // Empty list implies all
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:707
		{
			yyVAL.importClauses = yyDollar[2].importClauses
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:710
		{
			yyVAL.importClauses = yyDollar[1].importClauses
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:716
		{
			yyVAL.intVal = 1
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:719
		{
			yyVAL.intVal = yyDollar[1].intVal + 1
		}
	case 99:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line py.y:725
		{
			yyVAL.statements = []Statement{
				&FromStmt{
					DotPrefixCount: 0,
					ModulePath:     yyDollar[2].tokens,
					Imports:        yyDollar[4].importClauses,
				},
			}
		}
	case 100:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line py.y:734
		{
			yyVAL.statements = []Statement{
				&FromStmt{
					DotPrefixCount: yyDollar[2].intVal,
					Imports:        yyDollar[4].importClauses,
				},
			}
		}
	case 101:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line py.y:742
		{
			yyVAL.statements = []Statement{
				&FromStmt{
					DotPrefixCount: yyDollar[2].intVal,
					ModulePath:     yyDollar[3].tokens,
					Imports:        yyDollar[5].importClauses,
				},
			}
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:754
		{
			yyVAL.importClause = &ImportClause{
				Name: yyDollar[1].token,
			}
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:759
		{
			yyVAL.importClause = &ImportClause{
				Name:  yyDollar[1].token,
				Alias: yyDollar[2].token,
			}
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:768
		{
			yyVAL.statement = &ImportStmt{ModulePath: yyDollar[1].tokens}
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:771
		{
			yyVAL.statement = &ImportStmt{ModulePath: yyDollar[1].tokens, Alias: yyDollar[2].token}
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:777
		{
			yyVAL.importClauses = []*ImportClause{yyDollar[1].importClause}
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:780
		{
			yyVAL.importClauses = append(yyDollar[1].importClauses, yyDollar[3].importClause)
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:786
		{
			yyVAL.importClauses = yyDollar[1].importClauses
		}
	case 109:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:789
		{
			yyVAL.importClauses = yyDollar[1].importClauses
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:795
		{
			yyVAL.statements = []Statement{yyDollar[1].statement}
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:798
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[3].statement)
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:804
		{
			yyVAL.tokens = []*Token{yyDollar[1].token}
		}
	case 113:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:807
		{
			yyVAL.tokens = append(yyDollar[1].tokens, yyDollar[3].token)
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:813
		{
			yyVAL.tokens = []*Token{yyDollar[1].token}
		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:816
		{
			yyVAL.tokens = append(yyDollar[1].tokens, yyDollar[3].token)
		}
	case 116:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:821
		{
			yyVAL.statement = &GlobalStmt{Names: yyDollar[2].tokens}
		}
	case 117:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:827
		{
			yyVAL.statement = &ExecStmt{Expr: yyDollar[2].expression}
		}
	case 118:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line py.y:830
		{
			yyVAL.statement = &ExecStmt{Expr: yyDollar[2].expression, Global: yyDollar[4].expression}
		}
	case 119:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line py.y:833
		{
			yyVAL.statement = &ExecStmt{Expr: yyDollar[2].expression, Global: yyDollar[4].expression, Local: yyDollar[6].expression}
		}
	case 120:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:839
		{
			yyVAL.statement = &AssertStmt{
				Expr: yyDollar[2].expression,
			}
		}
	case 121:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line py.y:844
		{
			yyVAL.statement = &AssertStmt{
				Expr:  yyDollar[2].expression,
				Debug: yyDollar[4].expression,
			}
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:853
		{
			yyVAL.statement = yyDollar[1].statement
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:856
		{
			yyVAL.statement = yyDollar[1].statement
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:859
		{
			yyVAL.statement = yyDollar[1].statement
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:862
		{
			yyVAL.statement = yyDollar[1].statement
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:865
		{
			yyVAL.statement = yyDollar[1].statement
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:868
		{
			yyVAL.statement = yyDollar[1].statement
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:871
		{
			yyVAL.statement = yyDollar[1].statement
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:874
		{
			yyVAL.statement = yyDollar[1].statement
		}
	case 130:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line py.y:880
		{
			yyVAL.conditionClauses = []*ConditionClause{
				&ConditionClause{
					Matches: yyDollar[2].expression,
					Branch:  yyDollar[4].statements,
				},
			}
		}
	case 131:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line py.y:888
		{
			yyVAL.conditionClauses = append(
				yyDollar[1].conditionClauses,
				&ConditionClause{
					Matches: yyDollar[3].expression,
					Branch:  yyDollar[5].statements,
				})
		}
	case 132:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line py.y:899
		{
			yyVAL.statement = &ConditionStmt{
				If: &ConditionClause{
					Matches: yyDollar[2].expression,
					Branch:  yyDollar[4].statements,
				},
			}
		}
	case 133:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line py.y:907
		{
			yyVAL.statement = &ConditionStmt{
				If: &ConditionClause{
					Matches: yyDollar[2].expression,
					Branch:  yyDollar[4].statements,
				},
				Elif: yyDollar[5].conditionClauses,
			}
		}
	case 134:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line py.y:916
		{
			yyVAL.statement = &ConditionStmt{
				If: &ConditionClause{
					Matches: yyDollar[2].expression,
					Branch:  yyDollar[4].statements,
				},
				Elif: yyDollar[5].conditionClauses,
				Else: yyDollar[8].statements,
			}
		}
	case 135:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line py.y:926
		{
			yyVAL.statement = &ConditionStmt{
				If: &ConditionClause{
					Matches: yyDollar[2].expression,
					Branch:  yyDollar[4].statements,
				},
				Else: yyDollar[7].statements,
			}
		}
	case 136:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line py.y:938
		{
			yyVAL.statement = &WhileStmt{
				Loop: &ConditionClause{
					Matches: yyDollar[2].expression,
					Branch:  yyDollar[4].statements,
				},
			}
		}
	case 137:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line py.y:946
		{
			yyVAL.statement = &WhileStmt{
				Loop: &ConditionClause{
					Matches: yyDollar[2].expression,
					Branch:  yyDollar[4].statements,
				},
				Else: yyDollar[7].statements,
			}
		}
	case 138:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line py.y:958
		{
			yyVAL.statement = &ForStmt{
				Iterator: &Iterator{
					BoundVariables: yyDollar[2].expressions.ConvertToExpr(),
					Source:         yyDollar[4].expressions.ConvertToExpr(),
				},
				Loop: yyDollar[6].statements,
			}
		}
	case 139:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line py.y:967
		{
			yyVAL.statement = &ForStmt{
				Iterator: &Iterator{
					BoundVariables: yyDollar[2].expressions.ConvertToExpr(),
					Source:         yyDollar[4].expressions.ConvertToExpr(),
				},
				Loop: yyDollar[6].statements,
				Else: yyDollar[9].statements,
			}
		}
	case 140:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:980
		{
			yyDollar[1].conditionClause.Branch = yyDollar[3].statements
			yyVAL.conditionClauses = []*ConditionClause{yyDollar[1].conditionClause}
		}
	case 141:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line py.y:984
		{
			yyDollar[2].conditionClause.Branch = yyDollar[4].statements
			yyVAL.conditionClauses = append(yyDollar[1].conditionClauses, yyDollar[2].conditionClause)
		}
	case 142:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line py.y:990
		{
			yyVAL.statement = &TryStmt{
				Try:    yyDollar[3].statements,
				Except: yyDollar[4].conditionClauses,
			}
		}
	case 143:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line py.y:996
		{
			yyVAL.statement = &TryStmt{
				Try:    yyDollar[3].statements,
				Except: yyDollar[4].conditionClauses,
				Else:   yyDollar[7].statements,
			}
		}
	case 144:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line py.y:1003
		{
			yyVAL.statement = &TryStmt{
				Try:     yyDollar[3].statements,
				Except:  yyDollar[4].conditionClauses,
				Finally: yyDollar[7].statements,
			}
		}
	case 145:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line py.y:1010
		{
			yyVAL.statement = &TryStmt{
				Try:     yyDollar[3].statements,
				Except:  yyDollar[4].conditionClauses,
				Else:    yyDollar[7].statements,
				Finally: yyDollar[10].statements,
			}
		}
	case 146:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line py.y:1018
		{
			yyVAL.statement = &TryStmt{
				Try:     yyDollar[3].statements,
				Finally: yyDollar[6].statements,
			}
		}
	case 147:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line py.y:1027
		{
			yyVAL.statement = &WithStmt{
				WithClauses: yyDollar[2].withClauses,
				Statements:  yyDollar[4].statements,
			}
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1036
		{
			yyVAL.withClauses = []*WithClause{yyDollar[1].withClause}
		}
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1039
		{
			yyVAL.withClauses = append(yyDollar[1].withClauses, yyDollar[3].withClause)
		}
	case 150:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1045
		{
			yyVAL.withClause = &WithClause{
				Value: yyDollar[1].expression,
			}
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1050
		{
			yyVAL.withClause = &WithClause{
				Value:         yyDollar[1].expression,
				BoundVariable: yyDollar[3].expression,
			}
		}
	case 152:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1059
		{
			yyVAL.conditionClause = &ConditionClause{}
		}
	case 153:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:1062
		{
			yyVAL.conditionClause = &ConditionClause{
				Matches: yyDollar[2].expression,
			}
		}
	case 154:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line py.y:1067
		{
			yyVAL.conditionClause = &ConditionClause{
				Matches: yyDollar[2].expression,
				Alias:   yyDollar[4].expression,
			}
		}
	case 155:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line py.y:1073
		{
			yyVAL.conditionClause = &ConditionClause{
				Matches: yyDollar[2].expression,
				Alias:   yyDollar[4].expression,
			}
		}
	case 156:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1082
		{
			yyVAL.statements = yyDollar[1].statements
		}
	case 157:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line py.y:1085
		{
			yyVAL.statements = yyDollar[3].statements
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1091
		{
			yyVAL.expression = yyDollar[1].expression
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1094
		{
			yyVAL.expression = yyDollar[1].expression
		}
	case 160:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1100
		{
			yyVAL.expression = &LambdaExpr{
				Value: yyDollar[3].expression,
			}
		}
	case 161:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line py.y:1105
		{
			yyVAL.expression = &LambdaExpr{
				Arguments: yyDollar[2].arguments,
				Value:     yyDollar[4].expression,
			}
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1114
		{
			yyVAL.expression = yyDollar[1].expression
		}
	case 163:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line py.y:1117
		{
			yyVAL.expression = &ConditionExpr{
				True:      yyDollar[1].expression,
				Predicate: yyDollar[3].expression,
				False:     yyDollar[5].expression,
			}
		}
	case 164:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1124
		{
			yyVAL.expression = yyDollar[1].expression
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1130
		{
			yyVAL.expression = yyDollar[1].expression
		}
	case 166:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1133
		{
			yyVAL.expression = NewBinaryExpr(yyDollar[1].expression, yyDollar[2].token, yyDollar[3].expression)
		}
	case 167:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1139
		{
			yyVAL.expression = yyDollar[1].expression
		}
	case 168:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1142
		{
			yyVAL.expression = NewBinaryExpr(yyDollar[1].expression, yyDollar[2].token, yyDollar[3].expression)
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1148
		{
			yyVAL.expression = yyDollar[1].expression
		}
	case 170:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:1151
		{
			yyVAL.expression = &UnaryExpr{
				Op:    yyDollar[1].token,
				Value: yyDollar[2].expression,
			}
		}
	case 171:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1160
		{
			yyVAL.expression = yyDollar[1].expression
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1163
		{
			yyVAL.expression = NewBinaryExpr(yyDollar[1].expression, yyDollar[2].token, yyDollar[3].expression)
		}
	case 173:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1169
		{
			yyVAL.token = yyDollar[1].token
		}
	case 174:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1172
		{
			yyVAL.token = yyDollar[1].token
		}
	case 175:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1175
		{
			yyVAL.token = yyDollar[1].token
		}
	case 176:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1178
		{
			yyVAL.token = yyDollar[1].token
		}
	case 177:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1181
		{
			yyVAL.token = yyDollar[1].token
		}
	case 178:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1184
		{
			yyVAL.token = yyDollar[1].token
		}
	case 179:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1187
		{
			yyVAL.token = yyDollar[1].token
		}
	case 180:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1190
		{
			yyVAL.token = yyDollar[1].token
		}
	case 181:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:1193
		{
			yyVAL.token = yyDollar[1].token
			yyVAL.token.TokenType = NOT_IN
			// TODO pull IN's comment into NOT_IN
		}
	case 182:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:1198
		{
			yyVAL.token = yyDollar[1].token
			yyVAL.token.TokenType = IS_NOT
			// TODO pull NOT's comment into NOT_IN
		}
	case 183:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1206
		{
			yyVAL.expression = yyDollar[1].expression
		}
	case 184:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1209
		{
			yyVAL.expression = NewBinaryExpr(yyDollar[1].expression, yyDollar[2].token, yyDollar[3].expression)
		}
	case 185:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1215
		{
			yyVAL.expression = yyDollar[1].expression
		}
	case 186:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1218
		{
			yyVAL.expression = NewBinaryExpr(yyDollar[1].expression, yyDollar[2].token, yyDollar[3].expression)
		}
	case 187:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1224
		{
			yyVAL.expression = yyDollar[1].expression
		}
	case 188:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1227
		{
			yyVAL.expression = NewBinaryExpr(yyDollar[1].expression, yyDollar[2].token, yyDollar[3].expression)
		}
	case 189:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1233
		{
			yyVAL.expression = yyDollar[1].expression
		}
	case 190:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1236
		{
			yyVAL.expression = NewBinaryExpr(yyDollar[1].expression, yyDollar[2].token, yyDollar[3].expression)
		}
	case 191:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1239
		{
			yyVAL.expression = NewBinaryExpr(yyDollar[1].expression, yyDollar[2].token, yyDollar[3].expression)
		}
	case 192:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1245
		{
			yyVAL.expression = yyDollar[1].expression
		}
	case 193:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1248
		{
			yyVAL.expression = NewBinaryExpr(yyDollar[1].expression, yyDollar[2].token, yyDollar[3].expression)
		}
	case 194:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1251
		{
			yyVAL.expression = NewBinaryExpr(yyDollar[1].expression, yyDollar[2].token, yyDollar[3].expression)
		}
	case 195:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1257
		{
			yyVAL.expression = yyDollar[1].expression
		}
	case 196:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1260
		{
			yyVAL.expression = NewBinaryExpr(yyDollar[1].expression, yyDollar[2].token, yyDollar[3].expression)
		}
	case 197:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1263
		{
			yyVAL.expression = NewBinaryExpr(yyDollar[1].expression, yyDollar[2].token, yyDollar[3].expression)
		}
	case 198:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1266
		{
			yyVAL.expression = NewBinaryExpr(yyDollar[1].expression, yyDollar[2].token, yyDollar[3].expression)
		}
	case 199:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1269
		{
			yyVAL.expression = NewBinaryExpr(yyDollar[1].expression, yyDollar[2].token, yyDollar[3].expression)
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1275
		{
			yyVAL.token = yyDollar[1].token
		}
	case 201:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1278
		{
			yyVAL.token = yyDollar[1].token
		}
	case 202:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1281
		{
			yyVAL.token = yyDollar[1].token
		}
	case 203:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1287
		{
			yyVAL.tokens = []*Token{yyDollar[1].token}
		}
	case 204:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:1290
		{
			yyVAL.tokens = append(yyDollar[1].tokens, yyDollar[2].token)
		}
	case 205:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1296
		{
			yyVAL.expression = yyDollar[1].expression
		}
	case 206:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:1299
		{
			yyVAL.expression = NewFactorExpr(yyDollar[1].tokens, yyDollar[2].expression)
		}
	case 207:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1305
		{
			yyVAL.expression = yyDollar[1].expression
		}
	case 208:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1308
		{
			yyVAL.expression = NewBinaryExpr(yyDollar[1].expression, yyDollar[2].token, yyDollar[3].expression)
		}
	case 209:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:1311
		{
			yyVAL.expression = NewExpressionFromTrailers(yyDollar[1].expression, yyDollar[2].trailers)
		}
	case 210:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line py.y:1314
		{
			yyVAL.expression = NewBinaryExpr(NewExpressionFromTrailers(yyDollar[1].expression, yyDollar[2].trailers), yyDollar[3].token, yyDollar[4].expression)
		}
	case 211:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1320
		{
			yyVAL.trailers = []interface{}{yyDollar[1].trailer}
		}
	case 212:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:1323
		{
			yyVAL.trailers = append(yyDollar[1].trailers, yyDollar[2].trailer)
		}
	case 213:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:1329
		{
			yyVAL.expression = &CollectionExpr{
				Type: TupleCollection,
			}
		}
	case 214:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1334
		{
			yyVAL.expression = yyDollar[2].expression
		}
	case 215:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1337
		{
			yyVAL.expression = yyDollar[2].expression
		}
	case 216:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:1340
		{
			yyVAL.expression = &CollectionExpr{
				Type: ListCollection,
			}
		}
	case 217:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1345
		{
			yyVAL.expression = yyDollar[2].expression
		}
	case 218:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:1348
		{
			yyVAL.expression = &CollectionExpr{
				Type: DictCollection,
			}
		}
	case 219:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1353
		{
			yyVAL.expression = yyDollar[2].expression
		}
	case 220:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1356
		{
			yyVAL.expression = &EvalExpr{
				Expression: yyDollar[2].expressions.ConvertToExpr(),
			}
		}
	case 221:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1361
		{
			yyVAL.expression = NewNone(yyDollar[1].token)
		}
	case 222:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1364
		{
			// NOTE: True / False are classified as NAME because they are
			// reassignable variables, not constants ...
			//
			// >>> True = False
			// >>> True
			// False
			yyVAL.expression = NewIdentifier(yyDollar[1].token)
		}
	case 223:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1373
		{
			yyVAL.expression = NewNumber(yyDollar[1].token)
		}
	case 224:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1376
		{
			yyVAL.expression = NewNumber(yyDollar[1].token)
		}
	case 225:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1379
		{
			yyVAL.expression = &String{
				Pieces: yyDollar[1].tokens,
			}
		}
	case 226:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1387
		{
			yyVAL.tokens = []*Token{yyDollar[1].token}
		}
	case 227:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:1390
		{
			yyVAL.tokens = append(yyDollar[1].tokens, yyDollar[2].token)
		}
	case 228:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1396
		{
			yyDollar[1].expressions.ExplicitCollection = true
			yyVAL.expression = yyDollar[1].expressions.ConvertToExpr()
			yyVAL.expression.(*CollectionExpr).Type = ListCollection
		}
	case 229:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:1401
		{
			yyVAL.expression = &ComprehensionExpr{
				Type:      ListComprehension,
				Value:     yyDollar[1].expression,
				Iterators: yyDollar[2].iterators,
			}
		}
	case 230:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1411
		{
			yyVAL.expression = yyDollar[1].expressions.ConvertToExpr()
		}
	case 231:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:1414
		{
			yyVAL.expression = &ComprehensionExpr{
				Type:      GeneratorComprehension,
				Value:     yyDollar[1].expression,
				Iterators: yyDollar[2].iterators,
			}
		}
	case 232:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1424
		{
			yyVAL.expression = &LambdaExpr{
				Value: yyDollar[3].expression,
			}
		}
	case 233:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line py.y:1429
		{
			yyVAL.expression = &LambdaExpr{
				Arguments: yyDollar[2].arguments,
				Value:     yyDollar[4].expression,
			}
		}
	case 234:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:1438
		{
			yyVAL.trailer = []*Argument{}
		}
	case 235:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1441
		{
			yyVAL.trailer = yyDollar[2].arguments
		}
	case 236:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1444
		{
			yyVAL.trailer = yyDollar[2].subscripts
		}
	case 237:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:1447
		{
			yyVAL.trailer = yyDollar[2].token
		}
	case 238:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1453
		{
			yyVAL.subscripts = []*Subscript{yyDollar[1].subscript}
		}
	case 239:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1456
		{
			yyVAL.subscripts = append(yyDollar[1].subscripts, yyDollar[3].subscript)
		}
	case 240:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1462
		{
			yyVAL.subscripts = yyDollar[1].subscripts
		}
	case 241:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:1465
		{
			yyVAL.subscripts = yyDollar[1].subscripts
		}
	case 242:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1471
		{
			yyVAL.subscript = &Subscript{Ellipsis: true}
		}
	case 243:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1474
		{
			yyVAL.subscript = &Subscript{Index: yyDollar[1].expression}
		}
	case 244:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1477
		{
			yyVAL.subscript = &Subscript{}
		}
	case 245:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:1480
		{
			yyVAL.subscript = &Subscript{Left: yyDollar[1].expression}
		}
	case 246:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:1483
		{
			yyVAL.subscript = &Subscript{Middle: yyDollar[2].expression}
		}
	case 247:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1486
		{
			yyVAL.subscript = &Subscript{Left: yyDollar[1].expression, Middle: yyDollar[3].expression}
		}
	case 248:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:1489
		{
			yyVAL.subscript = yyDollar[2].subscript
		}
	case 249:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1492
		{
			yyVAL.subscript = yyDollar[3].subscript
			yyVAL.subscript.Left = yyDollar[1].expression
		}
	case 250:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1496
		{
			yyVAL.subscript = yyDollar[3].subscript
			yyVAL.subscript.Middle = yyDollar[2].expression
		}
	case 251:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line py.y:1500
		{
			yyVAL.subscript = yyDollar[4].subscript
			yyVAL.subscript.Left = yyDollar[1].expression
			yyVAL.subscript.Middle = yyDollar[3].expression
		}
	case 252:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1508
		{
			yyVAL.subscript = &Subscript{}
		}
	case 253:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:1511
		{
			yyVAL.subscript = &Subscript{Right: yyDollar[2].expression}
		}
	case 254:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1517
		{
			yyVAL.expressions = &ExprList{
				Expressions: []Expression{yyDollar[1].expression},
			}
		}
	case 255:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1522
		{
			yyVAL.expressions = yyDollar[1].expressions
			yyVAL.expressions.Expressions = append(yyVAL.expressions.Expressions, yyDollar[3].expression)
			yyVAL.expressions.ExplicitCollection = true
		}
	case 256:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1530
		{
			yyVAL.expressions = yyDollar[1].expressions
		}
	case 257:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:1533
		{
			yyVAL.expressions = yyDollar[1].expressions
			yyVAL.expressions.ExplicitCollection = true
		}
	case 258:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1539
		{
			yyVAL.expressions = &ExprList{
				Expressions: []Expression{yyDollar[1].expression},
			}
		}
	case 259:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1544
		{
			yyVAL.expressions = yyDollar[1].expressions
			yyVAL.expressions.Expressions = append(yyVAL.expressions.Expressions, yyDollar[3].expression)
			yyVAL.expressions.ExplicitCollection = true
		}
	case 260:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1552
		{
			yyVAL.expressions = yyDollar[1].expressions
		}
	case 261:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:1555
		{
			yyVAL.expressions = yyDollar[1].expressions
			yyVAL.expressions.ExplicitCollection = true
		}
	case 262:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1563
		{
			yyVAL.expressions = &ExprList{
				Expressions: []Expression{yyDollar[1].expression},
			}
		}
	case 263:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1568
		{
			yyVAL.expressions = yyDollar[1].expressions
			yyVAL.expressions.Expressions = append(yyVAL.expressions.Expressions, yyDollar[3].expression)
		}
	case 264:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1575
		{
			yyVAL.expressions = yyDollar[1].expressions
		}
	case 265:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:1578
		{
			yyVAL.expressions = yyDollar[1].expressions
			yyVAL.expressions.ExplicitCollection = true
		}
	case 266:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1585
		{
			yyVAL.expressions = &ExprList{
				Expressions: []Expression{
					&CollectionExpr{
						Type:  TupleCollection,
						Items: []Expression{yyDollar[1].expression, yyDollar[3].expression},
					},
				},
				ExplicitCollection: true,
			}
		}
	case 267:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line py.y:1596
		{
			yyVAL.expressions = yyDollar[1].expressions
			yyVAL.expressions.Expressions = append(
				yyVAL.expressions.Expressions,
				&CollectionExpr{
					Type:  TupleCollection,
					Items: []Expression{yyDollar[3].expression, yyDollar[5].expression},
				})
		}
	case 268:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1607
		{
			yyVAL.expressions = yyDollar[1].expressions
		}
	case 269:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:1610
		{
			yyVAL.expressions = yyDollar[1].expressions
		}
	case 270:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1616
		{
			yyVAL.expression = yyDollar[1].expressions.ConvertToExpr()
			yyVAL.expression.(*CollectionExpr).Type = DictCollection
		}
	case 271:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line py.y:1620
		{
			yyVAL.expression = &ComprehensionExpr{
				Type:      DictComprehension,
				Key:       yyDollar[1].expression,
				Value:     yyDollar[3].expression,
				Iterators: yyDollar[4].iterators,
			}
		}
	case 272:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1628
		{
			yyDollar[1].expressions.ExplicitCollection = true
			yyVAL.expression = yyDollar[1].expressions.ConvertToExpr()
			yyVAL.expression.(*CollectionExpr).Type = SetCollection

		}
	case 273:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:1634
		{
			yyVAL.expression = &ComprehensionExpr{
				Type:      SetComprehension,
				Key:       yyDollar[1].expression,
				Iterators: yyDollar[2].iterators,
			}
		}
	case 274:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line py.y:1644
		{
			yyVAL.statement = &ClassDef{
				Name:       yyDollar[2].token,
				Statements: yyDollar[4].statements,
			}
		}
	case 275:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line py.y:1650
		{
			yyVAL.statement = &ClassDef{
				Name:       yyDollar[2].token,
				Statements: yyDollar[6].statements,
			}
		}
	case 276:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line py.y:1656
		{
			yyVAL.statement = &ClassDef{
				Name:          yyDollar[2].token,
				ParentClasses: yyDollar[4].expressions.Expressions,
				Statements:    yyDollar[7].statements,
			}
		}
	case 277:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1666
		{
			yyVAL.arguments = []*Argument{yyDollar[1].argument}
		}
	case 278:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1669
		{
			yyVAL.arguments = append(yyDollar[1].arguments, yyDollar[3].argument)
		}
	case 279:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1675
		{
			yyVAL.arguments = yyDollar[1].arguments
		}
	case 280:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:1678
		{
			yyVAL.arguments = yyDollar[1].arguments
		}
	case 281:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1684
		{
			yyVAL.arguments = yyDollar[1].arguments
		}
	case 282:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line py.y:1687
		{
			yyDollar[1].arguments[len(yyDollar[1].arguments)-1].MergeFrom(&yyDollar[2].token.Node)

			yyVAL.arguments = append(yyDollar[1].arguments, NewPositionVarArg(yyDollar[3].token, yyDollar[4].expression))
		}
	case 283:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line py.y:1692
		{
			yyDollar[1].arguments[len(yyDollar[1].arguments)-1].MergeFrom(&yyDollar[2].token.Node)

			arg := NewPositionVarArg(yyDollar[3].token, yyDollar[4].expression)
			arg.MergeFrom(&yyDollar[5].token.Node)

			yyVAL.arguments = append(yyDollar[1].arguments, arg)
			yyVAL.arguments = append(yyVAL.arguments, yyDollar[6].arguments...)
		}
	case 284:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line py.y:1701
		{
			yyDollar[1].arguments[len(yyDollar[1].arguments)-1].MergeFrom(&yyDollar[2].token.Node)

			arg := NewPositionVarArg(yyDollar[3].token, yyDollar[4].expression)
			arg.MergeFrom(&yyDollar[5].token.Node)

			yyVAL.arguments = append(yyDollar[1].arguments, arg)
			yyVAL.arguments = append(yyVAL.arguments, NewKeywordVarArg(yyDollar[6].token, yyDollar[7].expression))
		}
	case 285:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line py.y:1710
		{
			yyDollar[1].arguments[len(yyDollar[1].arguments)-1].MergeFrom(&yyDollar[2].token.Node)

			arg := NewPositionVarArg(yyDollar[3].token, yyDollar[4].expression)
			arg.MergeFrom(&yyDollar[5].token.Node)

			yyVAL.arguments = append(yyDollar[1].arguments, arg)
			yyVAL.arguments = append(yyVAL.arguments, yyDollar[6].arguments...)

			yyVAL.arguments[len(yyVAL.arguments)-1].MergeFrom(&yyDollar[7].token.Node)

			yyVAL.arguments = append(yyVAL.arguments, NewKeywordVarArg(yyDollar[8].token, yyDollar[9].expression))
		}
	case 286:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line py.y:1723
		{
			yyDollar[1].arguments[len(yyDollar[1].arguments)-1].MergeFrom(&yyDollar[2].token.Node)
			yyVAL.arguments = append(yyDollar[1].arguments, NewKeywordVarArg(yyDollar[3].token, yyDollar[4].expression))
		}
	case 287:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:1727
		{
			yyVAL.arguments = []*Argument{NewPositionVarArg(yyDollar[1].token, yyDollar[2].expression)}
		}
	case 288:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line py.y:1730
		{
			arg := NewPositionVarArg(yyDollar[1].token, yyDollar[2].expression)
			arg.MergeFrom(&yyDollar[3].token.Node)

			yyVAL.arguments = append([]*Argument{arg}, yyDollar[4].arguments...)
		}
	case 289:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line py.y:1736
		{
			arg := NewPositionVarArg(yyDollar[1].token, yyDollar[2].expression)
			arg.MergeFrom(&yyDollar[3].token.Node)

			yyVAL.arguments = []*Argument{arg, NewKeywordVarArg(yyDollar[4].token, yyDollar[5].expression)}
		}
	case 290:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line py.y:1742
		{
			arg := NewPositionVarArg(yyDollar[1].token, yyDollar[2].expression)
			arg.MergeFrom(&yyDollar[3].token.Node)

			yyVAL.arguments = append([]*Argument{arg}, yyDollar[4].arguments...)
			yyVAL.arguments[len(yyVAL.arguments)-1].MergeFrom(&yyDollar[5].token.Node)

			yyVAL.arguments = append(yyVAL.arguments, NewKeywordVarArg(yyDollar[6].token, yyDollar[7].expression))
		}
	case 291:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:1751
		{
			yyVAL.arguments = []*Argument{NewKeywordVarArg(yyDollar[1].token, yyDollar[2].expression)}
		}
	case 292:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1758
		{
			yyVAL.argument = NewArgument(nil, nil, yyDollar[1].expression)
		}
	case 293:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:1761
		{
			yyVAL.argument = NewArgument(
				nil,
				nil,
				&ComprehensionExpr{
					Type:      GeneratorComprehension,
					Value:     yyDollar[1].expression,
					Iterators: yyDollar[2].iterators,
				})
		}
	case 294:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1771
		{
			yyVAL.argument = NewArgument(yyDollar[1].expression, yyDollar[2].token, yyDollar[3].expression)
		}
	case 295:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line py.y:1777
		{
			yyVAL.iterator = &Iterator{
				BoundVariables: yyDollar[2].expressions.ConvertToExpr(),
				Source:         yyDollar[4].expressions.ConvertToExpr(),
			}
		}
	case 296:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line py.y:1783
		{
			yyVAL.iterator = &Iterator{
				BoundVariables: yyDollar[2].expressions.ConvertToExpr(),
				Source:         yyDollar[4].expressions.ConvertToExpr(),
				Filters:        []Expression{yyDollar[6].expression},
			}
		}
	case 297:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1793
		{
			yyVAL.iterators = []*Iterator{yyDollar[1].iterator}
		}
	case 298:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:1796
		{
			yyVAL.iterators = append(yyDollar[1].iterators, yyDollar[2].iterator)
		}
	case 299:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1799
		{
			yyVAL.iterators = yyDollar[1].iterators
			last := yyVAL.iterators[len(yyVAL.iterators)-1]
			last.Filters = append(last.Filters, yyDollar[3].expression)
		}
	case 300:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line py.y:1807
		{
			yyVAL.iterator = &Iterator{
				BoundVariables: yyDollar[2].expressions.ConvertToExpr(),
				Source:         yyDollar[4].expression,
			}
		}
	case 301:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line py.y:1813
		{
			yyVAL.iterator = &Iterator{
				BoundVariables: yyDollar[2].expressions.ConvertToExpr(),
				Source:         yyDollar[4].expression,
				Filters:        []Expression{yyDollar[6].expression},
			}
		}
	case 302:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1823
		{
			yyVAL.iterators = []*Iterator{yyDollar[1].iterator}
		}
	case 303:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:1826
		{
			yyVAL.iterators = append(yyDollar[1].iterators, yyDollar[2].iterator)
		}
	case 304:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line py.y:1829
		{
			yyVAL.iterators = yyDollar[1].iterators
			last := yyVAL.iterators[len(yyVAL.iterators)-1]
			last.Filters = append(last.Filters, yyDollar[3].expression)
		}
	case 305:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line py.y:1837
		{
			yyVAL.expression = &YieldExpr{}
		}
	case 306:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line py.y:1840
		{
			yyVAL.expression = &YieldExpr{Expression: yyDollar[2].expressions.ConvertToExpr()}
		}
	}
	goto yystack /* stack new state and value */
}
