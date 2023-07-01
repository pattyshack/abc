// Auto-generated from source: ansi-c.lr

package ansi_c

import (
	fmt "fmt"
	io "io"
	sort "sort"
)

type CSymbolId int

const (
	CIdentifierToken    = CSymbolId(256)
	CConstantToken      = CSymbolId(257)
	CStringLiteralToken = CSymbolId(258)
	CSizeofToken        = CSymbolId(259)
	CPtrOpToken         = CSymbolId(260)
	CIncOpToken         = CSymbolId(261)
	CDecOpToken         = CSymbolId(262)
	CLeftOpToken        = CSymbolId(263)
	CRightOpToken       = CSymbolId(264)
	CLeOpToken          = CSymbolId(265)
	CGeOpToken          = CSymbolId(266)
	CEqOpToken          = CSymbolId(267)
	CNeOpToken          = CSymbolId(268)
	CAndOpToken         = CSymbolId(269)
	COrOpToken          = CSymbolId(270)
	CMulAssignToken     = CSymbolId(271)
	CDivAssignToken     = CSymbolId(272)
	CModAssignToken     = CSymbolId(273)
	CAddAssignToken     = CSymbolId(274)
	CSubAssignToken     = CSymbolId(275)
	CLeftAssignToken    = CSymbolId(276)
	CRightAssignToken   = CSymbolId(277)
	CAndAssignToken     = CSymbolId(278)
	CXorAssignToken     = CSymbolId(279)
	COrAssignToken      = CSymbolId(280)
	CTypeNameToken      = CSymbolId(281)
	CTypedefToken       = CSymbolId(282)
	CExternToken        = CSymbolId(283)
	CStaticToken        = CSymbolId(284)
	CAutoToken          = CSymbolId(285)
	CRegisterToken      = CSymbolId(286)
	CCharToken          = CSymbolId(287)
	CShortToken         = CSymbolId(288)
	CIntToken           = CSymbolId(289)
	CLongToken          = CSymbolId(290)
	CSignedToken        = CSymbolId(291)
	CUnsignedToken      = CSymbolId(292)
	CFloatToken         = CSymbolId(293)
	CDoubleToken        = CSymbolId(294)
	CConstToken         = CSymbolId(295)
	CVolatileToken      = CSymbolId(296)
	CVoidToken          = CSymbolId(297)
	CStructToken        = CSymbolId(298)
	CUnionToken         = CSymbolId(299)
	CEnumToken          = CSymbolId(300)
	CEllipsisToken      = CSymbolId(301)
	CCaseToken          = CSymbolId(302)
	CDefaultToken       = CSymbolId(303)
	CIfToken            = CSymbolId(304)
	CElseToken          = CSymbolId(305)
	CSwitchToken        = CSymbolId(306)
	CWhileToken         = CSymbolId(307)
	CDoToken            = CSymbolId(308)
	CForToken           = CSymbolId(309)
	CGotoToken          = CSymbolId(310)
	CContinueToken      = CSymbolId(311)
	CBreakToken         = CSymbolId(312)
	CReturnToken        = CSymbolId(313)
)

type CLocation struct {
	FileName string
	Line     int
	Column   int
}

func (l CLocation) String() string {
	return fmt.Sprintf("%v:%v:%v", l.FileName, l.Line, l.Column)
}

func (l CLocation) ShortString() string {
	return fmt.Sprintf("%v:%v", l.Line, l.Column)
}

type CToken interface {
	Id() CSymbolId
	Loc() CLocation
}

type CGenericSymbol struct {
	CSymbolId
	CLocation
}

func (t *CGenericSymbol) Id() CSymbolId { return t.CSymbolId }

func (t *CGenericSymbol) Loc() CLocation { return t.CLocation }

type CLexer interface {
	// Note: Return io.EOF to indicate end of stream
	// Token with unspecified value type should return *CGenericSymbol
	Next() (CToken, error)

	CurrentLocation() CLocation
}

type CReducer interface {
	// 83:4: primary_expression -> a: ...
	AToPrimaryExpression(Identifier_ *CGenericSymbol) (*CGenericSymbol, error)

	// 84:4: primary_expression -> b: ...
	BToPrimaryExpression(Constant_ *CGenericSymbol) (*CGenericSymbol, error)

	// 85:4: primary_expression -> c: ...
	CToPrimaryExpression(StringLiteral_ *CGenericSymbol) (*CGenericSymbol, error)

	// 86:4: primary_expression -> d: ...
	DToPrimaryExpression(char *CGenericSymbol, Expression_ *CGenericSymbol, char2 *CGenericSymbol) (*CGenericSymbol, error)

	// 89:4: postfix_expression -> a: ...
	AToPostfixExpression(PrimaryExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 90:4: postfix_expression -> b: ...
	BToPostfixExpression(PostfixExpression_ *CGenericSymbol, char *CGenericSymbol, Expression_ *CGenericSymbol, char2 *CGenericSymbol) (*CGenericSymbol, error)

	// 91:4: postfix_expression -> c: ...
	CToPostfixExpression(PostfixExpression_ *CGenericSymbol, char *CGenericSymbol, char2 *CGenericSymbol) (*CGenericSymbol, error)

	// 92:4: postfix_expression -> d: ...
	DToPostfixExpression(PostfixExpression_ *CGenericSymbol, char *CGenericSymbol, ArgumentExpressionList_ *CGenericSymbol, char2 *CGenericSymbol) (*CGenericSymbol, error)

	// 93:4: postfix_expression -> e: ...
	EToPostfixExpression(PostfixExpression_ *CGenericSymbol, char *CGenericSymbol, Identifier_ *CGenericSymbol) (*CGenericSymbol, error)

	// 94:4: postfix_expression -> f: ...
	FToPostfixExpression(PostfixExpression_ *CGenericSymbol, PtrOp_ *CGenericSymbol, Identifier_ *CGenericSymbol) (*CGenericSymbol, error)

	// 95:4: postfix_expression -> g: ...
	GToPostfixExpression(PostfixExpression_ *CGenericSymbol, IncOp_ *CGenericSymbol) (*CGenericSymbol, error)

	// 96:4: postfix_expression -> h: ...
	HToPostfixExpression(PostfixExpression_ *CGenericSymbol, DecOp_ *CGenericSymbol) (*CGenericSymbol, error)

	// 99:4: argument_expression_list -> a: ...
	AToArgumentExpressionList(AssignmentExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 100:4: argument_expression_list -> b: ...
	BToArgumentExpressionList(ArgumentExpressionList_ *CGenericSymbol, char *CGenericSymbol, AssignmentExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 103:4: unary_expression -> a: ...
	AToUnaryExpression(PostfixExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 104:4: unary_expression -> b: ...
	BToUnaryExpression(IncOp_ *CGenericSymbol, UnaryExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 105:4: unary_expression -> c: ...
	CToUnaryExpression(DecOp_ *CGenericSymbol, UnaryExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 106:4: unary_expression -> d: ...
	DToUnaryExpression(UnaryOperator_ *CGenericSymbol, CastExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 107:4: unary_expression -> e: ...
	EToUnaryExpression(Sizeof_ *CGenericSymbol, UnaryExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 108:4: unary_expression -> f: ...
	FToUnaryExpression(Sizeof_ *CGenericSymbol, char *CGenericSymbol, TypeName_ *CGenericSymbol, char2 *CGenericSymbol) (*CGenericSymbol, error)

	// 111:4: unary_operator -> a: ...
	AToUnaryOperator(char *CGenericSymbol) (*CGenericSymbol, error)

	// 112:4: unary_operator -> b: ...
	BToUnaryOperator(char *CGenericSymbol) (*CGenericSymbol, error)

	// 113:4: unary_operator -> c: ...
	CToUnaryOperator(char *CGenericSymbol) (*CGenericSymbol, error)

	// 114:4: unary_operator -> d: ...
	DToUnaryOperator(char *CGenericSymbol) (*CGenericSymbol, error)

	// 115:4: unary_operator -> e: ...
	EToUnaryOperator(char *CGenericSymbol) (*CGenericSymbol, error)

	// 116:4: unary_operator -> f: ...
	FToUnaryOperator(char *CGenericSymbol) (*CGenericSymbol, error)

	// 119:4: cast_expression -> a: ...
	AToCastExpression(UnaryExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 120:4: cast_expression -> b: ...
	BToCastExpression(char *CGenericSymbol, TypeName_ *CGenericSymbol, char2 *CGenericSymbol, CastExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 123:4: multiplicative_expression -> a: ...
	AToMultiplicativeExpression(CastExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 124:4: multiplicative_expression -> b: ...
	BToMultiplicativeExpression(MultiplicativeExpression_ *CGenericSymbol, char *CGenericSymbol, CastExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 125:4: multiplicative_expression -> c: ...
	CToMultiplicativeExpression(MultiplicativeExpression_ *CGenericSymbol, char *CGenericSymbol, CastExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 126:4: multiplicative_expression -> d: ...
	DToMultiplicativeExpression(MultiplicativeExpression_ *CGenericSymbol, char *CGenericSymbol, CastExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 129:4: additive_expression -> a: ...
	AToAdditiveExpression(MultiplicativeExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 130:4: additive_expression -> b: ...
	BToAdditiveExpression(AdditiveExpression_ *CGenericSymbol, char *CGenericSymbol, MultiplicativeExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 131:4: additive_expression -> c: ...
	CToAdditiveExpression(AdditiveExpression_ *CGenericSymbol, char *CGenericSymbol, MultiplicativeExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 134:4: shift_expression -> a: ...
	AToShiftExpression(AdditiveExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 135:4: shift_expression -> b: ...
	BToShiftExpression(ShiftExpression_ *CGenericSymbol, LeftOp_ *CGenericSymbol, AdditiveExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 136:4: shift_expression -> c: ...
	CToShiftExpression(ShiftExpression_ *CGenericSymbol, RightOp_ *CGenericSymbol, AdditiveExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 139:4: relational_expression -> a: ...
	AToRelationalExpression(ShiftExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 140:4: relational_expression -> b: ...
	BToRelationalExpression(RelationalExpression_ *CGenericSymbol, char *CGenericSymbol, ShiftExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 141:4: relational_expression -> c: ...
	CToRelationalExpression(RelationalExpression_ *CGenericSymbol, char *CGenericSymbol, ShiftExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 142:4: relational_expression -> d: ...
	DToRelationalExpression(RelationalExpression_ *CGenericSymbol, LeOp_ *CGenericSymbol, ShiftExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 143:4: relational_expression -> e: ...
	EToRelationalExpression(RelationalExpression_ *CGenericSymbol, GeOp_ *CGenericSymbol, ShiftExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 146:4: equality_expression -> a: ...
	AToEqualityExpression(RelationalExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 147:4: equality_expression -> b: ...
	BToEqualityExpression(EqualityExpression_ *CGenericSymbol, EqOp_ *CGenericSymbol, RelationalExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 148:4: equality_expression -> c: ...
	CToEqualityExpression(EqualityExpression_ *CGenericSymbol, NeOp_ *CGenericSymbol, RelationalExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 151:4: and_expression -> a: ...
	AToAndExpression(EqualityExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 152:4: and_expression -> b: ...
	BToAndExpression(AndExpression_ *CGenericSymbol, char *CGenericSymbol, EqualityExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 155:4: exclusive_or_expression -> a: ...
	AToExclusiveOrExpression(AndExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 156:4: exclusive_or_expression -> b: ...
	BToExclusiveOrExpression(ExclusiveOrExpression_ *CGenericSymbol, char *CGenericSymbol, AndExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 159:4: inclusive_or_expression -> a: ...
	AToInclusiveOrExpression(ExclusiveOrExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 160:4: inclusive_or_expression -> b: ...
	BToInclusiveOrExpression(InclusiveOrExpression_ *CGenericSymbol, char *CGenericSymbol, ExclusiveOrExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 163:4: logical_and_expression -> a: ...
	AToLogicalAndExpression(InclusiveOrExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 164:4: logical_and_expression -> b: ...
	BToLogicalAndExpression(LogicalAndExpression_ *CGenericSymbol, AndOp_ *CGenericSymbol, InclusiveOrExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 167:4: logical_or_expression -> a: ...
	AToLogicalOrExpression(LogicalAndExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 168:4: logical_or_expression -> b: ...
	BToLogicalOrExpression(LogicalOrExpression_ *CGenericSymbol, OrOp_ *CGenericSymbol, LogicalAndExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 171:4: conditional_expression -> a: ...
	AToConditionalExpression(LogicalOrExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 172:4: conditional_expression -> b: ...
	BToConditionalExpression(LogicalOrExpression_ *CGenericSymbol, char *CGenericSymbol, Expression_ *CGenericSymbol, char2 *CGenericSymbol, ConditionalExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 175:4: assignment_expression -> a: ...
	AToAssignmentExpression(ConditionalExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 176:4: assignment_expression -> b: ...
	BToAssignmentExpression(UnaryExpression_ *CGenericSymbol, AssignmentOperator_ *CGenericSymbol, AssignmentExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 179:4: assignment_operator -> a: ...
	AToAssignmentOperator(char *CGenericSymbol) (*CGenericSymbol, error)

	// 180:4: assignment_operator -> b: ...
	BToAssignmentOperator(MulAssign_ *CGenericSymbol) (*CGenericSymbol, error)

	// 181:4: assignment_operator -> c: ...
	CToAssignmentOperator(DivAssign_ *CGenericSymbol) (*CGenericSymbol, error)

	// 182:4: assignment_operator -> d: ...
	DToAssignmentOperator(ModAssign_ *CGenericSymbol) (*CGenericSymbol, error)

	// 183:4: assignment_operator -> e: ...
	EToAssignmentOperator(AddAssign_ *CGenericSymbol) (*CGenericSymbol, error)

	// 184:4: assignment_operator -> f: ...
	FToAssignmentOperator(SubAssign_ *CGenericSymbol) (*CGenericSymbol, error)

	// 185:4: assignment_operator -> g: ...
	GToAssignmentOperator(LeftAssign_ *CGenericSymbol) (*CGenericSymbol, error)

	// 186:4: assignment_operator -> h: ...
	HToAssignmentOperator(RightAssign_ *CGenericSymbol) (*CGenericSymbol, error)

	// 187:4: assignment_operator -> i: ...
	IToAssignmentOperator(AndAssign_ *CGenericSymbol) (*CGenericSymbol, error)

	// 188:4: assignment_operator -> j: ...
	JToAssignmentOperator(XorAssign_ *CGenericSymbol) (*CGenericSymbol, error)

	// 189:4: assignment_operator -> k: ...
	KToAssignmentOperator(OrAssign_ *CGenericSymbol) (*CGenericSymbol, error)

	// 192:4: expression -> a: ...
	AToExpression(AssignmentExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 193:4: expression -> b: ...
	BToExpression(Expression_ *CGenericSymbol, char *CGenericSymbol, AssignmentExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 196:4: constant_expression -> a: ...
	AToConstantExpression(ConditionalExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 199:4: declaration -> a: ...
	AToDeclaration(DeclarationSpecifiers_ *CGenericSymbol, char *CGenericSymbol) (*CGenericSymbol, error)

	// 200:4: declaration -> b: ...
	BToDeclaration(DeclarationSpecifiers_ *CGenericSymbol, InitDeclaratorList_ *CGenericSymbol, char *CGenericSymbol) (*CGenericSymbol, error)

	// 203:4: declaration_specifiers -> a: ...
	AToDeclarationSpecifiers(StorageClassSpecifier_ *CGenericSymbol) (*CGenericSymbol, error)

	// 204:4: declaration_specifiers -> b: ...
	BToDeclarationSpecifiers(StorageClassSpecifier_ *CGenericSymbol, DeclarationSpecifiers_ *CGenericSymbol) (*CGenericSymbol, error)

	// 205:4: declaration_specifiers -> c: ...
	CToDeclarationSpecifiers(TypeSpecifier_ *CGenericSymbol) (*CGenericSymbol, error)

	// 206:4: declaration_specifiers -> d: ...
	DToDeclarationSpecifiers(TypeSpecifier_ *CGenericSymbol, DeclarationSpecifiers_ *CGenericSymbol) (*CGenericSymbol, error)

	// 207:4: declaration_specifiers -> e: ...
	EToDeclarationSpecifiers(TypeQualifier_ *CGenericSymbol) (*CGenericSymbol, error)

	// 208:4: declaration_specifiers -> f: ...
	FToDeclarationSpecifiers(TypeQualifier_ *CGenericSymbol, DeclarationSpecifiers_ *CGenericSymbol) (*CGenericSymbol, error)

	// 211:4: init_declarator_list -> a: ...
	AToInitDeclaratorList(InitDeclarator_ *CGenericSymbol) (*CGenericSymbol, error)

	// 212:4: init_declarator_list -> b: ...
	BToInitDeclaratorList(InitDeclaratorList_ *CGenericSymbol, char *CGenericSymbol, InitDeclarator_ *CGenericSymbol) (*CGenericSymbol, error)

	// 215:4: init_declarator -> a: ...
	AToInitDeclarator(Declarator_ *CGenericSymbol) (*CGenericSymbol, error)

	// 216:4: init_declarator -> b: ...
	BToInitDeclarator(Declarator_ *CGenericSymbol, char *CGenericSymbol, Initializer_ *CGenericSymbol) (*CGenericSymbol, error)

	// 219:4: storage_class_specifier -> a: ...
	AToStorageClassSpecifier(Typedef_ *CGenericSymbol) (*CGenericSymbol, error)

	// 220:4: storage_class_specifier -> b: ...
	BToStorageClassSpecifier(Extern_ *CGenericSymbol) (*CGenericSymbol, error)

	// 221:4: storage_class_specifier -> c: ...
	CToStorageClassSpecifier(Static_ *CGenericSymbol) (*CGenericSymbol, error)

	// 222:4: storage_class_specifier -> d: ...
	DToStorageClassSpecifier(Auto_ *CGenericSymbol) (*CGenericSymbol, error)

	// 223:4: storage_class_specifier -> e: ...
	EToStorageClassSpecifier(Register_ *CGenericSymbol) (*CGenericSymbol, error)

	// 226:4: type_specifier -> a: ...
	AToTypeSpecifier(Void_ *CGenericSymbol) (*CGenericSymbol, error)

	// 227:4: type_specifier -> b: ...
	BToTypeSpecifier(Char_ *CGenericSymbol) (*CGenericSymbol, error)

	// 228:4: type_specifier -> c: ...
	CToTypeSpecifier(Short_ *CGenericSymbol) (*CGenericSymbol, error)

	// 229:4: type_specifier -> d: ...
	DToTypeSpecifier(Int_ *CGenericSymbol) (*CGenericSymbol, error)

	// 230:4: type_specifier -> e: ...
	EToTypeSpecifier(Long_ *CGenericSymbol) (*CGenericSymbol, error)

	// 231:4: type_specifier -> f: ...
	FToTypeSpecifier(Float_ *CGenericSymbol) (*CGenericSymbol, error)

	// 232:4: type_specifier -> g: ...
	GToTypeSpecifier(Double_ *CGenericSymbol) (*CGenericSymbol, error)

	// 233:4: type_specifier -> h: ...
	HToTypeSpecifier(Signed_ *CGenericSymbol) (*CGenericSymbol, error)

	// 234:4: type_specifier -> i: ...
	IToTypeSpecifier(Unsigned_ *CGenericSymbol) (*CGenericSymbol, error)

	// 235:4: type_specifier -> j: ...
	JToTypeSpecifier(StructOrUnionSpecifier_ *CGenericSymbol) (*CGenericSymbol, error)

	// 236:4: type_specifier -> k: ...
	KToTypeSpecifier(EnumSpecifier_ *CGenericSymbol) (*CGenericSymbol, error)

	// 237:4: type_specifier -> l: ...
	LToTypeSpecifier(TypeName_ *CGenericSymbol) (*CGenericSymbol, error)

	// 240:4: struct_or_union_specifier -> a: ...
	AToStructOrUnionSpecifier(StructOrUnion_ *CGenericSymbol, Identifier_ *CGenericSymbol, char *CGenericSymbol, StructDeclarationList_ *CGenericSymbol, char2 *CGenericSymbol) (*CGenericSymbol, error)

	// 241:4: struct_or_union_specifier -> b: ...
	BToStructOrUnionSpecifier(StructOrUnion_ *CGenericSymbol, char *CGenericSymbol, StructDeclarationList_ *CGenericSymbol, char2 *CGenericSymbol) (*CGenericSymbol, error)

	// 242:4: struct_or_union_specifier -> c: ...
	CToStructOrUnionSpecifier(StructOrUnion_ *CGenericSymbol, Identifier_ *CGenericSymbol) (*CGenericSymbol, error)

	// 245:4: struct_or_union -> a: ...
	AToStructOrUnion(Struct_ *CGenericSymbol) (*CGenericSymbol, error)

	// 246:4: struct_or_union -> b: ...
	BToStructOrUnion(Union_ *CGenericSymbol) (*CGenericSymbol, error)

	// 249:4: struct_declaration_list -> a: ...
	AToStructDeclarationList(StructDeclaration_ *CGenericSymbol) (*CGenericSymbol, error)

	// 250:4: struct_declaration_list -> b: ...
	BToStructDeclarationList(StructDeclarationList_ *CGenericSymbol, StructDeclaration_ *CGenericSymbol) (*CGenericSymbol, error)

	// 253:4: struct_declaration -> a: ...
	AToStructDeclaration(SpecifierQualifierList_ *CGenericSymbol, StructDeclaratorList_ *CGenericSymbol, char *CGenericSymbol) (*CGenericSymbol, error)

	// 256:4: specifier_qualifier_list -> a: ...
	AToSpecifierQualifierList(TypeSpecifier_ *CGenericSymbol, SpecifierQualifierList_ *CGenericSymbol) (*CGenericSymbol, error)

	// 257:4: specifier_qualifier_list -> b: ...
	BToSpecifierQualifierList(TypeSpecifier_ *CGenericSymbol) (*CGenericSymbol, error)

	// 258:4: specifier_qualifier_list -> c: ...
	CToSpecifierQualifierList(TypeQualifier_ *CGenericSymbol, SpecifierQualifierList_ *CGenericSymbol) (*CGenericSymbol, error)

	// 259:4: specifier_qualifier_list -> d: ...
	DToSpecifierQualifierList(TypeQualifier_ *CGenericSymbol) (*CGenericSymbol, error)

	// 262:4: struct_declarator_list -> a: ...
	AToStructDeclaratorList(StructDeclarator_ *CGenericSymbol) (*CGenericSymbol, error)

	// 263:4: struct_declarator_list -> b: ...
	BToStructDeclaratorList(StructDeclaratorList_ *CGenericSymbol, char *CGenericSymbol, StructDeclarator_ *CGenericSymbol) (*CGenericSymbol, error)

	// 266:4: struct_declarator -> a: ...
	AToStructDeclarator(Declarator_ *CGenericSymbol) (*CGenericSymbol, error)

	// 267:4: struct_declarator -> b: ...
	BToStructDeclarator(char *CGenericSymbol, ConstantExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 268:4: struct_declarator -> c: ...
	CToStructDeclarator(Declarator_ *CGenericSymbol, char *CGenericSymbol, ConstantExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 271:4: enum_specifier -> a: ...
	AToEnumSpecifier(Enum_ *CGenericSymbol, char *CGenericSymbol, EnumeratorList_ *CGenericSymbol, char2 *CGenericSymbol) (*CGenericSymbol, error)

	// 272:4: enum_specifier -> b: ...
	BToEnumSpecifier(Enum_ *CGenericSymbol, Identifier_ *CGenericSymbol, char *CGenericSymbol, EnumeratorList_ *CGenericSymbol, char2 *CGenericSymbol) (*CGenericSymbol, error)

	// 273:4: enum_specifier -> c: ...
	CToEnumSpecifier(Enum_ *CGenericSymbol, Identifier_ *CGenericSymbol) (*CGenericSymbol, error)

	// 276:4: enumerator_list -> a: ...
	AToEnumeratorList(Enumerator_ *CGenericSymbol) (*CGenericSymbol, error)

	// 277:4: enumerator_list -> b: ...
	BToEnumeratorList(EnumeratorList_ *CGenericSymbol, char *CGenericSymbol, Enumerator_ *CGenericSymbol) (*CGenericSymbol, error)

	// 280:4: enumerator -> a: ...
	AToEnumerator(Identifier_ *CGenericSymbol) (*CGenericSymbol, error)

	// 281:4: enumerator -> b: ...
	BToEnumerator(Identifier_ *CGenericSymbol, char *CGenericSymbol, ConstantExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 284:4: type_qualifier -> a: ...
	AToTypeQualifier(Const_ *CGenericSymbol) (*CGenericSymbol, error)

	// 285:4: type_qualifier -> b: ...
	BToTypeQualifier(Volatile_ *CGenericSymbol) (*CGenericSymbol, error)

	// 288:4: declarator -> a: ...
	AToDeclarator(Pointer_ *CGenericSymbol, DirectDeclarator_ *CGenericSymbol) (*CGenericSymbol, error)

	// 289:4: declarator -> b: ...
	BToDeclarator(DirectDeclarator_ *CGenericSymbol) (*CGenericSymbol, error)

	// 292:4: direct_declarator -> a: ...
	AToDirectDeclarator(Identifier_ *CGenericSymbol) (*CGenericSymbol, error)

	// 293:4: direct_declarator -> b: ...
	BToDirectDeclarator(char *CGenericSymbol, Declarator_ *CGenericSymbol, char2 *CGenericSymbol) (*CGenericSymbol, error)

	// 294:4: direct_declarator -> c: ...
	CToDirectDeclarator(DirectDeclarator_ *CGenericSymbol, char *CGenericSymbol, ConstantExpression_ *CGenericSymbol, char2 *CGenericSymbol) (*CGenericSymbol, error)

	// 295:4: direct_declarator -> d: ...
	DToDirectDeclarator(DirectDeclarator_ *CGenericSymbol, char *CGenericSymbol, char2 *CGenericSymbol) (*CGenericSymbol, error)

	// 296:4: direct_declarator -> e: ...
	EToDirectDeclarator(DirectDeclarator_ *CGenericSymbol, char *CGenericSymbol, ParameterTypeList_ *CGenericSymbol, char2 *CGenericSymbol) (*CGenericSymbol, error)

	// 297:4: direct_declarator -> f: ...
	FToDirectDeclarator(DirectDeclarator_ *CGenericSymbol, char *CGenericSymbol, IdentifierList_ *CGenericSymbol, char2 *CGenericSymbol) (*CGenericSymbol, error)

	// 298:4: direct_declarator -> g: ...
	GToDirectDeclarator(DirectDeclarator_ *CGenericSymbol, char *CGenericSymbol, char2 *CGenericSymbol) (*CGenericSymbol, error)

	// 301:4: pointer -> a: ...
	AToPointer(char *CGenericSymbol) (*CGenericSymbol, error)

	// 302:4: pointer -> b: ...
	BToPointer(char *CGenericSymbol, TypeQualifierList_ *CGenericSymbol) (*CGenericSymbol, error)

	// 303:4: pointer -> c: ...
	CToPointer(char *CGenericSymbol, Pointer_ *CGenericSymbol) (*CGenericSymbol, error)

	// 304:4: pointer -> d: ...
	DToPointer(char *CGenericSymbol, TypeQualifierList_ *CGenericSymbol, Pointer_ *CGenericSymbol) (*CGenericSymbol, error)

	// 307:4: type_qualifier_list -> a: ...
	AToTypeQualifierList(TypeQualifier_ *CGenericSymbol) (*CGenericSymbol, error)

	// 308:4: type_qualifier_list -> b: ...
	BToTypeQualifierList(TypeQualifierList_ *CGenericSymbol, TypeQualifier_ *CGenericSymbol) (*CGenericSymbol, error)

	// 312:4: parameter_type_list -> a: ...
	AToParameterTypeList(ParameterList_ *CGenericSymbol) (*CGenericSymbol, error)

	// 313:4: parameter_type_list -> b: ...
	BToParameterTypeList(ParameterList_ *CGenericSymbol, char *CGenericSymbol, Ellipsis_ *CGenericSymbol) (*CGenericSymbol, error)

	// 316:4: parameter_list -> a: ...
	AToParameterList(ParameterDeclaration_ *CGenericSymbol) (*CGenericSymbol, error)

	// 317:4: parameter_list -> b: ...
	BToParameterList(ParameterList_ *CGenericSymbol, char *CGenericSymbol, ParameterDeclaration_ *CGenericSymbol) (*CGenericSymbol, error)

	// 320:4: parameter_declaration -> a: ...
	AToParameterDeclaration(DeclarationSpecifiers_ *CGenericSymbol, Declarator_ *CGenericSymbol) (*CGenericSymbol, error)

	// 321:4: parameter_declaration -> b: ...
	BToParameterDeclaration(DeclarationSpecifiers_ *CGenericSymbol, AbstractDeclarator_ *CGenericSymbol) (*CGenericSymbol, error)

	// 322:4: parameter_declaration -> c: ...
	CToParameterDeclaration(DeclarationSpecifiers_ *CGenericSymbol) (*CGenericSymbol, error)

	// 325:4: identifier_list -> a: ...
	AToIdentifierList(Identifier_ *CGenericSymbol) (*CGenericSymbol, error)

	// 326:4: identifier_list -> b: ...
	BToIdentifierList(IdentifierList_ *CGenericSymbol, char *CGenericSymbol, Identifier_ *CGenericSymbol) (*CGenericSymbol, error)

	// 329:4: type_name -> a: ...
	AToTypeName(SpecifierQualifierList_ *CGenericSymbol) (*CGenericSymbol, error)

	// 330:4: type_name -> b: ...
	BToTypeName(SpecifierQualifierList_ *CGenericSymbol, AbstractDeclarator_ *CGenericSymbol) (*CGenericSymbol, error)

	// 333:4: abstract_declarator -> a: ...
	AToAbstractDeclarator(Pointer_ *CGenericSymbol) (*CGenericSymbol, error)

	// 334:4: abstract_declarator -> b: ...
	BToAbstractDeclarator(DirectAbstractDeclarator_ *CGenericSymbol) (*CGenericSymbol, error)

	// 335:4: abstract_declarator -> c: ...
	CToAbstractDeclarator(Pointer_ *CGenericSymbol, DirectAbstractDeclarator_ *CGenericSymbol) (*CGenericSymbol, error)

	// 338:4: direct_abstract_declarator -> a: ...
	AToDirectAbstractDeclarator(char *CGenericSymbol, AbstractDeclarator_ *CGenericSymbol, char2 *CGenericSymbol) (*CGenericSymbol, error)

	// 339:4: direct_abstract_declarator -> b: ...
	BToDirectAbstractDeclarator(char *CGenericSymbol, char2 *CGenericSymbol) (*CGenericSymbol, error)

	// 340:4: direct_abstract_declarator -> c: ...
	CToDirectAbstractDeclarator(char *CGenericSymbol, ConstantExpression_ *CGenericSymbol, char2 *CGenericSymbol) (*CGenericSymbol, error)

	// 341:4: direct_abstract_declarator -> d: ...
	DToDirectAbstractDeclarator(DirectAbstractDeclarator_ *CGenericSymbol, char *CGenericSymbol, char2 *CGenericSymbol) (*CGenericSymbol, error)

	// 342:4: direct_abstract_declarator -> e: ...
	EToDirectAbstractDeclarator(DirectAbstractDeclarator_ *CGenericSymbol, char *CGenericSymbol, ConstantExpression_ *CGenericSymbol, char2 *CGenericSymbol) (*CGenericSymbol, error)

	// 343:4: direct_abstract_declarator -> f: ...
	FToDirectAbstractDeclarator(char *CGenericSymbol, char2 *CGenericSymbol) (*CGenericSymbol, error)

	// 344:4: direct_abstract_declarator -> g: ...
	GToDirectAbstractDeclarator(char *CGenericSymbol, ParameterTypeList_ *CGenericSymbol, char2 *CGenericSymbol) (*CGenericSymbol, error)

	// 345:4: direct_abstract_declarator -> h: ...
	HToDirectAbstractDeclarator(DirectAbstractDeclarator_ *CGenericSymbol, char *CGenericSymbol, char2 *CGenericSymbol) (*CGenericSymbol, error)

	// 346:4: direct_abstract_declarator -> i: ...
	IToDirectAbstractDeclarator(DirectAbstractDeclarator_ *CGenericSymbol, char *CGenericSymbol, ParameterTypeList_ *CGenericSymbol, char2 *CGenericSymbol) (*CGenericSymbol, error)

	// 349:4: initializer -> a: ...
	AToInitializer(AssignmentExpression_ *CGenericSymbol) (*CGenericSymbol, error)

	// 350:4: initializer -> b: ...
	BToInitializer(char *CGenericSymbol, InitializerList_ *CGenericSymbol, char2 *CGenericSymbol) (*CGenericSymbol, error)

	// 351:4: initializer -> c: ...
	CToInitializer(char *CGenericSymbol, InitializerList_ *CGenericSymbol, char2 *CGenericSymbol, char3 *CGenericSymbol) (*CGenericSymbol, error)

	// 354:4: initializer_list -> a: ...
	AToInitializerList(Initializer_ *CGenericSymbol) (*CGenericSymbol, error)

	// 355:4: initializer_list -> b: ...
	BToInitializerList(InitializerList_ *CGenericSymbol, char *CGenericSymbol, Initializer_ *CGenericSymbol) (*CGenericSymbol, error)

	// 358:4: statement -> a: ...
	AToStatement(LabeledStatement_ *CGenericSymbol) (*CGenericSymbol, error)

	// 359:4: statement -> b: ...
	BToStatement(CompoundStatement_ *CGenericSymbol) (*CGenericSymbol, error)

	// 360:4: statement -> c: ...
	CToStatement(ExpressionStatement_ *CGenericSymbol) (*CGenericSymbol, error)

	// 361:4: statement -> d: ...
	DToStatement(SelectionStatement_ *CGenericSymbol) (*CGenericSymbol, error)

	// 362:4: statement -> e: ...
	EToStatement(IterationStatement_ *CGenericSymbol) (*CGenericSymbol, error)

	// 363:4: statement -> f: ...
	FToStatement(JumpStatement_ *CGenericSymbol) (*CGenericSymbol, error)

	// 366:4: labeled_statement -> a: ...
	AToLabeledStatement(Identifier_ *CGenericSymbol, char *CGenericSymbol, Statement_ *CGenericSymbol) (*CGenericSymbol, error)

	// 367:4: labeled_statement -> b: ...
	BToLabeledStatement(Case_ *CGenericSymbol, ConstantExpression_ *CGenericSymbol, char *CGenericSymbol, Statement_ *CGenericSymbol) (*CGenericSymbol, error)

	// 368:4: labeled_statement -> c: ...
	CToLabeledStatement(Default_ *CGenericSymbol, char *CGenericSymbol, Statement_ *CGenericSymbol) (*CGenericSymbol, error)

	// 371:4: compound_statement -> a: ...
	AToCompoundStatement(char *CGenericSymbol, char2 *CGenericSymbol) (*CGenericSymbol, error)

	// 372:4: compound_statement -> b: ...
	BToCompoundStatement(char *CGenericSymbol, StatementList_ *CGenericSymbol, char2 *CGenericSymbol) (*CGenericSymbol, error)

	// 373:4: compound_statement -> c: ...
	CToCompoundStatement(char *CGenericSymbol, DeclarationList_ *CGenericSymbol, char2 *CGenericSymbol) (*CGenericSymbol, error)

	// 374:4: compound_statement -> d: ...
	DToCompoundStatement(char *CGenericSymbol, DeclarationList_ *CGenericSymbol, StatementList_ *CGenericSymbol, char2 *CGenericSymbol) (*CGenericSymbol, error)

	// 377:4: declaration_list -> a: ...
	AToDeclarationList(Declaration_ *CGenericSymbol) (*CGenericSymbol, error)

	// 378:4: declaration_list -> b: ...
	BToDeclarationList(DeclarationList_ *CGenericSymbol, Declaration_ *CGenericSymbol) (*CGenericSymbol, error)

	// 381:4: statement_list -> a: ...
	AToStatementList(Statement_ *CGenericSymbol) (*CGenericSymbol, error)

	// 382:4: statement_list -> b: ...
	BToStatementList(StatementList_ *CGenericSymbol, Statement_ *CGenericSymbol) (*CGenericSymbol, error)

	// 385:4: expression_statement -> a: ...
	AToExpressionStatement(char *CGenericSymbol) (*CGenericSymbol, error)

	// 386:4: expression_statement -> b: ...
	BToExpressionStatement(Expression_ *CGenericSymbol, char *CGenericSymbol) (*CGenericSymbol, error)

	// 389:4: selection_statement -> a: ...
	AToSelectionStatement(If_ *CGenericSymbol, char *CGenericSymbol, Expression_ *CGenericSymbol, char2 *CGenericSymbol, Statement_ *CGenericSymbol) (*CGenericSymbol, error)

	// 390:4: selection_statement -> b: ...
	BToSelectionStatement(If_ *CGenericSymbol, char *CGenericSymbol, Expression_ *CGenericSymbol, char2 *CGenericSymbol, Statement_ *CGenericSymbol, Else_ *CGenericSymbol, Statement_2 *CGenericSymbol) (*CGenericSymbol, error)

	// 391:4: selection_statement -> c: ...
	CToSelectionStatement(Switch_ *CGenericSymbol, char *CGenericSymbol, Expression_ *CGenericSymbol, char2 *CGenericSymbol, Statement_ *CGenericSymbol) (*CGenericSymbol, error)

	// 394:4: iteration_statement -> a: ...
	AToIterationStatement(While_ *CGenericSymbol, char *CGenericSymbol, Expression_ *CGenericSymbol, char2 *CGenericSymbol, Statement_ *CGenericSymbol) (*CGenericSymbol, error)

	// 395:4: iteration_statement -> b: ...
	BToIterationStatement(Do_ *CGenericSymbol, Statement_ *CGenericSymbol, While_ *CGenericSymbol, char *CGenericSymbol, Expression_ *CGenericSymbol, char2 *CGenericSymbol, char3 *CGenericSymbol) (*CGenericSymbol, error)

	// 396:4: iteration_statement -> c: ...
	CToIterationStatement(For_ *CGenericSymbol, char *CGenericSymbol, ExpressionStatement_ *CGenericSymbol, ExpressionStatement_2 *CGenericSymbol, char2 *CGenericSymbol, Statement_ *CGenericSymbol) (*CGenericSymbol, error)

	// 397:4: iteration_statement -> d: ...
	DToIterationStatement(For_ *CGenericSymbol, char *CGenericSymbol, ExpressionStatement_ *CGenericSymbol, ExpressionStatement_2 *CGenericSymbol, Expression_ *CGenericSymbol, char2 *CGenericSymbol, Statement_ *CGenericSymbol) (*CGenericSymbol, error)

	// 400:4: jump_statement -> a: ...
	AToJumpStatement(Goto_ *CGenericSymbol, Identifier_ *CGenericSymbol, char *CGenericSymbol) (*CGenericSymbol, error)

	// 401:4: jump_statement -> b: ...
	BToJumpStatement(Continue_ *CGenericSymbol, char *CGenericSymbol) (*CGenericSymbol, error)

	// 402:4: jump_statement -> c: ...
	CToJumpStatement(Break_ *CGenericSymbol, char *CGenericSymbol) (*CGenericSymbol, error)

	// 403:4: jump_statement -> d: ...
	DToJumpStatement(Return_ *CGenericSymbol, char *CGenericSymbol) (*CGenericSymbol, error)

	// 404:4: jump_statement -> e: ...
	EToJumpStatement(Return_ *CGenericSymbol, Expression_ *CGenericSymbol, char *CGenericSymbol) (*CGenericSymbol, error)

	// 407:4: translation_unit -> a: ...
	AToTranslationUnit(ExternalDeclaration_ *CGenericSymbol) (*CGenericSymbol, error)

	// 408:4: translation_unit -> b: ...
	BToTranslationUnit(TranslationUnit_ *CGenericSymbol, ExternalDeclaration_ *CGenericSymbol) (*CGenericSymbol, error)

	// 411:4: external_declaration -> a: ...
	AToExternalDeclaration(FunctionDefinition_ *CGenericSymbol) (*CGenericSymbol, error)

	// 412:4: external_declaration -> b: ...
	BToExternalDeclaration(Declaration_ *CGenericSymbol) (*CGenericSymbol, error)

	// 415:4: function_definition -> a: ...
	AToFunctionDefinition(DeclarationSpecifiers_ *CGenericSymbol, Declarator_ *CGenericSymbol, DeclarationList_ *CGenericSymbol, CompoundStatement_ *CGenericSymbol) (*CGenericSymbol, error)

	// 416:4: function_definition -> b: ...
	BToFunctionDefinition(DeclarationSpecifiers_ *CGenericSymbol, Declarator_ *CGenericSymbol, CompoundStatement_ *CGenericSymbol) (*CGenericSymbol, error)

	// 417:4: function_definition -> c: ...
	CToFunctionDefinition(Declarator_ *CGenericSymbol, DeclarationList_ *CGenericSymbol, CompoundStatement_ *CGenericSymbol) (*CGenericSymbol, error)

	// 418:4: function_definition -> d: ...
	DToFunctionDefinition(Declarator_ *CGenericSymbol, CompoundStatement_ *CGenericSymbol) (*CGenericSymbol, error)
}

type CParseErrorHandler interface {
	Error(nextToken CToken, parseStack _CStack) error
}

type CDefaultParseErrorHandler struct{}

func (CDefaultParseErrorHandler) Error(nextToken CToken, stack _CStack) error {
	return fmt.Errorf(
		"Syntax error: unexpected symbol %v. Expecting %v (%v)",
		nextToken.Id(),
		CExpectedTerminals(stack[len(stack)-1].StateId),
		nextToken.Loc())
}

func CExpectedTerminals(id _CStateId) []CSymbolId {
	result := []CSymbolId{}
	for key, _ := range _CActionTable {
		if key._CStateId != id {
			continue
		}
		result = append(result, key.CSymbolId)
	}

	sort.Slice(result, func(i int, j int) bool { return result[i] < result[j] })
	return result
}

func CParse(lexer CLexer, reducer CReducer) (*CGenericSymbol, error) {

	return CParseWithCustomErrorHandler(
		lexer,
		reducer,
		CDefaultParseErrorHandler{})
}

func CParseWithCustomErrorHandler(
	lexer CLexer,
	reducer CReducer,
	errHandler CParseErrorHandler) (
	*CGenericSymbol,
	error) {

	item, err := _CParse(lexer, reducer, errHandler, _CState1)
	if err != nil {
		var errRetVal *CGenericSymbol
		return errRetVal, err
	}
	return item.Generic_, nil
}

// ================================================================
// Parser internal implementation
// User should normally avoid directly accessing the following code
// ================================================================

func _CParse(
	lexer CLexer,
	reducer CReducer,
	errHandler CParseErrorHandler,
	startState _CStateId) (
	*_CStackItem,
	error) {

	stateStack := _CStack{
		// Note: we don't have to populate the start symbol since its value
		// is never accessed.
		&_CStackItem{startState, nil},
	}

	symbolStack := &_CPseudoSymbolStack{lexer: lexer}

	for {
		nextSymbol, err := symbolStack.Top()
		if err != nil {
			return nil, err
		}

		action, ok := _CActionTable.Get(
			stateStack[len(stateStack)-1].StateId,
			nextSymbol.Id())
		if !ok {
			return nil, errHandler.Error(nextSymbol, stateStack)
		}

		if action.ActionType == _CShiftAction {
			stateStack = append(stateStack, action.ShiftItem(nextSymbol))

			_, err = symbolStack.Pop()
			if err != nil {
				return nil, err
			}
		} else if action.ActionType == _CReduceAction {
			var reduceSymbol *CSymbol
			stateStack, reduceSymbol, err = action.ReduceSymbol(
				reducer,
				stateStack)
			if err != nil {
				return nil, err
			}

			symbolStack.Push(reduceSymbol)
		} else if action.ActionType == _CAcceptAction {
			if len(stateStack) != 2 {
				panic("This should never happen")
			}
			return stateStack[1], nil
		} else {
			panic("Unknown action type: " + action.ActionType.String())
		}
	}
}

func (i CSymbolId) String() string {
	switch i {
	case _CEndMarker:
		return "$"
	case _CWildcardMarker:
		return "*"
	case CIdentifierToken:
		return "IDENTIFIER"
	case CConstantToken:
		return "CONSTANT"
	case CStringLiteralToken:
		return "STRING_LITERAL"
	case CSizeofToken:
		return "SIZEOF"
	case CPtrOpToken:
		return "PTR_OP"
	case CIncOpToken:
		return "INC_OP"
	case CDecOpToken:
		return "DEC_OP"
	case CLeftOpToken:
		return "LEFT_OP"
	case CRightOpToken:
		return "RIGHT_OP"
	case CLeOpToken:
		return "LE_OP"
	case CGeOpToken:
		return "GE_OP"
	case CEqOpToken:
		return "EQ_OP"
	case CNeOpToken:
		return "NE_OP"
	case CAndOpToken:
		return "AND_OP"
	case COrOpToken:
		return "OR_OP"
	case CMulAssignToken:
		return "MUL_ASSIGN"
	case CDivAssignToken:
		return "DIV_ASSIGN"
	case CModAssignToken:
		return "MOD_ASSIGN"
	case CAddAssignToken:
		return "ADD_ASSIGN"
	case CSubAssignToken:
		return "SUB_ASSIGN"
	case CLeftAssignToken:
		return "LEFT_ASSIGN"
	case CRightAssignToken:
		return "RIGHT_ASSIGN"
	case CAndAssignToken:
		return "AND_ASSIGN"
	case CXorAssignToken:
		return "XOR_ASSIGN"
	case COrAssignToken:
		return "OR_ASSIGN"
	case CTypeNameToken:
		return "TYPE_NAME"
	case CTypedefToken:
		return "TYPEDEF"
	case CExternToken:
		return "EXTERN"
	case CStaticToken:
		return "STATIC"
	case CAutoToken:
		return "AUTO"
	case CRegisterToken:
		return "REGISTER"
	case CCharToken:
		return "CHAR"
	case CShortToken:
		return "SHORT"
	case CIntToken:
		return "INT"
	case CLongToken:
		return "LONG"
	case CSignedToken:
		return "SIGNED"
	case CUnsignedToken:
		return "UNSIGNED"
	case CFloatToken:
		return "FLOAT"
	case CDoubleToken:
		return "DOUBLE"
	case CConstToken:
		return "CONST"
	case CVolatileToken:
		return "VOLATILE"
	case CVoidToken:
		return "VOID"
	case CStructToken:
		return "STRUCT"
	case CUnionToken:
		return "UNION"
	case CEnumToken:
		return "ENUM"
	case CEllipsisToken:
		return "ELLIPSIS"
	case CCaseToken:
		return "CASE"
	case CDefaultToken:
		return "DEFAULT"
	case CIfToken:
		return "IF"
	case CElseToken:
		return "ELSE"
	case CSwitchToken:
		return "SWITCH"
	case CWhileToken:
		return "WHILE"
	case CDoToken:
		return "DO"
	case CForToken:
		return "FOR"
	case CGotoToken:
		return "GOTO"
	case CContinueToken:
		return "CONTINUE"
	case CBreakToken:
		return "BREAK"
	case CReturnToken:
		return "RETURN"
	case '(':
		return "'('"
	case ')':
		return "')'"
	case '{':
		return "'{'"
	case '}':
		return "'}'"
	case '[':
		return "'['"
	case ']':
		return "']'"
	case ';':
		return "';'"
	case ':':
		return "':'"
	case ',':
		return "','"
	case '=':
		return "'='"
	case '?':
		return "'?'"
	case '*':
		return "'*'"
	case '/':
		return "'/'"
	case '-':
		return "'-'"
	case '+':
		return "'+'"
	case '%':
		return "'%'"
	case '&':
		return "'&'"
	case '|':
		return "'|'"
	case '!':
		return "'!'"
	case '.':
		return "'.'"
	case '^':
		return "'^'"
	case '<':
		return "'<'"
	case '>':
		return "'>'"
	case '~':
		return "'~'"
	case CPrimaryExpressionType:
		return "primary_expression"
	case CPostfixExpressionType:
		return "postfix_expression"
	case CArgumentExpressionListType:
		return "argument_expression_list"
	case CUnaryExpressionType:
		return "unary_expression"
	case CUnaryOperatorType:
		return "unary_operator"
	case CCastExpressionType:
		return "cast_expression"
	case CMultiplicativeExpressionType:
		return "multiplicative_expression"
	case CAdditiveExpressionType:
		return "additive_expression"
	case CShiftExpressionType:
		return "shift_expression"
	case CRelationalExpressionType:
		return "relational_expression"
	case CEqualityExpressionType:
		return "equality_expression"
	case CAndExpressionType:
		return "and_expression"
	case CExclusiveOrExpressionType:
		return "exclusive_or_expression"
	case CInclusiveOrExpressionType:
		return "inclusive_or_expression"
	case CLogicalAndExpressionType:
		return "logical_and_expression"
	case CLogicalOrExpressionType:
		return "logical_or_expression"
	case CConditionalExpressionType:
		return "conditional_expression"
	case CAssignmentExpressionType:
		return "assignment_expression"
	case CAssignmentOperatorType:
		return "assignment_operator"
	case CExpressionType:
		return "expression"
	case CConstantExpressionType:
		return "constant_expression"
	case CDeclarationType:
		return "declaration"
	case CDeclarationSpecifiersType:
		return "declaration_specifiers"
	case CInitDeclaratorListType:
		return "init_declarator_list"
	case CInitDeclaratorType:
		return "init_declarator"
	case CStorageClassSpecifierType:
		return "storage_class_specifier"
	case CTypeSpecifierType:
		return "type_specifier"
	case CStructOrUnionSpecifierType:
		return "struct_or_union_specifier"
	case CStructOrUnionType:
		return "struct_or_union"
	case CStructDeclarationListType:
		return "struct_declaration_list"
	case CStructDeclarationType:
		return "struct_declaration"
	case CSpecifierQualifierListType:
		return "specifier_qualifier_list"
	case CStructDeclaratorListType:
		return "struct_declarator_list"
	case CStructDeclaratorType:
		return "struct_declarator"
	case CEnumSpecifierType:
		return "enum_specifier"
	case CEnumeratorListType:
		return "enumerator_list"
	case CEnumeratorType:
		return "enumerator"
	case CTypeQualifierType:
		return "type_qualifier"
	case CDeclaratorType:
		return "declarator"
	case CDirectDeclaratorType:
		return "direct_declarator"
	case CPointerType:
		return "pointer"
	case CTypeQualifierListType:
		return "type_qualifier_list"
	case CParameterTypeListType:
		return "parameter_type_list"
	case CParameterListType:
		return "parameter_list"
	case CParameterDeclarationType:
		return "parameter_declaration"
	case CIdentifierListType:
		return "identifier_list"
	case CTypeNameType:
		return "type_name"
	case CAbstractDeclaratorType:
		return "abstract_declarator"
	case CDirectAbstractDeclaratorType:
		return "direct_abstract_declarator"
	case CInitializerType:
		return "initializer"
	case CInitializerListType:
		return "initializer_list"
	case CStatementType:
		return "statement"
	case CLabeledStatementType:
		return "labeled_statement"
	case CCompoundStatementType:
		return "compound_statement"
	case CDeclarationListType:
		return "declaration_list"
	case CStatementListType:
		return "statement_list"
	case CExpressionStatementType:
		return "expression_statement"
	case CSelectionStatementType:
		return "selection_statement"
	case CIterationStatementType:
		return "iteration_statement"
	case CJumpStatementType:
		return "jump_statement"
	case CTranslationUnitType:
		return "translation_unit"
	case CExternalDeclarationType:
		return "external_declaration"
	case CFunctionDefinitionType:
		return "function_definition"
	default:
		return fmt.Sprintf("?unknown symbol %d?", int(i))
	}
}

const (
	_CEndMarker      = CSymbolId(0)
	_CWildcardMarker = CSymbolId(-1)

	CPrimaryExpressionType        = CSymbolId(338)
	CPostfixExpressionType        = CSymbolId(339)
	CArgumentExpressionListType   = CSymbolId(340)
	CUnaryExpressionType          = CSymbolId(341)
	CUnaryOperatorType            = CSymbolId(342)
	CCastExpressionType           = CSymbolId(343)
	CMultiplicativeExpressionType = CSymbolId(344)
	CAdditiveExpressionType       = CSymbolId(345)
	CShiftExpressionType          = CSymbolId(346)
	CRelationalExpressionType     = CSymbolId(347)
	CEqualityExpressionType       = CSymbolId(348)
	CAndExpressionType            = CSymbolId(349)
	CExclusiveOrExpressionType    = CSymbolId(350)
	CInclusiveOrExpressionType    = CSymbolId(351)
	CLogicalAndExpressionType     = CSymbolId(352)
	CLogicalOrExpressionType      = CSymbolId(353)
	CConditionalExpressionType    = CSymbolId(354)
	CAssignmentExpressionType     = CSymbolId(355)
	CAssignmentOperatorType       = CSymbolId(356)
	CExpressionType               = CSymbolId(357)
	CConstantExpressionType       = CSymbolId(358)
	CDeclarationType              = CSymbolId(359)
	CDeclarationSpecifiersType    = CSymbolId(360)
	CInitDeclaratorListType       = CSymbolId(361)
	CInitDeclaratorType           = CSymbolId(362)
	CStorageClassSpecifierType    = CSymbolId(363)
	CTypeSpecifierType            = CSymbolId(364)
	CStructOrUnionSpecifierType   = CSymbolId(365)
	CStructOrUnionType            = CSymbolId(366)
	CStructDeclarationListType    = CSymbolId(367)
	CStructDeclarationType        = CSymbolId(368)
	CSpecifierQualifierListType   = CSymbolId(369)
	CStructDeclaratorListType     = CSymbolId(370)
	CStructDeclaratorType         = CSymbolId(371)
	CEnumSpecifierType            = CSymbolId(372)
	CEnumeratorListType           = CSymbolId(373)
	CEnumeratorType               = CSymbolId(374)
	CTypeQualifierType            = CSymbolId(375)
	CDeclaratorType               = CSymbolId(376)
	CDirectDeclaratorType         = CSymbolId(377)
	CPointerType                  = CSymbolId(378)
	CTypeQualifierListType        = CSymbolId(379)
	CParameterTypeListType        = CSymbolId(380)
	CParameterListType            = CSymbolId(381)
	CParameterDeclarationType     = CSymbolId(382)
	CIdentifierListType           = CSymbolId(383)
	CTypeNameType                 = CSymbolId(384)
	CAbstractDeclaratorType       = CSymbolId(385)
	CDirectAbstractDeclaratorType = CSymbolId(386)
	CInitializerType              = CSymbolId(387)
	CInitializerListType          = CSymbolId(388)
	CStatementType                = CSymbolId(389)
	CLabeledStatementType         = CSymbolId(390)
	CCompoundStatementType        = CSymbolId(391)
	CDeclarationListType          = CSymbolId(392)
	CStatementListType            = CSymbolId(393)
	CExpressionStatementType      = CSymbolId(394)
	CSelectionStatementType       = CSymbolId(395)
	CIterationStatementType       = CSymbolId(396)
	CJumpStatementType            = CSymbolId(397)
	CTranslationUnitType          = CSymbolId(398)
	CExternalDeclarationType      = CSymbolId(399)
	CFunctionDefinitionType       = CSymbolId(400)
)

type _CActionType int

const (
	// NOTE: error action is implicit
	_CShiftAction  = _CActionType(0)
	_CReduceAction = _CActionType(1)
	_CAcceptAction = _CActionType(2)
)

func (i _CActionType) String() string {
	switch i {
	case _CShiftAction:
		return "shift"
	case _CReduceAction:
		return "reduce"
	case _CAcceptAction:
		return "accept"
	default:
		return fmt.Sprintf("?Unknown action %d?", int(i))
	}
}

type _CReduceType int

const (
	_CReduceAToPrimaryExpression        = _CReduceType(1)
	_CReduceBToPrimaryExpression        = _CReduceType(2)
	_CReduceCToPrimaryExpression        = _CReduceType(3)
	_CReduceDToPrimaryExpression        = _CReduceType(4)
	_CReduceAToPostfixExpression        = _CReduceType(5)
	_CReduceBToPostfixExpression        = _CReduceType(6)
	_CReduceCToPostfixExpression        = _CReduceType(7)
	_CReduceDToPostfixExpression        = _CReduceType(8)
	_CReduceEToPostfixExpression        = _CReduceType(9)
	_CReduceFToPostfixExpression        = _CReduceType(10)
	_CReduceGToPostfixExpression        = _CReduceType(11)
	_CReduceHToPostfixExpression        = _CReduceType(12)
	_CReduceAToArgumentExpressionList   = _CReduceType(13)
	_CReduceBToArgumentExpressionList   = _CReduceType(14)
	_CReduceAToUnaryExpression          = _CReduceType(15)
	_CReduceBToUnaryExpression          = _CReduceType(16)
	_CReduceCToUnaryExpression          = _CReduceType(17)
	_CReduceDToUnaryExpression          = _CReduceType(18)
	_CReduceEToUnaryExpression          = _CReduceType(19)
	_CReduceFToUnaryExpression          = _CReduceType(20)
	_CReduceAToUnaryOperator            = _CReduceType(21)
	_CReduceBToUnaryOperator            = _CReduceType(22)
	_CReduceCToUnaryOperator            = _CReduceType(23)
	_CReduceDToUnaryOperator            = _CReduceType(24)
	_CReduceEToUnaryOperator            = _CReduceType(25)
	_CReduceFToUnaryOperator            = _CReduceType(26)
	_CReduceAToCastExpression           = _CReduceType(27)
	_CReduceBToCastExpression           = _CReduceType(28)
	_CReduceAToMultiplicativeExpression = _CReduceType(29)
	_CReduceBToMultiplicativeExpression = _CReduceType(30)
	_CReduceCToMultiplicativeExpression = _CReduceType(31)
	_CReduceDToMultiplicativeExpression = _CReduceType(32)
	_CReduceAToAdditiveExpression       = _CReduceType(33)
	_CReduceBToAdditiveExpression       = _CReduceType(34)
	_CReduceCToAdditiveExpression       = _CReduceType(35)
	_CReduceAToShiftExpression          = _CReduceType(36)
	_CReduceBToShiftExpression          = _CReduceType(37)
	_CReduceCToShiftExpression          = _CReduceType(38)
	_CReduceAToRelationalExpression     = _CReduceType(39)
	_CReduceBToRelationalExpression     = _CReduceType(40)
	_CReduceCToRelationalExpression     = _CReduceType(41)
	_CReduceDToRelationalExpression     = _CReduceType(42)
	_CReduceEToRelationalExpression     = _CReduceType(43)
	_CReduceAToEqualityExpression       = _CReduceType(44)
	_CReduceBToEqualityExpression       = _CReduceType(45)
	_CReduceCToEqualityExpression       = _CReduceType(46)
	_CReduceAToAndExpression            = _CReduceType(47)
	_CReduceBToAndExpression            = _CReduceType(48)
	_CReduceAToExclusiveOrExpression    = _CReduceType(49)
	_CReduceBToExclusiveOrExpression    = _CReduceType(50)
	_CReduceAToInclusiveOrExpression    = _CReduceType(51)
	_CReduceBToInclusiveOrExpression    = _CReduceType(52)
	_CReduceAToLogicalAndExpression     = _CReduceType(53)
	_CReduceBToLogicalAndExpression     = _CReduceType(54)
	_CReduceAToLogicalOrExpression      = _CReduceType(55)
	_CReduceBToLogicalOrExpression      = _CReduceType(56)
	_CReduceAToConditionalExpression    = _CReduceType(57)
	_CReduceBToConditionalExpression    = _CReduceType(58)
	_CReduceAToAssignmentExpression     = _CReduceType(59)
	_CReduceBToAssignmentExpression     = _CReduceType(60)
	_CReduceAToAssignmentOperator       = _CReduceType(61)
	_CReduceBToAssignmentOperator       = _CReduceType(62)
	_CReduceCToAssignmentOperator       = _CReduceType(63)
	_CReduceDToAssignmentOperator       = _CReduceType(64)
	_CReduceEToAssignmentOperator       = _CReduceType(65)
	_CReduceFToAssignmentOperator       = _CReduceType(66)
	_CReduceGToAssignmentOperator       = _CReduceType(67)
	_CReduceHToAssignmentOperator       = _CReduceType(68)
	_CReduceIToAssignmentOperator       = _CReduceType(69)
	_CReduceJToAssignmentOperator       = _CReduceType(70)
	_CReduceKToAssignmentOperator       = _CReduceType(71)
	_CReduceAToExpression               = _CReduceType(72)
	_CReduceBToExpression               = _CReduceType(73)
	_CReduceAToConstantExpression       = _CReduceType(74)
	_CReduceAToDeclaration              = _CReduceType(75)
	_CReduceBToDeclaration              = _CReduceType(76)
	_CReduceAToDeclarationSpecifiers    = _CReduceType(77)
	_CReduceBToDeclarationSpecifiers    = _CReduceType(78)
	_CReduceCToDeclarationSpecifiers    = _CReduceType(79)
	_CReduceDToDeclarationSpecifiers    = _CReduceType(80)
	_CReduceEToDeclarationSpecifiers    = _CReduceType(81)
	_CReduceFToDeclarationSpecifiers    = _CReduceType(82)
	_CReduceAToInitDeclaratorList       = _CReduceType(83)
	_CReduceBToInitDeclaratorList       = _CReduceType(84)
	_CReduceAToInitDeclarator           = _CReduceType(85)
	_CReduceBToInitDeclarator           = _CReduceType(86)
	_CReduceAToStorageClassSpecifier    = _CReduceType(87)
	_CReduceBToStorageClassSpecifier    = _CReduceType(88)
	_CReduceCToStorageClassSpecifier    = _CReduceType(89)
	_CReduceDToStorageClassSpecifier    = _CReduceType(90)
	_CReduceEToStorageClassSpecifier    = _CReduceType(91)
	_CReduceAToTypeSpecifier            = _CReduceType(92)
	_CReduceBToTypeSpecifier            = _CReduceType(93)
	_CReduceCToTypeSpecifier            = _CReduceType(94)
	_CReduceDToTypeSpecifier            = _CReduceType(95)
	_CReduceEToTypeSpecifier            = _CReduceType(96)
	_CReduceFToTypeSpecifier            = _CReduceType(97)
	_CReduceGToTypeSpecifier            = _CReduceType(98)
	_CReduceHToTypeSpecifier            = _CReduceType(99)
	_CReduceIToTypeSpecifier            = _CReduceType(100)
	_CReduceJToTypeSpecifier            = _CReduceType(101)
	_CReduceKToTypeSpecifier            = _CReduceType(102)
	_CReduceLToTypeSpecifier            = _CReduceType(103)
	_CReduceAToStructOrUnionSpecifier   = _CReduceType(104)
	_CReduceBToStructOrUnionSpecifier   = _CReduceType(105)
	_CReduceCToStructOrUnionSpecifier   = _CReduceType(106)
	_CReduceAToStructOrUnion            = _CReduceType(107)
	_CReduceBToStructOrUnion            = _CReduceType(108)
	_CReduceAToStructDeclarationList    = _CReduceType(109)
	_CReduceBToStructDeclarationList    = _CReduceType(110)
	_CReduceAToStructDeclaration        = _CReduceType(111)
	_CReduceAToSpecifierQualifierList   = _CReduceType(112)
	_CReduceBToSpecifierQualifierList   = _CReduceType(113)
	_CReduceCToSpecifierQualifierList   = _CReduceType(114)
	_CReduceDToSpecifierQualifierList   = _CReduceType(115)
	_CReduceAToStructDeclaratorList     = _CReduceType(116)
	_CReduceBToStructDeclaratorList     = _CReduceType(117)
	_CReduceAToStructDeclarator         = _CReduceType(118)
	_CReduceBToStructDeclarator         = _CReduceType(119)
	_CReduceCToStructDeclarator         = _CReduceType(120)
	_CReduceAToEnumSpecifier            = _CReduceType(121)
	_CReduceBToEnumSpecifier            = _CReduceType(122)
	_CReduceCToEnumSpecifier            = _CReduceType(123)
	_CReduceAToEnumeratorList           = _CReduceType(124)
	_CReduceBToEnumeratorList           = _CReduceType(125)
	_CReduceAToEnumerator               = _CReduceType(126)
	_CReduceBToEnumerator               = _CReduceType(127)
	_CReduceAToTypeQualifier            = _CReduceType(128)
	_CReduceBToTypeQualifier            = _CReduceType(129)
	_CReduceAToDeclarator               = _CReduceType(130)
	_CReduceBToDeclarator               = _CReduceType(131)
	_CReduceAToDirectDeclarator         = _CReduceType(132)
	_CReduceBToDirectDeclarator         = _CReduceType(133)
	_CReduceCToDirectDeclarator         = _CReduceType(134)
	_CReduceDToDirectDeclarator         = _CReduceType(135)
	_CReduceEToDirectDeclarator         = _CReduceType(136)
	_CReduceFToDirectDeclarator         = _CReduceType(137)
	_CReduceGToDirectDeclarator         = _CReduceType(138)
	_CReduceAToPointer                  = _CReduceType(139)
	_CReduceBToPointer                  = _CReduceType(140)
	_CReduceCToPointer                  = _CReduceType(141)
	_CReduceDToPointer                  = _CReduceType(142)
	_CReduceAToTypeQualifierList        = _CReduceType(143)
	_CReduceBToTypeQualifierList        = _CReduceType(144)
	_CReduceAToParameterTypeList        = _CReduceType(145)
	_CReduceBToParameterTypeList        = _CReduceType(146)
	_CReduceAToParameterList            = _CReduceType(147)
	_CReduceBToParameterList            = _CReduceType(148)
	_CReduceAToParameterDeclaration     = _CReduceType(149)
	_CReduceBToParameterDeclaration     = _CReduceType(150)
	_CReduceCToParameterDeclaration     = _CReduceType(151)
	_CReduceAToIdentifierList           = _CReduceType(152)
	_CReduceBToIdentifierList           = _CReduceType(153)
	_CReduceAToTypeName                 = _CReduceType(154)
	_CReduceBToTypeName                 = _CReduceType(155)
	_CReduceAToAbstractDeclarator       = _CReduceType(156)
	_CReduceBToAbstractDeclarator       = _CReduceType(157)
	_CReduceCToAbstractDeclarator       = _CReduceType(158)
	_CReduceAToDirectAbstractDeclarator = _CReduceType(159)
	_CReduceBToDirectAbstractDeclarator = _CReduceType(160)
	_CReduceCToDirectAbstractDeclarator = _CReduceType(161)
	_CReduceDToDirectAbstractDeclarator = _CReduceType(162)
	_CReduceEToDirectAbstractDeclarator = _CReduceType(163)
	_CReduceFToDirectAbstractDeclarator = _CReduceType(164)
	_CReduceGToDirectAbstractDeclarator = _CReduceType(165)
	_CReduceHToDirectAbstractDeclarator = _CReduceType(166)
	_CReduceIToDirectAbstractDeclarator = _CReduceType(167)
	_CReduceAToInitializer              = _CReduceType(168)
	_CReduceBToInitializer              = _CReduceType(169)
	_CReduceCToInitializer              = _CReduceType(170)
	_CReduceAToInitializerList          = _CReduceType(171)
	_CReduceBToInitializerList          = _CReduceType(172)
	_CReduceAToStatement                = _CReduceType(173)
	_CReduceBToStatement                = _CReduceType(174)
	_CReduceCToStatement                = _CReduceType(175)
	_CReduceDToStatement                = _CReduceType(176)
	_CReduceEToStatement                = _CReduceType(177)
	_CReduceFToStatement                = _CReduceType(178)
	_CReduceAToLabeledStatement         = _CReduceType(179)
	_CReduceBToLabeledStatement         = _CReduceType(180)
	_CReduceCToLabeledStatement         = _CReduceType(181)
	_CReduceAToCompoundStatement        = _CReduceType(182)
	_CReduceBToCompoundStatement        = _CReduceType(183)
	_CReduceCToCompoundStatement        = _CReduceType(184)
	_CReduceDToCompoundStatement        = _CReduceType(185)
	_CReduceAToDeclarationList          = _CReduceType(186)
	_CReduceBToDeclarationList          = _CReduceType(187)
	_CReduceAToStatementList            = _CReduceType(188)
	_CReduceBToStatementList            = _CReduceType(189)
	_CReduceAToExpressionStatement      = _CReduceType(190)
	_CReduceBToExpressionStatement      = _CReduceType(191)
	_CReduceAToSelectionStatement       = _CReduceType(192)
	_CReduceBToSelectionStatement       = _CReduceType(193)
	_CReduceCToSelectionStatement       = _CReduceType(194)
	_CReduceAToIterationStatement       = _CReduceType(195)
	_CReduceBToIterationStatement       = _CReduceType(196)
	_CReduceCToIterationStatement       = _CReduceType(197)
	_CReduceDToIterationStatement       = _CReduceType(198)
	_CReduceAToJumpStatement            = _CReduceType(199)
	_CReduceBToJumpStatement            = _CReduceType(200)
	_CReduceCToJumpStatement            = _CReduceType(201)
	_CReduceDToJumpStatement            = _CReduceType(202)
	_CReduceEToJumpStatement            = _CReduceType(203)
	_CReduceAToTranslationUnit          = _CReduceType(204)
	_CReduceBToTranslationUnit          = _CReduceType(205)
	_CReduceAToExternalDeclaration      = _CReduceType(206)
	_CReduceBToExternalDeclaration      = _CReduceType(207)
	_CReduceAToFunctionDefinition       = _CReduceType(208)
	_CReduceBToFunctionDefinition       = _CReduceType(209)
	_CReduceCToFunctionDefinition       = _CReduceType(210)
	_CReduceDToFunctionDefinition       = _CReduceType(211)
)

func (i _CReduceType) String() string {
	switch i {
	case _CReduceAToPrimaryExpression:
		return "AToPrimaryExpression"
	case _CReduceBToPrimaryExpression:
		return "BToPrimaryExpression"
	case _CReduceCToPrimaryExpression:
		return "CToPrimaryExpression"
	case _CReduceDToPrimaryExpression:
		return "DToPrimaryExpression"
	case _CReduceAToPostfixExpression:
		return "AToPostfixExpression"
	case _CReduceBToPostfixExpression:
		return "BToPostfixExpression"
	case _CReduceCToPostfixExpression:
		return "CToPostfixExpression"
	case _CReduceDToPostfixExpression:
		return "DToPostfixExpression"
	case _CReduceEToPostfixExpression:
		return "EToPostfixExpression"
	case _CReduceFToPostfixExpression:
		return "FToPostfixExpression"
	case _CReduceGToPostfixExpression:
		return "GToPostfixExpression"
	case _CReduceHToPostfixExpression:
		return "HToPostfixExpression"
	case _CReduceAToArgumentExpressionList:
		return "AToArgumentExpressionList"
	case _CReduceBToArgumentExpressionList:
		return "BToArgumentExpressionList"
	case _CReduceAToUnaryExpression:
		return "AToUnaryExpression"
	case _CReduceBToUnaryExpression:
		return "BToUnaryExpression"
	case _CReduceCToUnaryExpression:
		return "CToUnaryExpression"
	case _CReduceDToUnaryExpression:
		return "DToUnaryExpression"
	case _CReduceEToUnaryExpression:
		return "EToUnaryExpression"
	case _CReduceFToUnaryExpression:
		return "FToUnaryExpression"
	case _CReduceAToUnaryOperator:
		return "AToUnaryOperator"
	case _CReduceBToUnaryOperator:
		return "BToUnaryOperator"
	case _CReduceCToUnaryOperator:
		return "CToUnaryOperator"
	case _CReduceDToUnaryOperator:
		return "DToUnaryOperator"
	case _CReduceEToUnaryOperator:
		return "EToUnaryOperator"
	case _CReduceFToUnaryOperator:
		return "FToUnaryOperator"
	case _CReduceAToCastExpression:
		return "AToCastExpression"
	case _CReduceBToCastExpression:
		return "BToCastExpression"
	case _CReduceAToMultiplicativeExpression:
		return "AToMultiplicativeExpression"
	case _CReduceBToMultiplicativeExpression:
		return "BToMultiplicativeExpression"
	case _CReduceCToMultiplicativeExpression:
		return "CToMultiplicativeExpression"
	case _CReduceDToMultiplicativeExpression:
		return "DToMultiplicativeExpression"
	case _CReduceAToAdditiveExpression:
		return "AToAdditiveExpression"
	case _CReduceBToAdditiveExpression:
		return "BToAdditiveExpression"
	case _CReduceCToAdditiveExpression:
		return "CToAdditiveExpression"
	case _CReduceAToShiftExpression:
		return "AToShiftExpression"
	case _CReduceBToShiftExpression:
		return "BToShiftExpression"
	case _CReduceCToShiftExpression:
		return "CToShiftExpression"
	case _CReduceAToRelationalExpression:
		return "AToRelationalExpression"
	case _CReduceBToRelationalExpression:
		return "BToRelationalExpression"
	case _CReduceCToRelationalExpression:
		return "CToRelationalExpression"
	case _CReduceDToRelationalExpression:
		return "DToRelationalExpression"
	case _CReduceEToRelationalExpression:
		return "EToRelationalExpression"
	case _CReduceAToEqualityExpression:
		return "AToEqualityExpression"
	case _CReduceBToEqualityExpression:
		return "BToEqualityExpression"
	case _CReduceCToEqualityExpression:
		return "CToEqualityExpression"
	case _CReduceAToAndExpression:
		return "AToAndExpression"
	case _CReduceBToAndExpression:
		return "BToAndExpression"
	case _CReduceAToExclusiveOrExpression:
		return "AToExclusiveOrExpression"
	case _CReduceBToExclusiveOrExpression:
		return "BToExclusiveOrExpression"
	case _CReduceAToInclusiveOrExpression:
		return "AToInclusiveOrExpression"
	case _CReduceBToInclusiveOrExpression:
		return "BToInclusiveOrExpression"
	case _CReduceAToLogicalAndExpression:
		return "AToLogicalAndExpression"
	case _CReduceBToLogicalAndExpression:
		return "BToLogicalAndExpression"
	case _CReduceAToLogicalOrExpression:
		return "AToLogicalOrExpression"
	case _CReduceBToLogicalOrExpression:
		return "BToLogicalOrExpression"
	case _CReduceAToConditionalExpression:
		return "AToConditionalExpression"
	case _CReduceBToConditionalExpression:
		return "BToConditionalExpression"
	case _CReduceAToAssignmentExpression:
		return "AToAssignmentExpression"
	case _CReduceBToAssignmentExpression:
		return "BToAssignmentExpression"
	case _CReduceAToAssignmentOperator:
		return "AToAssignmentOperator"
	case _CReduceBToAssignmentOperator:
		return "BToAssignmentOperator"
	case _CReduceCToAssignmentOperator:
		return "CToAssignmentOperator"
	case _CReduceDToAssignmentOperator:
		return "DToAssignmentOperator"
	case _CReduceEToAssignmentOperator:
		return "EToAssignmentOperator"
	case _CReduceFToAssignmentOperator:
		return "FToAssignmentOperator"
	case _CReduceGToAssignmentOperator:
		return "GToAssignmentOperator"
	case _CReduceHToAssignmentOperator:
		return "HToAssignmentOperator"
	case _CReduceIToAssignmentOperator:
		return "IToAssignmentOperator"
	case _CReduceJToAssignmentOperator:
		return "JToAssignmentOperator"
	case _CReduceKToAssignmentOperator:
		return "KToAssignmentOperator"
	case _CReduceAToExpression:
		return "AToExpression"
	case _CReduceBToExpression:
		return "BToExpression"
	case _CReduceAToConstantExpression:
		return "AToConstantExpression"
	case _CReduceAToDeclaration:
		return "AToDeclaration"
	case _CReduceBToDeclaration:
		return "BToDeclaration"
	case _CReduceAToDeclarationSpecifiers:
		return "AToDeclarationSpecifiers"
	case _CReduceBToDeclarationSpecifiers:
		return "BToDeclarationSpecifiers"
	case _CReduceCToDeclarationSpecifiers:
		return "CToDeclarationSpecifiers"
	case _CReduceDToDeclarationSpecifiers:
		return "DToDeclarationSpecifiers"
	case _CReduceEToDeclarationSpecifiers:
		return "EToDeclarationSpecifiers"
	case _CReduceFToDeclarationSpecifiers:
		return "FToDeclarationSpecifiers"
	case _CReduceAToInitDeclaratorList:
		return "AToInitDeclaratorList"
	case _CReduceBToInitDeclaratorList:
		return "BToInitDeclaratorList"
	case _CReduceAToInitDeclarator:
		return "AToInitDeclarator"
	case _CReduceBToInitDeclarator:
		return "BToInitDeclarator"
	case _CReduceAToStorageClassSpecifier:
		return "AToStorageClassSpecifier"
	case _CReduceBToStorageClassSpecifier:
		return "BToStorageClassSpecifier"
	case _CReduceCToStorageClassSpecifier:
		return "CToStorageClassSpecifier"
	case _CReduceDToStorageClassSpecifier:
		return "DToStorageClassSpecifier"
	case _CReduceEToStorageClassSpecifier:
		return "EToStorageClassSpecifier"
	case _CReduceAToTypeSpecifier:
		return "AToTypeSpecifier"
	case _CReduceBToTypeSpecifier:
		return "BToTypeSpecifier"
	case _CReduceCToTypeSpecifier:
		return "CToTypeSpecifier"
	case _CReduceDToTypeSpecifier:
		return "DToTypeSpecifier"
	case _CReduceEToTypeSpecifier:
		return "EToTypeSpecifier"
	case _CReduceFToTypeSpecifier:
		return "FToTypeSpecifier"
	case _CReduceGToTypeSpecifier:
		return "GToTypeSpecifier"
	case _CReduceHToTypeSpecifier:
		return "HToTypeSpecifier"
	case _CReduceIToTypeSpecifier:
		return "IToTypeSpecifier"
	case _CReduceJToTypeSpecifier:
		return "JToTypeSpecifier"
	case _CReduceKToTypeSpecifier:
		return "KToTypeSpecifier"
	case _CReduceLToTypeSpecifier:
		return "LToTypeSpecifier"
	case _CReduceAToStructOrUnionSpecifier:
		return "AToStructOrUnionSpecifier"
	case _CReduceBToStructOrUnionSpecifier:
		return "BToStructOrUnionSpecifier"
	case _CReduceCToStructOrUnionSpecifier:
		return "CToStructOrUnionSpecifier"
	case _CReduceAToStructOrUnion:
		return "AToStructOrUnion"
	case _CReduceBToStructOrUnion:
		return "BToStructOrUnion"
	case _CReduceAToStructDeclarationList:
		return "AToStructDeclarationList"
	case _CReduceBToStructDeclarationList:
		return "BToStructDeclarationList"
	case _CReduceAToStructDeclaration:
		return "AToStructDeclaration"
	case _CReduceAToSpecifierQualifierList:
		return "AToSpecifierQualifierList"
	case _CReduceBToSpecifierQualifierList:
		return "BToSpecifierQualifierList"
	case _CReduceCToSpecifierQualifierList:
		return "CToSpecifierQualifierList"
	case _CReduceDToSpecifierQualifierList:
		return "DToSpecifierQualifierList"
	case _CReduceAToStructDeclaratorList:
		return "AToStructDeclaratorList"
	case _CReduceBToStructDeclaratorList:
		return "BToStructDeclaratorList"
	case _CReduceAToStructDeclarator:
		return "AToStructDeclarator"
	case _CReduceBToStructDeclarator:
		return "BToStructDeclarator"
	case _CReduceCToStructDeclarator:
		return "CToStructDeclarator"
	case _CReduceAToEnumSpecifier:
		return "AToEnumSpecifier"
	case _CReduceBToEnumSpecifier:
		return "BToEnumSpecifier"
	case _CReduceCToEnumSpecifier:
		return "CToEnumSpecifier"
	case _CReduceAToEnumeratorList:
		return "AToEnumeratorList"
	case _CReduceBToEnumeratorList:
		return "BToEnumeratorList"
	case _CReduceAToEnumerator:
		return "AToEnumerator"
	case _CReduceBToEnumerator:
		return "BToEnumerator"
	case _CReduceAToTypeQualifier:
		return "AToTypeQualifier"
	case _CReduceBToTypeQualifier:
		return "BToTypeQualifier"
	case _CReduceAToDeclarator:
		return "AToDeclarator"
	case _CReduceBToDeclarator:
		return "BToDeclarator"
	case _CReduceAToDirectDeclarator:
		return "AToDirectDeclarator"
	case _CReduceBToDirectDeclarator:
		return "BToDirectDeclarator"
	case _CReduceCToDirectDeclarator:
		return "CToDirectDeclarator"
	case _CReduceDToDirectDeclarator:
		return "DToDirectDeclarator"
	case _CReduceEToDirectDeclarator:
		return "EToDirectDeclarator"
	case _CReduceFToDirectDeclarator:
		return "FToDirectDeclarator"
	case _CReduceGToDirectDeclarator:
		return "GToDirectDeclarator"
	case _CReduceAToPointer:
		return "AToPointer"
	case _CReduceBToPointer:
		return "BToPointer"
	case _CReduceCToPointer:
		return "CToPointer"
	case _CReduceDToPointer:
		return "DToPointer"
	case _CReduceAToTypeQualifierList:
		return "AToTypeQualifierList"
	case _CReduceBToTypeQualifierList:
		return "BToTypeQualifierList"
	case _CReduceAToParameterTypeList:
		return "AToParameterTypeList"
	case _CReduceBToParameterTypeList:
		return "BToParameterTypeList"
	case _CReduceAToParameterList:
		return "AToParameterList"
	case _CReduceBToParameterList:
		return "BToParameterList"
	case _CReduceAToParameterDeclaration:
		return "AToParameterDeclaration"
	case _CReduceBToParameterDeclaration:
		return "BToParameterDeclaration"
	case _CReduceCToParameterDeclaration:
		return "CToParameterDeclaration"
	case _CReduceAToIdentifierList:
		return "AToIdentifierList"
	case _CReduceBToIdentifierList:
		return "BToIdentifierList"
	case _CReduceAToTypeName:
		return "AToTypeName"
	case _CReduceBToTypeName:
		return "BToTypeName"
	case _CReduceAToAbstractDeclarator:
		return "AToAbstractDeclarator"
	case _CReduceBToAbstractDeclarator:
		return "BToAbstractDeclarator"
	case _CReduceCToAbstractDeclarator:
		return "CToAbstractDeclarator"
	case _CReduceAToDirectAbstractDeclarator:
		return "AToDirectAbstractDeclarator"
	case _CReduceBToDirectAbstractDeclarator:
		return "BToDirectAbstractDeclarator"
	case _CReduceCToDirectAbstractDeclarator:
		return "CToDirectAbstractDeclarator"
	case _CReduceDToDirectAbstractDeclarator:
		return "DToDirectAbstractDeclarator"
	case _CReduceEToDirectAbstractDeclarator:
		return "EToDirectAbstractDeclarator"
	case _CReduceFToDirectAbstractDeclarator:
		return "FToDirectAbstractDeclarator"
	case _CReduceGToDirectAbstractDeclarator:
		return "GToDirectAbstractDeclarator"
	case _CReduceHToDirectAbstractDeclarator:
		return "HToDirectAbstractDeclarator"
	case _CReduceIToDirectAbstractDeclarator:
		return "IToDirectAbstractDeclarator"
	case _CReduceAToInitializer:
		return "AToInitializer"
	case _CReduceBToInitializer:
		return "BToInitializer"
	case _CReduceCToInitializer:
		return "CToInitializer"
	case _CReduceAToInitializerList:
		return "AToInitializerList"
	case _CReduceBToInitializerList:
		return "BToInitializerList"
	case _CReduceAToStatement:
		return "AToStatement"
	case _CReduceBToStatement:
		return "BToStatement"
	case _CReduceCToStatement:
		return "CToStatement"
	case _CReduceDToStatement:
		return "DToStatement"
	case _CReduceEToStatement:
		return "EToStatement"
	case _CReduceFToStatement:
		return "FToStatement"
	case _CReduceAToLabeledStatement:
		return "AToLabeledStatement"
	case _CReduceBToLabeledStatement:
		return "BToLabeledStatement"
	case _CReduceCToLabeledStatement:
		return "CToLabeledStatement"
	case _CReduceAToCompoundStatement:
		return "AToCompoundStatement"
	case _CReduceBToCompoundStatement:
		return "BToCompoundStatement"
	case _CReduceCToCompoundStatement:
		return "CToCompoundStatement"
	case _CReduceDToCompoundStatement:
		return "DToCompoundStatement"
	case _CReduceAToDeclarationList:
		return "AToDeclarationList"
	case _CReduceBToDeclarationList:
		return "BToDeclarationList"
	case _CReduceAToStatementList:
		return "AToStatementList"
	case _CReduceBToStatementList:
		return "BToStatementList"
	case _CReduceAToExpressionStatement:
		return "AToExpressionStatement"
	case _CReduceBToExpressionStatement:
		return "BToExpressionStatement"
	case _CReduceAToSelectionStatement:
		return "AToSelectionStatement"
	case _CReduceBToSelectionStatement:
		return "BToSelectionStatement"
	case _CReduceCToSelectionStatement:
		return "CToSelectionStatement"
	case _CReduceAToIterationStatement:
		return "AToIterationStatement"
	case _CReduceBToIterationStatement:
		return "BToIterationStatement"
	case _CReduceCToIterationStatement:
		return "CToIterationStatement"
	case _CReduceDToIterationStatement:
		return "DToIterationStatement"
	case _CReduceAToJumpStatement:
		return "AToJumpStatement"
	case _CReduceBToJumpStatement:
		return "BToJumpStatement"
	case _CReduceCToJumpStatement:
		return "CToJumpStatement"
	case _CReduceDToJumpStatement:
		return "DToJumpStatement"
	case _CReduceEToJumpStatement:
		return "EToJumpStatement"
	case _CReduceAToTranslationUnit:
		return "AToTranslationUnit"
	case _CReduceBToTranslationUnit:
		return "BToTranslationUnit"
	case _CReduceAToExternalDeclaration:
		return "AToExternalDeclaration"
	case _CReduceBToExternalDeclaration:
		return "BToExternalDeclaration"
	case _CReduceAToFunctionDefinition:
		return "AToFunctionDefinition"
	case _CReduceBToFunctionDefinition:
		return "BToFunctionDefinition"
	case _CReduceCToFunctionDefinition:
		return "CToFunctionDefinition"
	case _CReduceDToFunctionDefinition:
		return "DToFunctionDefinition"
	default:
		return fmt.Sprintf("?unknown reduce type %d?", int(i))
	}
}

type _CStateId int

func (id _CStateId) String() string {
	return fmt.Sprintf("State %d", int(id))
}

const (
	_CState1   = _CStateId(1)
	_CState2   = _CStateId(2)
	_CState3   = _CStateId(3)
	_CState4   = _CStateId(4)
	_CState5   = _CStateId(5)
	_CState6   = _CStateId(6)
	_CState7   = _CStateId(7)
	_CState8   = _CStateId(8)
	_CState9   = _CStateId(9)
	_CState10  = _CStateId(10)
	_CState11  = _CStateId(11)
	_CState12  = _CStateId(12)
	_CState13  = _CStateId(13)
	_CState14  = _CStateId(14)
	_CState15  = _CStateId(15)
	_CState16  = _CStateId(16)
	_CState17  = _CStateId(17)
	_CState18  = _CStateId(18)
	_CState19  = _CStateId(19)
	_CState20  = _CStateId(20)
	_CState21  = _CStateId(21)
	_CState22  = _CStateId(22)
	_CState23  = _CStateId(23)
	_CState24  = _CStateId(24)
	_CState25  = _CStateId(25)
	_CState26  = _CStateId(26)
	_CState27  = _CStateId(27)
	_CState28  = _CStateId(28)
	_CState29  = _CStateId(29)
	_CState30  = _CStateId(30)
	_CState31  = _CStateId(31)
	_CState32  = _CStateId(32)
	_CState33  = _CStateId(33)
	_CState34  = _CStateId(34)
	_CState35  = _CStateId(35)
	_CState36  = _CStateId(36)
	_CState37  = _CStateId(37)
	_CState38  = _CStateId(38)
	_CState39  = _CStateId(39)
	_CState40  = _CStateId(40)
	_CState41  = _CStateId(41)
	_CState42  = _CStateId(42)
	_CState43  = _CStateId(43)
	_CState44  = _CStateId(44)
	_CState45  = _CStateId(45)
	_CState46  = _CStateId(46)
	_CState47  = _CStateId(47)
	_CState48  = _CStateId(48)
	_CState49  = _CStateId(49)
	_CState50  = _CStateId(50)
	_CState51  = _CStateId(51)
	_CState52  = _CStateId(52)
	_CState53  = _CStateId(53)
	_CState54  = _CStateId(54)
	_CState55  = _CStateId(55)
	_CState56  = _CStateId(56)
	_CState57  = _CStateId(57)
	_CState58  = _CStateId(58)
	_CState59  = _CStateId(59)
	_CState60  = _CStateId(60)
	_CState61  = _CStateId(61)
	_CState62  = _CStateId(62)
	_CState63  = _CStateId(63)
	_CState64  = _CStateId(64)
	_CState65  = _CStateId(65)
	_CState66  = _CStateId(66)
	_CState67  = _CStateId(67)
	_CState68  = _CStateId(68)
	_CState69  = _CStateId(69)
	_CState70  = _CStateId(70)
	_CState71  = _CStateId(71)
	_CState72  = _CStateId(72)
	_CState73  = _CStateId(73)
	_CState74  = _CStateId(74)
	_CState75  = _CStateId(75)
	_CState76  = _CStateId(76)
	_CState77  = _CStateId(77)
	_CState78  = _CStateId(78)
	_CState79  = _CStateId(79)
	_CState80  = _CStateId(80)
	_CState81  = _CStateId(81)
	_CState82  = _CStateId(82)
	_CState83  = _CStateId(83)
	_CState84  = _CStateId(84)
	_CState85  = _CStateId(85)
	_CState86  = _CStateId(86)
	_CState87  = _CStateId(87)
	_CState88  = _CStateId(88)
	_CState89  = _CStateId(89)
	_CState90  = _CStateId(90)
	_CState91  = _CStateId(91)
	_CState92  = _CStateId(92)
	_CState93  = _CStateId(93)
	_CState94  = _CStateId(94)
	_CState95  = _CStateId(95)
	_CState96  = _CStateId(96)
	_CState97  = _CStateId(97)
	_CState98  = _CStateId(98)
	_CState99  = _CStateId(99)
	_CState100 = _CStateId(100)
	_CState101 = _CStateId(101)
	_CState102 = _CStateId(102)
	_CState103 = _CStateId(103)
	_CState104 = _CStateId(104)
	_CState105 = _CStateId(105)
	_CState106 = _CStateId(106)
	_CState107 = _CStateId(107)
	_CState108 = _CStateId(108)
	_CState109 = _CStateId(109)
	_CState110 = _CStateId(110)
	_CState111 = _CStateId(111)
	_CState112 = _CStateId(112)
	_CState113 = _CStateId(113)
	_CState114 = _CStateId(114)
	_CState115 = _CStateId(115)
	_CState116 = _CStateId(116)
	_CState117 = _CStateId(117)
	_CState118 = _CStateId(118)
	_CState119 = _CStateId(119)
	_CState120 = _CStateId(120)
	_CState121 = _CStateId(121)
	_CState122 = _CStateId(122)
	_CState123 = _CStateId(123)
	_CState124 = _CStateId(124)
	_CState125 = _CStateId(125)
	_CState126 = _CStateId(126)
	_CState127 = _CStateId(127)
	_CState128 = _CStateId(128)
	_CState129 = _CStateId(129)
	_CState130 = _CStateId(130)
	_CState131 = _CStateId(131)
	_CState132 = _CStateId(132)
	_CState133 = _CStateId(133)
	_CState134 = _CStateId(134)
	_CState135 = _CStateId(135)
	_CState136 = _CStateId(136)
	_CState137 = _CStateId(137)
	_CState138 = _CStateId(138)
	_CState139 = _CStateId(139)
	_CState140 = _CStateId(140)
	_CState141 = _CStateId(141)
	_CState142 = _CStateId(142)
	_CState143 = _CStateId(143)
	_CState144 = _CStateId(144)
	_CState145 = _CStateId(145)
	_CState146 = _CStateId(146)
	_CState147 = _CStateId(147)
	_CState148 = _CStateId(148)
	_CState149 = _CStateId(149)
	_CState150 = _CStateId(150)
	_CState151 = _CStateId(151)
	_CState152 = _CStateId(152)
	_CState153 = _CStateId(153)
	_CState154 = _CStateId(154)
	_CState155 = _CStateId(155)
	_CState156 = _CStateId(156)
	_CState157 = _CStateId(157)
	_CState158 = _CStateId(158)
	_CState159 = _CStateId(159)
	_CState160 = _CStateId(160)
	_CState161 = _CStateId(161)
	_CState162 = _CStateId(162)
	_CState163 = _CStateId(163)
	_CState164 = _CStateId(164)
	_CState165 = _CStateId(165)
	_CState166 = _CStateId(166)
	_CState167 = _CStateId(167)
	_CState168 = _CStateId(168)
	_CState169 = _CStateId(169)
	_CState170 = _CStateId(170)
	_CState171 = _CStateId(171)
	_CState172 = _CStateId(172)
	_CState173 = _CStateId(173)
	_CState174 = _CStateId(174)
	_CState175 = _CStateId(175)
	_CState176 = _CStateId(176)
	_CState177 = _CStateId(177)
	_CState178 = _CStateId(178)
	_CState179 = _CStateId(179)
	_CState180 = _CStateId(180)
	_CState181 = _CStateId(181)
	_CState182 = _CStateId(182)
	_CState183 = _CStateId(183)
	_CState184 = _CStateId(184)
	_CState185 = _CStateId(185)
	_CState186 = _CStateId(186)
	_CState187 = _CStateId(187)
	_CState188 = _CStateId(188)
	_CState189 = _CStateId(189)
	_CState190 = _CStateId(190)
	_CState191 = _CStateId(191)
	_CState192 = _CStateId(192)
	_CState193 = _CStateId(193)
	_CState194 = _CStateId(194)
	_CState195 = _CStateId(195)
	_CState196 = _CStateId(196)
	_CState197 = _CStateId(197)
	_CState198 = _CStateId(198)
	_CState199 = _CStateId(199)
	_CState200 = _CStateId(200)
	_CState201 = _CStateId(201)
	_CState202 = _CStateId(202)
	_CState203 = _CStateId(203)
	_CState204 = _CStateId(204)
	_CState205 = _CStateId(205)
	_CState206 = _CStateId(206)
	_CState207 = _CStateId(207)
	_CState208 = _CStateId(208)
	_CState209 = _CStateId(209)
	_CState210 = _CStateId(210)
	_CState211 = _CStateId(211)
	_CState212 = _CStateId(212)
	_CState213 = _CStateId(213)
	_CState214 = _CStateId(214)
	_CState215 = _CStateId(215)
	_CState216 = _CStateId(216)
	_CState217 = _CStateId(217)
	_CState218 = _CStateId(218)
	_CState219 = _CStateId(219)
	_CState220 = _CStateId(220)
	_CState221 = _CStateId(221)
	_CState222 = _CStateId(222)
	_CState223 = _CStateId(223)
	_CState224 = _CStateId(224)
	_CState225 = _CStateId(225)
	_CState226 = _CStateId(226)
	_CState227 = _CStateId(227)
	_CState228 = _CStateId(228)
	_CState229 = _CStateId(229)
	_CState230 = _CStateId(230)
	_CState231 = _CStateId(231)
	_CState232 = _CStateId(232)
	_CState233 = _CStateId(233)
	_CState234 = _CStateId(234)
	_CState235 = _CStateId(235)
	_CState236 = _CStateId(236)
	_CState237 = _CStateId(237)
	_CState238 = _CStateId(238)
	_CState239 = _CStateId(239)
	_CState240 = _CStateId(240)
	_CState241 = _CStateId(241)
	_CState242 = _CStateId(242)
	_CState243 = _CStateId(243)
	_CState244 = _CStateId(244)
	_CState245 = _CStateId(245)
	_CState246 = _CStateId(246)
	_CState247 = _CStateId(247)
	_CState248 = _CStateId(248)
	_CState249 = _CStateId(249)
	_CState250 = _CStateId(250)
	_CState251 = _CStateId(251)
	_CState252 = _CStateId(252)
	_CState253 = _CStateId(253)
	_CState254 = _CStateId(254)
	_CState255 = _CStateId(255)
	_CState256 = _CStateId(256)
	_CState257 = _CStateId(257)
	_CState258 = _CStateId(258)
	_CState259 = _CStateId(259)
	_CState260 = _CStateId(260)
	_CState261 = _CStateId(261)
	_CState262 = _CStateId(262)
	_CState263 = _CStateId(263)
	_CState264 = _CStateId(264)
	_CState265 = _CStateId(265)
	_CState266 = _CStateId(266)
	_CState267 = _CStateId(267)
	_CState268 = _CStateId(268)
	_CState269 = _CStateId(269)
	_CState270 = _CStateId(270)
	_CState271 = _CStateId(271)
	_CState272 = _CStateId(272)
	_CState273 = _CStateId(273)
	_CState274 = _CStateId(274)
	_CState275 = _CStateId(275)
	_CState276 = _CStateId(276)
	_CState277 = _CStateId(277)
	_CState278 = _CStateId(278)
	_CState279 = _CStateId(279)
	_CState280 = _CStateId(280)
	_CState281 = _CStateId(281)
	_CState282 = _CStateId(282)
	_CState283 = _CStateId(283)
	_CState284 = _CStateId(284)
	_CState285 = _CStateId(285)
	_CState286 = _CStateId(286)
	_CState287 = _CStateId(287)
	_CState288 = _CStateId(288)
	_CState289 = _CStateId(289)
	_CState290 = _CStateId(290)
	_CState291 = _CStateId(291)
	_CState292 = _CStateId(292)
	_CState293 = _CStateId(293)
	_CState294 = _CStateId(294)
	_CState295 = _CStateId(295)
	_CState296 = _CStateId(296)
	_CState297 = _CStateId(297)
	_CState298 = _CStateId(298)
	_CState299 = _CStateId(299)
	_CState300 = _CStateId(300)
	_CState301 = _CStateId(301)
	_CState302 = _CStateId(302)
	_CState303 = _CStateId(303)
	_CState304 = _CStateId(304)
	_CState305 = _CStateId(305)
	_CState306 = _CStateId(306)
	_CState307 = _CStateId(307)
	_CState308 = _CStateId(308)
	_CState309 = _CStateId(309)
	_CState310 = _CStateId(310)
	_CState311 = _CStateId(311)
	_CState312 = _CStateId(312)
	_CState313 = _CStateId(313)
	_CState314 = _CStateId(314)
	_CState315 = _CStateId(315)
	_CState316 = _CStateId(316)
	_CState317 = _CStateId(317)
	_CState318 = _CStateId(318)
	_CState319 = _CStateId(319)
	_CState320 = _CStateId(320)
	_CState321 = _CStateId(321)
	_CState322 = _CStateId(322)
	_CState323 = _CStateId(323)
	_CState324 = _CStateId(324)
	_CState325 = _CStateId(325)
	_CState326 = _CStateId(326)
	_CState327 = _CStateId(327)
	_CState328 = _CStateId(328)
	_CState329 = _CStateId(329)
	_CState330 = _CStateId(330)
	_CState331 = _CStateId(331)
	_CState332 = _CStateId(332)
	_CState333 = _CStateId(333)
	_CState334 = _CStateId(334)
	_CState335 = _CStateId(335)
	_CState336 = _CStateId(336)
	_CState337 = _CStateId(337)
	_CState338 = _CStateId(338)
	_CState339 = _CStateId(339)
	_CState340 = _CStateId(340)
	_CState341 = _CStateId(341)
	_CState342 = _CStateId(342)
	_CState343 = _CStateId(343)
	_CState344 = _CStateId(344)
	_CState345 = _CStateId(345)
	_CState346 = _CStateId(346)
	_CState347 = _CStateId(347)
	_CState348 = _CStateId(348)
	_CState349 = _CStateId(349)
	_CState350 = _CStateId(350)
	_CState351 = _CStateId(351)
)

type CSymbol struct {
	SymbolId_ CSymbolId

	Generic_ *CGenericSymbol
}

func NewSymbol(token CToken) (*CSymbol, error) {
	symbol, ok := token.(*CSymbol)
	if ok {
		return symbol, nil
	}

	symbol = &CSymbol{SymbolId_: token.Id()}
	switch token.Id() {
	case _CEndMarker, CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CPtrOpToken, CIncOpToken, CDecOpToken, CLeftOpToken, CRightOpToken, CLeOpToken, CGeOpToken, CEqOpToken, CNeOpToken, CAndOpToken, COrOpToken, CMulAssignToken, CDivAssignToken, CModAssignToken, CAddAssignToken, CSubAssignToken, CLeftAssignToken, CRightAssignToken, CAndAssignToken, CXorAssignToken, COrAssignToken, CTypeNameToken, CTypedefToken, CExternToken, CStaticToken, CAutoToken, CRegisterToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken, CEllipsisToken, CCaseToken, CDefaultToken, CIfToken, CElseToken, CSwitchToken, CWhileToken, CDoToken, CForToken, CGotoToken, CContinueToken, CBreakToken, CReturnToken, '(', ')', '{', '}', '[', ']', ';', ':', ',', '=', '?', '*', '/', '-', '+', '%', '&', '|', '!', '.', '^', '<', '>', '~':
		val, ok := token.(*CGenericSymbol)
		if !ok {
			return nil, fmt.Errorf(
				"Invalid value type for token %s.  "+
					"Expecting *CGenericSymbol (%v)",
				token.Id(),
				token.Loc())
		}
		symbol.Generic_ = val
	default:
		return nil, fmt.Errorf("Unexpected token type: %s", symbol.Id())
	}
	return symbol, nil
}

func (s *CSymbol) Id() CSymbolId {
	return s.SymbolId_
}

func (s *CSymbol) Loc() CLocation {
	type locator interface{ Loc() CLocation }
	switch s.SymbolId_ {
	}
	if s.Generic_ != nil {
		return s.Generic_.Loc()
	}
	return CLocation{}
}

type _CPseudoSymbolStack struct {
	lexer CLexer
	top   []*CSymbol
}

func (stack *_CPseudoSymbolStack) Top() (*CSymbol, error) {
	if len(stack.top) == 0 {
		token, err := stack.lexer.Next()
		if err != nil {
			if err != io.EOF {
				return nil, fmt.Errorf("Unexpected lex error: %s", err)
			}
			token = &CGenericSymbol{_CEndMarker, stack.lexer.CurrentLocation()}
		}
		item, err := NewSymbol(token)
		if err != nil {
			return nil, err
		}
		stack.top = append(stack.top, item)
	}
	return stack.top[len(stack.top)-1], nil
}

func (stack *_CPseudoSymbolStack) Push(symbol *CSymbol) {
	stack.top = append(stack.top, symbol)
}

func (stack *_CPseudoSymbolStack) Pop() (*CSymbol, error) {
	if len(stack.top) == 0 {
		return nil, fmt.Errorf("internal error: cannot pop an empty top")
	}
	ret := stack.top[len(stack.top)-1]
	stack.top = stack.top[:len(stack.top)-1]
	return ret, nil
}

type _CStackItem struct {
	StateId _CStateId

	*CSymbol
}

type _CStack []*_CStackItem

type _CAction struct {
	ActionType _CActionType

	ShiftStateId _CStateId
	ReduceType   _CReduceType
}

func (act *_CAction) ShiftItem(symbol *CSymbol) *_CStackItem {
	return &_CStackItem{StateId: act.ShiftStateId, CSymbol: symbol}
}

func (act *_CAction) ReduceSymbol(
	reducer CReducer,
	stack _CStack) (
	_CStack,
	*CSymbol,
	error) {

	var err error
	symbol := &CSymbol{}
	switch act.ReduceType {
	case _CReduceAToPrimaryExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CPrimaryExpressionType
		symbol.Generic_, err = reducer.AToPrimaryExpression(args[0].Generic_)
	case _CReduceBToPrimaryExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CPrimaryExpressionType
		symbol.Generic_, err = reducer.BToPrimaryExpression(args[0].Generic_)
	case _CReduceCToPrimaryExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CPrimaryExpressionType
		symbol.Generic_, err = reducer.CToPrimaryExpression(args[0].Generic_)
	case _CReduceDToPrimaryExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CPrimaryExpressionType
		symbol.Generic_, err = reducer.DToPrimaryExpression(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceAToPostfixExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CPostfixExpressionType
		symbol.Generic_, err = reducer.AToPostfixExpression(args[0].Generic_)
	case _CReduceBToPostfixExpression:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CPostfixExpressionType
		symbol.Generic_, err = reducer.BToPostfixExpression(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _CReduceCToPostfixExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CPostfixExpressionType
		symbol.Generic_, err = reducer.CToPostfixExpression(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceDToPostfixExpression:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CPostfixExpressionType
		symbol.Generic_, err = reducer.DToPostfixExpression(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _CReduceEToPostfixExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CPostfixExpressionType
		symbol.Generic_, err = reducer.EToPostfixExpression(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceFToPostfixExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CPostfixExpressionType
		symbol.Generic_, err = reducer.FToPostfixExpression(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceGToPostfixExpression:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CPostfixExpressionType
		symbol.Generic_, err = reducer.GToPostfixExpression(args[0].Generic_, args[1].Generic_)
	case _CReduceHToPostfixExpression:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CPostfixExpressionType
		symbol.Generic_, err = reducer.HToPostfixExpression(args[0].Generic_, args[1].Generic_)
	case _CReduceAToArgumentExpressionList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CArgumentExpressionListType
		symbol.Generic_, err = reducer.AToArgumentExpressionList(args[0].Generic_)
	case _CReduceBToArgumentExpressionList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CArgumentExpressionListType
		symbol.Generic_, err = reducer.BToArgumentExpressionList(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceAToUnaryExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CUnaryExpressionType
		symbol.Generic_, err = reducer.AToUnaryExpression(args[0].Generic_)
	case _CReduceBToUnaryExpression:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CUnaryExpressionType
		symbol.Generic_, err = reducer.BToUnaryExpression(args[0].Generic_, args[1].Generic_)
	case _CReduceCToUnaryExpression:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CUnaryExpressionType
		symbol.Generic_, err = reducer.CToUnaryExpression(args[0].Generic_, args[1].Generic_)
	case _CReduceDToUnaryExpression:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CUnaryExpressionType
		symbol.Generic_, err = reducer.DToUnaryExpression(args[0].Generic_, args[1].Generic_)
	case _CReduceEToUnaryExpression:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CUnaryExpressionType
		symbol.Generic_, err = reducer.EToUnaryExpression(args[0].Generic_, args[1].Generic_)
	case _CReduceFToUnaryExpression:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CUnaryExpressionType
		symbol.Generic_, err = reducer.FToUnaryExpression(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _CReduceAToUnaryOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CUnaryOperatorType
		symbol.Generic_, err = reducer.AToUnaryOperator(args[0].Generic_)
	case _CReduceBToUnaryOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CUnaryOperatorType
		symbol.Generic_, err = reducer.BToUnaryOperator(args[0].Generic_)
	case _CReduceCToUnaryOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CUnaryOperatorType
		symbol.Generic_, err = reducer.CToUnaryOperator(args[0].Generic_)
	case _CReduceDToUnaryOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CUnaryOperatorType
		symbol.Generic_, err = reducer.DToUnaryOperator(args[0].Generic_)
	case _CReduceEToUnaryOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CUnaryOperatorType
		symbol.Generic_, err = reducer.EToUnaryOperator(args[0].Generic_)
	case _CReduceFToUnaryOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CUnaryOperatorType
		symbol.Generic_, err = reducer.FToUnaryOperator(args[0].Generic_)
	case _CReduceAToCastExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CCastExpressionType
		symbol.Generic_, err = reducer.AToCastExpression(args[0].Generic_)
	case _CReduceBToCastExpression:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CCastExpressionType
		symbol.Generic_, err = reducer.BToCastExpression(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _CReduceAToMultiplicativeExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CMultiplicativeExpressionType
		symbol.Generic_, err = reducer.AToMultiplicativeExpression(args[0].Generic_)
	case _CReduceBToMultiplicativeExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CMultiplicativeExpressionType
		symbol.Generic_, err = reducer.BToMultiplicativeExpression(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceCToMultiplicativeExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CMultiplicativeExpressionType
		symbol.Generic_, err = reducer.CToMultiplicativeExpression(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceDToMultiplicativeExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CMultiplicativeExpressionType
		symbol.Generic_, err = reducer.DToMultiplicativeExpression(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceAToAdditiveExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CAdditiveExpressionType
		symbol.Generic_, err = reducer.AToAdditiveExpression(args[0].Generic_)
	case _CReduceBToAdditiveExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CAdditiveExpressionType
		symbol.Generic_, err = reducer.BToAdditiveExpression(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceCToAdditiveExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CAdditiveExpressionType
		symbol.Generic_, err = reducer.CToAdditiveExpression(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceAToShiftExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CShiftExpressionType
		symbol.Generic_, err = reducer.AToShiftExpression(args[0].Generic_)
	case _CReduceBToShiftExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CShiftExpressionType
		symbol.Generic_, err = reducer.BToShiftExpression(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceCToShiftExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CShiftExpressionType
		symbol.Generic_, err = reducer.CToShiftExpression(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceAToRelationalExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CRelationalExpressionType
		symbol.Generic_, err = reducer.AToRelationalExpression(args[0].Generic_)
	case _CReduceBToRelationalExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CRelationalExpressionType
		symbol.Generic_, err = reducer.BToRelationalExpression(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceCToRelationalExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CRelationalExpressionType
		symbol.Generic_, err = reducer.CToRelationalExpression(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceDToRelationalExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CRelationalExpressionType
		symbol.Generic_, err = reducer.DToRelationalExpression(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceEToRelationalExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CRelationalExpressionType
		symbol.Generic_, err = reducer.EToRelationalExpression(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceAToEqualityExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CEqualityExpressionType
		symbol.Generic_, err = reducer.AToEqualityExpression(args[0].Generic_)
	case _CReduceBToEqualityExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CEqualityExpressionType
		symbol.Generic_, err = reducer.BToEqualityExpression(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceCToEqualityExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CEqualityExpressionType
		symbol.Generic_, err = reducer.CToEqualityExpression(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceAToAndExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CAndExpressionType
		symbol.Generic_, err = reducer.AToAndExpression(args[0].Generic_)
	case _CReduceBToAndExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CAndExpressionType
		symbol.Generic_, err = reducer.BToAndExpression(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceAToExclusiveOrExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CExclusiveOrExpressionType
		symbol.Generic_, err = reducer.AToExclusiveOrExpression(args[0].Generic_)
	case _CReduceBToExclusiveOrExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CExclusiveOrExpressionType
		symbol.Generic_, err = reducer.BToExclusiveOrExpression(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceAToInclusiveOrExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CInclusiveOrExpressionType
		symbol.Generic_, err = reducer.AToInclusiveOrExpression(args[0].Generic_)
	case _CReduceBToInclusiveOrExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CInclusiveOrExpressionType
		symbol.Generic_, err = reducer.BToInclusiveOrExpression(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceAToLogicalAndExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CLogicalAndExpressionType
		symbol.Generic_, err = reducer.AToLogicalAndExpression(args[0].Generic_)
	case _CReduceBToLogicalAndExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CLogicalAndExpressionType
		symbol.Generic_, err = reducer.BToLogicalAndExpression(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceAToLogicalOrExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CLogicalOrExpressionType
		symbol.Generic_, err = reducer.AToLogicalOrExpression(args[0].Generic_)
	case _CReduceBToLogicalOrExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CLogicalOrExpressionType
		symbol.Generic_, err = reducer.BToLogicalOrExpression(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceAToConditionalExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CConditionalExpressionType
		symbol.Generic_, err = reducer.AToConditionalExpression(args[0].Generic_)
	case _CReduceBToConditionalExpression:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = CConditionalExpressionType
		symbol.Generic_, err = reducer.BToConditionalExpression(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_)
	case _CReduceAToAssignmentExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CAssignmentExpressionType
		symbol.Generic_, err = reducer.AToAssignmentExpression(args[0].Generic_)
	case _CReduceBToAssignmentExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CAssignmentExpressionType
		symbol.Generic_, err = reducer.BToAssignmentExpression(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceAToAssignmentOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CAssignmentOperatorType
		symbol.Generic_, err = reducer.AToAssignmentOperator(args[0].Generic_)
	case _CReduceBToAssignmentOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CAssignmentOperatorType
		symbol.Generic_, err = reducer.BToAssignmentOperator(args[0].Generic_)
	case _CReduceCToAssignmentOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CAssignmentOperatorType
		symbol.Generic_, err = reducer.CToAssignmentOperator(args[0].Generic_)
	case _CReduceDToAssignmentOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CAssignmentOperatorType
		symbol.Generic_, err = reducer.DToAssignmentOperator(args[0].Generic_)
	case _CReduceEToAssignmentOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CAssignmentOperatorType
		symbol.Generic_, err = reducer.EToAssignmentOperator(args[0].Generic_)
	case _CReduceFToAssignmentOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CAssignmentOperatorType
		symbol.Generic_, err = reducer.FToAssignmentOperator(args[0].Generic_)
	case _CReduceGToAssignmentOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CAssignmentOperatorType
		symbol.Generic_, err = reducer.GToAssignmentOperator(args[0].Generic_)
	case _CReduceHToAssignmentOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CAssignmentOperatorType
		symbol.Generic_, err = reducer.HToAssignmentOperator(args[0].Generic_)
	case _CReduceIToAssignmentOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CAssignmentOperatorType
		symbol.Generic_, err = reducer.IToAssignmentOperator(args[0].Generic_)
	case _CReduceJToAssignmentOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CAssignmentOperatorType
		symbol.Generic_, err = reducer.JToAssignmentOperator(args[0].Generic_)
	case _CReduceKToAssignmentOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CAssignmentOperatorType
		symbol.Generic_, err = reducer.KToAssignmentOperator(args[0].Generic_)
	case _CReduceAToExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CExpressionType
		symbol.Generic_, err = reducer.AToExpression(args[0].Generic_)
	case _CReduceBToExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CExpressionType
		symbol.Generic_, err = reducer.BToExpression(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceAToConstantExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CConstantExpressionType
		symbol.Generic_, err = reducer.AToConstantExpression(args[0].Generic_)
	case _CReduceAToDeclaration:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CDeclarationType
		symbol.Generic_, err = reducer.AToDeclaration(args[0].Generic_, args[1].Generic_)
	case _CReduceBToDeclaration:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CDeclarationType
		symbol.Generic_, err = reducer.BToDeclaration(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceAToDeclarationSpecifiers:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CDeclarationSpecifiersType
		symbol.Generic_, err = reducer.AToDeclarationSpecifiers(args[0].Generic_)
	case _CReduceBToDeclarationSpecifiers:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CDeclarationSpecifiersType
		symbol.Generic_, err = reducer.BToDeclarationSpecifiers(args[0].Generic_, args[1].Generic_)
	case _CReduceCToDeclarationSpecifiers:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CDeclarationSpecifiersType
		symbol.Generic_, err = reducer.CToDeclarationSpecifiers(args[0].Generic_)
	case _CReduceDToDeclarationSpecifiers:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CDeclarationSpecifiersType
		symbol.Generic_, err = reducer.DToDeclarationSpecifiers(args[0].Generic_, args[1].Generic_)
	case _CReduceEToDeclarationSpecifiers:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CDeclarationSpecifiersType
		symbol.Generic_, err = reducer.EToDeclarationSpecifiers(args[0].Generic_)
	case _CReduceFToDeclarationSpecifiers:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CDeclarationSpecifiersType
		symbol.Generic_, err = reducer.FToDeclarationSpecifiers(args[0].Generic_, args[1].Generic_)
	case _CReduceAToInitDeclaratorList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CInitDeclaratorListType
		symbol.Generic_, err = reducer.AToInitDeclaratorList(args[0].Generic_)
	case _CReduceBToInitDeclaratorList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CInitDeclaratorListType
		symbol.Generic_, err = reducer.BToInitDeclaratorList(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceAToInitDeclarator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CInitDeclaratorType
		symbol.Generic_, err = reducer.AToInitDeclarator(args[0].Generic_)
	case _CReduceBToInitDeclarator:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CInitDeclaratorType
		symbol.Generic_, err = reducer.BToInitDeclarator(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceAToStorageClassSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CStorageClassSpecifierType
		symbol.Generic_, err = reducer.AToStorageClassSpecifier(args[0].Generic_)
	case _CReduceBToStorageClassSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CStorageClassSpecifierType
		symbol.Generic_, err = reducer.BToStorageClassSpecifier(args[0].Generic_)
	case _CReduceCToStorageClassSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CStorageClassSpecifierType
		symbol.Generic_, err = reducer.CToStorageClassSpecifier(args[0].Generic_)
	case _CReduceDToStorageClassSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CStorageClassSpecifierType
		symbol.Generic_, err = reducer.DToStorageClassSpecifier(args[0].Generic_)
	case _CReduceEToStorageClassSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CStorageClassSpecifierType
		symbol.Generic_, err = reducer.EToStorageClassSpecifier(args[0].Generic_)
	case _CReduceAToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CTypeSpecifierType
		symbol.Generic_, err = reducer.AToTypeSpecifier(args[0].Generic_)
	case _CReduceBToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CTypeSpecifierType
		symbol.Generic_, err = reducer.BToTypeSpecifier(args[0].Generic_)
	case _CReduceCToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CTypeSpecifierType
		symbol.Generic_, err = reducer.CToTypeSpecifier(args[0].Generic_)
	case _CReduceDToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CTypeSpecifierType
		symbol.Generic_, err = reducer.DToTypeSpecifier(args[0].Generic_)
	case _CReduceEToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CTypeSpecifierType
		symbol.Generic_, err = reducer.EToTypeSpecifier(args[0].Generic_)
	case _CReduceFToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CTypeSpecifierType
		symbol.Generic_, err = reducer.FToTypeSpecifier(args[0].Generic_)
	case _CReduceGToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CTypeSpecifierType
		symbol.Generic_, err = reducer.GToTypeSpecifier(args[0].Generic_)
	case _CReduceHToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CTypeSpecifierType
		symbol.Generic_, err = reducer.HToTypeSpecifier(args[0].Generic_)
	case _CReduceIToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CTypeSpecifierType
		symbol.Generic_, err = reducer.IToTypeSpecifier(args[0].Generic_)
	case _CReduceJToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CTypeSpecifierType
		symbol.Generic_, err = reducer.JToTypeSpecifier(args[0].Generic_)
	case _CReduceKToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CTypeSpecifierType
		symbol.Generic_, err = reducer.KToTypeSpecifier(args[0].Generic_)
	case _CReduceLToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CTypeSpecifierType
		symbol.Generic_, err = reducer.LToTypeSpecifier(args[0].Generic_)
	case _CReduceAToStructOrUnionSpecifier:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = CStructOrUnionSpecifierType
		symbol.Generic_, err = reducer.AToStructOrUnionSpecifier(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_)
	case _CReduceBToStructOrUnionSpecifier:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CStructOrUnionSpecifierType
		symbol.Generic_, err = reducer.BToStructOrUnionSpecifier(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _CReduceCToStructOrUnionSpecifier:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CStructOrUnionSpecifierType
		symbol.Generic_, err = reducer.CToStructOrUnionSpecifier(args[0].Generic_, args[1].Generic_)
	case _CReduceAToStructOrUnion:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CStructOrUnionType
		symbol.Generic_, err = reducer.AToStructOrUnion(args[0].Generic_)
	case _CReduceBToStructOrUnion:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CStructOrUnionType
		symbol.Generic_, err = reducer.BToStructOrUnion(args[0].Generic_)
	case _CReduceAToStructDeclarationList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CStructDeclarationListType
		symbol.Generic_, err = reducer.AToStructDeclarationList(args[0].Generic_)
	case _CReduceBToStructDeclarationList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CStructDeclarationListType
		symbol.Generic_, err = reducer.BToStructDeclarationList(args[0].Generic_, args[1].Generic_)
	case _CReduceAToStructDeclaration:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CStructDeclarationType
		symbol.Generic_, err = reducer.AToStructDeclaration(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceAToSpecifierQualifierList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CSpecifierQualifierListType
		symbol.Generic_, err = reducer.AToSpecifierQualifierList(args[0].Generic_, args[1].Generic_)
	case _CReduceBToSpecifierQualifierList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CSpecifierQualifierListType
		symbol.Generic_, err = reducer.BToSpecifierQualifierList(args[0].Generic_)
	case _CReduceCToSpecifierQualifierList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CSpecifierQualifierListType
		symbol.Generic_, err = reducer.CToSpecifierQualifierList(args[0].Generic_, args[1].Generic_)
	case _CReduceDToSpecifierQualifierList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CSpecifierQualifierListType
		symbol.Generic_, err = reducer.DToSpecifierQualifierList(args[0].Generic_)
	case _CReduceAToStructDeclaratorList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CStructDeclaratorListType
		symbol.Generic_, err = reducer.AToStructDeclaratorList(args[0].Generic_)
	case _CReduceBToStructDeclaratorList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CStructDeclaratorListType
		symbol.Generic_, err = reducer.BToStructDeclaratorList(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceAToStructDeclarator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CStructDeclaratorType
		symbol.Generic_, err = reducer.AToStructDeclarator(args[0].Generic_)
	case _CReduceBToStructDeclarator:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CStructDeclaratorType
		symbol.Generic_, err = reducer.BToStructDeclarator(args[0].Generic_, args[1].Generic_)
	case _CReduceCToStructDeclarator:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CStructDeclaratorType
		symbol.Generic_, err = reducer.CToStructDeclarator(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceAToEnumSpecifier:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CEnumSpecifierType
		symbol.Generic_, err = reducer.AToEnumSpecifier(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _CReduceBToEnumSpecifier:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = CEnumSpecifierType
		symbol.Generic_, err = reducer.BToEnumSpecifier(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_)
	case _CReduceCToEnumSpecifier:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CEnumSpecifierType
		symbol.Generic_, err = reducer.CToEnumSpecifier(args[0].Generic_, args[1].Generic_)
	case _CReduceAToEnumeratorList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CEnumeratorListType
		symbol.Generic_, err = reducer.AToEnumeratorList(args[0].Generic_)
	case _CReduceBToEnumeratorList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CEnumeratorListType
		symbol.Generic_, err = reducer.BToEnumeratorList(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceAToEnumerator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CEnumeratorType
		symbol.Generic_, err = reducer.AToEnumerator(args[0].Generic_)
	case _CReduceBToEnumerator:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CEnumeratorType
		symbol.Generic_, err = reducer.BToEnumerator(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceAToTypeQualifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CTypeQualifierType
		symbol.Generic_, err = reducer.AToTypeQualifier(args[0].Generic_)
	case _CReduceBToTypeQualifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CTypeQualifierType
		symbol.Generic_, err = reducer.BToTypeQualifier(args[0].Generic_)
	case _CReduceAToDeclarator:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CDeclaratorType
		symbol.Generic_, err = reducer.AToDeclarator(args[0].Generic_, args[1].Generic_)
	case _CReduceBToDeclarator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CDeclaratorType
		symbol.Generic_, err = reducer.BToDeclarator(args[0].Generic_)
	case _CReduceAToDirectDeclarator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CDirectDeclaratorType
		symbol.Generic_, err = reducer.AToDirectDeclarator(args[0].Generic_)
	case _CReduceBToDirectDeclarator:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CDirectDeclaratorType
		symbol.Generic_, err = reducer.BToDirectDeclarator(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceCToDirectDeclarator:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CDirectDeclaratorType
		symbol.Generic_, err = reducer.CToDirectDeclarator(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _CReduceDToDirectDeclarator:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CDirectDeclaratorType
		symbol.Generic_, err = reducer.DToDirectDeclarator(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceEToDirectDeclarator:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CDirectDeclaratorType
		symbol.Generic_, err = reducer.EToDirectDeclarator(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _CReduceFToDirectDeclarator:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CDirectDeclaratorType
		symbol.Generic_, err = reducer.FToDirectDeclarator(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _CReduceGToDirectDeclarator:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CDirectDeclaratorType
		symbol.Generic_, err = reducer.GToDirectDeclarator(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceAToPointer:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CPointerType
		symbol.Generic_, err = reducer.AToPointer(args[0].Generic_)
	case _CReduceBToPointer:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CPointerType
		symbol.Generic_, err = reducer.BToPointer(args[0].Generic_, args[1].Generic_)
	case _CReduceCToPointer:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CPointerType
		symbol.Generic_, err = reducer.CToPointer(args[0].Generic_, args[1].Generic_)
	case _CReduceDToPointer:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CPointerType
		symbol.Generic_, err = reducer.DToPointer(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceAToTypeQualifierList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CTypeQualifierListType
		symbol.Generic_, err = reducer.AToTypeQualifierList(args[0].Generic_)
	case _CReduceBToTypeQualifierList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CTypeQualifierListType
		symbol.Generic_, err = reducer.BToTypeQualifierList(args[0].Generic_, args[1].Generic_)
	case _CReduceAToParameterTypeList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CParameterTypeListType
		symbol.Generic_, err = reducer.AToParameterTypeList(args[0].Generic_)
	case _CReduceBToParameterTypeList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CParameterTypeListType
		symbol.Generic_, err = reducer.BToParameterTypeList(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceAToParameterList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CParameterListType
		symbol.Generic_, err = reducer.AToParameterList(args[0].Generic_)
	case _CReduceBToParameterList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CParameterListType
		symbol.Generic_, err = reducer.BToParameterList(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceAToParameterDeclaration:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CParameterDeclarationType
		symbol.Generic_, err = reducer.AToParameterDeclaration(args[0].Generic_, args[1].Generic_)
	case _CReduceBToParameterDeclaration:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CParameterDeclarationType
		symbol.Generic_, err = reducer.BToParameterDeclaration(args[0].Generic_, args[1].Generic_)
	case _CReduceCToParameterDeclaration:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CParameterDeclarationType
		symbol.Generic_, err = reducer.CToParameterDeclaration(args[0].Generic_)
	case _CReduceAToIdentifierList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CIdentifierListType
		symbol.Generic_, err = reducer.AToIdentifierList(args[0].Generic_)
	case _CReduceBToIdentifierList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CIdentifierListType
		symbol.Generic_, err = reducer.BToIdentifierList(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceAToTypeName:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CTypeNameType
		symbol.Generic_, err = reducer.AToTypeName(args[0].Generic_)
	case _CReduceBToTypeName:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CTypeNameType
		symbol.Generic_, err = reducer.BToTypeName(args[0].Generic_, args[1].Generic_)
	case _CReduceAToAbstractDeclarator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CAbstractDeclaratorType
		symbol.Generic_, err = reducer.AToAbstractDeclarator(args[0].Generic_)
	case _CReduceBToAbstractDeclarator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CAbstractDeclaratorType
		symbol.Generic_, err = reducer.BToAbstractDeclarator(args[0].Generic_)
	case _CReduceCToAbstractDeclarator:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CAbstractDeclaratorType
		symbol.Generic_, err = reducer.CToAbstractDeclarator(args[0].Generic_, args[1].Generic_)
	case _CReduceAToDirectAbstractDeclarator:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CDirectAbstractDeclaratorType
		symbol.Generic_, err = reducer.AToDirectAbstractDeclarator(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceBToDirectAbstractDeclarator:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CDirectAbstractDeclaratorType
		symbol.Generic_, err = reducer.BToDirectAbstractDeclarator(args[0].Generic_, args[1].Generic_)
	case _CReduceCToDirectAbstractDeclarator:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CDirectAbstractDeclaratorType
		symbol.Generic_, err = reducer.CToDirectAbstractDeclarator(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceDToDirectAbstractDeclarator:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CDirectAbstractDeclaratorType
		symbol.Generic_, err = reducer.DToDirectAbstractDeclarator(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceEToDirectAbstractDeclarator:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CDirectAbstractDeclaratorType
		symbol.Generic_, err = reducer.EToDirectAbstractDeclarator(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _CReduceFToDirectAbstractDeclarator:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CDirectAbstractDeclaratorType
		symbol.Generic_, err = reducer.FToDirectAbstractDeclarator(args[0].Generic_, args[1].Generic_)
	case _CReduceGToDirectAbstractDeclarator:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CDirectAbstractDeclaratorType
		symbol.Generic_, err = reducer.GToDirectAbstractDeclarator(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceHToDirectAbstractDeclarator:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CDirectAbstractDeclaratorType
		symbol.Generic_, err = reducer.HToDirectAbstractDeclarator(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceIToDirectAbstractDeclarator:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CDirectAbstractDeclaratorType
		symbol.Generic_, err = reducer.IToDirectAbstractDeclarator(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _CReduceAToInitializer:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CInitializerType
		symbol.Generic_, err = reducer.AToInitializer(args[0].Generic_)
	case _CReduceBToInitializer:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CInitializerType
		symbol.Generic_, err = reducer.BToInitializer(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceCToInitializer:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CInitializerType
		symbol.Generic_, err = reducer.CToInitializer(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _CReduceAToInitializerList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CInitializerListType
		symbol.Generic_, err = reducer.AToInitializerList(args[0].Generic_)
	case _CReduceBToInitializerList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CInitializerListType
		symbol.Generic_, err = reducer.BToInitializerList(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceAToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CStatementType
		symbol.Generic_, err = reducer.AToStatement(args[0].Generic_)
	case _CReduceBToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CStatementType
		symbol.Generic_, err = reducer.BToStatement(args[0].Generic_)
	case _CReduceCToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CStatementType
		symbol.Generic_, err = reducer.CToStatement(args[0].Generic_)
	case _CReduceDToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CStatementType
		symbol.Generic_, err = reducer.DToStatement(args[0].Generic_)
	case _CReduceEToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CStatementType
		symbol.Generic_, err = reducer.EToStatement(args[0].Generic_)
	case _CReduceFToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CStatementType
		symbol.Generic_, err = reducer.FToStatement(args[0].Generic_)
	case _CReduceAToLabeledStatement:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CLabeledStatementType
		symbol.Generic_, err = reducer.AToLabeledStatement(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceBToLabeledStatement:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CLabeledStatementType
		symbol.Generic_, err = reducer.BToLabeledStatement(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _CReduceCToLabeledStatement:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CLabeledStatementType
		symbol.Generic_, err = reducer.CToLabeledStatement(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceAToCompoundStatement:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CCompoundStatementType
		symbol.Generic_, err = reducer.AToCompoundStatement(args[0].Generic_, args[1].Generic_)
	case _CReduceBToCompoundStatement:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CCompoundStatementType
		symbol.Generic_, err = reducer.BToCompoundStatement(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceCToCompoundStatement:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CCompoundStatementType
		symbol.Generic_, err = reducer.CToCompoundStatement(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceDToCompoundStatement:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CCompoundStatementType
		symbol.Generic_, err = reducer.DToCompoundStatement(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _CReduceAToDeclarationList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CDeclarationListType
		symbol.Generic_, err = reducer.AToDeclarationList(args[0].Generic_)
	case _CReduceBToDeclarationList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CDeclarationListType
		symbol.Generic_, err = reducer.BToDeclarationList(args[0].Generic_, args[1].Generic_)
	case _CReduceAToStatementList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CStatementListType
		symbol.Generic_, err = reducer.AToStatementList(args[0].Generic_)
	case _CReduceBToStatementList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CStatementListType
		symbol.Generic_, err = reducer.BToStatementList(args[0].Generic_, args[1].Generic_)
	case _CReduceAToExpressionStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CExpressionStatementType
		symbol.Generic_, err = reducer.AToExpressionStatement(args[0].Generic_)
	case _CReduceBToExpressionStatement:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CExpressionStatementType
		symbol.Generic_, err = reducer.BToExpressionStatement(args[0].Generic_, args[1].Generic_)
	case _CReduceAToSelectionStatement:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = CSelectionStatementType
		symbol.Generic_, err = reducer.AToSelectionStatement(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_)
	case _CReduceBToSelectionStatement:
		args := stack[len(stack)-7:]
		stack = stack[:len(stack)-7]
		symbol.SymbolId_ = CSelectionStatementType
		symbol.Generic_, err = reducer.BToSelectionStatement(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_, args[5].Generic_, args[6].Generic_)
	case _CReduceCToSelectionStatement:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = CSelectionStatementType
		symbol.Generic_, err = reducer.CToSelectionStatement(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_)
	case _CReduceAToIterationStatement:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = CIterationStatementType
		symbol.Generic_, err = reducer.AToIterationStatement(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_)
	case _CReduceBToIterationStatement:
		args := stack[len(stack)-7:]
		stack = stack[:len(stack)-7]
		symbol.SymbolId_ = CIterationStatementType
		symbol.Generic_, err = reducer.BToIterationStatement(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_, args[5].Generic_, args[6].Generic_)
	case _CReduceCToIterationStatement:
		args := stack[len(stack)-6:]
		stack = stack[:len(stack)-6]
		symbol.SymbolId_ = CIterationStatementType
		symbol.Generic_, err = reducer.CToIterationStatement(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_, args[5].Generic_)
	case _CReduceDToIterationStatement:
		args := stack[len(stack)-7:]
		stack = stack[:len(stack)-7]
		symbol.SymbolId_ = CIterationStatementType
		symbol.Generic_, err = reducer.DToIterationStatement(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_, args[5].Generic_, args[6].Generic_)
	case _CReduceAToJumpStatement:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CJumpStatementType
		symbol.Generic_, err = reducer.AToJumpStatement(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceBToJumpStatement:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CJumpStatementType
		symbol.Generic_, err = reducer.BToJumpStatement(args[0].Generic_, args[1].Generic_)
	case _CReduceCToJumpStatement:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CJumpStatementType
		symbol.Generic_, err = reducer.CToJumpStatement(args[0].Generic_, args[1].Generic_)
	case _CReduceDToJumpStatement:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CJumpStatementType
		symbol.Generic_, err = reducer.DToJumpStatement(args[0].Generic_, args[1].Generic_)
	case _CReduceEToJumpStatement:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CJumpStatementType
		symbol.Generic_, err = reducer.EToJumpStatement(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceAToTranslationUnit:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CTranslationUnitType
		symbol.Generic_, err = reducer.AToTranslationUnit(args[0].Generic_)
	case _CReduceBToTranslationUnit:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CTranslationUnitType
		symbol.Generic_, err = reducer.BToTranslationUnit(args[0].Generic_, args[1].Generic_)
	case _CReduceAToExternalDeclaration:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CExternalDeclarationType
		symbol.Generic_, err = reducer.AToExternalDeclaration(args[0].Generic_)
	case _CReduceBToExternalDeclaration:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CExternalDeclarationType
		symbol.Generic_, err = reducer.BToExternalDeclaration(args[0].Generic_)
	case _CReduceAToFunctionDefinition:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CFunctionDefinitionType
		symbol.Generic_, err = reducer.AToFunctionDefinition(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _CReduceBToFunctionDefinition:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CFunctionDefinitionType
		symbol.Generic_, err = reducer.BToFunctionDefinition(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceCToFunctionDefinition:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CFunctionDefinitionType
		symbol.Generic_, err = reducer.CToFunctionDefinition(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _CReduceDToFunctionDefinition:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CFunctionDefinitionType
		symbol.Generic_, err = reducer.DToFunctionDefinition(args[0].Generic_, args[1].Generic_)
	default:
		panic("Unknown reduce type: " + act.ReduceType.String())
	}

	if err != nil {
		err = fmt.Errorf("Unexpected %s reduce error: %s", act.ReduceType, err)
	}

	return stack, symbol, err
}

type _CActionTableKey struct {
	_CStateId
	CSymbolId
}

type _CActionTableType map[_CActionTableKey]*_CAction

func (table _CActionTableType) Get(
	stateId _CStateId,
	symbolId CSymbolId) (
	*_CAction,
	bool) {

	action, ok := table[_CActionTableKey{stateId, symbolId}]
	if ok {
		return action, ok
	}

	action, ok = table[_CActionTableKey{stateId, _CWildcardMarker}]
	return action, ok
}

var (
	_CGotoState1Action                        = &_CAction{_CShiftAction, _CState1, 0}
	_CGotoState2Action                        = &_CAction{_CShiftAction, _CState2, 0}
	_CGotoState3Action                        = &_CAction{_CShiftAction, _CState3, 0}
	_CGotoState4Action                        = &_CAction{_CShiftAction, _CState4, 0}
	_CGotoState5Action                        = &_CAction{_CShiftAction, _CState5, 0}
	_CGotoState6Action                        = &_CAction{_CShiftAction, _CState6, 0}
	_CGotoState7Action                        = &_CAction{_CShiftAction, _CState7, 0}
	_CGotoState8Action                        = &_CAction{_CShiftAction, _CState8, 0}
	_CGotoState9Action                        = &_CAction{_CShiftAction, _CState9, 0}
	_CGotoState10Action                       = &_CAction{_CShiftAction, _CState10, 0}
	_CGotoState11Action                       = &_CAction{_CShiftAction, _CState11, 0}
	_CGotoState12Action                       = &_CAction{_CShiftAction, _CState12, 0}
	_CGotoState13Action                       = &_CAction{_CShiftAction, _CState13, 0}
	_CGotoState14Action                       = &_CAction{_CShiftAction, _CState14, 0}
	_CGotoState15Action                       = &_CAction{_CShiftAction, _CState15, 0}
	_CGotoState16Action                       = &_CAction{_CShiftAction, _CState16, 0}
	_CGotoState17Action                       = &_CAction{_CShiftAction, _CState17, 0}
	_CGotoState18Action                       = &_CAction{_CShiftAction, _CState18, 0}
	_CGotoState19Action                       = &_CAction{_CShiftAction, _CState19, 0}
	_CGotoState20Action                       = &_CAction{_CShiftAction, _CState20, 0}
	_CGotoState21Action                       = &_CAction{_CShiftAction, _CState21, 0}
	_CGotoState22Action                       = &_CAction{_CShiftAction, _CState22, 0}
	_CGotoState23Action                       = &_CAction{_CShiftAction, _CState23, 0}
	_CGotoState24Action                       = &_CAction{_CShiftAction, _CState24, 0}
	_CGotoState25Action                       = &_CAction{_CShiftAction, _CState25, 0}
	_CGotoState26Action                       = &_CAction{_CShiftAction, _CState26, 0}
	_CGotoState27Action                       = &_CAction{_CShiftAction, _CState27, 0}
	_CGotoState28Action                       = &_CAction{_CShiftAction, _CState28, 0}
	_CGotoState29Action                       = &_CAction{_CShiftAction, _CState29, 0}
	_CGotoState30Action                       = &_CAction{_CShiftAction, _CState30, 0}
	_CGotoState31Action                       = &_CAction{_CShiftAction, _CState31, 0}
	_CGotoState32Action                       = &_CAction{_CShiftAction, _CState32, 0}
	_CGotoState33Action                       = &_CAction{_CShiftAction, _CState33, 0}
	_CGotoState34Action                       = &_CAction{_CShiftAction, _CState34, 0}
	_CGotoState35Action                       = &_CAction{_CShiftAction, _CState35, 0}
	_CGotoState36Action                       = &_CAction{_CShiftAction, _CState36, 0}
	_CGotoState37Action                       = &_CAction{_CShiftAction, _CState37, 0}
	_CGotoState38Action                       = &_CAction{_CShiftAction, _CState38, 0}
	_CGotoState39Action                       = &_CAction{_CShiftAction, _CState39, 0}
	_CGotoState40Action                       = &_CAction{_CShiftAction, _CState40, 0}
	_CGotoState41Action                       = &_CAction{_CShiftAction, _CState41, 0}
	_CGotoState42Action                       = &_CAction{_CShiftAction, _CState42, 0}
	_CGotoState43Action                       = &_CAction{_CShiftAction, _CState43, 0}
	_CGotoState44Action                       = &_CAction{_CShiftAction, _CState44, 0}
	_CGotoState45Action                       = &_CAction{_CShiftAction, _CState45, 0}
	_CGotoState46Action                       = &_CAction{_CShiftAction, _CState46, 0}
	_CGotoState47Action                       = &_CAction{_CShiftAction, _CState47, 0}
	_CGotoState48Action                       = &_CAction{_CShiftAction, _CState48, 0}
	_CGotoState49Action                       = &_CAction{_CShiftAction, _CState49, 0}
	_CGotoState50Action                       = &_CAction{_CShiftAction, _CState50, 0}
	_CGotoState51Action                       = &_CAction{_CShiftAction, _CState51, 0}
	_CGotoState52Action                       = &_CAction{_CShiftAction, _CState52, 0}
	_CGotoState53Action                       = &_CAction{_CShiftAction, _CState53, 0}
	_CGotoState54Action                       = &_CAction{_CShiftAction, _CState54, 0}
	_CGotoState55Action                       = &_CAction{_CShiftAction, _CState55, 0}
	_CGotoState56Action                       = &_CAction{_CShiftAction, _CState56, 0}
	_CGotoState57Action                       = &_CAction{_CShiftAction, _CState57, 0}
	_CGotoState58Action                       = &_CAction{_CShiftAction, _CState58, 0}
	_CGotoState59Action                       = &_CAction{_CShiftAction, _CState59, 0}
	_CGotoState60Action                       = &_CAction{_CShiftAction, _CState60, 0}
	_CGotoState61Action                       = &_CAction{_CShiftAction, _CState61, 0}
	_CGotoState62Action                       = &_CAction{_CShiftAction, _CState62, 0}
	_CGotoState63Action                       = &_CAction{_CShiftAction, _CState63, 0}
	_CGotoState64Action                       = &_CAction{_CShiftAction, _CState64, 0}
	_CGotoState65Action                       = &_CAction{_CShiftAction, _CState65, 0}
	_CGotoState66Action                       = &_CAction{_CShiftAction, _CState66, 0}
	_CGotoState67Action                       = &_CAction{_CShiftAction, _CState67, 0}
	_CGotoState68Action                       = &_CAction{_CShiftAction, _CState68, 0}
	_CGotoState69Action                       = &_CAction{_CShiftAction, _CState69, 0}
	_CGotoState70Action                       = &_CAction{_CShiftAction, _CState70, 0}
	_CGotoState71Action                       = &_CAction{_CShiftAction, _CState71, 0}
	_CGotoState72Action                       = &_CAction{_CShiftAction, _CState72, 0}
	_CGotoState73Action                       = &_CAction{_CShiftAction, _CState73, 0}
	_CGotoState74Action                       = &_CAction{_CShiftAction, _CState74, 0}
	_CGotoState75Action                       = &_CAction{_CShiftAction, _CState75, 0}
	_CGotoState76Action                       = &_CAction{_CShiftAction, _CState76, 0}
	_CGotoState77Action                       = &_CAction{_CShiftAction, _CState77, 0}
	_CGotoState78Action                       = &_CAction{_CShiftAction, _CState78, 0}
	_CGotoState79Action                       = &_CAction{_CShiftAction, _CState79, 0}
	_CGotoState80Action                       = &_CAction{_CShiftAction, _CState80, 0}
	_CGotoState81Action                       = &_CAction{_CShiftAction, _CState81, 0}
	_CGotoState82Action                       = &_CAction{_CShiftAction, _CState82, 0}
	_CGotoState83Action                       = &_CAction{_CShiftAction, _CState83, 0}
	_CGotoState84Action                       = &_CAction{_CShiftAction, _CState84, 0}
	_CGotoState85Action                       = &_CAction{_CShiftAction, _CState85, 0}
	_CGotoState86Action                       = &_CAction{_CShiftAction, _CState86, 0}
	_CGotoState87Action                       = &_CAction{_CShiftAction, _CState87, 0}
	_CGotoState88Action                       = &_CAction{_CShiftAction, _CState88, 0}
	_CGotoState89Action                       = &_CAction{_CShiftAction, _CState89, 0}
	_CGotoState90Action                       = &_CAction{_CShiftAction, _CState90, 0}
	_CGotoState91Action                       = &_CAction{_CShiftAction, _CState91, 0}
	_CGotoState92Action                       = &_CAction{_CShiftAction, _CState92, 0}
	_CGotoState93Action                       = &_CAction{_CShiftAction, _CState93, 0}
	_CGotoState94Action                       = &_CAction{_CShiftAction, _CState94, 0}
	_CGotoState95Action                       = &_CAction{_CShiftAction, _CState95, 0}
	_CGotoState96Action                       = &_CAction{_CShiftAction, _CState96, 0}
	_CGotoState97Action                       = &_CAction{_CShiftAction, _CState97, 0}
	_CGotoState98Action                       = &_CAction{_CShiftAction, _CState98, 0}
	_CGotoState99Action                       = &_CAction{_CShiftAction, _CState99, 0}
	_CGotoState100Action                      = &_CAction{_CShiftAction, _CState100, 0}
	_CGotoState101Action                      = &_CAction{_CShiftAction, _CState101, 0}
	_CGotoState102Action                      = &_CAction{_CShiftAction, _CState102, 0}
	_CGotoState103Action                      = &_CAction{_CShiftAction, _CState103, 0}
	_CGotoState104Action                      = &_CAction{_CShiftAction, _CState104, 0}
	_CGotoState105Action                      = &_CAction{_CShiftAction, _CState105, 0}
	_CGotoState106Action                      = &_CAction{_CShiftAction, _CState106, 0}
	_CGotoState107Action                      = &_CAction{_CShiftAction, _CState107, 0}
	_CGotoState108Action                      = &_CAction{_CShiftAction, _CState108, 0}
	_CGotoState109Action                      = &_CAction{_CShiftAction, _CState109, 0}
	_CGotoState110Action                      = &_CAction{_CShiftAction, _CState110, 0}
	_CGotoState111Action                      = &_CAction{_CShiftAction, _CState111, 0}
	_CGotoState112Action                      = &_CAction{_CShiftAction, _CState112, 0}
	_CGotoState113Action                      = &_CAction{_CShiftAction, _CState113, 0}
	_CGotoState114Action                      = &_CAction{_CShiftAction, _CState114, 0}
	_CGotoState115Action                      = &_CAction{_CShiftAction, _CState115, 0}
	_CGotoState116Action                      = &_CAction{_CShiftAction, _CState116, 0}
	_CGotoState117Action                      = &_CAction{_CShiftAction, _CState117, 0}
	_CGotoState118Action                      = &_CAction{_CShiftAction, _CState118, 0}
	_CGotoState119Action                      = &_CAction{_CShiftAction, _CState119, 0}
	_CGotoState120Action                      = &_CAction{_CShiftAction, _CState120, 0}
	_CGotoState121Action                      = &_CAction{_CShiftAction, _CState121, 0}
	_CGotoState122Action                      = &_CAction{_CShiftAction, _CState122, 0}
	_CGotoState123Action                      = &_CAction{_CShiftAction, _CState123, 0}
	_CGotoState124Action                      = &_CAction{_CShiftAction, _CState124, 0}
	_CGotoState125Action                      = &_CAction{_CShiftAction, _CState125, 0}
	_CGotoState126Action                      = &_CAction{_CShiftAction, _CState126, 0}
	_CGotoState127Action                      = &_CAction{_CShiftAction, _CState127, 0}
	_CGotoState128Action                      = &_CAction{_CShiftAction, _CState128, 0}
	_CGotoState129Action                      = &_CAction{_CShiftAction, _CState129, 0}
	_CGotoState130Action                      = &_CAction{_CShiftAction, _CState130, 0}
	_CGotoState131Action                      = &_CAction{_CShiftAction, _CState131, 0}
	_CGotoState132Action                      = &_CAction{_CShiftAction, _CState132, 0}
	_CGotoState133Action                      = &_CAction{_CShiftAction, _CState133, 0}
	_CGotoState134Action                      = &_CAction{_CShiftAction, _CState134, 0}
	_CGotoState135Action                      = &_CAction{_CShiftAction, _CState135, 0}
	_CGotoState136Action                      = &_CAction{_CShiftAction, _CState136, 0}
	_CGotoState137Action                      = &_CAction{_CShiftAction, _CState137, 0}
	_CGotoState138Action                      = &_CAction{_CShiftAction, _CState138, 0}
	_CGotoState139Action                      = &_CAction{_CShiftAction, _CState139, 0}
	_CGotoState140Action                      = &_CAction{_CShiftAction, _CState140, 0}
	_CGotoState141Action                      = &_CAction{_CShiftAction, _CState141, 0}
	_CGotoState142Action                      = &_CAction{_CShiftAction, _CState142, 0}
	_CGotoState143Action                      = &_CAction{_CShiftAction, _CState143, 0}
	_CGotoState144Action                      = &_CAction{_CShiftAction, _CState144, 0}
	_CGotoState145Action                      = &_CAction{_CShiftAction, _CState145, 0}
	_CGotoState146Action                      = &_CAction{_CShiftAction, _CState146, 0}
	_CGotoState147Action                      = &_CAction{_CShiftAction, _CState147, 0}
	_CGotoState148Action                      = &_CAction{_CShiftAction, _CState148, 0}
	_CGotoState149Action                      = &_CAction{_CShiftAction, _CState149, 0}
	_CGotoState150Action                      = &_CAction{_CShiftAction, _CState150, 0}
	_CGotoState151Action                      = &_CAction{_CShiftAction, _CState151, 0}
	_CGotoState152Action                      = &_CAction{_CShiftAction, _CState152, 0}
	_CGotoState153Action                      = &_CAction{_CShiftAction, _CState153, 0}
	_CGotoState154Action                      = &_CAction{_CShiftAction, _CState154, 0}
	_CGotoState155Action                      = &_CAction{_CShiftAction, _CState155, 0}
	_CGotoState156Action                      = &_CAction{_CShiftAction, _CState156, 0}
	_CGotoState157Action                      = &_CAction{_CShiftAction, _CState157, 0}
	_CGotoState158Action                      = &_CAction{_CShiftAction, _CState158, 0}
	_CGotoState159Action                      = &_CAction{_CShiftAction, _CState159, 0}
	_CGotoState160Action                      = &_CAction{_CShiftAction, _CState160, 0}
	_CGotoState161Action                      = &_CAction{_CShiftAction, _CState161, 0}
	_CGotoState162Action                      = &_CAction{_CShiftAction, _CState162, 0}
	_CGotoState163Action                      = &_CAction{_CShiftAction, _CState163, 0}
	_CGotoState164Action                      = &_CAction{_CShiftAction, _CState164, 0}
	_CGotoState165Action                      = &_CAction{_CShiftAction, _CState165, 0}
	_CGotoState166Action                      = &_CAction{_CShiftAction, _CState166, 0}
	_CGotoState167Action                      = &_CAction{_CShiftAction, _CState167, 0}
	_CGotoState168Action                      = &_CAction{_CShiftAction, _CState168, 0}
	_CGotoState169Action                      = &_CAction{_CShiftAction, _CState169, 0}
	_CGotoState170Action                      = &_CAction{_CShiftAction, _CState170, 0}
	_CGotoState171Action                      = &_CAction{_CShiftAction, _CState171, 0}
	_CGotoState172Action                      = &_CAction{_CShiftAction, _CState172, 0}
	_CGotoState173Action                      = &_CAction{_CShiftAction, _CState173, 0}
	_CGotoState174Action                      = &_CAction{_CShiftAction, _CState174, 0}
	_CGotoState175Action                      = &_CAction{_CShiftAction, _CState175, 0}
	_CGotoState176Action                      = &_CAction{_CShiftAction, _CState176, 0}
	_CGotoState177Action                      = &_CAction{_CShiftAction, _CState177, 0}
	_CGotoState178Action                      = &_CAction{_CShiftAction, _CState178, 0}
	_CGotoState179Action                      = &_CAction{_CShiftAction, _CState179, 0}
	_CGotoState180Action                      = &_CAction{_CShiftAction, _CState180, 0}
	_CGotoState181Action                      = &_CAction{_CShiftAction, _CState181, 0}
	_CGotoState182Action                      = &_CAction{_CShiftAction, _CState182, 0}
	_CGotoState183Action                      = &_CAction{_CShiftAction, _CState183, 0}
	_CGotoState184Action                      = &_CAction{_CShiftAction, _CState184, 0}
	_CGotoState185Action                      = &_CAction{_CShiftAction, _CState185, 0}
	_CGotoState186Action                      = &_CAction{_CShiftAction, _CState186, 0}
	_CGotoState187Action                      = &_CAction{_CShiftAction, _CState187, 0}
	_CGotoState188Action                      = &_CAction{_CShiftAction, _CState188, 0}
	_CGotoState189Action                      = &_CAction{_CShiftAction, _CState189, 0}
	_CGotoState190Action                      = &_CAction{_CShiftAction, _CState190, 0}
	_CGotoState191Action                      = &_CAction{_CShiftAction, _CState191, 0}
	_CGotoState192Action                      = &_CAction{_CShiftAction, _CState192, 0}
	_CGotoState193Action                      = &_CAction{_CShiftAction, _CState193, 0}
	_CGotoState194Action                      = &_CAction{_CShiftAction, _CState194, 0}
	_CGotoState195Action                      = &_CAction{_CShiftAction, _CState195, 0}
	_CGotoState196Action                      = &_CAction{_CShiftAction, _CState196, 0}
	_CGotoState197Action                      = &_CAction{_CShiftAction, _CState197, 0}
	_CGotoState198Action                      = &_CAction{_CShiftAction, _CState198, 0}
	_CGotoState199Action                      = &_CAction{_CShiftAction, _CState199, 0}
	_CGotoState200Action                      = &_CAction{_CShiftAction, _CState200, 0}
	_CGotoState201Action                      = &_CAction{_CShiftAction, _CState201, 0}
	_CGotoState202Action                      = &_CAction{_CShiftAction, _CState202, 0}
	_CGotoState203Action                      = &_CAction{_CShiftAction, _CState203, 0}
	_CGotoState204Action                      = &_CAction{_CShiftAction, _CState204, 0}
	_CGotoState205Action                      = &_CAction{_CShiftAction, _CState205, 0}
	_CGotoState206Action                      = &_CAction{_CShiftAction, _CState206, 0}
	_CGotoState207Action                      = &_CAction{_CShiftAction, _CState207, 0}
	_CGotoState208Action                      = &_CAction{_CShiftAction, _CState208, 0}
	_CGotoState209Action                      = &_CAction{_CShiftAction, _CState209, 0}
	_CGotoState210Action                      = &_CAction{_CShiftAction, _CState210, 0}
	_CGotoState211Action                      = &_CAction{_CShiftAction, _CState211, 0}
	_CGotoState212Action                      = &_CAction{_CShiftAction, _CState212, 0}
	_CGotoState213Action                      = &_CAction{_CShiftAction, _CState213, 0}
	_CGotoState214Action                      = &_CAction{_CShiftAction, _CState214, 0}
	_CGotoState215Action                      = &_CAction{_CShiftAction, _CState215, 0}
	_CGotoState216Action                      = &_CAction{_CShiftAction, _CState216, 0}
	_CGotoState217Action                      = &_CAction{_CShiftAction, _CState217, 0}
	_CGotoState218Action                      = &_CAction{_CShiftAction, _CState218, 0}
	_CGotoState219Action                      = &_CAction{_CShiftAction, _CState219, 0}
	_CGotoState220Action                      = &_CAction{_CShiftAction, _CState220, 0}
	_CGotoState221Action                      = &_CAction{_CShiftAction, _CState221, 0}
	_CGotoState222Action                      = &_CAction{_CShiftAction, _CState222, 0}
	_CGotoState223Action                      = &_CAction{_CShiftAction, _CState223, 0}
	_CGotoState224Action                      = &_CAction{_CShiftAction, _CState224, 0}
	_CGotoState225Action                      = &_CAction{_CShiftAction, _CState225, 0}
	_CGotoState226Action                      = &_CAction{_CShiftAction, _CState226, 0}
	_CGotoState227Action                      = &_CAction{_CShiftAction, _CState227, 0}
	_CGotoState228Action                      = &_CAction{_CShiftAction, _CState228, 0}
	_CGotoState229Action                      = &_CAction{_CShiftAction, _CState229, 0}
	_CGotoState230Action                      = &_CAction{_CShiftAction, _CState230, 0}
	_CGotoState231Action                      = &_CAction{_CShiftAction, _CState231, 0}
	_CGotoState232Action                      = &_CAction{_CShiftAction, _CState232, 0}
	_CGotoState233Action                      = &_CAction{_CShiftAction, _CState233, 0}
	_CGotoState234Action                      = &_CAction{_CShiftAction, _CState234, 0}
	_CGotoState235Action                      = &_CAction{_CShiftAction, _CState235, 0}
	_CGotoState236Action                      = &_CAction{_CShiftAction, _CState236, 0}
	_CGotoState237Action                      = &_CAction{_CShiftAction, _CState237, 0}
	_CGotoState238Action                      = &_CAction{_CShiftAction, _CState238, 0}
	_CGotoState239Action                      = &_CAction{_CShiftAction, _CState239, 0}
	_CGotoState240Action                      = &_CAction{_CShiftAction, _CState240, 0}
	_CGotoState241Action                      = &_CAction{_CShiftAction, _CState241, 0}
	_CGotoState242Action                      = &_CAction{_CShiftAction, _CState242, 0}
	_CGotoState243Action                      = &_CAction{_CShiftAction, _CState243, 0}
	_CGotoState244Action                      = &_CAction{_CShiftAction, _CState244, 0}
	_CGotoState245Action                      = &_CAction{_CShiftAction, _CState245, 0}
	_CGotoState246Action                      = &_CAction{_CShiftAction, _CState246, 0}
	_CGotoState247Action                      = &_CAction{_CShiftAction, _CState247, 0}
	_CGotoState248Action                      = &_CAction{_CShiftAction, _CState248, 0}
	_CGotoState249Action                      = &_CAction{_CShiftAction, _CState249, 0}
	_CGotoState250Action                      = &_CAction{_CShiftAction, _CState250, 0}
	_CGotoState251Action                      = &_CAction{_CShiftAction, _CState251, 0}
	_CGotoState252Action                      = &_CAction{_CShiftAction, _CState252, 0}
	_CGotoState253Action                      = &_CAction{_CShiftAction, _CState253, 0}
	_CGotoState254Action                      = &_CAction{_CShiftAction, _CState254, 0}
	_CGotoState255Action                      = &_CAction{_CShiftAction, _CState255, 0}
	_CGotoState256Action                      = &_CAction{_CShiftAction, _CState256, 0}
	_CGotoState257Action                      = &_CAction{_CShiftAction, _CState257, 0}
	_CGotoState258Action                      = &_CAction{_CShiftAction, _CState258, 0}
	_CGotoState259Action                      = &_CAction{_CShiftAction, _CState259, 0}
	_CGotoState260Action                      = &_CAction{_CShiftAction, _CState260, 0}
	_CGotoState261Action                      = &_CAction{_CShiftAction, _CState261, 0}
	_CGotoState262Action                      = &_CAction{_CShiftAction, _CState262, 0}
	_CGotoState263Action                      = &_CAction{_CShiftAction, _CState263, 0}
	_CGotoState264Action                      = &_CAction{_CShiftAction, _CState264, 0}
	_CGotoState265Action                      = &_CAction{_CShiftAction, _CState265, 0}
	_CGotoState266Action                      = &_CAction{_CShiftAction, _CState266, 0}
	_CGotoState267Action                      = &_CAction{_CShiftAction, _CState267, 0}
	_CGotoState268Action                      = &_CAction{_CShiftAction, _CState268, 0}
	_CGotoState269Action                      = &_CAction{_CShiftAction, _CState269, 0}
	_CGotoState270Action                      = &_CAction{_CShiftAction, _CState270, 0}
	_CGotoState271Action                      = &_CAction{_CShiftAction, _CState271, 0}
	_CGotoState272Action                      = &_CAction{_CShiftAction, _CState272, 0}
	_CGotoState273Action                      = &_CAction{_CShiftAction, _CState273, 0}
	_CGotoState274Action                      = &_CAction{_CShiftAction, _CState274, 0}
	_CGotoState275Action                      = &_CAction{_CShiftAction, _CState275, 0}
	_CGotoState276Action                      = &_CAction{_CShiftAction, _CState276, 0}
	_CGotoState277Action                      = &_CAction{_CShiftAction, _CState277, 0}
	_CGotoState278Action                      = &_CAction{_CShiftAction, _CState278, 0}
	_CGotoState279Action                      = &_CAction{_CShiftAction, _CState279, 0}
	_CGotoState280Action                      = &_CAction{_CShiftAction, _CState280, 0}
	_CGotoState281Action                      = &_CAction{_CShiftAction, _CState281, 0}
	_CGotoState282Action                      = &_CAction{_CShiftAction, _CState282, 0}
	_CGotoState283Action                      = &_CAction{_CShiftAction, _CState283, 0}
	_CGotoState284Action                      = &_CAction{_CShiftAction, _CState284, 0}
	_CGotoState285Action                      = &_CAction{_CShiftAction, _CState285, 0}
	_CGotoState286Action                      = &_CAction{_CShiftAction, _CState286, 0}
	_CGotoState287Action                      = &_CAction{_CShiftAction, _CState287, 0}
	_CGotoState288Action                      = &_CAction{_CShiftAction, _CState288, 0}
	_CGotoState289Action                      = &_CAction{_CShiftAction, _CState289, 0}
	_CGotoState290Action                      = &_CAction{_CShiftAction, _CState290, 0}
	_CGotoState291Action                      = &_CAction{_CShiftAction, _CState291, 0}
	_CGotoState292Action                      = &_CAction{_CShiftAction, _CState292, 0}
	_CGotoState293Action                      = &_CAction{_CShiftAction, _CState293, 0}
	_CGotoState294Action                      = &_CAction{_CShiftAction, _CState294, 0}
	_CGotoState295Action                      = &_CAction{_CShiftAction, _CState295, 0}
	_CGotoState296Action                      = &_CAction{_CShiftAction, _CState296, 0}
	_CGotoState297Action                      = &_CAction{_CShiftAction, _CState297, 0}
	_CGotoState298Action                      = &_CAction{_CShiftAction, _CState298, 0}
	_CGotoState299Action                      = &_CAction{_CShiftAction, _CState299, 0}
	_CGotoState300Action                      = &_CAction{_CShiftAction, _CState300, 0}
	_CGotoState301Action                      = &_CAction{_CShiftAction, _CState301, 0}
	_CGotoState302Action                      = &_CAction{_CShiftAction, _CState302, 0}
	_CGotoState303Action                      = &_CAction{_CShiftAction, _CState303, 0}
	_CGotoState304Action                      = &_CAction{_CShiftAction, _CState304, 0}
	_CGotoState305Action                      = &_CAction{_CShiftAction, _CState305, 0}
	_CGotoState306Action                      = &_CAction{_CShiftAction, _CState306, 0}
	_CGotoState307Action                      = &_CAction{_CShiftAction, _CState307, 0}
	_CGotoState308Action                      = &_CAction{_CShiftAction, _CState308, 0}
	_CGotoState309Action                      = &_CAction{_CShiftAction, _CState309, 0}
	_CGotoState310Action                      = &_CAction{_CShiftAction, _CState310, 0}
	_CGotoState311Action                      = &_CAction{_CShiftAction, _CState311, 0}
	_CGotoState312Action                      = &_CAction{_CShiftAction, _CState312, 0}
	_CGotoState313Action                      = &_CAction{_CShiftAction, _CState313, 0}
	_CGotoState314Action                      = &_CAction{_CShiftAction, _CState314, 0}
	_CGotoState315Action                      = &_CAction{_CShiftAction, _CState315, 0}
	_CGotoState316Action                      = &_CAction{_CShiftAction, _CState316, 0}
	_CGotoState317Action                      = &_CAction{_CShiftAction, _CState317, 0}
	_CGotoState318Action                      = &_CAction{_CShiftAction, _CState318, 0}
	_CGotoState319Action                      = &_CAction{_CShiftAction, _CState319, 0}
	_CGotoState320Action                      = &_CAction{_CShiftAction, _CState320, 0}
	_CGotoState321Action                      = &_CAction{_CShiftAction, _CState321, 0}
	_CGotoState322Action                      = &_CAction{_CShiftAction, _CState322, 0}
	_CGotoState323Action                      = &_CAction{_CShiftAction, _CState323, 0}
	_CGotoState324Action                      = &_CAction{_CShiftAction, _CState324, 0}
	_CGotoState325Action                      = &_CAction{_CShiftAction, _CState325, 0}
	_CGotoState326Action                      = &_CAction{_CShiftAction, _CState326, 0}
	_CGotoState327Action                      = &_CAction{_CShiftAction, _CState327, 0}
	_CGotoState328Action                      = &_CAction{_CShiftAction, _CState328, 0}
	_CGotoState329Action                      = &_CAction{_CShiftAction, _CState329, 0}
	_CGotoState330Action                      = &_CAction{_CShiftAction, _CState330, 0}
	_CGotoState331Action                      = &_CAction{_CShiftAction, _CState331, 0}
	_CGotoState332Action                      = &_CAction{_CShiftAction, _CState332, 0}
	_CGotoState333Action                      = &_CAction{_CShiftAction, _CState333, 0}
	_CGotoState334Action                      = &_CAction{_CShiftAction, _CState334, 0}
	_CGotoState335Action                      = &_CAction{_CShiftAction, _CState335, 0}
	_CGotoState336Action                      = &_CAction{_CShiftAction, _CState336, 0}
	_CGotoState337Action                      = &_CAction{_CShiftAction, _CState337, 0}
	_CGotoState338Action                      = &_CAction{_CShiftAction, _CState338, 0}
	_CGotoState339Action                      = &_CAction{_CShiftAction, _CState339, 0}
	_CGotoState340Action                      = &_CAction{_CShiftAction, _CState340, 0}
	_CGotoState341Action                      = &_CAction{_CShiftAction, _CState341, 0}
	_CGotoState342Action                      = &_CAction{_CShiftAction, _CState342, 0}
	_CGotoState343Action                      = &_CAction{_CShiftAction, _CState343, 0}
	_CGotoState344Action                      = &_CAction{_CShiftAction, _CState344, 0}
	_CGotoState345Action                      = &_CAction{_CShiftAction, _CState345, 0}
	_CGotoState346Action                      = &_CAction{_CShiftAction, _CState346, 0}
	_CGotoState347Action                      = &_CAction{_CShiftAction, _CState347, 0}
	_CGotoState348Action                      = &_CAction{_CShiftAction, _CState348, 0}
	_CGotoState349Action                      = &_CAction{_CShiftAction, _CState349, 0}
	_CGotoState350Action                      = &_CAction{_CShiftAction, _CState350, 0}
	_CGotoState351Action                      = &_CAction{_CShiftAction, _CState351, 0}
	_CReduceAToPrimaryExpressionAction        = &_CAction{_CReduceAction, 0, _CReduceAToPrimaryExpression}
	_CReduceBToPrimaryExpressionAction        = &_CAction{_CReduceAction, 0, _CReduceBToPrimaryExpression}
	_CReduceCToPrimaryExpressionAction        = &_CAction{_CReduceAction, 0, _CReduceCToPrimaryExpression}
	_CReduceDToPrimaryExpressionAction        = &_CAction{_CReduceAction, 0, _CReduceDToPrimaryExpression}
	_CReduceAToPostfixExpressionAction        = &_CAction{_CReduceAction, 0, _CReduceAToPostfixExpression}
	_CReduceBToPostfixExpressionAction        = &_CAction{_CReduceAction, 0, _CReduceBToPostfixExpression}
	_CReduceCToPostfixExpressionAction        = &_CAction{_CReduceAction, 0, _CReduceCToPostfixExpression}
	_CReduceDToPostfixExpressionAction        = &_CAction{_CReduceAction, 0, _CReduceDToPostfixExpression}
	_CReduceEToPostfixExpressionAction        = &_CAction{_CReduceAction, 0, _CReduceEToPostfixExpression}
	_CReduceFToPostfixExpressionAction        = &_CAction{_CReduceAction, 0, _CReduceFToPostfixExpression}
	_CReduceGToPostfixExpressionAction        = &_CAction{_CReduceAction, 0, _CReduceGToPostfixExpression}
	_CReduceHToPostfixExpressionAction        = &_CAction{_CReduceAction, 0, _CReduceHToPostfixExpression}
	_CReduceAToArgumentExpressionListAction   = &_CAction{_CReduceAction, 0, _CReduceAToArgumentExpressionList}
	_CReduceBToArgumentExpressionListAction   = &_CAction{_CReduceAction, 0, _CReduceBToArgumentExpressionList}
	_CReduceAToUnaryExpressionAction          = &_CAction{_CReduceAction, 0, _CReduceAToUnaryExpression}
	_CReduceBToUnaryExpressionAction          = &_CAction{_CReduceAction, 0, _CReduceBToUnaryExpression}
	_CReduceCToUnaryExpressionAction          = &_CAction{_CReduceAction, 0, _CReduceCToUnaryExpression}
	_CReduceDToUnaryExpressionAction          = &_CAction{_CReduceAction, 0, _CReduceDToUnaryExpression}
	_CReduceEToUnaryExpressionAction          = &_CAction{_CReduceAction, 0, _CReduceEToUnaryExpression}
	_CReduceFToUnaryExpressionAction          = &_CAction{_CReduceAction, 0, _CReduceFToUnaryExpression}
	_CReduceAToUnaryOperatorAction            = &_CAction{_CReduceAction, 0, _CReduceAToUnaryOperator}
	_CReduceBToUnaryOperatorAction            = &_CAction{_CReduceAction, 0, _CReduceBToUnaryOperator}
	_CReduceCToUnaryOperatorAction            = &_CAction{_CReduceAction, 0, _CReduceCToUnaryOperator}
	_CReduceDToUnaryOperatorAction            = &_CAction{_CReduceAction, 0, _CReduceDToUnaryOperator}
	_CReduceEToUnaryOperatorAction            = &_CAction{_CReduceAction, 0, _CReduceEToUnaryOperator}
	_CReduceFToUnaryOperatorAction            = &_CAction{_CReduceAction, 0, _CReduceFToUnaryOperator}
	_CReduceAToCastExpressionAction           = &_CAction{_CReduceAction, 0, _CReduceAToCastExpression}
	_CReduceBToCastExpressionAction           = &_CAction{_CReduceAction, 0, _CReduceBToCastExpression}
	_CReduceAToMultiplicativeExpressionAction = &_CAction{_CReduceAction, 0, _CReduceAToMultiplicativeExpression}
	_CReduceBToMultiplicativeExpressionAction = &_CAction{_CReduceAction, 0, _CReduceBToMultiplicativeExpression}
	_CReduceCToMultiplicativeExpressionAction = &_CAction{_CReduceAction, 0, _CReduceCToMultiplicativeExpression}
	_CReduceDToMultiplicativeExpressionAction = &_CAction{_CReduceAction, 0, _CReduceDToMultiplicativeExpression}
	_CReduceAToAdditiveExpressionAction       = &_CAction{_CReduceAction, 0, _CReduceAToAdditiveExpression}
	_CReduceBToAdditiveExpressionAction       = &_CAction{_CReduceAction, 0, _CReduceBToAdditiveExpression}
	_CReduceCToAdditiveExpressionAction       = &_CAction{_CReduceAction, 0, _CReduceCToAdditiveExpression}
	_CReduceAToShiftExpressionAction          = &_CAction{_CReduceAction, 0, _CReduceAToShiftExpression}
	_CReduceBToShiftExpressionAction          = &_CAction{_CReduceAction, 0, _CReduceBToShiftExpression}
	_CReduceCToShiftExpressionAction          = &_CAction{_CReduceAction, 0, _CReduceCToShiftExpression}
	_CReduceAToRelationalExpressionAction     = &_CAction{_CReduceAction, 0, _CReduceAToRelationalExpression}
	_CReduceBToRelationalExpressionAction     = &_CAction{_CReduceAction, 0, _CReduceBToRelationalExpression}
	_CReduceCToRelationalExpressionAction     = &_CAction{_CReduceAction, 0, _CReduceCToRelationalExpression}
	_CReduceDToRelationalExpressionAction     = &_CAction{_CReduceAction, 0, _CReduceDToRelationalExpression}
	_CReduceEToRelationalExpressionAction     = &_CAction{_CReduceAction, 0, _CReduceEToRelationalExpression}
	_CReduceAToEqualityExpressionAction       = &_CAction{_CReduceAction, 0, _CReduceAToEqualityExpression}
	_CReduceBToEqualityExpressionAction       = &_CAction{_CReduceAction, 0, _CReduceBToEqualityExpression}
	_CReduceCToEqualityExpressionAction       = &_CAction{_CReduceAction, 0, _CReduceCToEqualityExpression}
	_CReduceAToAndExpressionAction            = &_CAction{_CReduceAction, 0, _CReduceAToAndExpression}
	_CReduceBToAndExpressionAction            = &_CAction{_CReduceAction, 0, _CReduceBToAndExpression}
	_CReduceAToExclusiveOrExpressionAction    = &_CAction{_CReduceAction, 0, _CReduceAToExclusiveOrExpression}
	_CReduceBToExclusiveOrExpressionAction    = &_CAction{_CReduceAction, 0, _CReduceBToExclusiveOrExpression}
	_CReduceAToInclusiveOrExpressionAction    = &_CAction{_CReduceAction, 0, _CReduceAToInclusiveOrExpression}
	_CReduceBToInclusiveOrExpressionAction    = &_CAction{_CReduceAction, 0, _CReduceBToInclusiveOrExpression}
	_CReduceAToLogicalAndExpressionAction     = &_CAction{_CReduceAction, 0, _CReduceAToLogicalAndExpression}
	_CReduceBToLogicalAndExpressionAction     = &_CAction{_CReduceAction, 0, _CReduceBToLogicalAndExpression}
	_CReduceAToLogicalOrExpressionAction      = &_CAction{_CReduceAction, 0, _CReduceAToLogicalOrExpression}
	_CReduceBToLogicalOrExpressionAction      = &_CAction{_CReduceAction, 0, _CReduceBToLogicalOrExpression}
	_CReduceAToConditionalExpressionAction    = &_CAction{_CReduceAction, 0, _CReduceAToConditionalExpression}
	_CReduceBToConditionalExpressionAction    = &_CAction{_CReduceAction, 0, _CReduceBToConditionalExpression}
	_CReduceAToAssignmentExpressionAction     = &_CAction{_CReduceAction, 0, _CReduceAToAssignmentExpression}
	_CReduceBToAssignmentExpressionAction     = &_CAction{_CReduceAction, 0, _CReduceBToAssignmentExpression}
	_CReduceAToAssignmentOperatorAction       = &_CAction{_CReduceAction, 0, _CReduceAToAssignmentOperator}
	_CReduceBToAssignmentOperatorAction       = &_CAction{_CReduceAction, 0, _CReduceBToAssignmentOperator}
	_CReduceCToAssignmentOperatorAction       = &_CAction{_CReduceAction, 0, _CReduceCToAssignmentOperator}
	_CReduceDToAssignmentOperatorAction       = &_CAction{_CReduceAction, 0, _CReduceDToAssignmentOperator}
	_CReduceEToAssignmentOperatorAction       = &_CAction{_CReduceAction, 0, _CReduceEToAssignmentOperator}
	_CReduceFToAssignmentOperatorAction       = &_CAction{_CReduceAction, 0, _CReduceFToAssignmentOperator}
	_CReduceGToAssignmentOperatorAction       = &_CAction{_CReduceAction, 0, _CReduceGToAssignmentOperator}
	_CReduceHToAssignmentOperatorAction       = &_CAction{_CReduceAction, 0, _CReduceHToAssignmentOperator}
	_CReduceIToAssignmentOperatorAction       = &_CAction{_CReduceAction, 0, _CReduceIToAssignmentOperator}
	_CReduceJToAssignmentOperatorAction       = &_CAction{_CReduceAction, 0, _CReduceJToAssignmentOperator}
	_CReduceKToAssignmentOperatorAction       = &_CAction{_CReduceAction, 0, _CReduceKToAssignmentOperator}
	_CReduceAToExpressionAction               = &_CAction{_CReduceAction, 0, _CReduceAToExpression}
	_CReduceBToExpressionAction               = &_CAction{_CReduceAction, 0, _CReduceBToExpression}
	_CReduceAToConstantExpressionAction       = &_CAction{_CReduceAction, 0, _CReduceAToConstantExpression}
	_CReduceAToDeclarationAction              = &_CAction{_CReduceAction, 0, _CReduceAToDeclaration}
	_CReduceBToDeclarationAction              = &_CAction{_CReduceAction, 0, _CReduceBToDeclaration}
	_CReduceAToDeclarationSpecifiersAction    = &_CAction{_CReduceAction, 0, _CReduceAToDeclarationSpecifiers}
	_CReduceBToDeclarationSpecifiersAction    = &_CAction{_CReduceAction, 0, _CReduceBToDeclarationSpecifiers}
	_CReduceCToDeclarationSpecifiersAction    = &_CAction{_CReduceAction, 0, _CReduceCToDeclarationSpecifiers}
	_CReduceDToDeclarationSpecifiersAction    = &_CAction{_CReduceAction, 0, _CReduceDToDeclarationSpecifiers}
	_CReduceEToDeclarationSpecifiersAction    = &_CAction{_CReduceAction, 0, _CReduceEToDeclarationSpecifiers}
	_CReduceFToDeclarationSpecifiersAction    = &_CAction{_CReduceAction, 0, _CReduceFToDeclarationSpecifiers}
	_CReduceAToInitDeclaratorListAction       = &_CAction{_CReduceAction, 0, _CReduceAToInitDeclaratorList}
	_CReduceBToInitDeclaratorListAction       = &_CAction{_CReduceAction, 0, _CReduceBToInitDeclaratorList}
	_CReduceAToInitDeclaratorAction           = &_CAction{_CReduceAction, 0, _CReduceAToInitDeclarator}
	_CReduceBToInitDeclaratorAction           = &_CAction{_CReduceAction, 0, _CReduceBToInitDeclarator}
	_CReduceAToStorageClassSpecifierAction    = &_CAction{_CReduceAction, 0, _CReduceAToStorageClassSpecifier}
	_CReduceBToStorageClassSpecifierAction    = &_CAction{_CReduceAction, 0, _CReduceBToStorageClassSpecifier}
	_CReduceCToStorageClassSpecifierAction    = &_CAction{_CReduceAction, 0, _CReduceCToStorageClassSpecifier}
	_CReduceDToStorageClassSpecifierAction    = &_CAction{_CReduceAction, 0, _CReduceDToStorageClassSpecifier}
	_CReduceEToStorageClassSpecifierAction    = &_CAction{_CReduceAction, 0, _CReduceEToStorageClassSpecifier}
	_CReduceAToTypeSpecifierAction            = &_CAction{_CReduceAction, 0, _CReduceAToTypeSpecifier}
	_CReduceBToTypeSpecifierAction            = &_CAction{_CReduceAction, 0, _CReduceBToTypeSpecifier}
	_CReduceCToTypeSpecifierAction            = &_CAction{_CReduceAction, 0, _CReduceCToTypeSpecifier}
	_CReduceDToTypeSpecifierAction            = &_CAction{_CReduceAction, 0, _CReduceDToTypeSpecifier}
	_CReduceEToTypeSpecifierAction            = &_CAction{_CReduceAction, 0, _CReduceEToTypeSpecifier}
	_CReduceFToTypeSpecifierAction            = &_CAction{_CReduceAction, 0, _CReduceFToTypeSpecifier}
	_CReduceGToTypeSpecifierAction            = &_CAction{_CReduceAction, 0, _CReduceGToTypeSpecifier}
	_CReduceHToTypeSpecifierAction            = &_CAction{_CReduceAction, 0, _CReduceHToTypeSpecifier}
	_CReduceIToTypeSpecifierAction            = &_CAction{_CReduceAction, 0, _CReduceIToTypeSpecifier}
	_CReduceJToTypeSpecifierAction            = &_CAction{_CReduceAction, 0, _CReduceJToTypeSpecifier}
	_CReduceKToTypeSpecifierAction            = &_CAction{_CReduceAction, 0, _CReduceKToTypeSpecifier}
	_CReduceLToTypeSpecifierAction            = &_CAction{_CReduceAction, 0, _CReduceLToTypeSpecifier}
	_CReduceAToStructOrUnionSpecifierAction   = &_CAction{_CReduceAction, 0, _CReduceAToStructOrUnionSpecifier}
	_CReduceBToStructOrUnionSpecifierAction   = &_CAction{_CReduceAction, 0, _CReduceBToStructOrUnionSpecifier}
	_CReduceCToStructOrUnionSpecifierAction   = &_CAction{_CReduceAction, 0, _CReduceCToStructOrUnionSpecifier}
	_CReduceAToStructOrUnionAction            = &_CAction{_CReduceAction, 0, _CReduceAToStructOrUnion}
	_CReduceBToStructOrUnionAction            = &_CAction{_CReduceAction, 0, _CReduceBToStructOrUnion}
	_CReduceAToStructDeclarationListAction    = &_CAction{_CReduceAction, 0, _CReduceAToStructDeclarationList}
	_CReduceBToStructDeclarationListAction    = &_CAction{_CReduceAction, 0, _CReduceBToStructDeclarationList}
	_CReduceAToStructDeclarationAction        = &_CAction{_CReduceAction, 0, _CReduceAToStructDeclaration}
	_CReduceAToSpecifierQualifierListAction   = &_CAction{_CReduceAction, 0, _CReduceAToSpecifierQualifierList}
	_CReduceBToSpecifierQualifierListAction   = &_CAction{_CReduceAction, 0, _CReduceBToSpecifierQualifierList}
	_CReduceCToSpecifierQualifierListAction   = &_CAction{_CReduceAction, 0, _CReduceCToSpecifierQualifierList}
	_CReduceDToSpecifierQualifierListAction   = &_CAction{_CReduceAction, 0, _CReduceDToSpecifierQualifierList}
	_CReduceAToStructDeclaratorListAction     = &_CAction{_CReduceAction, 0, _CReduceAToStructDeclaratorList}
	_CReduceBToStructDeclaratorListAction     = &_CAction{_CReduceAction, 0, _CReduceBToStructDeclaratorList}
	_CReduceAToStructDeclaratorAction         = &_CAction{_CReduceAction, 0, _CReduceAToStructDeclarator}
	_CReduceBToStructDeclaratorAction         = &_CAction{_CReduceAction, 0, _CReduceBToStructDeclarator}
	_CReduceCToStructDeclaratorAction         = &_CAction{_CReduceAction, 0, _CReduceCToStructDeclarator}
	_CReduceAToEnumSpecifierAction            = &_CAction{_CReduceAction, 0, _CReduceAToEnumSpecifier}
	_CReduceBToEnumSpecifierAction            = &_CAction{_CReduceAction, 0, _CReduceBToEnumSpecifier}
	_CReduceCToEnumSpecifierAction            = &_CAction{_CReduceAction, 0, _CReduceCToEnumSpecifier}
	_CReduceAToEnumeratorListAction           = &_CAction{_CReduceAction, 0, _CReduceAToEnumeratorList}
	_CReduceBToEnumeratorListAction           = &_CAction{_CReduceAction, 0, _CReduceBToEnumeratorList}
	_CReduceAToEnumeratorAction               = &_CAction{_CReduceAction, 0, _CReduceAToEnumerator}
	_CReduceBToEnumeratorAction               = &_CAction{_CReduceAction, 0, _CReduceBToEnumerator}
	_CReduceAToTypeQualifierAction            = &_CAction{_CReduceAction, 0, _CReduceAToTypeQualifier}
	_CReduceBToTypeQualifierAction            = &_CAction{_CReduceAction, 0, _CReduceBToTypeQualifier}
	_CReduceAToDeclaratorAction               = &_CAction{_CReduceAction, 0, _CReduceAToDeclarator}
	_CReduceBToDeclaratorAction               = &_CAction{_CReduceAction, 0, _CReduceBToDeclarator}
	_CReduceAToDirectDeclaratorAction         = &_CAction{_CReduceAction, 0, _CReduceAToDirectDeclarator}
	_CReduceBToDirectDeclaratorAction         = &_CAction{_CReduceAction, 0, _CReduceBToDirectDeclarator}
	_CReduceCToDirectDeclaratorAction         = &_CAction{_CReduceAction, 0, _CReduceCToDirectDeclarator}
	_CReduceDToDirectDeclaratorAction         = &_CAction{_CReduceAction, 0, _CReduceDToDirectDeclarator}
	_CReduceEToDirectDeclaratorAction         = &_CAction{_CReduceAction, 0, _CReduceEToDirectDeclarator}
	_CReduceFToDirectDeclaratorAction         = &_CAction{_CReduceAction, 0, _CReduceFToDirectDeclarator}
	_CReduceGToDirectDeclaratorAction         = &_CAction{_CReduceAction, 0, _CReduceGToDirectDeclarator}
	_CReduceAToPointerAction                  = &_CAction{_CReduceAction, 0, _CReduceAToPointer}
	_CReduceBToPointerAction                  = &_CAction{_CReduceAction, 0, _CReduceBToPointer}
	_CReduceCToPointerAction                  = &_CAction{_CReduceAction, 0, _CReduceCToPointer}
	_CReduceDToPointerAction                  = &_CAction{_CReduceAction, 0, _CReduceDToPointer}
	_CReduceAToTypeQualifierListAction        = &_CAction{_CReduceAction, 0, _CReduceAToTypeQualifierList}
	_CReduceBToTypeQualifierListAction        = &_CAction{_CReduceAction, 0, _CReduceBToTypeQualifierList}
	_CReduceAToParameterTypeListAction        = &_CAction{_CReduceAction, 0, _CReduceAToParameterTypeList}
	_CReduceBToParameterTypeListAction        = &_CAction{_CReduceAction, 0, _CReduceBToParameterTypeList}
	_CReduceAToParameterListAction            = &_CAction{_CReduceAction, 0, _CReduceAToParameterList}
	_CReduceBToParameterListAction            = &_CAction{_CReduceAction, 0, _CReduceBToParameterList}
	_CReduceAToParameterDeclarationAction     = &_CAction{_CReduceAction, 0, _CReduceAToParameterDeclaration}
	_CReduceBToParameterDeclarationAction     = &_CAction{_CReduceAction, 0, _CReduceBToParameterDeclaration}
	_CReduceCToParameterDeclarationAction     = &_CAction{_CReduceAction, 0, _CReduceCToParameterDeclaration}
	_CReduceAToIdentifierListAction           = &_CAction{_CReduceAction, 0, _CReduceAToIdentifierList}
	_CReduceBToIdentifierListAction           = &_CAction{_CReduceAction, 0, _CReduceBToIdentifierList}
	_CReduceAToTypeNameAction                 = &_CAction{_CReduceAction, 0, _CReduceAToTypeName}
	_CReduceBToTypeNameAction                 = &_CAction{_CReduceAction, 0, _CReduceBToTypeName}
	_CReduceAToAbstractDeclaratorAction       = &_CAction{_CReduceAction, 0, _CReduceAToAbstractDeclarator}
	_CReduceBToAbstractDeclaratorAction       = &_CAction{_CReduceAction, 0, _CReduceBToAbstractDeclarator}
	_CReduceCToAbstractDeclaratorAction       = &_CAction{_CReduceAction, 0, _CReduceCToAbstractDeclarator}
	_CReduceAToDirectAbstractDeclaratorAction = &_CAction{_CReduceAction, 0, _CReduceAToDirectAbstractDeclarator}
	_CReduceBToDirectAbstractDeclaratorAction = &_CAction{_CReduceAction, 0, _CReduceBToDirectAbstractDeclarator}
	_CReduceCToDirectAbstractDeclaratorAction = &_CAction{_CReduceAction, 0, _CReduceCToDirectAbstractDeclarator}
	_CReduceDToDirectAbstractDeclaratorAction = &_CAction{_CReduceAction, 0, _CReduceDToDirectAbstractDeclarator}
	_CReduceEToDirectAbstractDeclaratorAction = &_CAction{_CReduceAction, 0, _CReduceEToDirectAbstractDeclarator}
	_CReduceFToDirectAbstractDeclaratorAction = &_CAction{_CReduceAction, 0, _CReduceFToDirectAbstractDeclarator}
	_CReduceGToDirectAbstractDeclaratorAction = &_CAction{_CReduceAction, 0, _CReduceGToDirectAbstractDeclarator}
	_CReduceHToDirectAbstractDeclaratorAction = &_CAction{_CReduceAction, 0, _CReduceHToDirectAbstractDeclarator}
	_CReduceIToDirectAbstractDeclaratorAction = &_CAction{_CReduceAction, 0, _CReduceIToDirectAbstractDeclarator}
	_CReduceAToInitializerAction              = &_CAction{_CReduceAction, 0, _CReduceAToInitializer}
	_CReduceBToInitializerAction              = &_CAction{_CReduceAction, 0, _CReduceBToInitializer}
	_CReduceCToInitializerAction              = &_CAction{_CReduceAction, 0, _CReduceCToInitializer}
	_CReduceAToInitializerListAction          = &_CAction{_CReduceAction, 0, _CReduceAToInitializerList}
	_CReduceBToInitializerListAction          = &_CAction{_CReduceAction, 0, _CReduceBToInitializerList}
	_CReduceAToStatementAction                = &_CAction{_CReduceAction, 0, _CReduceAToStatement}
	_CReduceBToStatementAction                = &_CAction{_CReduceAction, 0, _CReduceBToStatement}
	_CReduceCToStatementAction                = &_CAction{_CReduceAction, 0, _CReduceCToStatement}
	_CReduceDToStatementAction                = &_CAction{_CReduceAction, 0, _CReduceDToStatement}
	_CReduceEToStatementAction                = &_CAction{_CReduceAction, 0, _CReduceEToStatement}
	_CReduceFToStatementAction                = &_CAction{_CReduceAction, 0, _CReduceFToStatement}
	_CReduceAToLabeledStatementAction         = &_CAction{_CReduceAction, 0, _CReduceAToLabeledStatement}
	_CReduceBToLabeledStatementAction         = &_CAction{_CReduceAction, 0, _CReduceBToLabeledStatement}
	_CReduceCToLabeledStatementAction         = &_CAction{_CReduceAction, 0, _CReduceCToLabeledStatement}
	_CReduceAToCompoundStatementAction        = &_CAction{_CReduceAction, 0, _CReduceAToCompoundStatement}
	_CReduceBToCompoundStatementAction        = &_CAction{_CReduceAction, 0, _CReduceBToCompoundStatement}
	_CReduceCToCompoundStatementAction        = &_CAction{_CReduceAction, 0, _CReduceCToCompoundStatement}
	_CReduceDToCompoundStatementAction        = &_CAction{_CReduceAction, 0, _CReduceDToCompoundStatement}
	_CReduceAToDeclarationListAction          = &_CAction{_CReduceAction, 0, _CReduceAToDeclarationList}
	_CReduceBToDeclarationListAction          = &_CAction{_CReduceAction, 0, _CReduceBToDeclarationList}
	_CReduceAToStatementListAction            = &_CAction{_CReduceAction, 0, _CReduceAToStatementList}
	_CReduceBToStatementListAction            = &_CAction{_CReduceAction, 0, _CReduceBToStatementList}
	_CReduceAToExpressionStatementAction      = &_CAction{_CReduceAction, 0, _CReduceAToExpressionStatement}
	_CReduceBToExpressionStatementAction      = &_CAction{_CReduceAction, 0, _CReduceBToExpressionStatement}
	_CReduceAToSelectionStatementAction       = &_CAction{_CReduceAction, 0, _CReduceAToSelectionStatement}
	_CReduceBToSelectionStatementAction       = &_CAction{_CReduceAction, 0, _CReduceBToSelectionStatement}
	_CReduceCToSelectionStatementAction       = &_CAction{_CReduceAction, 0, _CReduceCToSelectionStatement}
	_CReduceAToIterationStatementAction       = &_CAction{_CReduceAction, 0, _CReduceAToIterationStatement}
	_CReduceBToIterationStatementAction       = &_CAction{_CReduceAction, 0, _CReduceBToIterationStatement}
	_CReduceCToIterationStatementAction       = &_CAction{_CReduceAction, 0, _CReduceCToIterationStatement}
	_CReduceDToIterationStatementAction       = &_CAction{_CReduceAction, 0, _CReduceDToIterationStatement}
	_CReduceAToJumpStatementAction            = &_CAction{_CReduceAction, 0, _CReduceAToJumpStatement}
	_CReduceBToJumpStatementAction            = &_CAction{_CReduceAction, 0, _CReduceBToJumpStatement}
	_CReduceCToJumpStatementAction            = &_CAction{_CReduceAction, 0, _CReduceCToJumpStatement}
	_CReduceDToJumpStatementAction            = &_CAction{_CReduceAction, 0, _CReduceDToJumpStatement}
	_CReduceEToJumpStatementAction            = &_CAction{_CReduceAction, 0, _CReduceEToJumpStatement}
	_CReduceAToTranslationUnitAction          = &_CAction{_CReduceAction, 0, _CReduceAToTranslationUnit}
	_CReduceBToTranslationUnitAction          = &_CAction{_CReduceAction, 0, _CReduceBToTranslationUnit}
	_CReduceAToExternalDeclarationAction      = &_CAction{_CReduceAction, 0, _CReduceAToExternalDeclaration}
	_CReduceBToExternalDeclarationAction      = &_CAction{_CReduceAction, 0, _CReduceBToExternalDeclaration}
	_CReduceAToFunctionDefinitionAction       = &_CAction{_CReduceAction, 0, _CReduceAToFunctionDefinition}
	_CReduceBToFunctionDefinitionAction       = &_CAction{_CReduceAction, 0, _CReduceBToFunctionDefinition}
	_CReduceCToFunctionDefinitionAction       = &_CAction{_CReduceAction, 0, _CReduceCToFunctionDefinition}
	_CReduceDToFunctionDefinitionAction       = &_CAction{_CReduceAction, 0, _CReduceDToFunctionDefinition}
)

var _CActionTable = _CActionTableType{
	{_CState2, _CEndMarker}:                     &_CAction{_CAcceptAction, 0, 0},
	{_CState1, CIdentifierToken}:                _CGotoState12Action,
	{_CState1, CTypeNameToken}:                  _CGotoState21Action,
	{_CState1, CTypedefToken}:                   _CGotoState20Action,
	{_CState1, CExternToken}:                    _CGotoState10Action,
	{_CState1, CStaticToken}:                    _CGotoState18Action,
	{_CState1, CAutoToken}:                      _CGotoState5Action,
	{_CState1, CRegisterToken}:                  _CGotoState15Action,
	{_CState1, CCharToken}:                      _CGotoState6Action,
	{_CState1, CShortToken}:                     _CGotoState16Action,
	{_CState1, CIntToken}:                       _CGotoState13Action,
	{_CState1, CLongToken}:                      _CGotoState14Action,
	{_CState1, CSignedToken}:                    _CGotoState17Action,
	{_CState1, CUnsignedToken}:                  _CGotoState23Action,
	{_CState1, CFloatToken}:                     _CGotoState11Action,
	{_CState1, CDoubleToken}:                    _CGotoState8Action,
	{_CState1, CConstToken}:                     _CGotoState7Action,
	{_CState1, CVolatileToken}:                  _CGotoState25Action,
	{_CState1, CVoidToken}:                      _CGotoState24Action,
	{_CState1, CStructToken}:                    _CGotoState19Action,
	{_CState1, CUnionToken}:                     _CGotoState22Action,
	{_CState1, CEnumToken}:                      _CGotoState9Action,
	{_CState1, '('}:                             _CGotoState3Action,
	{_CState1, '*'}:                             _CGotoState4Action,
	{_CState1, CDeclarationType}:                _CGotoState26Action,
	{_CState1, CDeclarationSpecifiersType}:      _CGotoState27Action,
	{_CState1, CStorageClassSpecifierType}:      _CGotoState34Action,
	{_CState1, CTypeSpecifierType}:              _CGotoState38Action,
	{_CState1, CStructOrUnionSpecifierType}:     _CGotoState36Action,
	{_CState1, CStructOrUnionType}:              _CGotoState35Action,
	{_CState1, CEnumSpecifierType}:              _CGotoState30Action,
	{_CState1, CTypeQualifierType}:              _CGotoState37Action,
	{_CState1, CDeclaratorType}:                 _CGotoState28Action,
	{_CState1, CDirectDeclaratorType}:           _CGotoState29Action,
	{_CState1, CPointerType}:                    _CGotoState33Action,
	{_CState1, CTranslationUnitType}:            _CGotoState2Action,
	{_CState1, CExternalDeclarationType}:        _CGotoState31Action,
	{_CState1, CFunctionDefinitionType}:         _CGotoState32Action,
	{_CState2, CIdentifierToken}:                _CGotoState12Action,
	{_CState2, CTypeNameToken}:                  _CGotoState21Action,
	{_CState2, CTypedefToken}:                   _CGotoState20Action,
	{_CState2, CExternToken}:                    _CGotoState10Action,
	{_CState2, CStaticToken}:                    _CGotoState18Action,
	{_CState2, CAutoToken}:                      _CGotoState5Action,
	{_CState2, CRegisterToken}:                  _CGotoState15Action,
	{_CState2, CCharToken}:                      _CGotoState6Action,
	{_CState2, CShortToken}:                     _CGotoState16Action,
	{_CState2, CIntToken}:                       _CGotoState13Action,
	{_CState2, CLongToken}:                      _CGotoState14Action,
	{_CState2, CSignedToken}:                    _CGotoState17Action,
	{_CState2, CUnsignedToken}:                  _CGotoState23Action,
	{_CState2, CFloatToken}:                     _CGotoState11Action,
	{_CState2, CDoubleToken}:                    _CGotoState8Action,
	{_CState2, CConstToken}:                     _CGotoState7Action,
	{_CState2, CVolatileToken}:                  _CGotoState25Action,
	{_CState2, CVoidToken}:                      _CGotoState24Action,
	{_CState2, CStructToken}:                    _CGotoState19Action,
	{_CState2, CUnionToken}:                     _CGotoState22Action,
	{_CState2, CEnumToken}:                      _CGotoState9Action,
	{_CState2, '('}:                             _CGotoState3Action,
	{_CState2, '*'}:                             _CGotoState4Action,
	{_CState2, CDeclarationType}:                _CGotoState26Action,
	{_CState2, CDeclarationSpecifiersType}:      _CGotoState27Action,
	{_CState2, CStorageClassSpecifierType}:      _CGotoState34Action,
	{_CState2, CTypeSpecifierType}:              _CGotoState38Action,
	{_CState2, CStructOrUnionSpecifierType}:     _CGotoState36Action,
	{_CState2, CStructOrUnionType}:              _CGotoState35Action,
	{_CState2, CEnumSpecifierType}:              _CGotoState30Action,
	{_CState2, CTypeQualifierType}:              _CGotoState37Action,
	{_CState2, CDeclaratorType}:                 _CGotoState28Action,
	{_CState2, CDirectDeclaratorType}:           _CGotoState29Action,
	{_CState2, CPointerType}:                    _CGotoState33Action,
	{_CState2, CExternalDeclarationType}:        _CGotoState60Action,
	{_CState2, CFunctionDefinitionType}:         _CGotoState32Action,
	{_CState3, CIdentifierToken}:                _CGotoState12Action,
	{_CState3, '('}:                             _CGotoState3Action,
	{_CState3, '*'}:                             _CGotoState4Action,
	{_CState3, CDeclaratorType}:                 _CGotoState39Action,
	{_CState3, CDirectDeclaratorType}:           _CGotoState29Action,
	{_CState3, CPointerType}:                    _CGotoState33Action,
	{_CState4, CConstToken}:                     _CGotoState7Action,
	{_CState4, CVolatileToken}:                  _CGotoState25Action,
	{_CState4, '*'}:                             _CGotoState4Action,
	{_CState4, CTypeQualifierType}:              _CGotoState41Action,
	{_CState4, CPointerType}:                    _CGotoState40Action,
	{_CState4, CTypeQualifierListType}:          _CGotoState42Action,
	{_CState9, CIdentifierToken}:                _CGotoState44Action,
	{_CState9, '{'}:                             _CGotoState43Action,
	{_CState27, CIdentifierToken}:               _CGotoState12Action,
	{_CState27, '('}:                            _CGotoState3Action,
	{_CState27, ';'}:                            _CGotoState45Action,
	{_CState27, '*'}:                            _CGotoState4Action,
	{_CState27, CInitDeclaratorListType}:        _CGotoState48Action,
	{_CState27, CInitDeclaratorType}:            _CGotoState47Action,
	{_CState27, CDeclaratorType}:                _CGotoState46Action,
	{_CState27, CDirectDeclaratorType}:          _CGotoState29Action,
	{_CState27, CPointerType}:                   _CGotoState33Action,
	{_CState28, CTypeNameToken}:                 _CGotoState21Action,
	{_CState28, CTypedefToken}:                  _CGotoState20Action,
	{_CState28, CExternToken}:                   _CGotoState10Action,
	{_CState28, CStaticToken}:                   _CGotoState18Action,
	{_CState28, CAutoToken}:                     _CGotoState5Action,
	{_CState28, CRegisterToken}:                 _CGotoState15Action,
	{_CState28, CCharToken}:                     _CGotoState6Action,
	{_CState28, CShortToken}:                    _CGotoState16Action,
	{_CState28, CIntToken}:                      _CGotoState13Action,
	{_CState28, CLongToken}:                     _CGotoState14Action,
	{_CState28, CSignedToken}:                   _CGotoState17Action,
	{_CState28, CUnsignedToken}:                 _CGotoState23Action,
	{_CState28, CFloatToken}:                    _CGotoState11Action,
	{_CState28, CDoubleToken}:                   _CGotoState8Action,
	{_CState28, CConstToken}:                    _CGotoState7Action,
	{_CState28, CVolatileToken}:                 _CGotoState25Action,
	{_CState28, CVoidToken}:                     _CGotoState24Action,
	{_CState28, CStructToken}:                   _CGotoState19Action,
	{_CState28, CUnionToken}:                    _CGotoState22Action,
	{_CState28, CEnumToken}:                     _CGotoState9Action,
	{_CState28, '{'}:                            _CGotoState49Action,
	{_CState28, CDeclarationType}:               _CGotoState51Action,
	{_CState28, CDeclarationSpecifiersType}:     _CGotoState53Action,
	{_CState28, CStorageClassSpecifierType}:     _CGotoState34Action,
	{_CState28, CTypeSpecifierType}:             _CGotoState38Action,
	{_CState28, CStructOrUnionSpecifierType}:    _CGotoState36Action,
	{_CState28, CStructOrUnionType}:             _CGotoState35Action,
	{_CState28, CEnumSpecifierType}:             _CGotoState30Action,
	{_CState28, CTypeQualifierType}:             _CGotoState37Action,
	{_CState28, CCompoundStatementType}:         _CGotoState50Action,
	{_CState28, CDeclarationListType}:           _CGotoState52Action,
	{_CState29, '('}:                            _CGotoState54Action,
	{_CState29, '['}:                            _CGotoState55Action,
	{_CState33, CIdentifierToken}:               _CGotoState12Action,
	{_CState33, '('}:                            _CGotoState3Action,
	{_CState33, CDirectDeclaratorType}:          _CGotoState56Action,
	{_CState34, CTypeNameToken}:                 _CGotoState21Action,
	{_CState34, CTypedefToken}:                  _CGotoState20Action,
	{_CState34, CExternToken}:                   _CGotoState10Action,
	{_CState34, CStaticToken}:                   _CGotoState18Action,
	{_CState34, CAutoToken}:                     _CGotoState5Action,
	{_CState34, CRegisterToken}:                 _CGotoState15Action,
	{_CState34, CCharToken}:                     _CGotoState6Action,
	{_CState34, CShortToken}:                    _CGotoState16Action,
	{_CState34, CIntToken}:                      _CGotoState13Action,
	{_CState34, CLongToken}:                     _CGotoState14Action,
	{_CState34, CSignedToken}:                   _CGotoState17Action,
	{_CState34, CUnsignedToken}:                 _CGotoState23Action,
	{_CState34, CFloatToken}:                    _CGotoState11Action,
	{_CState34, CDoubleToken}:                   _CGotoState8Action,
	{_CState34, CConstToken}:                    _CGotoState7Action,
	{_CState34, CVolatileToken}:                 _CGotoState25Action,
	{_CState34, CVoidToken}:                     _CGotoState24Action,
	{_CState34, CStructToken}:                   _CGotoState19Action,
	{_CState34, CUnionToken}:                    _CGotoState22Action,
	{_CState34, CEnumToken}:                     _CGotoState9Action,
	{_CState34, CDeclarationSpecifiersType}:     _CGotoState57Action,
	{_CState34, CStorageClassSpecifierType}:     _CGotoState34Action,
	{_CState34, CTypeSpecifierType}:             _CGotoState38Action,
	{_CState34, CStructOrUnionSpecifierType}:    _CGotoState36Action,
	{_CState34, CStructOrUnionType}:             _CGotoState35Action,
	{_CState34, CEnumSpecifierType}:             _CGotoState30Action,
	{_CState34, CTypeQualifierType}:             _CGotoState37Action,
	{_CState35, CIdentifierToken}:               _CGotoState59Action,
	{_CState35, '{'}:                            _CGotoState58Action,
	{_CState37, CTypeNameToken}:                 _CGotoState21Action,
	{_CState37, CTypedefToken}:                  _CGotoState20Action,
	{_CState37, CExternToken}:                   _CGotoState10Action,
	{_CState37, CStaticToken}:                   _CGotoState18Action,
	{_CState37, CAutoToken}:                     _CGotoState5Action,
	{_CState37, CRegisterToken}:                 _CGotoState15Action,
	{_CState37, CCharToken}:                     _CGotoState6Action,
	{_CState37, CShortToken}:                    _CGotoState16Action,
	{_CState37, CIntToken}:                      _CGotoState13Action,
	{_CState37, CLongToken}:                     _CGotoState14Action,
	{_CState37, CSignedToken}:                   _CGotoState17Action,
	{_CState37, CUnsignedToken}:                 _CGotoState23Action,
	{_CState37, CFloatToken}:                    _CGotoState11Action,
	{_CState37, CDoubleToken}:                   _CGotoState8Action,
	{_CState37, CConstToken}:                    _CGotoState7Action,
	{_CState37, CVolatileToken}:                 _CGotoState25Action,
	{_CState37, CVoidToken}:                     _CGotoState24Action,
	{_CState37, CStructToken}:                   _CGotoState19Action,
	{_CState37, CUnionToken}:                    _CGotoState22Action,
	{_CState37, CEnumToken}:                     _CGotoState9Action,
	{_CState37, CDeclarationSpecifiersType}:     _CGotoState61Action,
	{_CState37, CStorageClassSpecifierType}:     _CGotoState34Action,
	{_CState37, CTypeSpecifierType}:             _CGotoState38Action,
	{_CState37, CStructOrUnionSpecifierType}:    _CGotoState36Action,
	{_CState37, CStructOrUnionType}:             _CGotoState35Action,
	{_CState37, CEnumSpecifierType}:             _CGotoState30Action,
	{_CState37, CTypeQualifierType}:             _CGotoState37Action,
	{_CState38, CTypeNameToken}:                 _CGotoState21Action,
	{_CState38, CTypedefToken}:                  _CGotoState20Action,
	{_CState38, CExternToken}:                   _CGotoState10Action,
	{_CState38, CStaticToken}:                   _CGotoState18Action,
	{_CState38, CAutoToken}:                     _CGotoState5Action,
	{_CState38, CRegisterToken}:                 _CGotoState15Action,
	{_CState38, CCharToken}:                     _CGotoState6Action,
	{_CState38, CShortToken}:                    _CGotoState16Action,
	{_CState38, CIntToken}:                      _CGotoState13Action,
	{_CState38, CLongToken}:                     _CGotoState14Action,
	{_CState38, CSignedToken}:                   _CGotoState17Action,
	{_CState38, CUnsignedToken}:                 _CGotoState23Action,
	{_CState38, CFloatToken}:                    _CGotoState11Action,
	{_CState38, CDoubleToken}:                   _CGotoState8Action,
	{_CState38, CConstToken}:                    _CGotoState7Action,
	{_CState38, CVolatileToken}:                 _CGotoState25Action,
	{_CState38, CVoidToken}:                     _CGotoState24Action,
	{_CState38, CStructToken}:                   _CGotoState19Action,
	{_CState38, CUnionToken}:                    _CGotoState22Action,
	{_CState38, CEnumToken}:                     _CGotoState9Action,
	{_CState38, CDeclarationSpecifiersType}:     _CGotoState62Action,
	{_CState38, CStorageClassSpecifierType}:     _CGotoState34Action,
	{_CState38, CTypeSpecifierType}:             _CGotoState38Action,
	{_CState38, CStructOrUnionSpecifierType}:    _CGotoState36Action,
	{_CState38, CStructOrUnionType}:             _CGotoState35Action,
	{_CState38, CEnumSpecifierType}:             _CGotoState30Action,
	{_CState38, CTypeQualifierType}:             _CGotoState37Action,
	{_CState39, ')'}:                            _CGotoState63Action,
	{_CState42, CConstToken}:                    _CGotoState7Action,
	{_CState42, CVolatileToken}:                 _CGotoState25Action,
	{_CState42, '*'}:                            _CGotoState4Action,
	{_CState42, CTypeQualifierType}:             _CGotoState65Action,
	{_CState42, CPointerType}:                   _CGotoState64Action,
	{_CState43, CIdentifierToken}:               _CGotoState66Action,
	{_CState43, CEnumeratorListType}:            _CGotoState68Action,
	{_CState43, CEnumeratorType}:                _CGotoState67Action,
	{_CState44, '{'}:                            _CGotoState69Action,
	{_CState46, CTypeNameToken}:                 _CGotoState21Action,
	{_CState46, CTypedefToken}:                  _CGotoState20Action,
	{_CState46, CExternToken}:                   _CGotoState10Action,
	{_CState46, CStaticToken}:                   _CGotoState18Action,
	{_CState46, CAutoToken}:                     _CGotoState5Action,
	{_CState46, CRegisterToken}:                 _CGotoState15Action,
	{_CState46, CCharToken}:                     _CGotoState6Action,
	{_CState46, CShortToken}:                    _CGotoState16Action,
	{_CState46, CIntToken}:                      _CGotoState13Action,
	{_CState46, CLongToken}:                     _CGotoState14Action,
	{_CState46, CSignedToken}:                   _CGotoState17Action,
	{_CState46, CUnsignedToken}:                 _CGotoState23Action,
	{_CState46, CFloatToken}:                    _CGotoState11Action,
	{_CState46, CDoubleToken}:                   _CGotoState8Action,
	{_CState46, CConstToken}:                    _CGotoState7Action,
	{_CState46, CVolatileToken}:                 _CGotoState25Action,
	{_CState46, CVoidToken}:                     _CGotoState24Action,
	{_CState46, CStructToken}:                   _CGotoState19Action,
	{_CState46, CUnionToken}:                    _CGotoState22Action,
	{_CState46, CEnumToken}:                     _CGotoState9Action,
	{_CState46, '{'}:                            _CGotoState49Action,
	{_CState46, '='}:                            _CGotoState70Action,
	{_CState46, CDeclarationType}:               _CGotoState51Action,
	{_CState46, CDeclarationSpecifiersType}:     _CGotoState53Action,
	{_CState46, CStorageClassSpecifierType}:     _CGotoState34Action,
	{_CState46, CTypeSpecifierType}:             _CGotoState38Action,
	{_CState46, CStructOrUnionSpecifierType}:    _CGotoState36Action,
	{_CState46, CStructOrUnionType}:             _CGotoState35Action,
	{_CState46, CEnumSpecifierType}:             _CGotoState30Action,
	{_CState46, CTypeQualifierType}:             _CGotoState37Action,
	{_CState46, CCompoundStatementType}:         _CGotoState71Action,
	{_CState46, CDeclarationListType}:           _CGotoState72Action,
	{_CState48, ';'}:                            _CGotoState74Action,
	{_CState48, ','}:                            _CGotoState73Action,
	{_CState49, CIdentifierToken}:               _CGotoState93Action,
	{_CState49, CConstantToken}:                 _CGotoState86Action,
	{_CState49, CStringLiteralToken}:            _CGotoState98Action,
	{_CState49, CSizeofToken}:                   _CGotoState97Action,
	{_CState49, CIncOpToken}:                    _CGotoState95Action,
	{_CState49, CDecOpToken}:                    _CGotoState88Action,
	{_CState49, CTypeNameToken}:                 _CGotoState21Action,
	{_CState49, CTypedefToken}:                  _CGotoState20Action,
	{_CState49, CExternToken}:                   _CGotoState10Action,
	{_CState49, CStaticToken}:                   _CGotoState18Action,
	{_CState49, CAutoToken}:                     _CGotoState5Action,
	{_CState49, CRegisterToken}:                 _CGotoState15Action,
	{_CState49, CCharToken}:                     _CGotoState6Action,
	{_CState49, CShortToken}:                    _CGotoState16Action,
	{_CState49, CIntToken}:                      _CGotoState13Action,
	{_CState49, CLongToken}:                     _CGotoState14Action,
	{_CState49, CSignedToken}:                   _CGotoState17Action,
	{_CState49, CUnsignedToken}:                 _CGotoState23Action,
	{_CState49, CFloatToken}:                    _CGotoState11Action,
	{_CState49, CDoubleToken}:                   _CGotoState8Action,
	{_CState49, CConstToken}:                    _CGotoState7Action,
	{_CState49, CVolatileToken}:                 _CGotoState25Action,
	{_CState49, CVoidToken}:                     _CGotoState24Action,
	{_CState49, CStructToken}:                   _CGotoState19Action,
	{_CState49, CUnionToken}:                    _CGotoState22Action,
	{_CState49, CEnumToken}:                     _CGotoState9Action,
	{_CState49, CCaseToken}:                     _CGotoState85Action,
	{_CState49, CDefaultToken}:                  _CGotoState89Action,
	{_CState49, CIfToken}:                       _CGotoState94Action,
	{_CState49, CSwitchToken}:                   _CGotoState99Action,
	{_CState49, CWhileToken}:                    _CGotoState100Action,
	{_CState49, CDoToken}:                       _CGotoState90Action,
	{_CState49, CForToken}:                      _CGotoState91Action,
	{_CState49, CGotoToken}:                     _CGotoState92Action,
	{_CState49, CContinueToken}:                 _CGotoState87Action,
	{_CState49, CBreakToken}:                    _CGotoState84Action,
	{_CState49, CReturnToken}:                   _CGotoState96Action,
	{_CState49, '('}:                            _CGotoState77Action,
	{_CState49, '{'}:                            _CGotoState49Action,
	{_CState49, '}'}:                            _CGotoState82Action,
	{_CState49, ';'}:                            _CGotoState81Action,
	{_CState49, '*'}:                            _CGotoState78Action,
	{_CState49, '-'}:                            _CGotoState80Action,
	{_CState49, '+'}:                            _CGotoState79Action,
	{_CState49, '&'}:                            _CGotoState76Action,
	{_CState49, '!'}:                            _CGotoState75Action,
	{_CState49, '~'}:                            _CGotoState83Action,
	{_CState49, CPrimaryExpressionType}:         _CGotoState120Action,
	{_CState49, CPostfixExpressionType}:         _CGotoState119Action,
	{_CState49, CUnaryExpressionType}:           _CGotoState126Action,
	{_CState49, CUnaryOperatorType}:             _CGotoState127Action,
	{_CState49, CCastExpressionType}:            _CGotoState104Action,
	{_CState49, CMultiplicativeExpressionType}:  _CGotoState118Action,
	{_CState49, CAdditiveExpressionType}:        _CGotoState101Action,
	{_CState49, CShiftExpressionType}:           _CGotoState123Action,
	{_CState49, CRelationalExpressionType}:      _CGotoState121Action,
	{_CState49, CEqualityExpressionType}:        _CGotoState108Action,
	{_CState49, CAndExpressionType}:             _CGotoState102Action,
	{_CState49, CExclusiveOrExpressionType}:     _CGotoState109Action,
	{_CState49, CInclusiveOrExpressionType}:     _CGotoState112Action,
	{_CState49, CLogicalAndExpressionType}:      _CGotoState116Action,
	{_CState49, CLogicalOrExpressionType}:       _CGotoState117Action,
	{_CState49, CConditionalExpressionType}:     _CGotoState106Action,
	{_CState49, CAssignmentExpressionType}:      _CGotoState103Action,
	{_CState49, CExpressionType}:                _CGotoState110Action,
	{_CState49, CDeclarationType}:               _CGotoState51Action,
	{_CState49, CDeclarationSpecifiersType}:     _CGotoState53Action,
	{_CState49, CStorageClassSpecifierType}:     _CGotoState34Action,
	{_CState49, CTypeSpecifierType}:             _CGotoState38Action,
	{_CState49, CStructOrUnionSpecifierType}:    _CGotoState36Action,
	{_CState49, CStructOrUnionType}:             _CGotoState35Action,
	{_CState49, CEnumSpecifierType}:             _CGotoState30Action,
	{_CState49, CTypeQualifierType}:             _CGotoState37Action,
	{_CState49, CStatementType}:                 _CGotoState124Action,
	{_CState49, CLabeledStatementType}:          _CGotoState115Action,
	{_CState49, CCompoundStatementType}:         _CGotoState105Action,
	{_CState49, CDeclarationListType}:           _CGotoState107Action,
	{_CState49, CStatementListType}:             _CGotoState125Action,
	{_CState49, CExpressionStatementType}:       _CGotoState111Action,
	{_CState49, CSelectionStatementType}:        _CGotoState122Action,
	{_CState49, CIterationStatementType}:        _CGotoState113Action,
	{_CState49, CJumpStatementType}:             _CGotoState114Action,
	{_CState52, CTypeNameToken}:                 _CGotoState21Action,
	{_CState52, CTypedefToken}:                  _CGotoState20Action,
	{_CState52, CExternToken}:                   _CGotoState10Action,
	{_CState52, CStaticToken}:                   _CGotoState18Action,
	{_CState52, CAutoToken}:                     _CGotoState5Action,
	{_CState52, CRegisterToken}:                 _CGotoState15Action,
	{_CState52, CCharToken}:                     _CGotoState6Action,
	{_CState52, CShortToken}:                    _CGotoState16Action,
	{_CState52, CIntToken}:                      _CGotoState13Action,
	{_CState52, CLongToken}:                     _CGotoState14Action,
	{_CState52, CSignedToken}:                   _CGotoState17Action,
	{_CState52, CUnsignedToken}:                 _CGotoState23Action,
	{_CState52, CFloatToken}:                    _CGotoState11Action,
	{_CState52, CDoubleToken}:                   _CGotoState8Action,
	{_CState52, CConstToken}:                    _CGotoState7Action,
	{_CState52, CVolatileToken}:                 _CGotoState25Action,
	{_CState52, CVoidToken}:                     _CGotoState24Action,
	{_CState52, CStructToken}:                   _CGotoState19Action,
	{_CState52, CUnionToken}:                    _CGotoState22Action,
	{_CState52, CEnumToken}:                     _CGotoState9Action,
	{_CState52, '{'}:                            _CGotoState49Action,
	{_CState52, CDeclarationType}:               _CGotoState129Action,
	{_CState52, CDeclarationSpecifiersType}:     _CGotoState53Action,
	{_CState52, CStorageClassSpecifierType}:     _CGotoState34Action,
	{_CState52, CTypeSpecifierType}:             _CGotoState38Action,
	{_CState52, CStructOrUnionSpecifierType}:    _CGotoState36Action,
	{_CState52, CStructOrUnionType}:             _CGotoState35Action,
	{_CState52, CEnumSpecifierType}:             _CGotoState30Action,
	{_CState52, CTypeQualifierType}:             _CGotoState37Action,
	{_CState52, CCompoundStatementType}:         _CGotoState128Action,
	{_CState53, CIdentifierToken}:               _CGotoState12Action,
	{_CState53, '('}:                            _CGotoState3Action,
	{_CState53, ';'}:                            _CGotoState45Action,
	{_CState53, '*'}:                            _CGotoState4Action,
	{_CState53, CInitDeclaratorListType}:        _CGotoState48Action,
	{_CState53, CInitDeclaratorType}:            _CGotoState47Action,
	{_CState53, CDeclaratorType}:                _CGotoState130Action,
	{_CState53, CDirectDeclaratorType}:          _CGotoState29Action,
	{_CState53, CPointerType}:                   _CGotoState33Action,
	{_CState54, CIdentifierToken}:               _CGotoState132Action,
	{_CState54, CTypeNameToken}:                 _CGotoState21Action,
	{_CState54, CTypedefToken}:                  _CGotoState20Action,
	{_CState54, CExternToken}:                   _CGotoState10Action,
	{_CState54, CStaticToken}:                   _CGotoState18Action,
	{_CState54, CAutoToken}:                     _CGotoState5Action,
	{_CState54, CRegisterToken}:                 _CGotoState15Action,
	{_CState54, CCharToken}:                     _CGotoState6Action,
	{_CState54, CShortToken}:                    _CGotoState16Action,
	{_CState54, CIntToken}:                      _CGotoState13Action,
	{_CState54, CLongToken}:                     _CGotoState14Action,
	{_CState54, CSignedToken}:                   _CGotoState17Action,
	{_CState54, CUnsignedToken}:                 _CGotoState23Action,
	{_CState54, CFloatToken}:                    _CGotoState11Action,
	{_CState54, CDoubleToken}:                   _CGotoState8Action,
	{_CState54, CConstToken}:                    _CGotoState7Action,
	{_CState54, CVolatileToken}:                 _CGotoState25Action,
	{_CState54, CVoidToken}:                     _CGotoState24Action,
	{_CState54, CStructToken}:                   _CGotoState19Action,
	{_CState54, CUnionToken}:                    _CGotoState22Action,
	{_CState54, CEnumToken}:                     _CGotoState9Action,
	{_CState54, ')'}:                            _CGotoState131Action,
	{_CState54, CDeclarationSpecifiersType}:     _CGotoState133Action,
	{_CState54, CStorageClassSpecifierType}:     _CGotoState34Action,
	{_CState54, CTypeSpecifierType}:             _CGotoState38Action,
	{_CState54, CStructOrUnionSpecifierType}:    _CGotoState36Action,
	{_CState54, CStructOrUnionType}:             _CGotoState35Action,
	{_CState54, CEnumSpecifierType}:             _CGotoState30Action,
	{_CState54, CTypeQualifierType}:             _CGotoState37Action,
	{_CState54, CParameterTypeListType}:         _CGotoState137Action,
	{_CState54, CParameterListType}:             _CGotoState136Action,
	{_CState54, CParameterDeclarationType}:      _CGotoState135Action,
	{_CState54, CIdentifierListType}:            _CGotoState134Action,
	{_CState55, CIdentifierToken}:               _CGotoState139Action,
	{_CState55, CConstantToken}:                 _CGotoState86Action,
	{_CState55, CStringLiteralToken}:            _CGotoState98Action,
	{_CState55, CSizeofToken}:                   _CGotoState97Action,
	{_CState55, CIncOpToken}:                    _CGotoState95Action,
	{_CState55, CDecOpToken}:                    _CGotoState88Action,
	{_CState55, '('}:                            _CGotoState77Action,
	{_CState55, ']'}:                            _CGotoState138Action,
	{_CState55, '*'}:                            _CGotoState78Action,
	{_CState55, '-'}:                            _CGotoState80Action,
	{_CState55, '+'}:                            _CGotoState79Action,
	{_CState55, '&'}:                            _CGotoState76Action,
	{_CState55, '!'}:                            _CGotoState75Action,
	{_CState55, '~'}:                            _CGotoState83Action,
	{_CState55, CPrimaryExpressionType}:         _CGotoState120Action,
	{_CState55, CPostfixExpressionType}:         _CGotoState119Action,
	{_CState55, CUnaryExpressionType}:           _CGotoState142Action,
	{_CState55, CUnaryOperatorType}:             _CGotoState127Action,
	{_CState55, CCastExpressionType}:            _CGotoState104Action,
	{_CState55, CMultiplicativeExpressionType}:  _CGotoState118Action,
	{_CState55, CAdditiveExpressionType}:        _CGotoState101Action,
	{_CState55, CShiftExpressionType}:           _CGotoState123Action,
	{_CState55, CRelationalExpressionType}:      _CGotoState121Action,
	{_CState55, CEqualityExpressionType}:        _CGotoState108Action,
	{_CState55, CAndExpressionType}:             _CGotoState102Action,
	{_CState55, CExclusiveOrExpressionType}:     _CGotoState109Action,
	{_CState55, CInclusiveOrExpressionType}:     _CGotoState112Action,
	{_CState55, CLogicalAndExpressionType}:      _CGotoState116Action,
	{_CState55, CLogicalOrExpressionType}:       _CGotoState117Action,
	{_CState55, CConditionalExpressionType}:     _CGotoState140Action,
	{_CState55, CConstantExpressionType}:        _CGotoState141Action,
	{_CState56, '('}:                            _CGotoState54Action,
	{_CState56, '['}:                            _CGotoState55Action,
	{_CState58, CTypeNameToken}:                 _CGotoState21Action,
	{_CState58, CCharToken}:                     _CGotoState6Action,
	{_CState58, CShortToken}:                    _CGotoState16Action,
	{_CState58, CIntToken}:                      _CGotoState13Action,
	{_CState58, CLongToken}:                     _CGotoState14Action,
	{_CState58, CSignedToken}:                   _CGotoState17Action,
	{_CState58, CUnsignedToken}:                 _CGotoState23Action,
	{_CState58, CFloatToken}:                    _CGotoState11Action,
	{_CState58, CDoubleToken}:                   _CGotoState8Action,
	{_CState58, CConstToken}:                    _CGotoState7Action,
	{_CState58, CVolatileToken}:                 _CGotoState25Action,
	{_CState58, CVoidToken}:                     _CGotoState24Action,
	{_CState58, CStructToken}:                   _CGotoState19Action,
	{_CState58, CUnionToken}:                    _CGotoState22Action,
	{_CState58, CEnumToken}:                     _CGotoState9Action,
	{_CState58, CTypeSpecifierType}:             _CGotoState147Action,
	{_CState58, CStructOrUnionSpecifierType}:    _CGotoState36Action,
	{_CState58, CStructOrUnionType}:             _CGotoState35Action,
	{_CState58, CStructDeclarationListType}:     _CGotoState145Action,
	{_CState58, CStructDeclarationType}:         _CGotoState144Action,
	{_CState58, CSpecifierQualifierListType}:    _CGotoState143Action,
	{_CState58, CEnumSpecifierType}:             _CGotoState30Action,
	{_CState58, CTypeQualifierType}:             _CGotoState146Action,
	{_CState59, '{'}:                            _CGotoState148Action,
	{_CState66, '='}:                            _CGotoState149Action,
	{_CState68, '}'}:                            _CGotoState151Action,
	{_CState68, ','}:                            _CGotoState150Action,
	{_CState69, CIdentifierToken}:               _CGotoState66Action,
	{_CState69, CEnumeratorListType}:            _CGotoState152Action,
	{_CState69, CEnumeratorType}:                _CGotoState67Action,
	{_CState70, CIdentifierToken}:               _CGotoState139Action,
	{_CState70, CConstantToken}:                 _CGotoState86Action,
	{_CState70, CStringLiteralToken}:            _CGotoState98Action,
	{_CState70, CSizeofToken}:                   _CGotoState97Action,
	{_CState70, CIncOpToken}:                    _CGotoState95Action,
	{_CState70, CDecOpToken}:                    _CGotoState88Action,
	{_CState70, '('}:                            _CGotoState77Action,
	{_CState70, '{'}:                            _CGotoState153Action,
	{_CState70, '*'}:                            _CGotoState78Action,
	{_CState70, '-'}:                            _CGotoState80Action,
	{_CState70, '+'}:                            _CGotoState79Action,
	{_CState70, '&'}:                            _CGotoState76Action,
	{_CState70, '!'}:                            _CGotoState75Action,
	{_CState70, '~'}:                            _CGotoState83Action,
	{_CState70, CPrimaryExpressionType}:         _CGotoState120Action,
	{_CState70, CPostfixExpressionType}:         _CGotoState119Action,
	{_CState70, CUnaryExpressionType}:           _CGotoState126Action,
	{_CState70, CUnaryOperatorType}:             _CGotoState127Action,
	{_CState70, CCastExpressionType}:            _CGotoState104Action,
	{_CState70, CMultiplicativeExpressionType}:  _CGotoState118Action,
	{_CState70, CAdditiveExpressionType}:        _CGotoState101Action,
	{_CState70, CShiftExpressionType}:           _CGotoState123Action,
	{_CState70, CRelationalExpressionType}:      _CGotoState121Action,
	{_CState70, CEqualityExpressionType}:        _CGotoState108Action,
	{_CState70, CAndExpressionType}:             _CGotoState102Action,
	{_CState70, CExclusiveOrExpressionType}:     _CGotoState109Action,
	{_CState70, CInclusiveOrExpressionType}:     _CGotoState112Action,
	{_CState70, CLogicalAndExpressionType}:      _CGotoState116Action,
	{_CState70, CLogicalOrExpressionType}:       _CGotoState117Action,
	{_CState70, CConditionalExpressionType}:     _CGotoState106Action,
	{_CState70, CAssignmentExpressionType}:      _CGotoState154Action,
	{_CState70, CInitializerType}:               _CGotoState155Action,
	{_CState72, CTypeNameToken}:                 _CGotoState21Action,
	{_CState72, CTypedefToken}:                  _CGotoState20Action,
	{_CState72, CExternToken}:                   _CGotoState10Action,
	{_CState72, CStaticToken}:                   _CGotoState18Action,
	{_CState72, CAutoToken}:                     _CGotoState5Action,
	{_CState72, CRegisterToken}:                 _CGotoState15Action,
	{_CState72, CCharToken}:                     _CGotoState6Action,
	{_CState72, CShortToken}:                    _CGotoState16Action,
	{_CState72, CIntToken}:                      _CGotoState13Action,
	{_CState72, CLongToken}:                     _CGotoState14Action,
	{_CState72, CSignedToken}:                   _CGotoState17Action,
	{_CState72, CUnsignedToken}:                 _CGotoState23Action,
	{_CState72, CFloatToken}:                    _CGotoState11Action,
	{_CState72, CDoubleToken}:                   _CGotoState8Action,
	{_CState72, CConstToken}:                    _CGotoState7Action,
	{_CState72, CVolatileToken}:                 _CGotoState25Action,
	{_CState72, CVoidToken}:                     _CGotoState24Action,
	{_CState72, CStructToken}:                   _CGotoState19Action,
	{_CState72, CUnionToken}:                    _CGotoState22Action,
	{_CState72, CEnumToken}:                     _CGotoState9Action,
	{_CState72, '{'}:                            _CGotoState49Action,
	{_CState72, CDeclarationType}:               _CGotoState129Action,
	{_CState72, CDeclarationSpecifiersType}:     _CGotoState53Action,
	{_CState72, CStorageClassSpecifierType}:     _CGotoState34Action,
	{_CState72, CTypeSpecifierType}:             _CGotoState38Action,
	{_CState72, CStructOrUnionSpecifierType}:    _CGotoState36Action,
	{_CState72, CStructOrUnionType}:             _CGotoState35Action,
	{_CState72, CEnumSpecifierType}:             _CGotoState30Action,
	{_CState72, CTypeQualifierType}:             _CGotoState37Action,
	{_CState72, CCompoundStatementType}:         _CGotoState156Action,
	{_CState73, CIdentifierToken}:               _CGotoState12Action,
	{_CState73, '('}:                            _CGotoState3Action,
	{_CState73, '*'}:                            _CGotoState4Action,
	{_CState73, CInitDeclaratorType}:            _CGotoState157Action,
	{_CState73, CDeclaratorType}:                _CGotoState130Action,
	{_CState73, CDirectDeclaratorType}:          _CGotoState29Action,
	{_CState73, CPointerType}:                   _CGotoState33Action,
	{_CState77, CIdentifierToken}:               _CGotoState139Action,
	{_CState77, CConstantToken}:                 _CGotoState86Action,
	{_CState77, CStringLiteralToken}:            _CGotoState98Action,
	{_CState77, CSizeofToken}:                   _CGotoState97Action,
	{_CState77, CIncOpToken}:                    _CGotoState95Action,
	{_CState77, CDecOpToken}:                    _CGotoState88Action,
	{_CState77, CTypeNameToken}:                 _CGotoState21Action,
	{_CState77, CCharToken}:                     _CGotoState6Action,
	{_CState77, CShortToken}:                    _CGotoState16Action,
	{_CState77, CIntToken}:                      _CGotoState13Action,
	{_CState77, CLongToken}:                     _CGotoState14Action,
	{_CState77, CSignedToken}:                   _CGotoState17Action,
	{_CState77, CUnsignedToken}:                 _CGotoState23Action,
	{_CState77, CFloatToken}:                    _CGotoState11Action,
	{_CState77, CDoubleToken}:                   _CGotoState8Action,
	{_CState77, CConstToken}:                    _CGotoState7Action,
	{_CState77, CVolatileToken}:                 _CGotoState25Action,
	{_CState77, CVoidToken}:                     _CGotoState24Action,
	{_CState77, CStructToken}:                   _CGotoState19Action,
	{_CState77, CUnionToken}:                    _CGotoState22Action,
	{_CState77, CEnumToken}:                     _CGotoState9Action,
	{_CState77, '('}:                            _CGotoState77Action,
	{_CState77, '*'}:                            _CGotoState78Action,
	{_CState77, '-'}:                            _CGotoState80Action,
	{_CState77, '+'}:                            _CGotoState79Action,
	{_CState77, '&'}:                            _CGotoState76Action,
	{_CState77, '!'}:                            _CGotoState75Action,
	{_CState77, '~'}:                            _CGotoState83Action,
	{_CState77, CPrimaryExpressionType}:         _CGotoState120Action,
	{_CState77, CPostfixExpressionType}:         _CGotoState119Action,
	{_CState77, CUnaryExpressionType}:           _CGotoState126Action,
	{_CState77, CUnaryOperatorType}:             _CGotoState127Action,
	{_CState77, CCastExpressionType}:            _CGotoState104Action,
	{_CState77, CMultiplicativeExpressionType}:  _CGotoState118Action,
	{_CState77, CAdditiveExpressionType}:        _CGotoState101Action,
	{_CState77, CShiftExpressionType}:           _CGotoState123Action,
	{_CState77, CRelationalExpressionType}:      _CGotoState121Action,
	{_CState77, CEqualityExpressionType}:        _CGotoState108Action,
	{_CState77, CAndExpressionType}:             _CGotoState102Action,
	{_CState77, CExclusiveOrExpressionType}:     _CGotoState109Action,
	{_CState77, CInclusiveOrExpressionType}:     _CGotoState112Action,
	{_CState77, CLogicalAndExpressionType}:      _CGotoState116Action,
	{_CState77, CLogicalOrExpressionType}:       _CGotoState117Action,
	{_CState77, CConditionalExpressionType}:     _CGotoState106Action,
	{_CState77, CAssignmentExpressionType}:      _CGotoState103Action,
	{_CState77, CExpressionType}:                _CGotoState158Action,
	{_CState77, CTypeSpecifierType}:             _CGotoState147Action,
	{_CState77, CStructOrUnionSpecifierType}:    _CGotoState36Action,
	{_CState77, CStructOrUnionType}:             _CGotoState35Action,
	{_CState77, CSpecifierQualifierListType}:    _CGotoState159Action,
	{_CState77, CEnumSpecifierType}:             _CGotoState30Action,
	{_CState77, CTypeQualifierType}:             _CGotoState146Action,
	{_CState77, CTypeNameType}:                  _CGotoState160Action,
	{_CState84, ';'}:                            _CGotoState161Action,
	{_CState85, CIdentifierToken}:               _CGotoState139Action,
	{_CState85, CConstantToken}:                 _CGotoState86Action,
	{_CState85, CStringLiteralToken}:            _CGotoState98Action,
	{_CState85, CSizeofToken}:                   _CGotoState97Action,
	{_CState85, CIncOpToken}:                    _CGotoState95Action,
	{_CState85, CDecOpToken}:                    _CGotoState88Action,
	{_CState85, '('}:                            _CGotoState77Action,
	{_CState85, '*'}:                            _CGotoState78Action,
	{_CState85, '-'}:                            _CGotoState80Action,
	{_CState85, '+'}:                            _CGotoState79Action,
	{_CState85, '&'}:                            _CGotoState76Action,
	{_CState85, '!'}:                            _CGotoState75Action,
	{_CState85, '~'}:                            _CGotoState83Action,
	{_CState85, CPrimaryExpressionType}:         _CGotoState120Action,
	{_CState85, CPostfixExpressionType}:         _CGotoState119Action,
	{_CState85, CUnaryExpressionType}:           _CGotoState142Action,
	{_CState85, CUnaryOperatorType}:             _CGotoState127Action,
	{_CState85, CCastExpressionType}:            _CGotoState104Action,
	{_CState85, CMultiplicativeExpressionType}:  _CGotoState118Action,
	{_CState85, CAdditiveExpressionType}:        _CGotoState101Action,
	{_CState85, CShiftExpressionType}:           _CGotoState123Action,
	{_CState85, CRelationalExpressionType}:      _CGotoState121Action,
	{_CState85, CEqualityExpressionType}:        _CGotoState108Action,
	{_CState85, CAndExpressionType}:             _CGotoState102Action,
	{_CState85, CExclusiveOrExpressionType}:     _CGotoState109Action,
	{_CState85, CInclusiveOrExpressionType}:     _CGotoState112Action,
	{_CState85, CLogicalAndExpressionType}:      _CGotoState116Action,
	{_CState85, CLogicalOrExpressionType}:       _CGotoState117Action,
	{_CState85, CConditionalExpressionType}:     _CGotoState140Action,
	{_CState85, CConstantExpressionType}:        _CGotoState162Action,
	{_CState87, ';'}:                            _CGotoState163Action,
	{_CState88, CIdentifierToken}:               _CGotoState139Action,
	{_CState88, CConstantToken}:                 _CGotoState86Action,
	{_CState88, CStringLiteralToken}:            _CGotoState98Action,
	{_CState88, CSizeofToken}:                   _CGotoState97Action,
	{_CState88, CIncOpToken}:                    _CGotoState95Action,
	{_CState88, CDecOpToken}:                    _CGotoState88Action,
	{_CState88, '('}:                            _CGotoState164Action,
	{_CState88, '*'}:                            _CGotoState78Action,
	{_CState88, '-'}:                            _CGotoState80Action,
	{_CState88, '+'}:                            _CGotoState79Action,
	{_CState88, '&'}:                            _CGotoState76Action,
	{_CState88, '!'}:                            _CGotoState75Action,
	{_CState88, '~'}:                            _CGotoState83Action,
	{_CState88, CPrimaryExpressionType}:         _CGotoState120Action,
	{_CState88, CPostfixExpressionType}:         _CGotoState119Action,
	{_CState88, CUnaryExpressionType}:           _CGotoState165Action,
	{_CState88, CUnaryOperatorType}:             _CGotoState127Action,
	{_CState89, ':'}:                            _CGotoState166Action,
	{_CState90, CIdentifierToken}:               _CGotoState93Action,
	{_CState90, CConstantToken}:                 _CGotoState86Action,
	{_CState90, CStringLiteralToken}:            _CGotoState98Action,
	{_CState90, CSizeofToken}:                   _CGotoState97Action,
	{_CState90, CIncOpToken}:                    _CGotoState95Action,
	{_CState90, CDecOpToken}:                    _CGotoState88Action,
	{_CState90, CCaseToken}:                     _CGotoState85Action,
	{_CState90, CDefaultToken}:                  _CGotoState89Action,
	{_CState90, CIfToken}:                       _CGotoState94Action,
	{_CState90, CSwitchToken}:                   _CGotoState99Action,
	{_CState90, CWhileToken}:                    _CGotoState100Action,
	{_CState90, CDoToken}:                       _CGotoState90Action,
	{_CState90, CForToken}:                      _CGotoState91Action,
	{_CState90, CGotoToken}:                     _CGotoState92Action,
	{_CState90, CContinueToken}:                 _CGotoState87Action,
	{_CState90, CBreakToken}:                    _CGotoState84Action,
	{_CState90, CReturnToken}:                   _CGotoState96Action,
	{_CState90, '('}:                            _CGotoState77Action,
	{_CState90, '{'}:                            _CGotoState49Action,
	{_CState90, ';'}:                            _CGotoState81Action,
	{_CState90, '*'}:                            _CGotoState78Action,
	{_CState90, '-'}:                            _CGotoState80Action,
	{_CState90, '+'}:                            _CGotoState79Action,
	{_CState90, '&'}:                            _CGotoState76Action,
	{_CState90, '!'}:                            _CGotoState75Action,
	{_CState90, '~'}:                            _CGotoState83Action,
	{_CState90, CPrimaryExpressionType}:         _CGotoState120Action,
	{_CState90, CPostfixExpressionType}:         _CGotoState119Action,
	{_CState90, CUnaryExpressionType}:           _CGotoState126Action,
	{_CState90, CUnaryOperatorType}:             _CGotoState127Action,
	{_CState90, CCastExpressionType}:            _CGotoState104Action,
	{_CState90, CMultiplicativeExpressionType}:  _CGotoState118Action,
	{_CState90, CAdditiveExpressionType}:        _CGotoState101Action,
	{_CState90, CShiftExpressionType}:           _CGotoState123Action,
	{_CState90, CRelationalExpressionType}:      _CGotoState121Action,
	{_CState90, CEqualityExpressionType}:        _CGotoState108Action,
	{_CState90, CAndExpressionType}:             _CGotoState102Action,
	{_CState90, CExclusiveOrExpressionType}:     _CGotoState109Action,
	{_CState90, CInclusiveOrExpressionType}:     _CGotoState112Action,
	{_CState90, CLogicalAndExpressionType}:      _CGotoState116Action,
	{_CState90, CLogicalOrExpressionType}:       _CGotoState117Action,
	{_CState90, CConditionalExpressionType}:     _CGotoState106Action,
	{_CState90, CAssignmentExpressionType}:      _CGotoState103Action,
	{_CState90, CExpressionType}:                _CGotoState110Action,
	{_CState90, CStatementType}:                 _CGotoState167Action,
	{_CState90, CLabeledStatementType}:          _CGotoState115Action,
	{_CState90, CCompoundStatementType}:         _CGotoState105Action,
	{_CState90, CExpressionStatementType}:       _CGotoState111Action,
	{_CState90, CSelectionStatementType}:        _CGotoState122Action,
	{_CState90, CIterationStatementType}:        _CGotoState113Action,
	{_CState90, CJumpStatementType}:             _CGotoState114Action,
	{_CState91, '('}:                            _CGotoState168Action,
	{_CState92, CIdentifierToken}:               _CGotoState169Action,
	{_CState93, ':'}:                            _CGotoState170Action,
	{_CState94, '('}:                            _CGotoState171Action,
	{_CState95, CIdentifierToken}:               _CGotoState139Action,
	{_CState95, CConstantToken}:                 _CGotoState86Action,
	{_CState95, CStringLiteralToken}:            _CGotoState98Action,
	{_CState95, CSizeofToken}:                   _CGotoState97Action,
	{_CState95, CIncOpToken}:                    _CGotoState95Action,
	{_CState95, CDecOpToken}:                    _CGotoState88Action,
	{_CState95, '('}:                            _CGotoState164Action,
	{_CState95, '*'}:                            _CGotoState78Action,
	{_CState95, '-'}:                            _CGotoState80Action,
	{_CState95, '+'}:                            _CGotoState79Action,
	{_CState95, '&'}:                            _CGotoState76Action,
	{_CState95, '!'}:                            _CGotoState75Action,
	{_CState95, '~'}:                            _CGotoState83Action,
	{_CState95, CPrimaryExpressionType}:         _CGotoState120Action,
	{_CState95, CPostfixExpressionType}:         _CGotoState119Action,
	{_CState95, CUnaryExpressionType}:           _CGotoState172Action,
	{_CState95, CUnaryOperatorType}:             _CGotoState127Action,
	{_CState96, CIdentifierToken}:               _CGotoState139Action,
	{_CState96, CConstantToken}:                 _CGotoState86Action,
	{_CState96, CStringLiteralToken}:            _CGotoState98Action,
	{_CState96, CSizeofToken}:                   _CGotoState97Action,
	{_CState96, CIncOpToken}:                    _CGotoState95Action,
	{_CState96, CDecOpToken}:                    _CGotoState88Action,
	{_CState96, '('}:                            _CGotoState77Action,
	{_CState96, ';'}:                            _CGotoState173Action,
	{_CState96, '*'}:                            _CGotoState78Action,
	{_CState96, '-'}:                            _CGotoState80Action,
	{_CState96, '+'}:                            _CGotoState79Action,
	{_CState96, '&'}:                            _CGotoState76Action,
	{_CState96, '!'}:                            _CGotoState75Action,
	{_CState96, '~'}:                            _CGotoState83Action,
	{_CState96, CPrimaryExpressionType}:         _CGotoState120Action,
	{_CState96, CPostfixExpressionType}:         _CGotoState119Action,
	{_CState96, CUnaryExpressionType}:           _CGotoState126Action,
	{_CState96, CUnaryOperatorType}:             _CGotoState127Action,
	{_CState96, CCastExpressionType}:            _CGotoState104Action,
	{_CState96, CMultiplicativeExpressionType}:  _CGotoState118Action,
	{_CState96, CAdditiveExpressionType}:        _CGotoState101Action,
	{_CState96, CShiftExpressionType}:           _CGotoState123Action,
	{_CState96, CRelationalExpressionType}:      _CGotoState121Action,
	{_CState96, CEqualityExpressionType}:        _CGotoState108Action,
	{_CState96, CAndExpressionType}:             _CGotoState102Action,
	{_CState96, CExclusiveOrExpressionType}:     _CGotoState109Action,
	{_CState96, CInclusiveOrExpressionType}:     _CGotoState112Action,
	{_CState96, CLogicalAndExpressionType}:      _CGotoState116Action,
	{_CState96, CLogicalOrExpressionType}:       _CGotoState117Action,
	{_CState96, CConditionalExpressionType}:     _CGotoState106Action,
	{_CState96, CAssignmentExpressionType}:      _CGotoState103Action,
	{_CState96, CExpressionType}:                _CGotoState174Action,
	{_CState97, CIdentifierToken}:               _CGotoState139Action,
	{_CState97, CConstantToken}:                 _CGotoState86Action,
	{_CState97, CStringLiteralToken}:            _CGotoState98Action,
	{_CState97, CSizeofToken}:                   _CGotoState97Action,
	{_CState97, CIncOpToken}:                    _CGotoState95Action,
	{_CState97, CDecOpToken}:                    _CGotoState88Action,
	{_CState97, '('}:                            _CGotoState175Action,
	{_CState97, '*'}:                            _CGotoState78Action,
	{_CState97, '-'}:                            _CGotoState80Action,
	{_CState97, '+'}:                            _CGotoState79Action,
	{_CState97, '&'}:                            _CGotoState76Action,
	{_CState97, '!'}:                            _CGotoState75Action,
	{_CState97, '~'}:                            _CGotoState83Action,
	{_CState97, CPrimaryExpressionType}:         _CGotoState120Action,
	{_CState97, CPostfixExpressionType}:         _CGotoState119Action,
	{_CState97, CUnaryExpressionType}:           _CGotoState176Action,
	{_CState97, CUnaryOperatorType}:             _CGotoState127Action,
	{_CState99, '('}:                            _CGotoState177Action,
	{_CState100, '('}:                           _CGotoState178Action,
	{_CState101, '-'}:                           _CGotoState180Action,
	{_CState101, '+'}:                           _CGotoState179Action,
	{_CState102, '&'}:                           _CGotoState181Action,
	{_CState107, CIdentifierToken}:              _CGotoState93Action,
	{_CState107, CConstantToken}:                _CGotoState86Action,
	{_CState107, CStringLiteralToken}:           _CGotoState98Action,
	{_CState107, CSizeofToken}:                  _CGotoState97Action,
	{_CState107, CIncOpToken}:                   _CGotoState95Action,
	{_CState107, CDecOpToken}:                   _CGotoState88Action,
	{_CState107, CTypeNameToken}:                _CGotoState21Action,
	{_CState107, CTypedefToken}:                 _CGotoState20Action,
	{_CState107, CExternToken}:                  _CGotoState10Action,
	{_CState107, CStaticToken}:                  _CGotoState18Action,
	{_CState107, CAutoToken}:                    _CGotoState5Action,
	{_CState107, CRegisterToken}:                _CGotoState15Action,
	{_CState107, CCharToken}:                    _CGotoState6Action,
	{_CState107, CShortToken}:                   _CGotoState16Action,
	{_CState107, CIntToken}:                     _CGotoState13Action,
	{_CState107, CLongToken}:                    _CGotoState14Action,
	{_CState107, CSignedToken}:                  _CGotoState17Action,
	{_CState107, CUnsignedToken}:                _CGotoState23Action,
	{_CState107, CFloatToken}:                   _CGotoState11Action,
	{_CState107, CDoubleToken}:                  _CGotoState8Action,
	{_CState107, CConstToken}:                   _CGotoState7Action,
	{_CState107, CVolatileToken}:                _CGotoState25Action,
	{_CState107, CVoidToken}:                    _CGotoState24Action,
	{_CState107, CStructToken}:                  _CGotoState19Action,
	{_CState107, CUnionToken}:                   _CGotoState22Action,
	{_CState107, CEnumToken}:                    _CGotoState9Action,
	{_CState107, CCaseToken}:                    _CGotoState85Action,
	{_CState107, CDefaultToken}:                 _CGotoState89Action,
	{_CState107, CIfToken}:                      _CGotoState94Action,
	{_CState107, CSwitchToken}:                  _CGotoState99Action,
	{_CState107, CWhileToken}:                   _CGotoState100Action,
	{_CState107, CDoToken}:                      _CGotoState90Action,
	{_CState107, CForToken}:                     _CGotoState91Action,
	{_CState107, CGotoToken}:                    _CGotoState92Action,
	{_CState107, CContinueToken}:                _CGotoState87Action,
	{_CState107, CBreakToken}:                   _CGotoState84Action,
	{_CState107, CReturnToken}:                  _CGotoState96Action,
	{_CState107, '('}:                           _CGotoState77Action,
	{_CState107, '{'}:                           _CGotoState49Action,
	{_CState107, '}'}:                           _CGotoState182Action,
	{_CState107, ';'}:                           _CGotoState81Action,
	{_CState107, '*'}:                           _CGotoState78Action,
	{_CState107, '-'}:                           _CGotoState80Action,
	{_CState107, '+'}:                           _CGotoState79Action,
	{_CState107, '&'}:                           _CGotoState76Action,
	{_CState107, '!'}:                           _CGotoState75Action,
	{_CState107, '~'}:                           _CGotoState83Action,
	{_CState107, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState107, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState107, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState107, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState107, CCastExpressionType}:           _CGotoState104Action,
	{_CState107, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState107, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState107, CShiftExpressionType}:          _CGotoState123Action,
	{_CState107, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState107, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState107, CAndExpressionType}:            _CGotoState102Action,
	{_CState107, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState107, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState107, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState107, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState107, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState107, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState107, CExpressionType}:               _CGotoState110Action,
	{_CState107, CDeclarationType}:              _CGotoState129Action,
	{_CState107, CDeclarationSpecifiersType}:    _CGotoState53Action,
	{_CState107, CStorageClassSpecifierType}:    _CGotoState34Action,
	{_CState107, CTypeSpecifierType}:            _CGotoState38Action,
	{_CState107, CStructOrUnionSpecifierType}:   _CGotoState36Action,
	{_CState107, CStructOrUnionType}:            _CGotoState35Action,
	{_CState107, CEnumSpecifierType}:            _CGotoState30Action,
	{_CState107, CTypeQualifierType}:            _CGotoState37Action,
	{_CState107, CStatementType}:                _CGotoState124Action,
	{_CState107, CLabeledStatementType}:         _CGotoState115Action,
	{_CState107, CCompoundStatementType}:        _CGotoState105Action,
	{_CState107, CStatementListType}:            _CGotoState183Action,
	{_CState107, CExpressionStatementType}:      _CGotoState111Action,
	{_CState107, CSelectionStatementType}:       _CGotoState122Action,
	{_CState107, CIterationStatementType}:       _CGotoState113Action,
	{_CState107, CJumpStatementType}:            _CGotoState114Action,
	{_CState108, CEqOpToken}:                    _CGotoState184Action,
	{_CState108, CNeOpToken}:                    _CGotoState185Action,
	{_CState109, '^'}:                           _CGotoState186Action,
	{_CState110, ';'}:                           _CGotoState188Action,
	{_CState110, ','}:                           _CGotoState187Action,
	{_CState112, '|'}:                           _CGotoState189Action,
	{_CState116, CAndOpToken}:                   _CGotoState190Action,
	{_CState117, COrOpToken}:                    _CGotoState192Action,
	{_CState117, '?'}:                           _CGotoState191Action,
	{_CState118, '*'}:                           _CGotoState194Action,
	{_CState118, '/'}:                           _CGotoState195Action,
	{_CState118, '%'}:                           _CGotoState193Action,
	{_CState119, CPtrOpToken}:                   _CGotoState201Action,
	{_CState119, CIncOpToken}:                   _CGotoState200Action,
	{_CState119, CDecOpToken}:                   _CGotoState199Action,
	{_CState119, '('}:                           _CGotoState196Action,
	{_CState119, '['}:                           _CGotoState198Action,
	{_CState119, '.'}:                           _CGotoState197Action,
	{_CState121, CLeOpToken}:                    _CGotoState205Action,
	{_CState121, CGeOpToken}:                    _CGotoState204Action,
	{_CState121, '<'}:                           _CGotoState202Action,
	{_CState121, '>'}:                           _CGotoState203Action,
	{_CState123, CLeftOpToken}:                  _CGotoState206Action,
	{_CState123, CRightOpToken}:                 _CGotoState207Action,
	{_CState125, CIdentifierToken}:              _CGotoState93Action,
	{_CState125, CConstantToken}:                _CGotoState86Action,
	{_CState125, CStringLiteralToken}:           _CGotoState98Action,
	{_CState125, CSizeofToken}:                  _CGotoState97Action,
	{_CState125, CIncOpToken}:                   _CGotoState95Action,
	{_CState125, CDecOpToken}:                   _CGotoState88Action,
	{_CState125, CCaseToken}:                    _CGotoState85Action,
	{_CState125, CDefaultToken}:                 _CGotoState89Action,
	{_CState125, CIfToken}:                      _CGotoState94Action,
	{_CState125, CSwitchToken}:                  _CGotoState99Action,
	{_CState125, CWhileToken}:                   _CGotoState100Action,
	{_CState125, CDoToken}:                      _CGotoState90Action,
	{_CState125, CForToken}:                     _CGotoState91Action,
	{_CState125, CGotoToken}:                    _CGotoState92Action,
	{_CState125, CContinueToken}:                _CGotoState87Action,
	{_CState125, CBreakToken}:                   _CGotoState84Action,
	{_CState125, CReturnToken}:                  _CGotoState96Action,
	{_CState125, '('}:                           _CGotoState77Action,
	{_CState125, '{'}:                           _CGotoState49Action,
	{_CState125, '}'}:                           _CGotoState208Action,
	{_CState125, ';'}:                           _CGotoState81Action,
	{_CState125, '*'}:                           _CGotoState78Action,
	{_CState125, '-'}:                           _CGotoState80Action,
	{_CState125, '+'}:                           _CGotoState79Action,
	{_CState125, '&'}:                           _CGotoState76Action,
	{_CState125, '!'}:                           _CGotoState75Action,
	{_CState125, '~'}:                           _CGotoState83Action,
	{_CState125, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState125, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState125, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState125, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState125, CCastExpressionType}:           _CGotoState104Action,
	{_CState125, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState125, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState125, CShiftExpressionType}:          _CGotoState123Action,
	{_CState125, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState125, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState125, CAndExpressionType}:            _CGotoState102Action,
	{_CState125, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState125, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState125, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState125, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState125, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState125, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState125, CExpressionType}:               _CGotoState110Action,
	{_CState125, CStatementType}:                _CGotoState209Action,
	{_CState125, CLabeledStatementType}:         _CGotoState115Action,
	{_CState125, CCompoundStatementType}:        _CGotoState105Action,
	{_CState125, CExpressionStatementType}:      _CGotoState111Action,
	{_CState125, CSelectionStatementType}:       _CGotoState122Action,
	{_CState125, CIterationStatementType}:       _CGotoState113Action,
	{_CState125, CJumpStatementType}:            _CGotoState114Action,
	{_CState126, CMulAssignToken}:               _CGotoState216Action,
	{_CState126, CDivAssignToken}:               _CGotoState213Action,
	{_CState126, CModAssignToken}:               _CGotoState215Action,
	{_CState126, CAddAssignToken}:               _CGotoState211Action,
	{_CState126, CSubAssignToken}:               _CGotoState219Action,
	{_CState126, CLeftAssignToken}:              _CGotoState214Action,
	{_CState126, CRightAssignToken}:             _CGotoState218Action,
	{_CState126, CAndAssignToken}:               _CGotoState212Action,
	{_CState126, CXorAssignToken}:               _CGotoState220Action,
	{_CState126, COrAssignToken}:                _CGotoState217Action,
	{_CState126, '='}:                           _CGotoState210Action,
	{_CState126, CAssignmentOperatorType}:       _CGotoState221Action,
	{_CState127, CIdentifierToken}:              _CGotoState139Action,
	{_CState127, CConstantToken}:                _CGotoState86Action,
	{_CState127, CStringLiteralToken}:           _CGotoState98Action,
	{_CState127, CSizeofToken}:                  _CGotoState97Action,
	{_CState127, CIncOpToken}:                   _CGotoState95Action,
	{_CState127, CDecOpToken}:                   _CGotoState88Action,
	{_CState127, '('}:                           _CGotoState77Action,
	{_CState127, '*'}:                           _CGotoState78Action,
	{_CState127, '-'}:                           _CGotoState80Action,
	{_CState127, '+'}:                           _CGotoState79Action,
	{_CState127, '&'}:                           _CGotoState76Action,
	{_CState127, '!'}:                           _CGotoState75Action,
	{_CState127, '~'}:                           _CGotoState83Action,
	{_CState127, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState127, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState127, CUnaryExpressionType}:          _CGotoState142Action,
	{_CState127, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState127, CCastExpressionType}:           _CGotoState222Action,
	{_CState130, '='}:                           _CGotoState70Action,
	{_CState133, CIdentifierToken}:              _CGotoState12Action,
	{_CState133, '('}:                           _CGotoState223Action,
	{_CState133, '['}:                           _CGotoState224Action,
	{_CState133, '*'}:                           _CGotoState4Action,
	{_CState133, CDeclaratorType}:               _CGotoState226Action,
	{_CState133, CDirectDeclaratorType}:         _CGotoState29Action,
	{_CState133, CPointerType}:                  _CGotoState228Action,
	{_CState133, CAbstractDeclaratorType}:       _CGotoState225Action,
	{_CState133, CDirectAbstractDeclaratorType}: _CGotoState227Action,
	{_CState134, ')'}:                           _CGotoState229Action,
	{_CState134, ','}:                           _CGotoState230Action,
	{_CState136, ','}:                           _CGotoState231Action,
	{_CState137, ')'}:                           _CGotoState232Action,
	{_CState141, ']'}:                           _CGotoState233Action,
	{_CState143, CIdentifierToken}:              _CGotoState12Action,
	{_CState143, '('}:                           _CGotoState3Action,
	{_CState143, ':'}:                           _CGotoState234Action,
	{_CState143, '*'}:                           _CGotoState4Action,
	{_CState143, CStructDeclaratorListType}:     _CGotoState237Action,
	{_CState143, CStructDeclaratorType}:         _CGotoState236Action,
	{_CState143, CDeclaratorType}:               _CGotoState235Action,
	{_CState143, CDirectDeclaratorType}:         _CGotoState29Action,
	{_CState143, CPointerType}:                  _CGotoState33Action,
	{_CState145, CTypeNameToken}:                _CGotoState21Action,
	{_CState145, CCharToken}:                    _CGotoState6Action,
	{_CState145, CShortToken}:                   _CGotoState16Action,
	{_CState145, CIntToken}:                     _CGotoState13Action,
	{_CState145, CLongToken}:                    _CGotoState14Action,
	{_CState145, CSignedToken}:                  _CGotoState17Action,
	{_CState145, CUnsignedToken}:                _CGotoState23Action,
	{_CState145, CFloatToken}:                   _CGotoState11Action,
	{_CState145, CDoubleToken}:                  _CGotoState8Action,
	{_CState145, CConstToken}:                   _CGotoState7Action,
	{_CState145, CVolatileToken}:                _CGotoState25Action,
	{_CState145, CVoidToken}:                    _CGotoState24Action,
	{_CState145, CStructToken}:                  _CGotoState19Action,
	{_CState145, CUnionToken}:                   _CGotoState22Action,
	{_CState145, CEnumToken}:                    _CGotoState9Action,
	{_CState145, '}'}:                           _CGotoState238Action,
	{_CState145, CTypeSpecifierType}:            _CGotoState147Action,
	{_CState145, CStructOrUnionSpecifierType}:   _CGotoState36Action,
	{_CState145, CStructOrUnionType}:            _CGotoState35Action,
	{_CState145, CStructDeclarationType}:        _CGotoState239Action,
	{_CState145, CSpecifierQualifierListType}:   _CGotoState143Action,
	{_CState145, CEnumSpecifierType}:            _CGotoState30Action,
	{_CState145, CTypeQualifierType}:            _CGotoState146Action,
	{_CState146, CTypeNameToken}:                _CGotoState21Action,
	{_CState146, CCharToken}:                    _CGotoState6Action,
	{_CState146, CShortToken}:                   _CGotoState16Action,
	{_CState146, CIntToken}:                     _CGotoState13Action,
	{_CState146, CLongToken}:                    _CGotoState14Action,
	{_CState146, CSignedToken}:                  _CGotoState17Action,
	{_CState146, CUnsignedToken}:                _CGotoState23Action,
	{_CState146, CFloatToken}:                   _CGotoState11Action,
	{_CState146, CDoubleToken}:                  _CGotoState8Action,
	{_CState146, CConstToken}:                   _CGotoState7Action,
	{_CState146, CVolatileToken}:                _CGotoState25Action,
	{_CState146, CVoidToken}:                    _CGotoState24Action,
	{_CState146, CStructToken}:                  _CGotoState19Action,
	{_CState146, CUnionToken}:                   _CGotoState22Action,
	{_CState146, CEnumToken}:                    _CGotoState9Action,
	{_CState146, CTypeSpecifierType}:            _CGotoState147Action,
	{_CState146, CStructOrUnionSpecifierType}:   _CGotoState36Action,
	{_CState146, CStructOrUnionType}:            _CGotoState35Action,
	{_CState146, CSpecifierQualifierListType}:   _CGotoState240Action,
	{_CState146, CEnumSpecifierType}:            _CGotoState30Action,
	{_CState146, CTypeQualifierType}:            _CGotoState146Action,
	{_CState147, CTypeNameToken}:                _CGotoState21Action,
	{_CState147, CCharToken}:                    _CGotoState6Action,
	{_CState147, CShortToken}:                   _CGotoState16Action,
	{_CState147, CIntToken}:                     _CGotoState13Action,
	{_CState147, CLongToken}:                    _CGotoState14Action,
	{_CState147, CSignedToken}:                  _CGotoState17Action,
	{_CState147, CUnsignedToken}:                _CGotoState23Action,
	{_CState147, CFloatToken}:                   _CGotoState11Action,
	{_CState147, CDoubleToken}:                  _CGotoState8Action,
	{_CState147, CConstToken}:                   _CGotoState7Action,
	{_CState147, CVolatileToken}:                _CGotoState25Action,
	{_CState147, CVoidToken}:                    _CGotoState24Action,
	{_CState147, CStructToken}:                  _CGotoState19Action,
	{_CState147, CUnionToken}:                   _CGotoState22Action,
	{_CState147, CEnumToken}:                    _CGotoState9Action,
	{_CState147, CTypeSpecifierType}:            _CGotoState147Action,
	{_CState147, CStructOrUnionSpecifierType}:   _CGotoState36Action,
	{_CState147, CStructOrUnionType}:            _CGotoState35Action,
	{_CState147, CSpecifierQualifierListType}:   _CGotoState241Action,
	{_CState147, CEnumSpecifierType}:            _CGotoState30Action,
	{_CState147, CTypeQualifierType}:            _CGotoState146Action,
	{_CState148, CTypeNameToken}:                _CGotoState21Action,
	{_CState148, CCharToken}:                    _CGotoState6Action,
	{_CState148, CShortToken}:                   _CGotoState16Action,
	{_CState148, CIntToken}:                     _CGotoState13Action,
	{_CState148, CLongToken}:                    _CGotoState14Action,
	{_CState148, CSignedToken}:                  _CGotoState17Action,
	{_CState148, CUnsignedToken}:                _CGotoState23Action,
	{_CState148, CFloatToken}:                   _CGotoState11Action,
	{_CState148, CDoubleToken}:                  _CGotoState8Action,
	{_CState148, CConstToken}:                   _CGotoState7Action,
	{_CState148, CVolatileToken}:                _CGotoState25Action,
	{_CState148, CVoidToken}:                    _CGotoState24Action,
	{_CState148, CStructToken}:                  _CGotoState19Action,
	{_CState148, CUnionToken}:                   _CGotoState22Action,
	{_CState148, CEnumToken}:                    _CGotoState9Action,
	{_CState148, CTypeSpecifierType}:            _CGotoState147Action,
	{_CState148, CStructOrUnionSpecifierType}:   _CGotoState36Action,
	{_CState148, CStructOrUnionType}:            _CGotoState35Action,
	{_CState148, CStructDeclarationListType}:    _CGotoState242Action,
	{_CState148, CStructDeclarationType}:        _CGotoState144Action,
	{_CState148, CSpecifierQualifierListType}:   _CGotoState143Action,
	{_CState148, CEnumSpecifierType}:            _CGotoState30Action,
	{_CState148, CTypeQualifierType}:            _CGotoState146Action,
	{_CState149, CIdentifierToken}:              _CGotoState139Action,
	{_CState149, CConstantToken}:                _CGotoState86Action,
	{_CState149, CStringLiteralToken}:           _CGotoState98Action,
	{_CState149, CSizeofToken}:                  _CGotoState97Action,
	{_CState149, CIncOpToken}:                   _CGotoState95Action,
	{_CState149, CDecOpToken}:                   _CGotoState88Action,
	{_CState149, '('}:                           _CGotoState77Action,
	{_CState149, '*'}:                           _CGotoState78Action,
	{_CState149, '-'}:                           _CGotoState80Action,
	{_CState149, '+'}:                           _CGotoState79Action,
	{_CState149, '&'}:                           _CGotoState76Action,
	{_CState149, '!'}:                           _CGotoState75Action,
	{_CState149, '~'}:                           _CGotoState83Action,
	{_CState149, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState149, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState149, CUnaryExpressionType}:          _CGotoState142Action,
	{_CState149, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState149, CCastExpressionType}:           _CGotoState104Action,
	{_CState149, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState149, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState149, CShiftExpressionType}:          _CGotoState123Action,
	{_CState149, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState149, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState149, CAndExpressionType}:            _CGotoState102Action,
	{_CState149, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState149, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState149, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState149, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState149, CConditionalExpressionType}:    _CGotoState140Action,
	{_CState149, CConstantExpressionType}:       _CGotoState243Action,
	{_CState150, CIdentifierToken}:              _CGotoState66Action,
	{_CState150, CEnumeratorType}:               _CGotoState244Action,
	{_CState152, '}'}:                           _CGotoState245Action,
	{_CState152, ','}:                           _CGotoState150Action,
	{_CState153, CIdentifierToken}:              _CGotoState139Action,
	{_CState153, CConstantToken}:                _CGotoState86Action,
	{_CState153, CStringLiteralToken}:           _CGotoState98Action,
	{_CState153, CSizeofToken}:                  _CGotoState97Action,
	{_CState153, CIncOpToken}:                   _CGotoState95Action,
	{_CState153, CDecOpToken}:                   _CGotoState88Action,
	{_CState153, '('}:                           _CGotoState77Action,
	{_CState153, '{'}:                           _CGotoState153Action,
	{_CState153, '*'}:                           _CGotoState78Action,
	{_CState153, '-'}:                           _CGotoState80Action,
	{_CState153, '+'}:                           _CGotoState79Action,
	{_CState153, '&'}:                           _CGotoState76Action,
	{_CState153, '!'}:                           _CGotoState75Action,
	{_CState153, '~'}:                           _CGotoState83Action,
	{_CState153, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState153, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState153, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState153, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState153, CCastExpressionType}:           _CGotoState104Action,
	{_CState153, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState153, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState153, CShiftExpressionType}:          _CGotoState123Action,
	{_CState153, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState153, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState153, CAndExpressionType}:            _CGotoState102Action,
	{_CState153, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState153, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState153, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState153, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState153, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState153, CAssignmentExpressionType}:     _CGotoState154Action,
	{_CState153, CInitializerType}:              _CGotoState246Action,
	{_CState153, CInitializerListType}:          _CGotoState247Action,
	{_CState158, ')'}:                           _CGotoState248Action,
	{_CState158, ','}:                           _CGotoState187Action,
	{_CState159, '('}:                           _CGotoState249Action,
	{_CState159, '['}:                           _CGotoState224Action,
	{_CState159, '*'}:                           _CGotoState4Action,
	{_CState159, CPointerType}:                  _CGotoState251Action,
	{_CState159, CAbstractDeclaratorType}:       _CGotoState250Action,
	{_CState159, CDirectAbstractDeclaratorType}: _CGotoState227Action,
	{_CState160, ')'}:                           _CGotoState252Action,
	{_CState162, ':'}:                           _CGotoState253Action,
	{_CState164, CIdentifierToken}:              _CGotoState139Action,
	{_CState164, CConstantToken}:                _CGotoState86Action,
	{_CState164, CStringLiteralToken}:           _CGotoState98Action,
	{_CState164, CSizeofToken}:                  _CGotoState97Action,
	{_CState164, CIncOpToken}:                   _CGotoState95Action,
	{_CState164, CDecOpToken}:                   _CGotoState88Action,
	{_CState164, '('}:                           _CGotoState77Action,
	{_CState164, '*'}:                           _CGotoState78Action,
	{_CState164, '-'}:                           _CGotoState80Action,
	{_CState164, '+'}:                           _CGotoState79Action,
	{_CState164, '&'}:                           _CGotoState76Action,
	{_CState164, '!'}:                           _CGotoState75Action,
	{_CState164, '~'}:                           _CGotoState83Action,
	{_CState164, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState164, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState164, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState164, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState164, CCastExpressionType}:           _CGotoState104Action,
	{_CState164, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState164, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState164, CShiftExpressionType}:          _CGotoState123Action,
	{_CState164, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState164, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState164, CAndExpressionType}:            _CGotoState102Action,
	{_CState164, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState164, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState164, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState164, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState164, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState164, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState164, CExpressionType}:               _CGotoState158Action,
	{_CState166, CIdentifierToken}:              _CGotoState93Action,
	{_CState166, CConstantToken}:                _CGotoState86Action,
	{_CState166, CStringLiteralToken}:           _CGotoState98Action,
	{_CState166, CSizeofToken}:                  _CGotoState97Action,
	{_CState166, CIncOpToken}:                   _CGotoState95Action,
	{_CState166, CDecOpToken}:                   _CGotoState88Action,
	{_CState166, CCaseToken}:                    _CGotoState85Action,
	{_CState166, CDefaultToken}:                 _CGotoState89Action,
	{_CState166, CIfToken}:                      _CGotoState94Action,
	{_CState166, CSwitchToken}:                  _CGotoState99Action,
	{_CState166, CWhileToken}:                   _CGotoState100Action,
	{_CState166, CDoToken}:                      _CGotoState90Action,
	{_CState166, CForToken}:                     _CGotoState91Action,
	{_CState166, CGotoToken}:                    _CGotoState92Action,
	{_CState166, CContinueToken}:                _CGotoState87Action,
	{_CState166, CBreakToken}:                   _CGotoState84Action,
	{_CState166, CReturnToken}:                  _CGotoState96Action,
	{_CState166, '('}:                           _CGotoState77Action,
	{_CState166, '{'}:                           _CGotoState49Action,
	{_CState166, ';'}:                           _CGotoState81Action,
	{_CState166, '*'}:                           _CGotoState78Action,
	{_CState166, '-'}:                           _CGotoState80Action,
	{_CState166, '+'}:                           _CGotoState79Action,
	{_CState166, '&'}:                           _CGotoState76Action,
	{_CState166, '!'}:                           _CGotoState75Action,
	{_CState166, '~'}:                           _CGotoState83Action,
	{_CState166, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState166, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState166, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState166, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState166, CCastExpressionType}:           _CGotoState104Action,
	{_CState166, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState166, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState166, CShiftExpressionType}:          _CGotoState123Action,
	{_CState166, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState166, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState166, CAndExpressionType}:            _CGotoState102Action,
	{_CState166, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState166, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState166, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState166, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState166, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState166, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState166, CExpressionType}:               _CGotoState110Action,
	{_CState166, CStatementType}:                _CGotoState254Action,
	{_CState166, CLabeledStatementType}:         _CGotoState115Action,
	{_CState166, CCompoundStatementType}:        _CGotoState105Action,
	{_CState166, CExpressionStatementType}:      _CGotoState111Action,
	{_CState166, CSelectionStatementType}:       _CGotoState122Action,
	{_CState166, CIterationStatementType}:       _CGotoState113Action,
	{_CState166, CJumpStatementType}:            _CGotoState114Action,
	{_CState167, CWhileToken}:                   _CGotoState255Action,
	{_CState168, CIdentifierToken}:              _CGotoState139Action,
	{_CState168, CConstantToken}:                _CGotoState86Action,
	{_CState168, CStringLiteralToken}:           _CGotoState98Action,
	{_CState168, CSizeofToken}:                  _CGotoState97Action,
	{_CState168, CIncOpToken}:                   _CGotoState95Action,
	{_CState168, CDecOpToken}:                   _CGotoState88Action,
	{_CState168, '('}:                           _CGotoState77Action,
	{_CState168, ';'}:                           _CGotoState81Action,
	{_CState168, '*'}:                           _CGotoState78Action,
	{_CState168, '-'}:                           _CGotoState80Action,
	{_CState168, '+'}:                           _CGotoState79Action,
	{_CState168, '&'}:                           _CGotoState76Action,
	{_CState168, '!'}:                           _CGotoState75Action,
	{_CState168, '~'}:                           _CGotoState83Action,
	{_CState168, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState168, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState168, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState168, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState168, CCastExpressionType}:           _CGotoState104Action,
	{_CState168, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState168, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState168, CShiftExpressionType}:          _CGotoState123Action,
	{_CState168, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState168, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState168, CAndExpressionType}:            _CGotoState102Action,
	{_CState168, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState168, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState168, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState168, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState168, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState168, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState168, CExpressionType}:               _CGotoState110Action,
	{_CState168, CExpressionStatementType}:      _CGotoState256Action,
	{_CState169, ';'}:                           _CGotoState257Action,
	{_CState170, CIdentifierToken}:              _CGotoState93Action,
	{_CState170, CConstantToken}:                _CGotoState86Action,
	{_CState170, CStringLiteralToken}:           _CGotoState98Action,
	{_CState170, CSizeofToken}:                  _CGotoState97Action,
	{_CState170, CIncOpToken}:                   _CGotoState95Action,
	{_CState170, CDecOpToken}:                   _CGotoState88Action,
	{_CState170, CCaseToken}:                    _CGotoState85Action,
	{_CState170, CDefaultToken}:                 _CGotoState89Action,
	{_CState170, CIfToken}:                      _CGotoState94Action,
	{_CState170, CSwitchToken}:                  _CGotoState99Action,
	{_CState170, CWhileToken}:                   _CGotoState100Action,
	{_CState170, CDoToken}:                      _CGotoState90Action,
	{_CState170, CForToken}:                     _CGotoState91Action,
	{_CState170, CGotoToken}:                    _CGotoState92Action,
	{_CState170, CContinueToken}:                _CGotoState87Action,
	{_CState170, CBreakToken}:                   _CGotoState84Action,
	{_CState170, CReturnToken}:                  _CGotoState96Action,
	{_CState170, '('}:                           _CGotoState77Action,
	{_CState170, '{'}:                           _CGotoState49Action,
	{_CState170, ';'}:                           _CGotoState81Action,
	{_CState170, '*'}:                           _CGotoState78Action,
	{_CState170, '-'}:                           _CGotoState80Action,
	{_CState170, '+'}:                           _CGotoState79Action,
	{_CState170, '&'}:                           _CGotoState76Action,
	{_CState170, '!'}:                           _CGotoState75Action,
	{_CState170, '~'}:                           _CGotoState83Action,
	{_CState170, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState170, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState170, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState170, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState170, CCastExpressionType}:           _CGotoState104Action,
	{_CState170, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState170, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState170, CShiftExpressionType}:          _CGotoState123Action,
	{_CState170, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState170, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState170, CAndExpressionType}:            _CGotoState102Action,
	{_CState170, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState170, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState170, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState170, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState170, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState170, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState170, CExpressionType}:               _CGotoState110Action,
	{_CState170, CStatementType}:                _CGotoState258Action,
	{_CState170, CLabeledStatementType}:         _CGotoState115Action,
	{_CState170, CCompoundStatementType}:        _CGotoState105Action,
	{_CState170, CExpressionStatementType}:      _CGotoState111Action,
	{_CState170, CSelectionStatementType}:       _CGotoState122Action,
	{_CState170, CIterationStatementType}:       _CGotoState113Action,
	{_CState170, CJumpStatementType}:            _CGotoState114Action,
	{_CState171, CIdentifierToken}:              _CGotoState139Action,
	{_CState171, CConstantToken}:                _CGotoState86Action,
	{_CState171, CStringLiteralToken}:           _CGotoState98Action,
	{_CState171, CSizeofToken}:                  _CGotoState97Action,
	{_CState171, CIncOpToken}:                   _CGotoState95Action,
	{_CState171, CDecOpToken}:                   _CGotoState88Action,
	{_CState171, '('}:                           _CGotoState77Action,
	{_CState171, '*'}:                           _CGotoState78Action,
	{_CState171, '-'}:                           _CGotoState80Action,
	{_CState171, '+'}:                           _CGotoState79Action,
	{_CState171, '&'}:                           _CGotoState76Action,
	{_CState171, '!'}:                           _CGotoState75Action,
	{_CState171, '~'}:                           _CGotoState83Action,
	{_CState171, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState171, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState171, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState171, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState171, CCastExpressionType}:           _CGotoState104Action,
	{_CState171, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState171, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState171, CShiftExpressionType}:          _CGotoState123Action,
	{_CState171, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState171, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState171, CAndExpressionType}:            _CGotoState102Action,
	{_CState171, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState171, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState171, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState171, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState171, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState171, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState171, CExpressionType}:               _CGotoState259Action,
	{_CState174, ';'}:                           _CGotoState260Action,
	{_CState174, ','}:                           _CGotoState187Action,
	{_CState175, CIdentifierToken}:              _CGotoState139Action,
	{_CState175, CConstantToken}:                _CGotoState86Action,
	{_CState175, CStringLiteralToken}:           _CGotoState98Action,
	{_CState175, CSizeofToken}:                  _CGotoState97Action,
	{_CState175, CIncOpToken}:                   _CGotoState95Action,
	{_CState175, CDecOpToken}:                   _CGotoState88Action,
	{_CState175, CTypeNameToken}:                _CGotoState21Action,
	{_CState175, CCharToken}:                    _CGotoState6Action,
	{_CState175, CShortToken}:                   _CGotoState16Action,
	{_CState175, CIntToken}:                     _CGotoState13Action,
	{_CState175, CLongToken}:                    _CGotoState14Action,
	{_CState175, CSignedToken}:                  _CGotoState17Action,
	{_CState175, CUnsignedToken}:                _CGotoState23Action,
	{_CState175, CFloatToken}:                   _CGotoState11Action,
	{_CState175, CDoubleToken}:                  _CGotoState8Action,
	{_CState175, CConstToken}:                   _CGotoState7Action,
	{_CState175, CVolatileToken}:                _CGotoState25Action,
	{_CState175, CVoidToken}:                    _CGotoState24Action,
	{_CState175, CStructToken}:                  _CGotoState19Action,
	{_CState175, CUnionToken}:                   _CGotoState22Action,
	{_CState175, CEnumToken}:                    _CGotoState9Action,
	{_CState175, '('}:                           _CGotoState77Action,
	{_CState175, '*'}:                           _CGotoState78Action,
	{_CState175, '-'}:                           _CGotoState80Action,
	{_CState175, '+'}:                           _CGotoState79Action,
	{_CState175, '&'}:                           _CGotoState76Action,
	{_CState175, '!'}:                           _CGotoState75Action,
	{_CState175, '~'}:                           _CGotoState83Action,
	{_CState175, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState175, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState175, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState175, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState175, CCastExpressionType}:           _CGotoState104Action,
	{_CState175, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState175, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState175, CShiftExpressionType}:          _CGotoState123Action,
	{_CState175, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState175, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState175, CAndExpressionType}:            _CGotoState102Action,
	{_CState175, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState175, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState175, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState175, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState175, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState175, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState175, CExpressionType}:               _CGotoState158Action,
	{_CState175, CTypeSpecifierType}:            _CGotoState147Action,
	{_CState175, CStructOrUnionSpecifierType}:   _CGotoState36Action,
	{_CState175, CStructOrUnionType}:            _CGotoState35Action,
	{_CState175, CSpecifierQualifierListType}:   _CGotoState159Action,
	{_CState175, CEnumSpecifierType}:            _CGotoState30Action,
	{_CState175, CTypeQualifierType}:            _CGotoState146Action,
	{_CState175, CTypeNameType}:                 _CGotoState261Action,
	{_CState177, CIdentifierToken}:              _CGotoState139Action,
	{_CState177, CConstantToken}:                _CGotoState86Action,
	{_CState177, CStringLiteralToken}:           _CGotoState98Action,
	{_CState177, CSizeofToken}:                  _CGotoState97Action,
	{_CState177, CIncOpToken}:                   _CGotoState95Action,
	{_CState177, CDecOpToken}:                   _CGotoState88Action,
	{_CState177, '('}:                           _CGotoState77Action,
	{_CState177, '*'}:                           _CGotoState78Action,
	{_CState177, '-'}:                           _CGotoState80Action,
	{_CState177, '+'}:                           _CGotoState79Action,
	{_CState177, '&'}:                           _CGotoState76Action,
	{_CState177, '!'}:                           _CGotoState75Action,
	{_CState177, '~'}:                           _CGotoState83Action,
	{_CState177, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState177, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState177, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState177, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState177, CCastExpressionType}:           _CGotoState104Action,
	{_CState177, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState177, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState177, CShiftExpressionType}:          _CGotoState123Action,
	{_CState177, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState177, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState177, CAndExpressionType}:            _CGotoState102Action,
	{_CState177, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState177, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState177, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState177, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState177, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState177, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState177, CExpressionType}:               _CGotoState262Action,
	{_CState178, CIdentifierToken}:              _CGotoState139Action,
	{_CState178, CConstantToken}:                _CGotoState86Action,
	{_CState178, CStringLiteralToken}:           _CGotoState98Action,
	{_CState178, CSizeofToken}:                  _CGotoState97Action,
	{_CState178, CIncOpToken}:                   _CGotoState95Action,
	{_CState178, CDecOpToken}:                   _CGotoState88Action,
	{_CState178, '('}:                           _CGotoState77Action,
	{_CState178, '*'}:                           _CGotoState78Action,
	{_CState178, '-'}:                           _CGotoState80Action,
	{_CState178, '+'}:                           _CGotoState79Action,
	{_CState178, '&'}:                           _CGotoState76Action,
	{_CState178, '!'}:                           _CGotoState75Action,
	{_CState178, '~'}:                           _CGotoState83Action,
	{_CState178, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState178, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState178, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState178, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState178, CCastExpressionType}:           _CGotoState104Action,
	{_CState178, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState178, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState178, CShiftExpressionType}:          _CGotoState123Action,
	{_CState178, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState178, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState178, CAndExpressionType}:            _CGotoState102Action,
	{_CState178, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState178, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState178, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState178, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState178, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState178, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState178, CExpressionType}:               _CGotoState263Action,
	{_CState179, CIdentifierToken}:              _CGotoState139Action,
	{_CState179, CConstantToken}:                _CGotoState86Action,
	{_CState179, CStringLiteralToken}:           _CGotoState98Action,
	{_CState179, CSizeofToken}:                  _CGotoState97Action,
	{_CState179, CIncOpToken}:                   _CGotoState95Action,
	{_CState179, CDecOpToken}:                   _CGotoState88Action,
	{_CState179, '('}:                           _CGotoState77Action,
	{_CState179, '*'}:                           _CGotoState78Action,
	{_CState179, '-'}:                           _CGotoState80Action,
	{_CState179, '+'}:                           _CGotoState79Action,
	{_CState179, '&'}:                           _CGotoState76Action,
	{_CState179, '!'}:                           _CGotoState75Action,
	{_CState179, '~'}:                           _CGotoState83Action,
	{_CState179, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState179, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState179, CUnaryExpressionType}:          _CGotoState142Action,
	{_CState179, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState179, CCastExpressionType}:           _CGotoState104Action,
	{_CState179, CMultiplicativeExpressionType}: _CGotoState264Action,
	{_CState180, CIdentifierToken}:              _CGotoState139Action,
	{_CState180, CConstantToken}:                _CGotoState86Action,
	{_CState180, CStringLiteralToken}:           _CGotoState98Action,
	{_CState180, CSizeofToken}:                  _CGotoState97Action,
	{_CState180, CIncOpToken}:                   _CGotoState95Action,
	{_CState180, CDecOpToken}:                   _CGotoState88Action,
	{_CState180, '('}:                           _CGotoState77Action,
	{_CState180, '*'}:                           _CGotoState78Action,
	{_CState180, '-'}:                           _CGotoState80Action,
	{_CState180, '+'}:                           _CGotoState79Action,
	{_CState180, '&'}:                           _CGotoState76Action,
	{_CState180, '!'}:                           _CGotoState75Action,
	{_CState180, '~'}:                           _CGotoState83Action,
	{_CState180, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState180, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState180, CUnaryExpressionType}:          _CGotoState142Action,
	{_CState180, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState180, CCastExpressionType}:           _CGotoState104Action,
	{_CState180, CMultiplicativeExpressionType}: _CGotoState265Action,
	{_CState181, CIdentifierToken}:              _CGotoState139Action,
	{_CState181, CConstantToken}:                _CGotoState86Action,
	{_CState181, CStringLiteralToken}:           _CGotoState98Action,
	{_CState181, CSizeofToken}:                  _CGotoState97Action,
	{_CState181, CIncOpToken}:                   _CGotoState95Action,
	{_CState181, CDecOpToken}:                   _CGotoState88Action,
	{_CState181, '('}:                           _CGotoState77Action,
	{_CState181, '*'}:                           _CGotoState78Action,
	{_CState181, '-'}:                           _CGotoState80Action,
	{_CState181, '+'}:                           _CGotoState79Action,
	{_CState181, '&'}:                           _CGotoState76Action,
	{_CState181, '!'}:                           _CGotoState75Action,
	{_CState181, '~'}:                           _CGotoState83Action,
	{_CState181, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState181, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState181, CUnaryExpressionType}:          _CGotoState142Action,
	{_CState181, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState181, CCastExpressionType}:           _CGotoState104Action,
	{_CState181, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState181, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState181, CShiftExpressionType}:          _CGotoState123Action,
	{_CState181, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState181, CEqualityExpressionType}:       _CGotoState266Action,
	{_CState183, CIdentifierToken}:              _CGotoState93Action,
	{_CState183, CConstantToken}:                _CGotoState86Action,
	{_CState183, CStringLiteralToken}:           _CGotoState98Action,
	{_CState183, CSizeofToken}:                  _CGotoState97Action,
	{_CState183, CIncOpToken}:                   _CGotoState95Action,
	{_CState183, CDecOpToken}:                   _CGotoState88Action,
	{_CState183, CCaseToken}:                    _CGotoState85Action,
	{_CState183, CDefaultToken}:                 _CGotoState89Action,
	{_CState183, CIfToken}:                      _CGotoState94Action,
	{_CState183, CSwitchToken}:                  _CGotoState99Action,
	{_CState183, CWhileToken}:                   _CGotoState100Action,
	{_CState183, CDoToken}:                      _CGotoState90Action,
	{_CState183, CForToken}:                     _CGotoState91Action,
	{_CState183, CGotoToken}:                    _CGotoState92Action,
	{_CState183, CContinueToken}:                _CGotoState87Action,
	{_CState183, CBreakToken}:                   _CGotoState84Action,
	{_CState183, CReturnToken}:                  _CGotoState96Action,
	{_CState183, '('}:                           _CGotoState77Action,
	{_CState183, '{'}:                           _CGotoState49Action,
	{_CState183, '}'}:                           _CGotoState267Action,
	{_CState183, ';'}:                           _CGotoState81Action,
	{_CState183, '*'}:                           _CGotoState78Action,
	{_CState183, '-'}:                           _CGotoState80Action,
	{_CState183, '+'}:                           _CGotoState79Action,
	{_CState183, '&'}:                           _CGotoState76Action,
	{_CState183, '!'}:                           _CGotoState75Action,
	{_CState183, '~'}:                           _CGotoState83Action,
	{_CState183, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState183, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState183, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState183, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState183, CCastExpressionType}:           _CGotoState104Action,
	{_CState183, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState183, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState183, CShiftExpressionType}:          _CGotoState123Action,
	{_CState183, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState183, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState183, CAndExpressionType}:            _CGotoState102Action,
	{_CState183, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState183, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState183, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState183, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState183, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState183, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState183, CExpressionType}:               _CGotoState110Action,
	{_CState183, CStatementType}:                _CGotoState209Action,
	{_CState183, CLabeledStatementType}:         _CGotoState115Action,
	{_CState183, CCompoundStatementType}:        _CGotoState105Action,
	{_CState183, CExpressionStatementType}:      _CGotoState111Action,
	{_CState183, CSelectionStatementType}:       _CGotoState122Action,
	{_CState183, CIterationStatementType}:       _CGotoState113Action,
	{_CState183, CJumpStatementType}:            _CGotoState114Action,
	{_CState184, CIdentifierToken}:              _CGotoState139Action,
	{_CState184, CConstantToken}:                _CGotoState86Action,
	{_CState184, CStringLiteralToken}:           _CGotoState98Action,
	{_CState184, CSizeofToken}:                  _CGotoState97Action,
	{_CState184, CIncOpToken}:                   _CGotoState95Action,
	{_CState184, CDecOpToken}:                   _CGotoState88Action,
	{_CState184, '('}:                           _CGotoState77Action,
	{_CState184, '*'}:                           _CGotoState78Action,
	{_CState184, '-'}:                           _CGotoState80Action,
	{_CState184, '+'}:                           _CGotoState79Action,
	{_CState184, '&'}:                           _CGotoState76Action,
	{_CState184, '!'}:                           _CGotoState75Action,
	{_CState184, '~'}:                           _CGotoState83Action,
	{_CState184, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState184, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState184, CUnaryExpressionType}:          _CGotoState142Action,
	{_CState184, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState184, CCastExpressionType}:           _CGotoState104Action,
	{_CState184, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState184, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState184, CShiftExpressionType}:          _CGotoState123Action,
	{_CState184, CRelationalExpressionType}:     _CGotoState268Action,
	{_CState185, CIdentifierToken}:              _CGotoState139Action,
	{_CState185, CConstantToken}:                _CGotoState86Action,
	{_CState185, CStringLiteralToken}:           _CGotoState98Action,
	{_CState185, CSizeofToken}:                  _CGotoState97Action,
	{_CState185, CIncOpToken}:                   _CGotoState95Action,
	{_CState185, CDecOpToken}:                   _CGotoState88Action,
	{_CState185, '('}:                           _CGotoState77Action,
	{_CState185, '*'}:                           _CGotoState78Action,
	{_CState185, '-'}:                           _CGotoState80Action,
	{_CState185, '+'}:                           _CGotoState79Action,
	{_CState185, '&'}:                           _CGotoState76Action,
	{_CState185, '!'}:                           _CGotoState75Action,
	{_CState185, '~'}:                           _CGotoState83Action,
	{_CState185, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState185, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState185, CUnaryExpressionType}:          _CGotoState142Action,
	{_CState185, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState185, CCastExpressionType}:           _CGotoState104Action,
	{_CState185, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState185, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState185, CShiftExpressionType}:          _CGotoState123Action,
	{_CState185, CRelationalExpressionType}:     _CGotoState269Action,
	{_CState186, CIdentifierToken}:              _CGotoState139Action,
	{_CState186, CConstantToken}:                _CGotoState86Action,
	{_CState186, CStringLiteralToken}:           _CGotoState98Action,
	{_CState186, CSizeofToken}:                  _CGotoState97Action,
	{_CState186, CIncOpToken}:                   _CGotoState95Action,
	{_CState186, CDecOpToken}:                   _CGotoState88Action,
	{_CState186, '('}:                           _CGotoState77Action,
	{_CState186, '*'}:                           _CGotoState78Action,
	{_CState186, '-'}:                           _CGotoState80Action,
	{_CState186, '+'}:                           _CGotoState79Action,
	{_CState186, '&'}:                           _CGotoState76Action,
	{_CState186, '!'}:                           _CGotoState75Action,
	{_CState186, '~'}:                           _CGotoState83Action,
	{_CState186, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState186, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState186, CUnaryExpressionType}:          _CGotoState142Action,
	{_CState186, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState186, CCastExpressionType}:           _CGotoState104Action,
	{_CState186, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState186, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState186, CShiftExpressionType}:          _CGotoState123Action,
	{_CState186, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState186, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState186, CAndExpressionType}:            _CGotoState270Action,
	{_CState187, CIdentifierToken}:              _CGotoState139Action,
	{_CState187, CConstantToken}:                _CGotoState86Action,
	{_CState187, CStringLiteralToken}:           _CGotoState98Action,
	{_CState187, CSizeofToken}:                  _CGotoState97Action,
	{_CState187, CIncOpToken}:                   _CGotoState95Action,
	{_CState187, CDecOpToken}:                   _CGotoState88Action,
	{_CState187, '('}:                           _CGotoState77Action,
	{_CState187, '*'}:                           _CGotoState78Action,
	{_CState187, '-'}:                           _CGotoState80Action,
	{_CState187, '+'}:                           _CGotoState79Action,
	{_CState187, '&'}:                           _CGotoState76Action,
	{_CState187, '!'}:                           _CGotoState75Action,
	{_CState187, '~'}:                           _CGotoState83Action,
	{_CState187, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState187, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState187, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState187, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState187, CCastExpressionType}:           _CGotoState104Action,
	{_CState187, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState187, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState187, CShiftExpressionType}:          _CGotoState123Action,
	{_CState187, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState187, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState187, CAndExpressionType}:            _CGotoState102Action,
	{_CState187, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState187, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState187, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState187, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState187, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState187, CAssignmentExpressionType}:     _CGotoState271Action,
	{_CState189, CIdentifierToken}:              _CGotoState139Action,
	{_CState189, CConstantToken}:                _CGotoState86Action,
	{_CState189, CStringLiteralToken}:           _CGotoState98Action,
	{_CState189, CSizeofToken}:                  _CGotoState97Action,
	{_CState189, CIncOpToken}:                   _CGotoState95Action,
	{_CState189, CDecOpToken}:                   _CGotoState88Action,
	{_CState189, '('}:                           _CGotoState77Action,
	{_CState189, '*'}:                           _CGotoState78Action,
	{_CState189, '-'}:                           _CGotoState80Action,
	{_CState189, '+'}:                           _CGotoState79Action,
	{_CState189, '&'}:                           _CGotoState76Action,
	{_CState189, '!'}:                           _CGotoState75Action,
	{_CState189, '~'}:                           _CGotoState83Action,
	{_CState189, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState189, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState189, CUnaryExpressionType}:          _CGotoState142Action,
	{_CState189, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState189, CCastExpressionType}:           _CGotoState104Action,
	{_CState189, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState189, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState189, CShiftExpressionType}:          _CGotoState123Action,
	{_CState189, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState189, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState189, CAndExpressionType}:            _CGotoState102Action,
	{_CState189, CExclusiveOrExpressionType}:    _CGotoState272Action,
	{_CState190, CIdentifierToken}:              _CGotoState139Action,
	{_CState190, CConstantToken}:                _CGotoState86Action,
	{_CState190, CStringLiteralToken}:           _CGotoState98Action,
	{_CState190, CSizeofToken}:                  _CGotoState97Action,
	{_CState190, CIncOpToken}:                   _CGotoState95Action,
	{_CState190, CDecOpToken}:                   _CGotoState88Action,
	{_CState190, '('}:                           _CGotoState77Action,
	{_CState190, '*'}:                           _CGotoState78Action,
	{_CState190, '-'}:                           _CGotoState80Action,
	{_CState190, '+'}:                           _CGotoState79Action,
	{_CState190, '&'}:                           _CGotoState76Action,
	{_CState190, '!'}:                           _CGotoState75Action,
	{_CState190, '~'}:                           _CGotoState83Action,
	{_CState190, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState190, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState190, CUnaryExpressionType}:          _CGotoState142Action,
	{_CState190, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState190, CCastExpressionType}:           _CGotoState104Action,
	{_CState190, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState190, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState190, CShiftExpressionType}:          _CGotoState123Action,
	{_CState190, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState190, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState190, CAndExpressionType}:            _CGotoState102Action,
	{_CState190, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState190, CInclusiveOrExpressionType}:    _CGotoState273Action,
	{_CState191, CIdentifierToken}:              _CGotoState139Action,
	{_CState191, CConstantToken}:                _CGotoState86Action,
	{_CState191, CStringLiteralToken}:           _CGotoState98Action,
	{_CState191, CSizeofToken}:                  _CGotoState97Action,
	{_CState191, CIncOpToken}:                   _CGotoState95Action,
	{_CState191, CDecOpToken}:                   _CGotoState88Action,
	{_CState191, '('}:                           _CGotoState77Action,
	{_CState191, '*'}:                           _CGotoState78Action,
	{_CState191, '-'}:                           _CGotoState80Action,
	{_CState191, '+'}:                           _CGotoState79Action,
	{_CState191, '&'}:                           _CGotoState76Action,
	{_CState191, '!'}:                           _CGotoState75Action,
	{_CState191, '~'}:                           _CGotoState83Action,
	{_CState191, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState191, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState191, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState191, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState191, CCastExpressionType}:           _CGotoState104Action,
	{_CState191, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState191, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState191, CShiftExpressionType}:          _CGotoState123Action,
	{_CState191, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState191, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState191, CAndExpressionType}:            _CGotoState102Action,
	{_CState191, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState191, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState191, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState191, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState191, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState191, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState191, CExpressionType}:               _CGotoState274Action,
	{_CState192, CIdentifierToken}:              _CGotoState139Action,
	{_CState192, CConstantToken}:                _CGotoState86Action,
	{_CState192, CStringLiteralToken}:           _CGotoState98Action,
	{_CState192, CSizeofToken}:                  _CGotoState97Action,
	{_CState192, CIncOpToken}:                   _CGotoState95Action,
	{_CState192, CDecOpToken}:                   _CGotoState88Action,
	{_CState192, '('}:                           _CGotoState77Action,
	{_CState192, '*'}:                           _CGotoState78Action,
	{_CState192, '-'}:                           _CGotoState80Action,
	{_CState192, '+'}:                           _CGotoState79Action,
	{_CState192, '&'}:                           _CGotoState76Action,
	{_CState192, '!'}:                           _CGotoState75Action,
	{_CState192, '~'}:                           _CGotoState83Action,
	{_CState192, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState192, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState192, CUnaryExpressionType}:          _CGotoState142Action,
	{_CState192, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState192, CCastExpressionType}:           _CGotoState104Action,
	{_CState192, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState192, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState192, CShiftExpressionType}:          _CGotoState123Action,
	{_CState192, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState192, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState192, CAndExpressionType}:            _CGotoState102Action,
	{_CState192, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState192, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState192, CLogicalAndExpressionType}:     _CGotoState275Action,
	{_CState193, CIdentifierToken}:              _CGotoState139Action,
	{_CState193, CConstantToken}:                _CGotoState86Action,
	{_CState193, CStringLiteralToken}:           _CGotoState98Action,
	{_CState193, CSizeofToken}:                  _CGotoState97Action,
	{_CState193, CIncOpToken}:                   _CGotoState95Action,
	{_CState193, CDecOpToken}:                   _CGotoState88Action,
	{_CState193, '('}:                           _CGotoState77Action,
	{_CState193, '*'}:                           _CGotoState78Action,
	{_CState193, '-'}:                           _CGotoState80Action,
	{_CState193, '+'}:                           _CGotoState79Action,
	{_CState193, '&'}:                           _CGotoState76Action,
	{_CState193, '!'}:                           _CGotoState75Action,
	{_CState193, '~'}:                           _CGotoState83Action,
	{_CState193, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState193, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState193, CUnaryExpressionType}:          _CGotoState142Action,
	{_CState193, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState193, CCastExpressionType}:           _CGotoState276Action,
	{_CState194, CIdentifierToken}:              _CGotoState139Action,
	{_CState194, CConstantToken}:                _CGotoState86Action,
	{_CState194, CStringLiteralToken}:           _CGotoState98Action,
	{_CState194, CSizeofToken}:                  _CGotoState97Action,
	{_CState194, CIncOpToken}:                   _CGotoState95Action,
	{_CState194, CDecOpToken}:                   _CGotoState88Action,
	{_CState194, '('}:                           _CGotoState77Action,
	{_CState194, '*'}:                           _CGotoState78Action,
	{_CState194, '-'}:                           _CGotoState80Action,
	{_CState194, '+'}:                           _CGotoState79Action,
	{_CState194, '&'}:                           _CGotoState76Action,
	{_CState194, '!'}:                           _CGotoState75Action,
	{_CState194, '~'}:                           _CGotoState83Action,
	{_CState194, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState194, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState194, CUnaryExpressionType}:          _CGotoState142Action,
	{_CState194, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState194, CCastExpressionType}:           _CGotoState277Action,
	{_CState195, CIdentifierToken}:              _CGotoState139Action,
	{_CState195, CConstantToken}:                _CGotoState86Action,
	{_CState195, CStringLiteralToken}:           _CGotoState98Action,
	{_CState195, CSizeofToken}:                  _CGotoState97Action,
	{_CState195, CIncOpToken}:                   _CGotoState95Action,
	{_CState195, CDecOpToken}:                   _CGotoState88Action,
	{_CState195, '('}:                           _CGotoState77Action,
	{_CState195, '*'}:                           _CGotoState78Action,
	{_CState195, '-'}:                           _CGotoState80Action,
	{_CState195, '+'}:                           _CGotoState79Action,
	{_CState195, '&'}:                           _CGotoState76Action,
	{_CState195, '!'}:                           _CGotoState75Action,
	{_CState195, '~'}:                           _CGotoState83Action,
	{_CState195, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState195, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState195, CUnaryExpressionType}:          _CGotoState142Action,
	{_CState195, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState195, CCastExpressionType}:           _CGotoState278Action,
	{_CState196, CIdentifierToken}:              _CGotoState139Action,
	{_CState196, CConstantToken}:                _CGotoState86Action,
	{_CState196, CStringLiteralToken}:           _CGotoState98Action,
	{_CState196, CSizeofToken}:                  _CGotoState97Action,
	{_CState196, CIncOpToken}:                   _CGotoState95Action,
	{_CState196, CDecOpToken}:                   _CGotoState88Action,
	{_CState196, '('}:                           _CGotoState77Action,
	{_CState196, ')'}:                           _CGotoState279Action,
	{_CState196, '*'}:                           _CGotoState78Action,
	{_CState196, '-'}:                           _CGotoState80Action,
	{_CState196, '+'}:                           _CGotoState79Action,
	{_CState196, '&'}:                           _CGotoState76Action,
	{_CState196, '!'}:                           _CGotoState75Action,
	{_CState196, '~'}:                           _CGotoState83Action,
	{_CState196, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState196, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState196, CArgumentExpressionListType}:   _CGotoState280Action,
	{_CState196, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState196, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState196, CCastExpressionType}:           _CGotoState104Action,
	{_CState196, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState196, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState196, CShiftExpressionType}:          _CGotoState123Action,
	{_CState196, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState196, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState196, CAndExpressionType}:            _CGotoState102Action,
	{_CState196, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState196, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState196, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState196, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState196, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState196, CAssignmentExpressionType}:     _CGotoState281Action,
	{_CState197, CIdentifierToken}:              _CGotoState282Action,
	{_CState198, CIdentifierToken}:              _CGotoState139Action,
	{_CState198, CConstantToken}:                _CGotoState86Action,
	{_CState198, CStringLiteralToken}:           _CGotoState98Action,
	{_CState198, CSizeofToken}:                  _CGotoState97Action,
	{_CState198, CIncOpToken}:                   _CGotoState95Action,
	{_CState198, CDecOpToken}:                   _CGotoState88Action,
	{_CState198, '('}:                           _CGotoState77Action,
	{_CState198, '*'}:                           _CGotoState78Action,
	{_CState198, '-'}:                           _CGotoState80Action,
	{_CState198, '+'}:                           _CGotoState79Action,
	{_CState198, '&'}:                           _CGotoState76Action,
	{_CState198, '!'}:                           _CGotoState75Action,
	{_CState198, '~'}:                           _CGotoState83Action,
	{_CState198, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState198, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState198, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState198, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState198, CCastExpressionType}:           _CGotoState104Action,
	{_CState198, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState198, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState198, CShiftExpressionType}:          _CGotoState123Action,
	{_CState198, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState198, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState198, CAndExpressionType}:            _CGotoState102Action,
	{_CState198, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState198, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState198, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState198, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState198, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState198, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState198, CExpressionType}:               _CGotoState283Action,
	{_CState201, CIdentifierToken}:              _CGotoState284Action,
	{_CState202, CIdentifierToken}:              _CGotoState139Action,
	{_CState202, CConstantToken}:                _CGotoState86Action,
	{_CState202, CStringLiteralToken}:           _CGotoState98Action,
	{_CState202, CSizeofToken}:                  _CGotoState97Action,
	{_CState202, CIncOpToken}:                   _CGotoState95Action,
	{_CState202, CDecOpToken}:                   _CGotoState88Action,
	{_CState202, '('}:                           _CGotoState77Action,
	{_CState202, '*'}:                           _CGotoState78Action,
	{_CState202, '-'}:                           _CGotoState80Action,
	{_CState202, '+'}:                           _CGotoState79Action,
	{_CState202, '&'}:                           _CGotoState76Action,
	{_CState202, '!'}:                           _CGotoState75Action,
	{_CState202, '~'}:                           _CGotoState83Action,
	{_CState202, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState202, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState202, CUnaryExpressionType}:          _CGotoState142Action,
	{_CState202, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState202, CCastExpressionType}:           _CGotoState104Action,
	{_CState202, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState202, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState202, CShiftExpressionType}:          _CGotoState285Action,
	{_CState203, CIdentifierToken}:              _CGotoState139Action,
	{_CState203, CConstantToken}:                _CGotoState86Action,
	{_CState203, CStringLiteralToken}:           _CGotoState98Action,
	{_CState203, CSizeofToken}:                  _CGotoState97Action,
	{_CState203, CIncOpToken}:                   _CGotoState95Action,
	{_CState203, CDecOpToken}:                   _CGotoState88Action,
	{_CState203, '('}:                           _CGotoState77Action,
	{_CState203, '*'}:                           _CGotoState78Action,
	{_CState203, '-'}:                           _CGotoState80Action,
	{_CState203, '+'}:                           _CGotoState79Action,
	{_CState203, '&'}:                           _CGotoState76Action,
	{_CState203, '!'}:                           _CGotoState75Action,
	{_CState203, '~'}:                           _CGotoState83Action,
	{_CState203, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState203, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState203, CUnaryExpressionType}:          _CGotoState142Action,
	{_CState203, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState203, CCastExpressionType}:           _CGotoState104Action,
	{_CState203, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState203, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState203, CShiftExpressionType}:          _CGotoState286Action,
	{_CState204, CIdentifierToken}:              _CGotoState139Action,
	{_CState204, CConstantToken}:                _CGotoState86Action,
	{_CState204, CStringLiteralToken}:           _CGotoState98Action,
	{_CState204, CSizeofToken}:                  _CGotoState97Action,
	{_CState204, CIncOpToken}:                   _CGotoState95Action,
	{_CState204, CDecOpToken}:                   _CGotoState88Action,
	{_CState204, '('}:                           _CGotoState77Action,
	{_CState204, '*'}:                           _CGotoState78Action,
	{_CState204, '-'}:                           _CGotoState80Action,
	{_CState204, '+'}:                           _CGotoState79Action,
	{_CState204, '&'}:                           _CGotoState76Action,
	{_CState204, '!'}:                           _CGotoState75Action,
	{_CState204, '~'}:                           _CGotoState83Action,
	{_CState204, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState204, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState204, CUnaryExpressionType}:          _CGotoState142Action,
	{_CState204, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState204, CCastExpressionType}:           _CGotoState104Action,
	{_CState204, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState204, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState204, CShiftExpressionType}:          _CGotoState287Action,
	{_CState205, CIdentifierToken}:              _CGotoState139Action,
	{_CState205, CConstantToken}:                _CGotoState86Action,
	{_CState205, CStringLiteralToken}:           _CGotoState98Action,
	{_CState205, CSizeofToken}:                  _CGotoState97Action,
	{_CState205, CIncOpToken}:                   _CGotoState95Action,
	{_CState205, CDecOpToken}:                   _CGotoState88Action,
	{_CState205, '('}:                           _CGotoState77Action,
	{_CState205, '*'}:                           _CGotoState78Action,
	{_CState205, '-'}:                           _CGotoState80Action,
	{_CState205, '+'}:                           _CGotoState79Action,
	{_CState205, '&'}:                           _CGotoState76Action,
	{_CState205, '!'}:                           _CGotoState75Action,
	{_CState205, '~'}:                           _CGotoState83Action,
	{_CState205, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState205, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState205, CUnaryExpressionType}:          _CGotoState142Action,
	{_CState205, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState205, CCastExpressionType}:           _CGotoState104Action,
	{_CState205, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState205, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState205, CShiftExpressionType}:          _CGotoState288Action,
	{_CState206, CIdentifierToken}:              _CGotoState139Action,
	{_CState206, CConstantToken}:                _CGotoState86Action,
	{_CState206, CStringLiteralToken}:           _CGotoState98Action,
	{_CState206, CSizeofToken}:                  _CGotoState97Action,
	{_CState206, CIncOpToken}:                   _CGotoState95Action,
	{_CState206, CDecOpToken}:                   _CGotoState88Action,
	{_CState206, '('}:                           _CGotoState77Action,
	{_CState206, '*'}:                           _CGotoState78Action,
	{_CState206, '-'}:                           _CGotoState80Action,
	{_CState206, '+'}:                           _CGotoState79Action,
	{_CState206, '&'}:                           _CGotoState76Action,
	{_CState206, '!'}:                           _CGotoState75Action,
	{_CState206, '~'}:                           _CGotoState83Action,
	{_CState206, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState206, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState206, CUnaryExpressionType}:          _CGotoState142Action,
	{_CState206, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState206, CCastExpressionType}:           _CGotoState104Action,
	{_CState206, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState206, CAdditiveExpressionType}:       _CGotoState289Action,
	{_CState207, CIdentifierToken}:              _CGotoState139Action,
	{_CState207, CConstantToken}:                _CGotoState86Action,
	{_CState207, CStringLiteralToken}:           _CGotoState98Action,
	{_CState207, CSizeofToken}:                  _CGotoState97Action,
	{_CState207, CIncOpToken}:                   _CGotoState95Action,
	{_CState207, CDecOpToken}:                   _CGotoState88Action,
	{_CState207, '('}:                           _CGotoState77Action,
	{_CState207, '*'}:                           _CGotoState78Action,
	{_CState207, '-'}:                           _CGotoState80Action,
	{_CState207, '+'}:                           _CGotoState79Action,
	{_CState207, '&'}:                           _CGotoState76Action,
	{_CState207, '!'}:                           _CGotoState75Action,
	{_CState207, '~'}:                           _CGotoState83Action,
	{_CState207, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState207, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState207, CUnaryExpressionType}:          _CGotoState142Action,
	{_CState207, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState207, CCastExpressionType}:           _CGotoState104Action,
	{_CState207, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState207, CAdditiveExpressionType}:       _CGotoState290Action,
	{_CState221, CIdentifierToken}:              _CGotoState139Action,
	{_CState221, CConstantToken}:                _CGotoState86Action,
	{_CState221, CStringLiteralToken}:           _CGotoState98Action,
	{_CState221, CSizeofToken}:                  _CGotoState97Action,
	{_CState221, CIncOpToken}:                   _CGotoState95Action,
	{_CState221, CDecOpToken}:                   _CGotoState88Action,
	{_CState221, '('}:                           _CGotoState77Action,
	{_CState221, '*'}:                           _CGotoState78Action,
	{_CState221, '-'}:                           _CGotoState80Action,
	{_CState221, '+'}:                           _CGotoState79Action,
	{_CState221, '&'}:                           _CGotoState76Action,
	{_CState221, '!'}:                           _CGotoState75Action,
	{_CState221, '~'}:                           _CGotoState83Action,
	{_CState221, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState221, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState221, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState221, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState221, CCastExpressionType}:           _CGotoState104Action,
	{_CState221, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState221, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState221, CShiftExpressionType}:          _CGotoState123Action,
	{_CState221, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState221, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState221, CAndExpressionType}:            _CGotoState102Action,
	{_CState221, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState221, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState221, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState221, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState221, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState221, CAssignmentExpressionType}:     _CGotoState291Action,
	{_CState223, CIdentifierToken}:              _CGotoState12Action,
	{_CState223, CTypeNameToken}:                _CGotoState21Action,
	{_CState223, CTypedefToken}:                 _CGotoState20Action,
	{_CState223, CExternToken}:                  _CGotoState10Action,
	{_CState223, CStaticToken}:                  _CGotoState18Action,
	{_CState223, CAutoToken}:                    _CGotoState5Action,
	{_CState223, CRegisterToken}:                _CGotoState15Action,
	{_CState223, CCharToken}:                    _CGotoState6Action,
	{_CState223, CShortToken}:                   _CGotoState16Action,
	{_CState223, CIntToken}:                     _CGotoState13Action,
	{_CState223, CLongToken}:                    _CGotoState14Action,
	{_CState223, CSignedToken}:                  _CGotoState17Action,
	{_CState223, CUnsignedToken}:                _CGotoState23Action,
	{_CState223, CFloatToken}:                   _CGotoState11Action,
	{_CState223, CDoubleToken}:                  _CGotoState8Action,
	{_CState223, CConstToken}:                   _CGotoState7Action,
	{_CState223, CVolatileToken}:                _CGotoState25Action,
	{_CState223, CVoidToken}:                    _CGotoState24Action,
	{_CState223, CStructToken}:                  _CGotoState19Action,
	{_CState223, CUnionToken}:                   _CGotoState22Action,
	{_CState223, CEnumToken}:                    _CGotoState9Action,
	{_CState223, '('}:                           _CGotoState223Action,
	{_CState223, ')'}:                           _CGotoState292Action,
	{_CState223, '['}:                           _CGotoState224Action,
	{_CState223, '*'}:                           _CGotoState4Action,
	{_CState223, CDeclarationSpecifiersType}:    _CGotoState133Action,
	{_CState223, CStorageClassSpecifierType}:    _CGotoState34Action,
	{_CState223, CTypeSpecifierType}:            _CGotoState38Action,
	{_CState223, CStructOrUnionSpecifierType}:   _CGotoState36Action,
	{_CState223, CStructOrUnionType}:            _CGotoState35Action,
	{_CState223, CEnumSpecifierType}:            _CGotoState30Action,
	{_CState223, CTypeQualifierType}:            _CGotoState37Action,
	{_CState223, CDeclaratorType}:               _CGotoState39Action,
	{_CState223, CDirectDeclaratorType}:         _CGotoState29Action,
	{_CState223, CPointerType}:                  _CGotoState228Action,
	{_CState223, CParameterTypeListType}:        _CGotoState294Action,
	{_CState223, CParameterListType}:            _CGotoState136Action,
	{_CState223, CParameterDeclarationType}:     _CGotoState135Action,
	{_CState223, CAbstractDeclaratorType}:       _CGotoState293Action,
	{_CState223, CDirectAbstractDeclaratorType}: _CGotoState227Action,
	{_CState224, CIdentifierToken}:              _CGotoState139Action,
	{_CState224, CConstantToken}:                _CGotoState86Action,
	{_CState224, CStringLiteralToken}:           _CGotoState98Action,
	{_CState224, CSizeofToken}:                  _CGotoState97Action,
	{_CState224, CIncOpToken}:                   _CGotoState95Action,
	{_CState224, CDecOpToken}:                   _CGotoState88Action,
	{_CState224, '('}:                           _CGotoState77Action,
	{_CState224, ']'}:                           _CGotoState295Action,
	{_CState224, '*'}:                           _CGotoState78Action,
	{_CState224, '-'}:                           _CGotoState80Action,
	{_CState224, '+'}:                           _CGotoState79Action,
	{_CState224, '&'}:                           _CGotoState76Action,
	{_CState224, '!'}:                           _CGotoState75Action,
	{_CState224, '~'}:                           _CGotoState83Action,
	{_CState224, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState224, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState224, CUnaryExpressionType}:          _CGotoState142Action,
	{_CState224, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState224, CCastExpressionType}:           _CGotoState104Action,
	{_CState224, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState224, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState224, CShiftExpressionType}:          _CGotoState123Action,
	{_CState224, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState224, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState224, CAndExpressionType}:            _CGotoState102Action,
	{_CState224, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState224, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState224, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState224, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState224, CConditionalExpressionType}:    _CGotoState140Action,
	{_CState224, CConstantExpressionType}:       _CGotoState296Action,
	{_CState227, '('}:                           _CGotoState297Action,
	{_CState227, '['}:                           _CGotoState298Action,
	{_CState228, CIdentifierToken}:              _CGotoState12Action,
	{_CState228, '('}:                           _CGotoState223Action,
	{_CState228, '['}:                           _CGotoState224Action,
	{_CState228, CDirectDeclaratorType}:         _CGotoState56Action,
	{_CState228, CDirectAbstractDeclaratorType}: _CGotoState299Action,
	{_CState230, CIdentifierToken}:              _CGotoState300Action,
	{_CState231, CTypeNameToken}:                _CGotoState21Action,
	{_CState231, CTypedefToken}:                 _CGotoState20Action,
	{_CState231, CExternToken}:                  _CGotoState10Action,
	{_CState231, CStaticToken}:                  _CGotoState18Action,
	{_CState231, CAutoToken}:                    _CGotoState5Action,
	{_CState231, CRegisterToken}:                _CGotoState15Action,
	{_CState231, CCharToken}:                    _CGotoState6Action,
	{_CState231, CShortToken}:                   _CGotoState16Action,
	{_CState231, CIntToken}:                     _CGotoState13Action,
	{_CState231, CLongToken}:                    _CGotoState14Action,
	{_CState231, CSignedToken}:                  _CGotoState17Action,
	{_CState231, CUnsignedToken}:                _CGotoState23Action,
	{_CState231, CFloatToken}:                   _CGotoState11Action,
	{_CState231, CDoubleToken}:                  _CGotoState8Action,
	{_CState231, CConstToken}:                   _CGotoState7Action,
	{_CState231, CVolatileToken}:                _CGotoState25Action,
	{_CState231, CVoidToken}:                    _CGotoState24Action,
	{_CState231, CStructToken}:                  _CGotoState19Action,
	{_CState231, CUnionToken}:                   _CGotoState22Action,
	{_CState231, CEnumToken}:                    _CGotoState9Action,
	{_CState231, CEllipsisToken}:                _CGotoState301Action,
	{_CState231, CDeclarationSpecifiersType}:    _CGotoState133Action,
	{_CState231, CStorageClassSpecifierType}:    _CGotoState34Action,
	{_CState231, CTypeSpecifierType}:            _CGotoState38Action,
	{_CState231, CStructOrUnionSpecifierType}:   _CGotoState36Action,
	{_CState231, CStructOrUnionType}:            _CGotoState35Action,
	{_CState231, CEnumSpecifierType}:            _CGotoState30Action,
	{_CState231, CTypeQualifierType}:            _CGotoState37Action,
	{_CState231, CParameterDeclarationType}:     _CGotoState302Action,
	{_CState234, CIdentifierToken}:              _CGotoState139Action,
	{_CState234, CConstantToken}:                _CGotoState86Action,
	{_CState234, CStringLiteralToken}:           _CGotoState98Action,
	{_CState234, CSizeofToken}:                  _CGotoState97Action,
	{_CState234, CIncOpToken}:                   _CGotoState95Action,
	{_CState234, CDecOpToken}:                   _CGotoState88Action,
	{_CState234, '('}:                           _CGotoState77Action,
	{_CState234, '*'}:                           _CGotoState78Action,
	{_CState234, '-'}:                           _CGotoState80Action,
	{_CState234, '+'}:                           _CGotoState79Action,
	{_CState234, '&'}:                           _CGotoState76Action,
	{_CState234, '!'}:                           _CGotoState75Action,
	{_CState234, '~'}:                           _CGotoState83Action,
	{_CState234, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState234, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState234, CUnaryExpressionType}:          _CGotoState142Action,
	{_CState234, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState234, CCastExpressionType}:           _CGotoState104Action,
	{_CState234, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState234, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState234, CShiftExpressionType}:          _CGotoState123Action,
	{_CState234, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState234, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState234, CAndExpressionType}:            _CGotoState102Action,
	{_CState234, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState234, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState234, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState234, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState234, CConditionalExpressionType}:    _CGotoState140Action,
	{_CState234, CConstantExpressionType}:       _CGotoState303Action,
	{_CState235, ':'}:                           _CGotoState304Action,
	{_CState237, ';'}:                           _CGotoState306Action,
	{_CState237, ','}:                           _CGotoState305Action,
	{_CState242, CTypeNameToken}:                _CGotoState21Action,
	{_CState242, CCharToken}:                    _CGotoState6Action,
	{_CState242, CShortToken}:                   _CGotoState16Action,
	{_CState242, CIntToken}:                     _CGotoState13Action,
	{_CState242, CLongToken}:                    _CGotoState14Action,
	{_CState242, CSignedToken}:                  _CGotoState17Action,
	{_CState242, CUnsignedToken}:                _CGotoState23Action,
	{_CState242, CFloatToken}:                   _CGotoState11Action,
	{_CState242, CDoubleToken}:                  _CGotoState8Action,
	{_CState242, CConstToken}:                   _CGotoState7Action,
	{_CState242, CVolatileToken}:                _CGotoState25Action,
	{_CState242, CVoidToken}:                    _CGotoState24Action,
	{_CState242, CStructToken}:                  _CGotoState19Action,
	{_CState242, CUnionToken}:                   _CGotoState22Action,
	{_CState242, CEnumToken}:                    _CGotoState9Action,
	{_CState242, '}'}:                           _CGotoState307Action,
	{_CState242, CTypeSpecifierType}:            _CGotoState147Action,
	{_CState242, CStructOrUnionSpecifierType}:   _CGotoState36Action,
	{_CState242, CStructOrUnionType}:            _CGotoState35Action,
	{_CState242, CStructDeclarationType}:        _CGotoState239Action,
	{_CState242, CSpecifierQualifierListType}:   _CGotoState143Action,
	{_CState242, CEnumSpecifierType}:            _CGotoState30Action,
	{_CState242, CTypeQualifierType}:            _CGotoState146Action,
	{_CState247, '}'}:                           _CGotoState309Action,
	{_CState247, ','}:                           _CGotoState308Action,
	{_CState249, CTypeNameToken}:                _CGotoState21Action,
	{_CState249, CTypedefToken}:                 _CGotoState20Action,
	{_CState249, CExternToken}:                  _CGotoState10Action,
	{_CState249, CStaticToken}:                  _CGotoState18Action,
	{_CState249, CAutoToken}:                    _CGotoState5Action,
	{_CState249, CRegisterToken}:                _CGotoState15Action,
	{_CState249, CCharToken}:                    _CGotoState6Action,
	{_CState249, CShortToken}:                   _CGotoState16Action,
	{_CState249, CIntToken}:                     _CGotoState13Action,
	{_CState249, CLongToken}:                    _CGotoState14Action,
	{_CState249, CSignedToken}:                  _CGotoState17Action,
	{_CState249, CUnsignedToken}:                _CGotoState23Action,
	{_CState249, CFloatToken}:                   _CGotoState11Action,
	{_CState249, CDoubleToken}:                  _CGotoState8Action,
	{_CState249, CConstToken}:                   _CGotoState7Action,
	{_CState249, CVolatileToken}:                _CGotoState25Action,
	{_CState249, CVoidToken}:                    _CGotoState24Action,
	{_CState249, CStructToken}:                  _CGotoState19Action,
	{_CState249, CUnionToken}:                   _CGotoState22Action,
	{_CState249, CEnumToken}:                    _CGotoState9Action,
	{_CState249, '('}:                           _CGotoState249Action,
	{_CState249, ')'}:                           _CGotoState292Action,
	{_CState249, '['}:                           _CGotoState224Action,
	{_CState249, '*'}:                           _CGotoState4Action,
	{_CState249, CDeclarationSpecifiersType}:    _CGotoState133Action,
	{_CState249, CStorageClassSpecifierType}:    _CGotoState34Action,
	{_CState249, CTypeSpecifierType}:            _CGotoState38Action,
	{_CState249, CStructOrUnionSpecifierType}:   _CGotoState36Action,
	{_CState249, CStructOrUnionType}:            _CGotoState35Action,
	{_CState249, CEnumSpecifierType}:            _CGotoState30Action,
	{_CState249, CTypeQualifierType}:            _CGotoState37Action,
	{_CState249, CPointerType}:                  _CGotoState251Action,
	{_CState249, CParameterTypeListType}:        _CGotoState294Action,
	{_CState249, CParameterListType}:            _CGotoState136Action,
	{_CState249, CParameterDeclarationType}:     _CGotoState135Action,
	{_CState249, CAbstractDeclaratorType}:       _CGotoState293Action,
	{_CState249, CDirectAbstractDeclaratorType}: _CGotoState227Action,
	{_CState251, '('}:                           _CGotoState249Action,
	{_CState251, '['}:                           _CGotoState224Action,
	{_CState251, CDirectAbstractDeclaratorType}: _CGotoState299Action,
	{_CState252, CIdentifierToken}:              _CGotoState139Action,
	{_CState252, CConstantToken}:                _CGotoState86Action,
	{_CState252, CStringLiteralToken}:           _CGotoState98Action,
	{_CState252, CSizeofToken}:                  _CGotoState97Action,
	{_CState252, CIncOpToken}:                   _CGotoState95Action,
	{_CState252, CDecOpToken}:                   _CGotoState88Action,
	{_CState252, '('}:                           _CGotoState77Action,
	{_CState252, '*'}:                           _CGotoState78Action,
	{_CState252, '-'}:                           _CGotoState80Action,
	{_CState252, '+'}:                           _CGotoState79Action,
	{_CState252, '&'}:                           _CGotoState76Action,
	{_CState252, '!'}:                           _CGotoState75Action,
	{_CState252, '~'}:                           _CGotoState83Action,
	{_CState252, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState252, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState252, CUnaryExpressionType}:          _CGotoState142Action,
	{_CState252, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState252, CCastExpressionType}:           _CGotoState310Action,
	{_CState253, CIdentifierToken}:              _CGotoState93Action,
	{_CState253, CConstantToken}:                _CGotoState86Action,
	{_CState253, CStringLiteralToken}:           _CGotoState98Action,
	{_CState253, CSizeofToken}:                  _CGotoState97Action,
	{_CState253, CIncOpToken}:                   _CGotoState95Action,
	{_CState253, CDecOpToken}:                   _CGotoState88Action,
	{_CState253, CCaseToken}:                    _CGotoState85Action,
	{_CState253, CDefaultToken}:                 _CGotoState89Action,
	{_CState253, CIfToken}:                      _CGotoState94Action,
	{_CState253, CSwitchToken}:                  _CGotoState99Action,
	{_CState253, CWhileToken}:                   _CGotoState100Action,
	{_CState253, CDoToken}:                      _CGotoState90Action,
	{_CState253, CForToken}:                     _CGotoState91Action,
	{_CState253, CGotoToken}:                    _CGotoState92Action,
	{_CState253, CContinueToken}:                _CGotoState87Action,
	{_CState253, CBreakToken}:                   _CGotoState84Action,
	{_CState253, CReturnToken}:                  _CGotoState96Action,
	{_CState253, '('}:                           _CGotoState77Action,
	{_CState253, '{'}:                           _CGotoState49Action,
	{_CState253, ';'}:                           _CGotoState81Action,
	{_CState253, '*'}:                           _CGotoState78Action,
	{_CState253, '-'}:                           _CGotoState80Action,
	{_CState253, '+'}:                           _CGotoState79Action,
	{_CState253, '&'}:                           _CGotoState76Action,
	{_CState253, '!'}:                           _CGotoState75Action,
	{_CState253, '~'}:                           _CGotoState83Action,
	{_CState253, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState253, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState253, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState253, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState253, CCastExpressionType}:           _CGotoState104Action,
	{_CState253, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState253, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState253, CShiftExpressionType}:          _CGotoState123Action,
	{_CState253, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState253, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState253, CAndExpressionType}:            _CGotoState102Action,
	{_CState253, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState253, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState253, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState253, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState253, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState253, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState253, CExpressionType}:               _CGotoState110Action,
	{_CState253, CStatementType}:                _CGotoState311Action,
	{_CState253, CLabeledStatementType}:         _CGotoState115Action,
	{_CState253, CCompoundStatementType}:        _CGotoState105Action,
	{_CState253, CExpressionStatementType}:      _CGotoState111Action,
	{_CState253, CSelectionStatementType}:       _CGotoState122Action,
	{_CState253, CIterationStatementType}:       _CGotoState113Action,
	{_CState253, CJumpStatementType}:            _CGotoState114Action,
	{_CState255, '('}:                           _CGotoState312Action,
	{_CState256, CIdentifierToken}:              _CGotoState139Action,
	{_CState256, CConstantToken}:                _CGotoState86Action,
	{_CState256, CStringLiteralToken}:           _CGotoState98Action,
	{_CState256, CSizeofToken}:                  _CGotoState97Action,
	{_CState256, CIncOpToken}:                   _CGotoState95Action,
	{_CState256, CDecOpToken}:                   _CGotoState88Action,
	{_CState256, '('}:                           _CGotoState77Action,
	{_CState256, ';'}:                           _CGotoState81Action,
	{_CState256, '*'}:                           _CGotoState78Action,
	{_CState256, '-'}:                           _CGotoState80Action,
	{_CState256, '+'}:                           _CGotoState79Action,
	{_CState256, '&'}:                           _CGotoState76Action,
	{_CState256, '!'}:                           _CGotoState75Action,
	{_CState256, '~'}:                           _CGotoState83Action,
	{_CState256, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState256, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState256, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState256, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState256, CCastExpressionType}:           _CGotoState104Action,
	{_CState256, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState256, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState256, CShiftExpressionType}:          _CGotoState123Action,
	{_CState256, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState256, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState256, CAndExpressionType}:            _CGotoState102Action,
	{_CState256, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState256, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState256, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState256, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState256, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState256, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState256, CExpressionType}:               _CGotoState110Action,
	{_CState256, CExpressionStatementType}:      _CGotoState313Action,
	{_CState259, ')'}:                           _CGotoState314Action,
	{_CState259, ','}:                           _CGotoState187Action,
	{_CState261, ')'}:                           _CGotoState315Action,
	{_CState262, ')'}:                           _CGotoState316Action,
	{_CState262, ','}:                           _CGotoState187Action,
	{_CState263, ')'}:                           _CGotoState317Action,
	{_CState263, ','}:                           _CGotoState187Action,
	{_CState264, '*'}:                           _CGotoState194Action,
	{_CState264, '/'}:                           _CGotoState195Action,
	{_CState264, '%'}:                           _CGotoState193Action,
	{_CState265, '*'}:                           _CGotoState194Action,
	{_CState265, '/'}:                           _CGotoState195Action,
	{_CState265, '%'}:                           _CGotoState193Action,
	{_CState266, CEqOpToken}:                    _CGotoState184Action,
	{_CState266, CNeOpToken}:                    _CGotoState185Action,
	{_CState268, CLeOpToken}:                    _CGotoState205Action,
	{_CState268, CGeOpToken}:                    _CGotoState204Action,
	{_CState268, '<'}:                           _CGotoState202Action,
	{_CState268, '>'}:                           _CGotoState203Action,
	{_CState269, CLeOpToken}:                    _CGotoState205Action,
	{_CState269, CGeOpToken}:                    _CGotoState204Action,
	{_CState269, '<'}:                           _CGotoState202Action,
	{_CState269, '>'}:                           _CGotoState203Action,
	{_CState270, '&'}:                           _CGotoState181Action,
	{_CState272, '^'}:                           _CGotoState186Action,
	{_CState273, '|'}:                           _CGotoState189Action,
	{_CState274, ':'}:                           _CGotoState318Action,
	{_CState274, ','}:                           _CGotoState187Action,
	{_CState275, CAndOpToken}:                   _CGotoState190Action,
	{_CState280, ')'}:                           _CGotoState319Action,
	{_CState280, ','}:                           _CGotoState320Action,
	{_CState283, ']'}:                           _CGotoState321Action,
	{_CState283, ','}:                           _CGotoState187Action,
	{_CState285, CLeftOpToken}:                  _CGotoState206Action,
	{_CState285, CRightOpToken}:                 _CGotoState207Action,
	{_CState286, CLeftOpToken}:                  _CGotoState206Action,
	{_CState286, CRightOpToken}:                 _CGotoState207Action,
	{_CState287, CLeftOpToken}:                  _CGotoState206Action,
	{_CState287, CRightOpToken}:                 _CGotoState207Action,
	{_CState288, CLeftOpToken}:                  _CGotoState206Action,
	{_CState288, CRightOpToken}:                 _CGotoState207Action,
	{_CState289, '-'}:                           _CGotoState180Action,
	{_CState289, '+'}:                           _CGotoState179Action,
	{_CState290, '-'}:                           _CGotoState180Action,
	{_CState290, '+'}:                           _CGotoState179Action,
	{_CState293, ')'}:                           _CGotoState322Action,
	{_CState294, ')'}:                           _CGotoState323Action,
	{_CState296, ']'}:                           _CGotoState324Action,
	{_CState297, CTypeNameToken}:                _CGotoState21Action,
	{_CState297, CTypedefToken}:                 _CGotoState20Action,
	{_CState297, CExternToken}:                  _CGotoState10Action,
	{_CState297, CStaticToken}:                  _CGotoState18Action,
	{_CState297, CAutoToken}:                    _CGotoState5Action,
	{_CState297, CRegisterToken}:                _CGotoState15Action,
	{_CState297, CCharToken}:                    _CGotoState6Action,
	{_CState297, CShortToken}:                   _CGotoState16Action,
	{_CState297, CIntToken}:                     _CGotoState13Action,
	{_CState297, CLongToken}:                    _CGotoState14Action,
	{_CState297, CSignedToken}:                  _CGotoState17Action,
	{_CState297, CUnsignedToken}:                _CGotoState23Action,
	{_CState297, CFloatToken}:                   _CGotoState11Action,
	{_CState297, CDoubleToken}:                  _CGotoState8Action,
	{_CState297, CConstToken}:                   _CGotoState7Action,
	{_CState297, CVolatileToken}:                _CGotoState25Action,
	{_CState297, CVoidToken}:                    _CGotoState24Action,
	{_CState297, CStructToken}:                  _CGotoState19Action,
	{_CState297, CUnionToken}:                   _CGotoState22Action,
	{_CState297, CEnumToken}:                    _CGotoState9Action,
	{_CState297, ')'}:                           _CGotoState325Action,
	{_CState297, CDeclarationSpecifiersType}:    _CGotoState133Action,
	{_CState297, CStorageClassSpecifierType}:    _CGotoState34Action,
	{_CState297, CTypeSpecifierType}:            _CGotoState38Action,
	{_CState297, CStructOrUnionSpecifierType}:   _CGotoState36Action,
	{_CState297, CStructOrUnionType}:            _CGotoState35Action,
	{_CState297, CEnumSpecifierType}:            _CGotoState30Action,
	{_CState297, CTypeQualifierType}:            _CGotoState37Action,
	{_CState297, CParameterTypeListType}:        _CGotoState326Action,
	{_CState297, CParameterListType}:            _CGotoState136Action,
	{_CState297, CParameterDeclarationType}:     _CGotoState135Action,
	{_CState298, CIdentifierToken}:              _CGotoState139Action,
	{_CState298, CConstantToken}:                _CGotoState86Action,
	{_CState298, CStringLiteralToken}:           _CGotoState98Action,
	{_CState298, CSizeofToken}:                  _CGotoState97Action,
	{_CState298, CIncOpToken}:                   _CGotoState95Action,
	{_CState298, CDecOpToken}:                   _CGotoState88Action,
	{_CState298, '('}:                           _CGotoState77Action,
	{_CState298, ']'}:                           _CGotoState327Action,
	{_CState298, '*'}:                           _CGotoState78Action,
	{_CState298, '-'}:                           _CGotoState80Action,
	{_CState298, '+'}:                           _CGotoState79Action,
	{_CState298, '&'}:                           _CGotoState76Action,
	{_CState298, '!'}:                           _CGotoState75Action,
	{_CState298, '~'}:                           _CGotoState83Action,
	{_CState298, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState298, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState298, CUnaryExpressionType}:          _CGotoState142Action,
	{_CState298, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState298, CCastExpressionType}:           _CGotoState104Action,
	{_CState298, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState298, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState298, CShiftExpressionType}:          _CGotoState123Action,
	{_CState298, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState298, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState298, CAndExpressionType}:            _CGotoState102Action,
	{_CState298, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState298, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState298, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState298, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState298, CConditionalExpressionType}:    _CGotoState140Action,
	{_CState298, CConstantExpressionType}:       _CGotoState328Action,
	{_CState299, '('}:                           _CGotoState297Action,
	{_CState299, '['}:                           _CGotoState298Action,
	{_CState304, CIdentifierToken}:              _CGotoState139Action,
	{_CState304, CConstantToken}:                _CGotoState86Action,
	{_CState304, CStringLiteralToken}:           _CGotoState98Action,
	{_CState304, CSizeofToken}:                  _CGotoState97Action,
	{_CState304, CIncOpToken}:                   _CGotoState95Action,
	{_CState304, CDecOpToken}:                   _CGotoState88Action,
	{_CState304, '('}:                           _CGotoState77Action,
	{_CState304, '*'}:                           _CGotoState78Action,
	{_CState304, '-'}:                           _CGotoState80Action,
	{_CState304, '+'}:                           _CGotoState79Action,
	{_CState304, '&'}:                           _CGotoState76Action,
	{_CState304, '!'}:                           _CGotoState75Action,
	{_CState304, '~'}:                           _CGotoState83Action,
	{_CState304, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState304, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState304, CUnaryExpressionType}:          _CGotoState142Action,
	{_CState304, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState304, CCastExpressionType}:           _CGotoState104Action,
	{_CState304, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState304, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState304, CShiftExpressionType}:          _CGotoState123Action,
	{_CState304, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState304, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState304, CAndExpressionType}:            _CGotoState102Action,
	{_CState304, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState304, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState304, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState304, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState304, CConditionalExpressionType}:    _CGotoState140Action,
	{_CState304, CConstantExpressionType}:       _CGotoState329Action,
	{_CState305, CIdentifierToken}:              _CGotoState12Action,
	{_CState305, '('}:                           _CGotoState3Action,
	{_CState305, ':'}:                           _CGotoState234Action,
	{_CState305, '*'}:                           _CGotoState4Action,
	{_CState305, CStructDeclaratorType}:         _CGotoState330Action,
	{_CState305, CDeclaratorType}:               _CGotoState235Action,
	{_CState305, CDirectDeclaratorType}:         _CGotoState29Action,
	{_CState305, CPointerType}:                  _CGotoState33Action,
	{_CState308, CIdentifierToken}:              _CGotoState139Action,
	{_CState308, CConstantToken}:                _CGotoState86Action,
	{_CState308, CStringLiteralToken}:           _CGotoState98Action,
	{_CState308, CSizeofToken}:                  _CGotoState97Action,
	{_CState308, CIncOpToken}:                   _CGotoState95Action,
	{_CState308, CDecOpToken}:                   _CGotoState88Action,
	{_CState308, '('}:                           _CGotoState77Action,
	{_CState308, '{'}:                           _CGotoState153Action,
	{_CState308, '}'}:                           _CGotoState331Action,
	{_CState308, '*'}:                           _CGotoState78Action,
	{_CState308, '-'}:                           _CGotoState80Action,
	{_CState308, '+'}:                           _CGotoState79Action,
	{_CState308, '&'}:                           _CGotoState76Action,
	{_CState308, '!'}:                           _CGotoState75Action,
	{_CState308, '~'}:                           _CGotoState83Action,
	{_CState308, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState308, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState308, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState308, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState308, CCastExpressionType}:           _CGotoState104Action,
	{_CState308, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState308, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState308, CShiftExpressionType}:          _CGotoState123Action,
	{_CState308, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState308, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState308, CAndExpressionType}:            _CGotoState102Action,
	{_CState308, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState308, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState308, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState308, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState308, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState308, CAssignmentExpressionType}:     _CGotoState154Action,
	{_CState308, CInitializerType}:              _CGotoState332Action,
	{_CState312, CIdentifierToken}:              _CGotoState139Action,
	{_CState312, CConstantToken}:                _CGotoState86Action,
	{_CState312, CStringLiteralToken}:           _CGotoState98Action,
	{_CState312, CSizeofToken}:                  _CGotoState97Action,
	{_CState312, CIncOpToken}:                   _CGotoState95Action,
	{_CState312, CDecOpToken}:                   _CGotoState88Action,
	{_CState312, '('}:                           _CGotoState77Action,
	{_CState312, '*'}:                           _CGotoState78Action,
	{_CState312, '-'}:                           _CGotoState80Action,
	{_CState312, '+'}:                           _CGotoState79Action,
	{_CState312, '&'}:                           _CGotoState76Action,
	{_CState312, '!'}:                           _CGotoState75Action,
	{_CState312, '~'}:                           _CGotoState83Action,
	{_CState312, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState312, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState312, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState312, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState312, CCastExpressionType}:           _CGotoState104Action,
	{_CState312, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState312, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState312, CShiftExpressionType}:          _CGotoState123Action,
	{_CState312, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState312, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState312, CAndExpressionType}:            _CGotoState102Action,
	{_CState312, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState312, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState312, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState312, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState312, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState312, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState312, CExpressionType}:               _CGotoState333Action,
	{_CState313, CIdentifierToken}:              _CGotoState139Action,
	{_CState313, CConstantToken}:                _CGotoState86Action,
	{_CState313, CStringLiteralToken}:           _CGotoState98Action,
	{_CState313, CSizeofToken}:                  _CGotoState97Action,
	{_CState313, CIncOpToken}:                   _CGotoState95Action,
	{_CState313, CDecOpToken}:                   _CGotoState88Action,
	{_CState313, '('}:                           _CGotoState77Action,
	{_CState313, ')'}:                           _CGotoState334Action,
	{_CState313, '*'}:                           _CGotoState78Action,
	{_CState313, '-'}:                           _CGotoState80Action,
	{_CState313, '+'}:                           _CGotoState79Action,
	{_CState313, '&'}:                           _CGotoState76Action,
	{_CState313, '!'}:                           _CGotoState75Action,
	{_CState313, '~'}:                           _CGotoState83Action,
	{_CState313, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState313, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState313, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState313, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState313, CCastExpressionType}:           _CGotoState104Action,
	{_CState313, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState313, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState313, CShiftExpressionType}:          _CGotoState123Action,
	{_CState313, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState313, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState313, CAndExpressionType}:            _CGotoState102Action,
	{_CState313, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState313, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState313, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState313, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState313, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState313, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState313, CExpressionType}:               _CGotoState335Action,
	{_CState314, CIdentifierToken}:              _CGotoState93Action,
	{_CState314, CConstantToken}:                _CGotoState86Action,
	{_CState314, CStringLiteralToken}:           _CGotoState98Action,
	{_CState314, CSizeofToken}:                  _CGotoState97Action,
	{_CState314, CIncOpToken}:                   _CGotoState95Action,
	{_CState314, CDecOpToken}:                   _CGotoState88Action,
	{_CState314, CCaseToken}:                    _CGotoState85Action,
	{_CState314, CDefaultToken}:                 _CGotoState89Action,
	{_CState314, CIfToken}:                      _CGotoState94Action,
	{_CState314, CSwitchToken}:                  _CGotoState99Action,
	{_CState314, CWhileToken}:                   _CGotoState100Action,
	{_CState314, CDoToken}:                      _CGotoState90Action,
	{_CState314, CForToken}:                     _CGotoState91Action,
	{_CState314, CGotoToken}:                    _CGotoState92Action,
	{_CState314, CContinueToken}:                _CGotoState87Action,
	{_CState314, CBreakToken}:                   _CGotoState84Action,
	{_CState314, CReturnToken}:                  _CGotoState96Action,
	{_CState314, '('}:                           _CGotoState77Action,
	{_CState314, '{'}:                           _CGotoState49Action,
	{_CState314, ';'}:                           _CGotoState81Action,
	{_CState314, '*'}:                           _CGotoState78Action,
	{_CState314, '-'}:                           _CGotoState80Action,
	{_CState314, '+'}:                           _CGotoState79Action,
	{_CState314, '&'}:                           _CGotoState76Action,
	{_CState314, '!'}:                           _CGotoState75Action,
	{_CState314, '~'}:                           _CGotoState83Action,
	{_CState314, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState314, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState314, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState314, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState314, CCastExpressionType}:           _CGotoState104Action,
	{_CState314, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState314, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState314, CShiftExpressionType}:          _CGotoState123Action,
	{_CState314, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState314, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState314, CAndExpressionType}:            _CGotoState102Action,
	{_CState314, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState314, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState314, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState314, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState314, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState314, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState314, CExpressionType}:               _CGotoState110Action,
	{_CState314, CStatementType}:                _CGotoState338Action,
	{_CState314, CLabeledStatementType}:         _CGotoState115Action,
	{_CState314, CCompoundStatementType}:        _CGotoState105Action,
	{_CState314, CExpressionStatementType}:      _CGotoState111Action,
	{_CState314, CSelectionStatementType}:       _CGotoState122Action,
	{_CState314, CIterationStatementType}:       _CGotoState113Action,
	{_CState314, CJumpStatementType}:            _CGotoState114Action,
	{_CState316, CIdentifierToken}:              _CGotoState93Action,
	{_CState316, CConstantToken}:                _CGotoState86Action,
	{_CState316, CStringLiteralToken}:           _CGotoState98Action,
	{_CState316, CSizeofToken}:                  _CGotoState97Action,
	{_CState316, CIncOpToken}:                   _CGotoState95Action,
	{_CState316, CDecOpToken}:                   _CGotoState88Action,
	{_CState316, CCaseToken}:                    _CGotoState85Action,
	{_CState316, CDefaultToken}:                 _CGotoState89Action,
	{_CState316, CIfToken}:                      _CGotoState94Action,
	{_CState316, CSwitchToken}:                  _CGotoState99Action,
	{_CState316, CWhileToken}:                   _CGotoState100Action,
	{_CState316, CDoToken}:                      _CGotoState90Action,
	{_CState316, CForToken}:                     _CGotoState91Action,
	{_CState316, CGotoToken}:                    _CGotoState92Action,
	{_CState316, CContinueToken}:                _CGotoState87Action,
	{_CState316, CBreakToken}:                   _CGotoState84Action,
	{_CState316, CReturnToken}:                  _CGotoState96Action,
	{_CState316, '('}:                           _CGotoState77Action,
	{_CState316, '{'}:                           _CGotoState49Action,
	{_CState316, ';'}:                           _CGotoState81Action,
	{_CState316, '*'}:                           _CGotoState78Action,
	{_CState316, '-'}:                           _CGotoState80Action,
	{_CState316, '+'}:                           _CGotoState79Action,
	{_CState316, '&'}:                           _CGotoState76Action,
	{_CState316, '!'}:                           _CGotoState75Action,
	{_CState316, '~'}:                           _CGotoState83Action,
	{_CState316, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState316, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState316, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState316, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState316, CCastExpressionType}:           _CGotoState104Action,
	{_CState316, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState316, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState316, CShiftExpressionType}:          _CGotoState123Action,
	{_CState316, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState316, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState316, CAndExpressionType}:            _CGotoState102Action,
	{_CState316, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState316, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState316, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState316, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState316, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState316, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState316, CExpressionType}:               _CGotoState110Action,
	{_CState316, CStatementType}:                _CGotoState339Action,
	{_CState316, CLabeledStatementType}:         _CGotoState115Action,
	{_CState316, CCompoundStatementType}:        _CGotoState105Action,
	{_CState316, CExpressionStatementType}:      _CGotoState111Action,
	{_CState316, CSelectionStatementType}:       _CGotoState122Action,
	{_CState316, CIterationStatementType}:       _CGotoState113Action,
	{_CState316, CJumpStatementType}:            _CGotoState114Action,
	{_CState317, CIdentifierToken}:              _CGotoState93Action,
	{_CState317, CConstantToken}:                _CGotoState86Action,
	{_CState317, CStringLiteralToken}:           _CGotoState98Action,
	{_CState317, CSizeofToken}:                  _CGotoState97Action,
	{_CState317, CIncOpToken}:                   _CGotoState95Action,
	{_CState317, CDecOpToken}:                   _CGotoState88Action,
	{_CState317, CCaseToken}:                    _CGotoState85Action,
	{_CState317, CDefaultToken}:                 _CGotoState89Action,
	{_CState317, CIfToken}:                      _CGotoState94Action,
	{_CState317, CSwitchToken}:                  _CGotoState99Action,
	{_CState317, CWhileToken}:                   _CGotoState100Action,
	{_CState317, CDoToken}:                      _CGotoState90Action,
	{_CState317, CForToken}:                     _CGotoState91Action,
	{_CState317, CGotoToken}:                    _CGotoState92Action,
	{_CState317, CContinueToken}:                _CGotoState87Action,
	{_CState317, CBreakToken}:                   _CGotoState84Action,
	{_CState317, CReturnToken}:                  _CGotoState96Action,
	{_CState317, '('}:                           _CGotoState77Action,
	{_CState317, '{'}:                           _CGotoState49Action,
	{_CState317, ';'}:                           _CGotoState81Action,
	{_CState317, '*'}:                           _CGotoState78Action,
	{_CState317, '-'}:                           _CGotoState80Action,
	{_CState317, '+'}:                           _CGotoState79Action,
	{_CState317, '&'}:                           _CGotoState76Action,
	{_CState317, '!'}:                           _CGotoState75Action,
	{_CState317, '~'}:                           _CGotoState83Action,
	{_CState317, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState317, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState317, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState317, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState317, CCastExpressionType}:           _CGotoState104Action,
	{_CState317, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState317, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState317, CShiftExpressionType}:          _CGotoState123Action,
	{_CState317, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState317, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState317, CAndExpressionType}:            _CGotoState102Action,
	{_CState317, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState317, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState317, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState317, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState317, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState317, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState317, CExpressionType}:               _CGotoState110Action,
	{_CState317, CStatementType}:                _CGotoState340Action,
	{_CState317, CLabeledStatementType}:         _CGotoState115Action,
	{_CState317, CCompoundStatementType}:        _CGotoState105Action,
	{_CState317, CExpressionStatementType}:      _CGotoState111Action,
	{_CState317, CSelectionStatementType}:       _CGotoState122Action,
	{_CState317, CIterationStatementType}:       _CGotoState113Action,
	{_CState317, CJumpStatementType}:            _CGotoState114Action,
	{_CState318, CIdentifierToken}:              _CGotoState139Action,
	{_CState318, CConstantToken}:                _CGotoState86Action,
	{_CState318, CStringLiteralToken}:           _CGotoState98Action,
	{_CState318, CSizeofToken}:                  _CGotoState97Action,
	{_CState318, CIncOpToken}:                   _CGotoState95Action,
	{_CState318, CDecOpToken}:                   _CGotoState88Action,
	{_CState318, '('}:                           _CGotoState77Action,
	{_CState318, '*'}:                           _CGotoState78Action,
	{_CState318, '-'}:                           _CGotoState80Action,
	{_CState318, '+'}:                           _CGotoState79Action,
	{_CState318, '&'}:                           _CGotoState76Action,
	{_CState318, '!'}:                           _CGotoState75Action,
	{_CState318, '~'}:                           _CGotoState83Action,
	{_CState318, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState318, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState318, CUnaryExpressionType}:          _CGotoState142Action,
	{_CState318, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState318, CCastExpressionType}:           _CGotoState104Action,
	{_CState318, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState318, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState318, CShiftExpressionType}:          _CGotoState123Action,
	{_CState318, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState318, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState318, CAndExpressionType}:            _CGotoState102Action,
	{_CState318, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState318, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState318, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState318, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState318, CConditionalExpressionType}:    _CGotoState341Action,
	{_CState320, CIdentifierToken}:              _CGotoState139Action,
	{_CState320, CConstantToken}:                _CGotoState86Action,
	{_CState320, CStringLiteralToken}:           _CGotoState98Action,
	{_CState320, CSizeofToken}:                  _CGotoState97Action,
	{_CState320, CIncOpToken}:                   _CGotoState95Action,
	{_CState320, CDecOpToken}:                   _CGotoState88Action,
	{_CState320, '('}:                           _CGotoState77Action,
	{_CState320, '*'}:                           _CGotoState78Action,
	{_CState320, '-'}:                           _CGotoState80Action,
	{_CState320, '+'}:                           _CGotoState79Action,
	{_CState320, '&'}:                           _CGotoState76Action,
	{_CState320, '!'}:                           _CGotoState75Action,
	{_CState320, '~'}:                           _CGotoState83Action,
	{_CState320, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState320, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState320, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState320, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState320, CCastExpressionType}:           _CGotoState104Action,
	{_CState320, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState320, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState320, CShiftExpressionType}:          _CGotoState123Action,
	{_CState320, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState320, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState320, CAndExpressionType}:            _CGotoState102Action,
	{_CState320, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState320, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState320, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState320, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState320, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState320, CAssignmentExpressionType}:     _CGotoState342Action,
	{_CState326, ')'}:                           _CGotoState343Action,
	{_CState328, ']'}:                           _CGotoState344Action,
	{_CState333, ')'}:                           _CGotoState345Action,
	{_CState333, ','}:                           _CGotoState187Action,
	{_CState334, CIdentifierToken}:              _CGotoState93Action,
	{_CState334, CConstantToken}:                _CGotoState86Action,
	{_CState334, CStringLiteralToken}:           _CGotoState98Action,
	{_CState334, CSizeofToken}:                  _CGotoState97Action,
	{_CState334, CIncOpToken}:                   _CGotoState95Action,
	{_CState334, CDecOpToken}:                   _CGotoState88Action,
	{_CState334, CCaseToken}:                    _CGotoState85Action,
	{_CState334, CDefaultToken}:                 _CGotoState89Action,
	{_CState334, CIfToken}:                      _CGotoState94Action,
	{_CState334, CSwitchToken}:                  _CGotoState99Action,
	{_CState334, CWhileToken}:                   _CGotoState100Action,
	{_CState334, CDoToken}:                      _CGotoState90Action,
	{_CState334, CForToken}:                     _CGotoState91Action,
	{_CState334, CGotoToken}:                    _CGotoState92Action,
	{_CState334, CContinueToken}:                _CGotoState87Action,
	{_CState334, CBreakToken}:                   _CGotoState84Action,
	{_CState334, CReturnToken}:                  _CGotoState96Action,
	{_CState334, '('}:                           _CGotoState77Action,
	{_CState334, '{'}:                           _CGotoState49Action,
	{_CState334, ';'}:                           _CGotoState81Action,
	{_CState334, '*'}:                           _CGotoState78Action,
	{_CState334, '-'}:                           _CGotoState80Action,
	{_CState334, '+'}:                           _CGotoState79Action,
	{_CState334, '&'}:                           _CGotoState76Action,
	{_CState334, '!'}:                           _CGotoState75Action,
	{_CState334, '~'}:                           _CGotoState83Action,
	{_CState334, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState334, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState334, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState334, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState334, CCastExpressionType}:           _CGotoState104Action,
	{_CState334, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState334, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState334, CShiftExpressionType}:          _CGotoState123Action,
	{_CState334, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState334, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState334, CAndExpressionType}:            _CGotoState102Action,
	{_CState334, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState334, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState334, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState334, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState334, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState334, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState334, CExpressionType}:               _CGotoState110Action,
	{_CState334, CStatementType}:                _CGotoState346Action,
	{_CState334, CLabeledStatementType}:         _CGotoState115Action,
	{_CState334, CCompoundStatementType}:        _CGotoState105Action,
	{_CState334, CExpressionStatementType}:      _CGotoState111Action,
	{_CState334, CSelectionStatementType}:       _CGotoState122Action,
	{_CState334, CIterationStatementType}:       _CGotoState113Action,
	{_CState334, CJumpStatementType}:            _CGotoState114Action,
	{_CState335, ')'}:                           _CGotoState347Action,
	{_CState335, ','}:                           _CGotoState187Action,
	{_CState336, CElseToken}:                    _CGotoState348Action,
	{_CState337, CElseToken}:                    _CGotoState348Action,
	{_CState338, CElseToken}:                    _CGotoState348Action,
	{_CState345, ';'}:                           _CGotoState349Action,
	{_CState347, CIdentifierToken}:              _CGotoState93Action,
	{_CState347, CConstantToken}:                _CGotoState86Action,
	{_CState347, CStringLiteralToken}:           _CGotoState98Action,
	{_CState347, CSizeofToken}:                  _CGotoState97Action,
	{_CState347, CIncOpToken}:                   _CGotoState95Action,
	{_CState347, CDecOpToken}:                   _CGotoState88Action,
	{_CState347, CCaseToken}:                    _CGotoState85Action,
	{_CState347, CDefaultToken}:                 _CGotoState89Action,
	{_CState347, CIfToken}:                      _CGotoState94Action,
	{_CState347, CSwitchToken}:                  _CGotoState99Action,
	{_CState347, CWhileToken}:                   _CGotoState100Action,
	{_CState347, CDoToken}:                      _CGotoState90Action,
	{_CState347, CForToken}:                     _CGotoState91Action,
	{_CState347, CGotoToken}:                    _CGotoState92Action,
	{_CState347, CContinueToken}:                _CGotoState87Action,
	{_CState347, CBreakToken}:                   _CGotoState84Action,
	{_CState347, CReturnToken}:                  _CGotoState96Action,
	{_CState347, '('}:                           _CGotoState77Action,
	{_CState347, '{'}:                           _CGotoState49Action,
	{_CState347, ';'}:                           _CGotoState81Action,
	{_CState347, '*'}:                           _CGotoState78Action,
	{_CState347, '-'}:                           _CGotoState80Action,
	{_CState347, '+'}:                           _CGotoState79Action,
	{_CState347, '&'}:                           _CGotoState76Action,
	{_CState347, '!'}:                           _CGotoState75Action,
	{_CState347, '~'}:                           _CGotoState83Action,
	{_CState347, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState347, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState347, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState347, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState347, CCastExpressionType}:           _CGotoState104Action,
	{_CState347, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState347, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState347, CShiftExpressionType}:          _CGotoState123Action,
	{_CState347, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState347, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState347, CAndExpressionType}:            _CGotoState102Action,
	{_CState347, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState347, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState347, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState347, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState347, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState347, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState347, CExpressionType}:               _CGotoState110Action,
	{_CState347, CStatementType}:                _CGotoState350Action,
	{_CState347, CLabeledStatementType}:         _CGotoState115Action,
	{_CState347, CCompoundStatementType}:        _CGotoState105Action,
	{_CState347, CExpressionStatementType}:      _CGotoState111Action,
	{_CState347, CSelectionStatementType}:       _CGotoState122Action,
	{_CState347, CIterationStatementType}:       _CGotoState113Action,
	{_CState347, CJumpStatementType}:            _CGotoState114Action,
	{_CState348, CIdentifierToken}:              _CGotoState93Action,
	{_CState348, CConstantToken}:                _CGotoState86Action,
	{_CState348, CStringLiteralToken}:           _CGotoState98Action,
	{_CState348, CSizeofToken}:                  _CGotoState97Action,
	{_CState348, CIncOpToken}:                   _CGotoState95Action,
	{_CState348, CDecOpToken}:                   _CGotoState88Action,
	{_CState348, CCaseToken}:                    _CGotoState85Action,
	{_CState348, CDefaultToken}:                 _CGotoState89Action,
	{_CState348, CIfToken}:                      _CGotoState94Action,
	{_CState348, CSwitchToken}:                  _CGotoState99Action,
	{_CState348, CWhileToken}:                   _CGotoState100Action,
	{_CState348, CDoToken}:                      _CGotoState90Action,
	{_CState348, CForToken}:                     _CGotoState91Action,
	{_CState348, CGotoToken}:                    _CGotoState92Action,
	{_CState348, CContinueToken}:                _CGotoState87Action,
	{_CState348, CBreakToken}:                   _CGotoState84Action,
	{_CState348, CReturnToken}:                  _CGotoState96Action,
	{_CState348, '('}:                           _CGotoState77Action,
	{_CState348, '{'}:                           _CGotoState49Action,
	{_CState348, ';'}:                           _CGotoState81Action,
	{_CState348, '*'}:                           _CGotoState78Action,
	{_CState348, '-'}:                           _CGotoState80Action,
	{_CState348, '+'}:                           _CGotoState79Action,
	{_CState348, '&'}:                           _CGotoState76Action,
	{_CState348, '!'}:                           _CGotoState75Action,
	{_CState348, '~'}:                           _CGotoState83Action,
	{_CState348, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState348, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState348, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState348, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState348, CCastExpressionType}:           _CGotoState104Action,
	{_CState348, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState348, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState348, CShiftExpressionType}:          _CGotoState123Action,
	{_CState348, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState348, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState348, CAndExpressionType}:            _CGotoState102Action,
	{_CState348, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState348, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState348, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState348, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState348, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState348, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState348, CExpressionType}:               _CGotoState110Action,
	{_CState348, CStatementType}:                _CGotoState351Action,
	{_CState348, CLabeledStatementType}:         _CGotoState115Action,
	{_CState348, CCompoundStatementType}:        _CGotoState105Action,
	{_CState348, CExpressionStatementType}:      _CGotoState111Action,
	{_CState348, CSelectionStatementType}:       _CGotoState122Action,
	{_CState348, CIterationStatementType}:       _CGotoState113Action,
	{_CState348, CJumpStatementType}:            _CGotoState114Action,
	{_CState4, _CWildcardMarker}:                _CReduceAToPointerAction,
	{_CState5, _CWildcardMarker}:                _CReduceDToStorageClassSpecifierAction,
	{_CState6, _CWildcardMarker}:                _CReduceBToTypeSpecifierAction,
	{_CState7, _CWildcardMarker}:                _CReduceAToTypeQualifierAction,
	{_CState8, _CWildcardMarker}:                _CReduceGToTypeSpecifierAction,
	{_CState10, _CWildcardMarker}:               _CReduceBToStorageClassSpecifierAction,
	{_CState11, _CWildcardMarker}:               _CReduceFToTypeSpecifierAction,
	{_CState12, _CWildcardMarker}:               _CReduceAToDirectDeclaratorAction,
	{_CState13, _CWildcardMarker}:               _CReduceDToTypeSpecifierAction,
	{_CState14, _CWildcardMarker}:               _CReduceEToTypeSpecifierAction,
	{_CState15, _CWildcardMarker}:               _CReduceEToStorageClassSpecifierAction,
	{_CState16, _CWildcardMarker}:               _CReduceCToTypeSpecifierAction,
	{_CState17, _CWildcardMarker}:               _CReduceHToTypeSpecifierAction,
	{_CState18, _CWildcardMarker}:               _CReduceCToStorageClassSpecifierAction,
	{_CState19, _CWildcardMarker}:               _CReduceAToStructOrUnionAction,
	{_CState20, _CWildcardMarker}:               _CReduceAToStorageClassSpecifierAction,
	{_CState21, _CWildcardMarker}:               _CReduceLToTypeSpecifierAction,
	{_CState22, _CWildcardMarker}:               _CReduceBToStructOrUnionAction,
	{_CState23, _CWildcardMarker}:               _CReduceIToTypeSpecifierAction,
	{_CState24, _CWildcardMarker}:               _CReduceAToTypeSpecifierAction,
	{_CState25, _CWildcardMarker}:               _CReduceBToTypeQualifierAction,
	{_CState26, _CWildcardMarker}:               _CReduceBToExternalDeclarationAction,
	{_CState29, _CWildcardMarker}:               _CReduceBToDeclaratorAction,
	{_CState30, _CWildcardMarker}:               _CReduceKToTypeSpecifierAction,
	{_CState31, _CWildcardMarker}:               _CReduceAToTranslationUnitAction,
	{_CState32, _CWildcardMarker}:               _CReduceAToExternalDeclarationAction,
	{_CState34, _CWildcardMarker}:               _CReduceAToDeclarationSpecifiersAction,
	{_CState36, _CWildcardMarker}:               _CReduceJToTypeSpecifierAction,
	{_CState37, _CWildcardMarker}:               _CReduceEToDeclarationSpecifiersAction,
	{_CState38, _CWildcardMarker}:               _CReduceCToDeclarationSpecifiersAction,
	{_CState40, _CWildcardMarker}:               _CReduceCToPointerAction,
	{_CState41, _CWildcardMarker}:               _CReduceAToTypeQualifierListAction,
	{_CState42, _CWildcardMarker}:               _CReduceBToPointerAction,
	{_CState44, _CWildcardMarker}:               _CReduceCToEnumSpecifierAction,
	{_CState45, _CWildcardMarker}:               _CReduceAToDeclarationAction,
	{_CState46, _CWildcardMarker}:               _CReduceAToInitDeclaratorAction,
	{_CState47, _CWildcardMarker}:               _CReduceAToInitDeclaratorListAction,
	{_CState50, _CWildcardMarker}:               _CReduceDToFunctionDefinitionAction,
	{_CState51, _CWildcardMarker}:               _CReduceAToDeclarationListAction,
	{_CState56, _CWildcardMarker}:               _CReduceAToDeclaratorAction,
	{_CState57, _CWildcardMarker}:               _CReduceBToDeclarationSpecifiersAction,
	{_CState59, _CWildcardMarker}:               _CReduceCToStructOrUnionSpecifierAction,
	{_CState60, _CWildcardMarker}:               _CReduceBToTranslationUnitAction,
	{_CState61, _CWildcardMarker}:               _CReduceFToDeclarationSpecifiersAction,
	{_CState62, _CWildcardMarker}:               _CReduceDToDeclarationSpecifiersAction,
	{_CState63, _CWildcardMarker}:               _CReduceBToDirectDeclaratorAction,
	{_CState64, _CWildcardMarker}:               _CReduceDToPointerAction,
	{_CState65, _CWildcardMarker}:               _CReduceBToTypeQualifierListAction,
	{_CState66, _CWildcardMarker}:               _CReduceAToEnumeratorAction,
	{_CState67, _CWildcardMarker}:               _CReduceAToEnumeratorListAction,
	{_CState71, _CWildcardMarker}:               _CReduceBToFunctionDefinitionAction,
	{_CState74, _CWildcardMarker}:               _CReduceBToDeclarationAction,
	{_CState75, _CWildcardMarker}:               _CReduceFToUnaryOperatorAction,
	{_CState76, _CWildcardMarker}:               _CReduceAToUnaryOperatorAction,
	{_CState78, _CWildcardMarker}:               _CReduceBToUnaryOperatorAction,
	{_CState79, _CWildcardMarker}:               _CReduceCToUnaryOperatorAction,
	{_CState80, _CWildcardMarker}:               _CReduceDToUnaryOperatorAction,
	{_CState81, _CWildcardMarker}:               _CReduceAToExpressionStatementAction,
	{_CState82, _CWildcardMarker}:               _CReduceAToCompoundStatementAction,
	{_CState83, _CWildcardMarker}:               _CReduceEToUnaryOperatorAction,
	{_CState86, _CWildcardMarker}:               _CReduceBToPrimaryExpressionAction,
	{_CState93, _CWildcardMarker}:               _CReduceAToPrimaryExpressionAction,
	{_CState98, _CWildcardMarker}:               _CReduceCToPrimaryExpressionAction,
	{_CState101, _CWildcardMarker}:              _CReduceAToShiftExpressionAction,
	{_CState102, _CWildcardMarker}:              _CReduceAToExclusiveOrExpressionAction,
	{_CState103, _CWildcardMarker}:              _CReduceAToExpressionAction,
	{_CState104, _CWildcardMarker}:              _CReduceAToMultiplicativeExpressionAction,
	{_CState105, _CWildcardMarker}:              _CReduceBToStatementAction,
	{_CState106, _CWildcardMarker}:              _CReduceAToAssignmentExpressionAction,
	{_CState108, _CWildcardMarker}:              _CReduceAToAndExpressionAction,
	{_CState109, _CWildcardMarker}:              _CReduceAToInclusiveOrExpressionAction,
	{_CState111, _CWildcardMarker}:              _CReduceCToStatementAction,
	{_CState112, _CWildcardMarker}:              _CReduceAToLogicalAndExpressionAction,
	{_CState113, _CWildcardMarker}:              _CReduceEToStatementAction,
	{_CState114, _CWildcardMarker}:              _CReduceFToStatementAction,
	{_CState115, _CWildcardMarker}:              _CReduceAToStatementAction,
	{_CState116, _CWildcardMarker}:              _CReduceAToLogicalOrExpressionAction,
	{_CState117, _CWildcardMarker}:              _CReduceAToConditionalExpressionAction,
	{_CState118, _CWildcardMarker}:              _CReduceAToAdditiveExpressionAction,
	{_CState119, _CWildcardMarker}:              _CReduceAToUnaryExpressionAction,
	{_CState120, _CWildcardMarker}:              _CReduceAToPostfixExpressionAction,
	{_CState121, _CWildcardMarker}:              _CReduceAToEqualityExpressionAction,
	{_CState122, _CWildcardMarker}:              _CReduceDToStatementAction,
	{_CState123, _CWildcardMarker}:              _CReduceAToRelationalExpressionAction,
	{_CState124, _CWildcardMarker}:              _CReduceAToStatementListAction,
	{_CState126, _CWildcardMarker}:              _CReduceAToCastExpressionAction,
	{_CState128, _CWildcardMarker}:              _CReduceCToFunctionDefinitionAction,
	{_CState129, _CWildcardMarker}:              _CReduceBToDeclarationListAction,
	{_CState130, _CWildcardMarker}:              _CReduceAToInitDeclaratorAction,
	{_CState131, _CWildcardMarker}:              _CReduceGToDirectDeclaratorAction,
	{_CState132, _CWildcardMarker}:              _CReduceAToIdentifierListAction,
	{_CState133, _CWildcardMarker}:              _CReduceCToParameterDeclarationAction,
	{_CState135, _CWildcardMarker}:              _CReduceAToParameterListAction,
	{_CState136, ')'}:                           _CReduceAToParameterTypeListAction,
	{_CState138, _CWildcardMarker}:              _CReduceDToDirectDeclaratorAction,
	{_CState139, _CWildcardMarker}:              _CReduceAToPrimaryExpressionAction,
	{_CState140, _CWildcardMarker}:              _CReduceAToConstantExpressionAction,
	{_CState142, _CWildcardMarker}:              _CReduceAToCastExpressionAction,
	{_CState144, _CWildcardMarker}:              _CReduceAToStructDeclarationListAction,
	{_CState146, _CWildcardMarker}:              _CReduceDToSpecifierQualifierListAction,
	{_CState147, _CWildcardMarker}:              _CReduceBToSpecifierQualifierListAction,
	{_CState151, _CWildcardMarker}:              _CReduceAToEnumSpecifierAction,
	{_CState154, _CWildcardMarker}:              _CReduceAToInitializerAction,
	{_CState155, _CWildcardMarker}:              _CReduceBToInitDeclaratorAction,
	{_CState156, _CWildcardMarker}:              _CReduceAToFunctionDefinitionAction,
	{_CState157, _CWildcardMarker}:              _CReduceBToInitDeclaratorListAction,
	{_CState159, ')'}:                           _CReduceAToTypeNameAction,
	{_CState161, _CWildcardMarker}:              _CReduceCToJumpStatementAction,
	{_CState163, _CWildcardMarker}:              _CReduceBToJumpStatementAction,
	{_CState165, _CWildcardMarker}:              _CReduceCToUnaryExpressionAction,
	{_CState172, _CWildcardMarker}:              _CReduceBToUnaryExpressionAction,
	{_CState173, _CWildcardMarker}:              _CReduceDToJumpStatementAction,
	{_CState176, _CWildcardMarker}:              _CReduceEToUnaryExpressionAction,
	{_CState182, _CWildcardMarker}:              _CReduceCToCompoundStatementAction,
	{_CState188, _CWildcardMarker}:              _CReduceBToExpressionStatementAction,
	{_CState199, _CWildcardMarker}:              _CReduceHToPostfixExpressionAction,
	{_CState200, _CWildcardMarker}:              _CReduceGToPostfixExpressionAction,
	{_CState208, _CWildcardMarker}:              _CReduceBToCompoundStatementAction,
	{_CState209, _CWildcardMarker}:              _CReduceBToStatementListAction,
	{_CState210, _CWildcardMarker}:              _CReduceAToAssignmentOperatorAction,
	{_CState211, _CWildcardMarker}:              _CReduceEToAssignmentOperatorAction,
	{_CState212, _CWildcardMarker}:              _CReduceIToAssignmentOperatorAction,
	{_CState213, _CWildcardMarker}:              _CReduceCToAssignmentOperatorAction,
	{_CState214, _CWildcardMarker}:              _CReduceGToAssignmentOperatorAction,
	{_CState215, _CWildcardMarker}:              _CReduceDToAssignmentOperatorAction,
	{_CState216, _CWildcardMarker}:              _CReduceBToAssignmentOperatorAction,
	{_CState217, _CWildcardMarker}:              _CReduceKToAssignmentOperatorAction,
	{_CState218, _CWildcardMarker}:              _CReduceHToAssignmentOperatorAction,
	{_CState219, _CWildcardMarker}:              _CReduceFToAssignmentOperatorAction,
	{_CState220, _CWildcardMarker}:              _CReduceJToAssignmentOperatorAction,
	{_CState222, _CWildcardMarker}:              _CReduceDToUnaryExpressionAction,
	{_CState225, _CWildcardMarker}:              _CReduceBToParameterDeclarationAction,
	{_CState226, _CWildcardMarker}:              _CReduceAToParameterDeclarationAction,
	{_CState227, _CWildcardMarker}:              _CReduceBToAbstractDeclaratorAction,
	{_CState228, _CWildcardMarker}:              _CReduceAToAbstractDeclaratorAction,
	{_CState229, _CWildcardMarker}:              _CReduceFToDirectDeclaratorAction,
	{_CState232, _CWildcardMarker}:              _CReduceEToDirectDeclaratorAction,
	{_CState233, _CWildcardMarker}:              _CReduceCToDirectDeclaratorAction,
	{_CState235, _CWildcardMarker}:              _CReduceAToStructDeclaratorAction,
	{_CState236, _CWildcardMarker}:              _CReduceAToStructDeclaratorListAction,
	{_CState238, _CWildcardMarker}:              _CReduceBToStructOrUnionSpecifierAction,
	{_CState239, _CWildcardMarker}:              _CReduceBToStructDeclarationListAction,
	{_CState240, _CWildcardMarker}:              _CReduceCToSpecifierQualifierListAction,
	{_CState241, _CWildcardMarker}:              _CReduceAToSpecifierQualifierListAction,
	{_CState243, _CWildcardMarker}:              _CReduceBToEnumeratorAction,
	{_CState244, _CWildcardMarker}:              _CReduceBToEnumeratorListAction,
	{_CState245, _CWildcardMarker}:              _CReduceBToEnumSpecifierAction,
	{_CState246, _CWildcardMarker}:              _CReduceAToInitializerListAction,
	{_CState248, _CWildcardMarker}:              _CReduceDToPrimaryExpressionAction,
	{_CState250, ')'}:                           _CReduceBToTypeNameAction,
	{_CState251, ')'}:                           _CReduceAToAbstractDeclaratorAction,
	{_CState254, _CWildcardMarker}:              _CReduceCToLabeledStatementAction,
	{_CState257, _CWildcardMarker}:              _CReduceAToJumpStatementAction,
	{_CState258, _CWildcardMarker}:              _CReduceAToLabeledStatementAction,
	{_CState260, _CWildcardMarker}:              _CReduceEToJumpStatementAction,
	{_CState264, _CWildcardMarker}:              _CReduceBToAdditiveExpressionAction,
	{_CState265, _CWildcardMarker}:              _CReduceCToAdditiveExpressionAction,
	{_CState266, _CWildcardMarker}:              _CReduceBToAndExpressionAction,
	{_CState267, _CWildcardMarker}:              _CReduceDToCompoundStatementAction,
	{_CState268, _CWildcardMarker}:              _CReduceBToEqualityExpressionAction,
	{_CState269, _CWildcardMarker}:              _CReduceCToEqualityExpressionAction,
	{_CState270, _CWildcardMarker}:              _CReduceBToExclusiveOrExpressionAction,
	{_CState271, _CWildcardMarker}:              _CReduceBToExpressionAction,
	{_CState272, _CWildcardMarker}:              _CReduceBToInclusiveOrExpressionAction,
	{_CState273, _CWildcardMarker}:              _CReduceBToLogicalAndExpressionAction,
	{_CState275, _CWildcardMarker}:              _CReduceBToLogicalOrExpressionAction,
	{_CState276, _CWildcardMarker}:              _CReduceDToMultiplicativeExpressionAction,
	{_CState277, _CWildcardMarker}:              _CReduceBToMultiplicativeExpressionAction,
	{_CState278, _CWildcardMarker}:              _CReduceCToMultiplicativeExpressionAction,
	{_CState279, _CWildcardMarker}:              _CReduceCToPostfixExpressionAction,
	{_CState281, _CWildcardMarker}:              _CReduceAToArgumentExpressionListAction,
	{_CState282, _CWildcardMarker}:              _CReduceEToPostfixExpressionAction,
	{_CState284, _CWildcardMarker}:              _CReduceFToPostfixExpressionAction,
	{_CState285, _CWildcardMarker}:              _CReduceBToRelationalExpressionAction,
	{_CState286, _CWildcardMarker}:              _CReduceCToRelationalExpressionAction,
	{_CState287, _CWildcardMarker}:              _CReduceEToRelationalExpressionAction,
	{_CState288, _CWildcardMarker}:              _CReduceDToRelationalExpressionAction,
	{_CState289, _CWildcardMarker}:              _CReduceBToShiftExpressionAction,
	{_CState290, _CWildcardMarker}:              _CReduceCToShiftExpressionAction,
	{_CState291, _CWildcardMarker}:              _CReduceBToAssignmentExpressionAction,
	{_CState292, _CWildcardMarker}:              _CReduceFToDirectAbstractDeclaratorAction,
	{_CState295, _CWildcardMarker}:              _CReduceBToDirectAbstractDeclaratorAction,
	{_CState299, _CWildcardMarker}:              _CReduceCToAbstractDeclaratorAction,
	{_CState300, _CWildcardMarker}:              _CReduceBToIdentifierListAction,
	{_CState301, ')'}:                           _CReduceBToParameterTypeListAction,
	{_CState302, _CWildcardMarker}:              _CReduceBToParameterListAction,
	{_CState303, _CWildcardMarker}:              _CReduceBToStructDeclaratorAction,
	{_CState306, _CWildcardMarker}:              _CReduceAToStructDeclarationAction,
	{_CState307, _CWildcardMarker}:              _CReduceAToStructOrUnionSpecifierAction,
	{_CState309, _CWildcardMarker}:              _CReduceBToInitializerAction,
	{_CState310, _CWildcardMarker}:              _CReduceBToCastExpressionAction,
	{_CState311, _CWildcardMarker}:              _CReduceBToLabeledStatementAction,
	{_CState315, _CWildcardMarker}:              _CReduceFToUnaryExpressionAction,
	{_CState319, _CWildcardMarker}:              _CReduceDToPostfixExpressionAction,
	{_CState321, _CWildcardMarker}:              _CReduceBToPostfixExpressionAction,
	{_CState322, _CWildcardMarker}:              _CReduceAToDirectAbstractDeclaratorAction,
	{_CState323, _CWildcardMarker}:              _CReduceGToDirectAbstractDeclaratorAction,
	{_CState324, _CWildcardMarker}:              _CReduceCToDirectAbstractDeclaratorAction,
	{_CState325, _CWildcardMarker}:              _CReduceHToDirectAbstractDeclaratorAction,
	{_CState327, _CWildcardMarker}:              _CReduceDToDirectAbstractDeclaratorAction,
	{_CState329, _CWildcardMarker}:              _CReduceCToStructDeclaratorAction,
	{_CState330, _CWildcardMarker}:              _CReduceBToStructDeclaratorListAction,
	{_CState331, _CWildcardMarker}:              _CReduceCToInitializerAction,
	{_CState332, _CWildcardMarker}:              _CReduceBToInitializerListAction,
	{_CState336, _CWildcardMarker}:              _CReduceAToSelectionStatementAction,
	{_CState337, '!'}:                           _CReduceAToSelectionStatementAction,
	{_CState337, '&'}:                           _CReduceAToSelectionStatementAction,
	{_CState337, '('}:                           _CReduceAToSelectionStatementAction,
	{_CState337, '*'}:                           _CReduceAToSelectionStatementAction,
	{_CState337, '+'}:                           _CReduceAToSelectionStatementAction,
	{_CState337, '-'}:                           _CReduceAToSelectionStatementAction,
	{_CState337, ';'}:                           _CReduceAToSelectionStatementAction,
	{_CState337, '{'}:                           _CReduceAToSelectionStatementAction,
	{_CState337, '}'}:                           _CReduceAToSelectionStatementAction,
	{_CState337, '~'}:                           _CReduceAToSelectionStatementAction,
	{_CState337, CBreakToken}:                   _CReduceAToSelectionStatementAction,
	{_CState337, CCaseToken}:                    _CReduceAToSelectionStatementAction,
	{_CState337, CConstantToken}:                _CReduceAToSelectionStatementAction,
	{_CState337, CContinueToken}:                _CReduceAToSelectionStatementAction,
	{_CState337, CDecOpToken}:                   _CReduceAToSelectionStatementAction,
	{_CState337, CDefaultToken}:                 _CReduceAToSelectionStatementAction,
	{_CState337, CDoToken}:                      _CReduceAToSelectionStatementAction,
	{_CState337, CElseToken}:                    _CReduceAToSelectionStatementAction,
	{_CState337, CForToken}:                     _CReduceAToSelectionStatementAction,
	{_CState337, CGotoToken}:                    _CReduceAToSelectionStatementAction,
	{_CState337, CIdentifierToken}:              _CReduceAToSelectionStatementAction,
	{_CState337, CIfToken}:                      _CReduceAToSelectionStatementAction,
	{_CState337, CIncOpToken}:                   _CReduceAToSelectionStatementAction,
	{_CState337, CReturnToken}:                  _CReduceAToSelectionStatementAction,
	{_CState337, CSizeofToken}:                  _CReduceAToSelectionStatementAction,
	{_CState337, CStringLiteralToken}:           _CReduceAToSelectionStatementAction,
	{_CState337, CSwitchToken}:                  _CReduceAToSelectionStatementAction,
	{_CState337, CWhileToken}:                   _CReduceAToSelectionStatementAction,
	{_CState338, CElseToken}:                    _CReduceAToSelectionStatementAction,
	{_CState338, CWhileToken}:                   _CReduceAToSelectionStatementAction,
	{_CState339, _CWildcardMarker}:              _CReduceCToSelectionStatementAction,
	{_CState340, _CWildcardMarker}:              _CReduceAToIterationStatementAction,
	{_CState341, _CWildcardMarker}:              _CReduceBToConditionalExpressionAction,
	{_CState342, _CWildcardMarker}:              _CReduceBToArgumentExpressionListAction,
	{_CState343, _CWildcardMarker}:              _CReduceIToDirectAbstractDeclaratorAction,
	{_CState344, _CWildcardMarker}:              _CReduceEToDirectAbstractDeclaratorAction,
	{_CState346, _CWildcardMarker}:              _CReduceCToIterationStatementAction,
	{_CState349, _CWildcardMarker}:              _CReduceBToIterationStatementAction,
	{_CState350, _CWildcardMarker}:              _CReduceDToIterationStatementAction,
	{_CState351, _CWildcardMarker}:              _CReduceBToSelectionStatementAction,
}

/*
Parser Debug States:
  State 1:
    Kernel Items:
      #accept: ^.translation_unit
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 12
      TYPE_NAME -> State 21
      TYPEDEF -> State 20
      EXTERN -> State 10
      STATIC -> State 18
      AUTO -> State 5
      REGISTER -> State 15
      CHAR -> State 6
      SHORT -> State 16
      INT -> State 13
      LONG -> State 14
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 11
      DOUBLE -> State 8
      CONST -> State 7
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 9
      '(' -> State 3
      '*' -> State 4
      declaration -> State 26
      declaration_specifiers -> State 27
      storage_class_specifier -> State 34
      type_specifier -> State 38
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      enum_specifier -> State 30
      type_qualifier -> State 37
      declarator -> State 28
      direct_declarator -> State 29
      pointer -> State 33
      translation_unit -> State 2
      external_declaration -> State 31
      function_definition -> State 32

  State 2:
    Kernel Items:
      #accept: ^ translation_unit., $
      translation_unit: translation_unit.external_declaration
    Reduce:
      $ -> [#accept]
    Goto:
      IDENTIFIER -> State 12
      TYPE_NAME -> State 21
      TYPEDEF -> State 20
      EXTERN -> State 10
      STATIC -> State 18
      AUTO -> State 5
      REGISTER -> State 15
      CHAR -> State 6
      SHORT -> State 16
      INT -> State 13
      LONG -> State 14
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 11
      DOUBLE -> State 8
      CONST -> State 7
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 9
      '(' -> State 3
      '*' -> State 4
      declaration -> State 26
      declaration_specifiers -> State 27
      storage_class_specifier -> State 34
      type_specifier -> State 38
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      enum_specifier -> State 30
      type_qualifier -> State 37
      declarator -> State 28
      direct_declarator -> State 29
      pointer -> State 33
      external_declaration -> State 60
      function_definition -> State 32

  State 3:
    Kernel Items:
      direct_declarator: '('.declarator ')'
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 12
      '(' -> State 3
      '*' -> State 4
      declarator -> State 39
      direct_declarator -> State 29
      pointer -> State 33

  State 4:
    Kernel Items:
      pointer: '*'., *
      pointer: '*'.type_qualifier_list
      pointer: '*'.pointer
      pointer: '*'.type_qualifier_list pointer
    Reduce:
      * -> [pointer]
    Goto:
      CONST -> State 7
      VOLATILE -> State 25
      '*' -> State 4
      type_qualifier -> State 41
      pointer -> State 40
      type_qualifier_list -> State 42

  State 5:
    Kernel Items:
      storage_class_specifier: AUTO., *
    Reduce:
      * -> [storage_class_specifier]
    Goto:
      (nil)

  State 6:
    Kernel Items:
      type_specifier: CHAR., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 7:
    Kernel Items:
      type_qualifier: CONST., *
    Reduce:
      * -> [type_qualifier]
    Goto:
      (nil)

  State 8:
    Kernel Items:
      type_specifier: DOUBLE., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 9:
    Kernel Items:
      enum_specifier: ENUM.'{' enumerator_list '}'
      enum_specifier: ENUM.IDENTIFIER '{' enumerator_list '}'
      enum_specifier: ENUM.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 44
      '{' -> State 43

  State 10:
    Kernel Items:
      storage_class_specifier: EXTERN., *
    Reduce:
      * -> [storage_class_specifier]
    Goto:
      (nil)

  State 11:
    Kernel Items:
      type_specifier: FLOAT., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 12:
    Kernel Items:
      direct_declarator: IDENTIFIER., *
    Reduce:
      * -> [direct_declarator]
    Goto:
      (nil)

  State 13:
    Kernel Items:
      type_specifier: INT., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 14:
    Kernel Items:
      type_specifier: LONG., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 15:
    Kernel Items:
      storage_class_specifier: REGISTER., *
    Reduce:
      * -> [storage_class_specifier]
    Goto:
      (nil)

  State 16:
    Kernel Items:
      type_specifier: SHORT., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 17:
    Kernel Items:
      type_specifier: SIGNED., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 18:
    Kernel Items:
      storage_class_specifier: STATIC., *
    Reduce:
      * -> [storage_class_specifier]
    Goto:
      (nil)

  State 19:
    Kernel Items:
      struct_or_union: STRUCT., *
    Reduce:
      * -> [struct_or_union]
    Goto:
      (nil)

  State 20:
    Kernel Items:
      storage_class_specifier: TYPEDEF., *
    Reduce:
      * -> [storage_class_specifier]
    Goto:
      (nil)

  State 21:
    Kernel Items:
      type_specifier: TYPE_NAME., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 22:
    Kernel Items:
      struct_or_union: UNION., *
    Reduce:
      * -> [struct_or_union]
    Goto:
      (nil)

  State 23:
    Kernel Items:
      type_specifier: UNSIGNED., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 24:
    Kernel Items:
      type_specifier: VOID., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 25:
    Kernel Items:
      type_qualifier: VOLATILE., *
    Reduce:
      * -> [type_qualifier]
    Goto:
      (nil)

  State 26:
    Kernel Items:
      external_declaration: declaration., *
    Reduce:
      * -> [external_declaration]
    Goto:
      (nil)

  State 27:
    Kernel Items:
      declaration: declaration_specifiers.';'
      declaration: declaration_specifiers.init_declarator_list ';'
      function_definition: declaration_specifiers.declarator declaration_list compound_statement
      function_definition: declaration_specifiers.declarator compound_statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 12
      '(' -> State 3
      ';' -> State 45
      '*' -> State 4
      init_declarator_list -> State 48
      init_declarator -> State 47
      declarator -> State 46
      direct_declarator -> State 29
      pointer -> State 33

  State 28:
    Kernel Items:
      function_definition: declarator.declaration_list compound_statement
      function_definition: declarator.compound_statement
    Reduce:
      (nil)
    Goto:
      TYPE_NAME -> State 21
      TYPEDEF -> State 20
      EXTERN -> State 10
      STATIC -> State 18
      AUTO -> State 5
      REGISTER -> State 15
      CHAR -> State 6
      SHORT -> State 16
      INT -> State 13
      LONG -> State 14
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 11
      DOUBLE -> State 8
      CONST -> State 7
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 9
      '{' -> State 49
      declaration -> State 51
      declaration_specifiers -> State 53
      storage_class_specifier -> State 34
      type_specifier -> State 38
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      enum_specifier -> State 30
      type_qualifier -> State 37
      compound_statement -> State 50
      declaration_list -> State 52

  State 29:
    Kernel Items:
      declarator: direct_declarator., *
      direct_declarator: direct_declarator.'[' constant_expression ']'
      direct_declarator: direct_declarator.'[' ']'
      direct_declarator: direct_declarator.'(' parameter_type_list ')'
      direct_declarator: direct_declarator.'(' identifier_list ')'
      direct_declarator: direct_declarator.'(' ')'
    Reduce:
      * -> [declarator]
    Goto:
      '(' -> State 54
      '[' -> State 55

  State 30:
    Kernel Items:
      type_specifier: enum_specifier., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 31:
    Kernel Items:
      translation_unit: external_declaration., *
    Reduce:
      * -> [translation_unit]
    Goto:
      (nil)

  State 32:
    Kernel Items:
      external_declaration: function_definition., *
    Reduce:
      * -> [external_declaration]
    Goto:
      (nil)

  State 33:
    Kernel Items:
      declarator: pointer.direct_declarator
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 12
      '(' -> State 3
      direct_declarator -> State 56

  State 34:
    Kernel Items:
      declaration_specifiers: storage_class_specifier., *
      declaration_specifiers: storage_class_specifier.declaration_specifiers
    Reduce:
      * -> [declaration_specifiers]
    Goto:
      TYPE_NAME -> State 21
      TYPEDEF -> State 20
      EXTERN -> State 10
      STATIC -> State 18
      AUTO -> State 5
      REGISTER -> State 15
      CHAR -> State 6
      SHORT -> State 16
      INT -> State 13
      LONG -> State 14
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 11
      DOUBLE -> State 8
      CONST -> State 7
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 9
      declaration_specifiers -> State 57
      storage_class_specifier -> State 34
      type_specifier -> State 38
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      enum_specifier -> State 30
      type_qualifier -> State 37

  State 35:
    Kernel Items:
      struct_or_union_specifier: struct_or_union.IDENTIFIER '{' struct_declaration_list '}'
      struct_or_union_specifier: struct_or_union.'{' struct_declaration_list '}'
      struct_or_union_specifier: struct_or_union.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 59
      '{' -> State 58

  State 36:
    Kernel Items:
      type_specifier: struct_or_union_specifier., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 37:
    Kernel Items:
      declaration_specifiers: type_qualifier., *
      declaration_specifiers: type_qualifier.declaration_specifiers
    Reduce:
      * -> [declaration_specifiers]
    Goto:
      TYPE_NAME -> State 21
      TYPEDEF -> State 20
      EXTERN -> State 10
      STATIC -> State 18
      AUTO -> State 5
      REGISTER -> State 15
      CHAR -> State 6
      SHORT -> State 16
      INT -> State 13
      LONG -> State 14
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 11
      DOUBLE -> State 8
      CONST -> State 7
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 9
      declaration_specifiers -> State 61
      storage_class_specifier -> State 34
      type_specifier -> State 38
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      enum_specifier -> State 30
      type_qualifier -> State 37

  State 38:
    Kernel Items:
      declaration_specifiers: type_specifier., *
      declaration_specifiers: type_specifier.declaration_specifiers
    Reduce:
      * -> [declaration_specifiers]
    Goto:
      TYPE_NAME -> State 21
      TYPEDEF -> State 20
      EXTERN -> State 10
      STATIC -> State 18
      AUTO -> State 5
      REGISTER -> State 15
      CHAR -> State 6
      SHORT -> State 16
      INT -> State 13
      LONG -> State 14
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 11
      DOUBLE -> State 8
      CONST -> State 7
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 9
      declaration_specifiers -> State 62
      storage_class_specifier -> State 34
      type_specifier -> State 38
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      enum_specifier -> State 30
      type_qualifier -> State 37

  State 39:
    Kernel Items:
      direct_declarator: '(' declarator.')'
    Reduce:
      (nil)
    Goto:
      ')' -> State 63

  State 40:
    Kernel Items:
      pointer: '*' pointer., *
    Reduce:
      * -> [pointer]
    Goto:
      (nil)

  State 41:
    Kernel Items:
      type_qualifier_list: type_qualifier., *
    Reduce:
      * -> [type_qualifier_list]
    Goto:
      (nil)

  State 42:
    Kernel Items:
      pointer: '*' type_qualifier_list., *
      pointer: '*' type_qualifier_list.pointer
      type_qualifier_list: type_qualifier_list.type_qualifier
    Reduce:
      * -> [pointer]
    Goto:
      CONST -> State 7
      VOLATILE -> State 25
      '*' -> State 4
      type_qualifier -> State 65
      pointer -> State 64

  State 43:
    Kernel Items:
      enum_specifier: ENUM '{'.enumerator_list '}'
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 66
      enumerator_list -> State 68
      enumerator -> State 67

  State 44:
    Kernel Items:
      enum_specifier: ENUM IDENTIFIER.'{' enumerator_list '}'
      enum_specifier: ENUM IDENTIFIER., *
    Reduce:
      * -> [enum_specifier]
    Goto:
      '{' -> State 69

  State 45:
    Kernel Items:
      declaration: declaration_specifiers ';'., *
    Reduce:
      * -> [declaration]
    Goto:
      (nil)

  State 46:
    Kernel Items:
      init_declarator: declarator., *
      init_declarator: declarator.'=' initializer
      function_definition: declaration_specifiers declarator.declaration_list compound_statement
      function_definition: declaration_specifiers declarator.compound_statement
    Reduce:
      * -> [init_declarator]
    Goto:
      TYPE_NAME -> State 21
      TYPEDEF -> State 20
      EXTERN -> State 10
      STATIC -> State 18
      AUTO -> State 5
      REGISTER -> State 15
      CHAR -> State 6
      SHORT -> State 16
      INT -> State 13
      LONG -> State 14
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 11
      DOUBLE -> State 8
      CONST -> State 7
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 9
      '{' -> State 49
      '=' -> State 70
      declaration -> State 51
      declaration_specifiers -> State 53
      storage_class_specifier -> State 34
      type_specifier -> State 38
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      enum_specifier -> State 30
      type_qualifier -> State 37
      compound_statement -> State 71
      declaration_list -> State 72

  State 47:
    Kernel Items:
      init_declarator_list: init_declarator., *
    Reduce:
      * -> [init_declarator_list]
    Goto:
      (nil)

  State 48:
    Kernel Items:
      declaration: declaration_specifiers init_declarator_list.';'
      init_declarator_list: init_declarator_list.',' init_declarator
    Reduce:
      (nil)
    Goto:
      ';' -> State 74
      ',' -> State 73

  State 49:
    Kernel Items:
      compound_statement: '{'.'}'
      compound_statement: '{'.statement_list '}'
      compound_statement: '{'.declaration_list '}'
      compound_statement: '{'.declaration_list statement_list '}'
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 93
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      TYPE_NAME -> State 21
      TYPEDEF -> State 20
      EXTERN -> State 10
      STATIC -> State 18
      AUTO -> State 5
      REGISTER -> State 15
      CHAR -> State 6
      SHORT -> State 16
      INT -> State 13
      LONG -> State 14
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 11
      DOUBLE -> State 8
      CONST -> State 7
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 9
      CASE -> State 85
      DEFAULT -> State 89
      IF -> State 94
      SWITCH -> State 99
      WHILE -> State 100
      DO -> State 90
      FOR -> State 91
      GOTO -> State 92
      CONTINUE -> State 87
      BREAK -> State 84
      RETURN -> State 96
      '(' -> State 77
      '{' -> State 49
      '}' -> State 82
      ';' -> State 81
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 110
      declaration -> State 51
      declaration_specifiers -> State 53
      storage_class_specifier -> State 34
      type_specifier -> State 38
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      enum_specifier -> State 30
      type_qualifier -> State 37
      statement -> State 124
      labeled_statement -> State 115
      compound_statement -> State 105
      declaration_list -> State 107
      statement_list -> State 125
      expression_statement -> State 111
      selection_statement -> State 122
      iteration_statement -> State 113
      jump_statement -> State 114

  State 50:
    Kernel Items:
      function_definition: declarator compound_statement., *
    Reduce:
      * -> [function_definition]
    Goto:
      (nil)

  State 51:
    Kernel Items:
      declaration_list: declaration., *
    Reduce:
      * -> [declaration_list]
    Goto:
      (nil)

  State 52:
    Kernel Items:
      declaration_list: declaration_list.declaration
      function_definition: declarator declaration_list.compound_statement
    Reduce:
      (nil)
    Goto:
      TYPE_NAME -> State 21
      TYPEDEF -> State 20
      EXTERN -> State 10
      STATIC -> State 18
      AUTO -> State 5
      REGISTER -> State 15
      CHAR -> State 6
      SHORT -> State 16
      INT -> State 13
      LONG -> State 14
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 11
      DOUBLE -> State 8
      CONST -> State 7
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 9
      '{' -> State 49
      declaration -> State 129
      declaration_specifiers -> State 53
      storage_class_specifier -> State 34
      type_specifier -> State 38
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      enum_specifier -> State 30
      type_qualifier -> State 37
      compound_statement -> State 128

  State 53:
    Kernel Items:
      declaration: declaration_specifiers.';'
      declaration: declaration_specifiers.init_declarator_list ';'
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 12
      '(' -> State 3
      ';' -> State 45
      '*' -> State 4
      init_declarator_list -> State 48
      init_declarator -> State 47
      declarator -> State 130
      direct_declarator -> State 29
      pointer -> State 33

  State 54:
    Kernel Items:
      direct_declarator: direct_declarator '('.parameter_type_list ')'
      direct_declarator: direct_declarator '('.identifier_list ')'
      direct_declarator: direct_declarator '('.')'
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 132
      TYPE_NAME -> State 21
      TYPEDEF -> State 20
      EXTERN -> State 10
      STATIC -> State 18
      AUTO -> State 5
      REGISTER -> State 15
      CHAR -> State 6
      SHORT -> State 16
      INT -> State 13
      LONG -> State 14
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 11
      DOUBLE -> State 8
      CONST -> State 7
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 9
      ')' -> State 131
      declaration_specifiers -> State 133
      storage_class_specifier -> State 34
      type_specifier -> State 38
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      enum_specifier -> State 30
      type_qualifier -> State 37
      parameter_type_list -> State 137
      parameter_list -> State 136
      parameter_declaration -> State 135
      identifier_list -> State 134

  State 55:
    Kernel Items:
      direct_declarator: direct_declarator '['.constant_expression ']'
      direct_declarator: direct_declarator '['.']'
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      ']' -> State 138
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 142
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 140
      constant_expression -> State 141

  State 56:
    Kernel Items:
      declarator: pointer direct_declarator., *
      direct_declarator: direct_declarator.'[' constant_expression ']'
      direct_declarator: direct_declarator.'[' ']'
      direct_declarator: direct_declarator.'(' parameter_type_list ')'
      direct_declarator: direct_declarator.'(' identifier_list ')'
      direct_declarator: direct_declarator.'(' ')'
    Reduce:
      * -> [declarator]
    Goto:
      '(' -> State 54
      '[' -> State 55

  State 57:
    Kernel Items:
      declaration_specifiers: storage_class_specifier declaration_specifiers., *
    Reduce:
      * -> [declaration_specifiers]
    Goto:
      (nil)

  State 58:
    Kernel Items:
      struct_or_union_specifier: struct_or_union '{'.struct_declaration_list '}'
    Reduce:
      (nil)
    Goto:
      TYPE_NAME -> State 21
      CHAR -> State 6
      SHORT -> State 16
      INT -> State 13
      LONG -> State 14
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 11
      DOUBLE -> State 8
      CONST -> State 7
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 9
      type_specifier -> State 147
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      struct_declaration_list -> State 145
      struct_declaration -> State 144
      specifier_qualifier_list -> State 143
      enum_specifier -> State 30
      type_qualifier -> State 146

  State 59:
    Kernel Items:
      struct_or_union_specifier: struct_or_union IDENTIFIER.'{' struct_declaration_list '}'
      struct_or_union_specifier: struct_or_union IDENTIFIER., *
    Reduce:
      * -> [struct_or_union_specifier]
    Goto:
      '{' -> State 148

  State 60:
    Kernel Items:
      translation_unit: translation_unit external_declaration., *
    Reduce:
      * -> [translation_unit]
    Goto:
      (nil)

  State 61:
    Kernel Items:
      declaration_specifiers: type_qualifier declaration_specifiers., *
    Reduce:
      * -> [declaration_specifiers]
    Goto:
      (nil)

  State 62:
    Kernel Items:
      declaration_specifiers: type_specifier declaration_specifiers., *
    Reduce:
      * -> [declaration_specifiers]
    Goto:
      (nil)

  State 63:
    Kernel Items:
      direct_declarator: '(' declarator ')'., *
    Reduce:
      * -> [direct_declarator]
    Goto:
      (nil)

  State 64:
    Kernel Items:
      pointer: '*' type_qualifier_list pointer., *
    Reduce:
      * -> [pointer]
    Goto:
      (nil)

  State 65:
    Kernel Items:
      type_qualifier_list: type_qualifier_list type_qualifier., *
    Reduce:
      * -> [type_qualifier_list]
    Goto:
      (nil)

  State 66:
    Kernel Items:
      enumerator: IDENTIFIER., *
      enumerator: IDENTIFIER.'=' constant_expression
    Reduce:
      * -> [enumerator]
    Goto:
      '=' -> State 149

  State 67:
    Kernel Items:
      enumerator_list: enumerator., *
    Reduce:
      * -> [enumerator_list]
    Goto:
      (nil)

  State 68:
    Kernel Items:
      enum_specifier: ENUM '{' enumerator_list.'}'
      enumerator_list: enumerator_list.',' enumerator
    Reduce:
      (nil)
    Goto:
      '}' -> State 151
      ',' -> State 150

  State 69:
    Kernel Items:
      enum_specifier: ENUM IDENTIFIER '{'.enumerator_list '}'
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 66
      enumerator_list -> State 152
      enumerator -> State 67

  State 70:
    Kernel Items:
      init_declarator: declarator '='.initializer
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '{' -> State 153
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 154
      initializer -> State 155

  State 71:
    Kernel Items:
      function_definition: declaration_specifiers declarator compound_statement., *
    Reduce:
      * -> [function_definition]
    Goto:
      (nil)

  State 72:
    Kernel Items:
      declaration_list: declaration_list.declaration
      function_definition: declaration_specifiers declarator declaration_list.compound_statement
    Reduce:
      (nil)
    Goto:
      TYPE_NAME -> State 21
      TYPEDEF -> State 20
      EXTERN -> State 10
      STATIC -> State 18
      AUTO -> State 5
      REGISTER -> State 15
      CHAR -> State 6
      SHORT -> State 16
      INT -> State 13
      LONG -> State 14
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 11
      DOUBLE -> State 8
      CONST -> State 7
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 9
      '{' -> State 49
      declaration -> State 129
      declaration_specifiers -> State 53
      storage_class_specifier -> State 34
      type_specifier -> State 38
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      enum_specifier -> State 30
      type_qualifier -> State 37
      compound_statement -> State 156

  State 73:
    Kernel Items:
      init_declarator_list: init_declarator_list ','.init_declarator
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 12
      '(' -> State 3
      '*' -> State 4
      init_declarator -> State 157
      declarator -> State 130
      direct_declarator -> State 29
      pointer -> State 33

  State 74:
    Kernel Items:
      declaration: declaration_specifiers init_declarator_list ';'., *
    Reduce:
      * -> [declaration]
    Goto:
      (nil)

  State 75:
    Kernel Items:
      unary_operator: '!'., *
    Reduce:
      * -> [unary_operator]
    Goto:
      (nil)

  State 76:
    Kernel Items:
      unary_operator: '&'., *
    Reduce:
      * -> [unary_operator]
    Goto:
      (nil)

  State 77:
    Kernel Items:
      primary_expression: '('.expression ')'
      cast_expression: '('.type_name ')' cast_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      TYPE_NAME -> State 21
      CHAR -> State 6
      SHORT -> State 16
      INT -> State 13
      LONG -> State 14
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 11
      DOUBLE -> State 8
      CONST -> State 7
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 9
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 158
      type_specifier -> State 147
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      specifier_qualifier_list -> State 159
      enum_specifier -> State 30
      type_qualifier -> State 146
      type_name -> State 160

  State 78:
    Kernel Items:
      unary_operator: '*'., *
    Reduce:
      * -> [unary_operator]
    Goto:
      (nil)

  State 79:
    Kernel Items:
      unary_operator: '+'., *
    Reduce:
      * -> [unary_operator]
    Goto:
      (nil)

  State 80:
    Kernel Items:
      unary_operator: '-'., *
    Reduce:
      * -> [unary_operator]
    Goto:
      (nil)

  State 81:
    Kernel Items:
      expression_statement: ';'., *
    Reduce:
      * -> [expression_statement]
    Goto:
      (nil)

  State 82:
    Kernel Items:
      compound_statement: '{' '}'., *
    Reduce:
      * -> [compound_statement]
    Goto:
      (nil)

  State 83:
    Kernel Items:
      unary_operator: '~'., *
    Reduce:
      * -> [unary_operator]
    Goto:
      (nil)

  State 84:
    Kernel Items:
      jump_statement: BREAK.';'
    Reduce:
      (nil)
    Goto:
      ';' -> State 161

  State 85:
    Kernel Items:
      labeled_statement: CASE.constant_expression ':' statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 142
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 140
      constant_expression -> State 162

  State 86:
    Kernel Items:
      primary_expression: CONSTANT., *
    Reduce:
      * -> [primary_expression]
    Goto:
      (nil)

  State 87:
    Kernel Items:
      jump_statement: CONTINUE.';'
    Reduce:
      (nil)
    Goto:
      ';' -> State 163

  State 88:
    Kernel Items:
      unary_expression: DEC_OP.unary_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 164
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 165
      unary_operator -> State 127

  State 89:
    Kernel Items:
      labeled_statement: DEFAULT.':' statement
    Reduce:
      (nil)
    Goto:
      ':' -> State 166

  State 90:
    Kernel Items:
      iteration_statement: DO.statement WHILE '(' expression ')' ';'
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 93
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      CASE -> State 85
      DEFAULT -> State 89
      IF -> State 94
      SWITCH -> State 99
      WHILE -> State 100
      DO -> State 90
      FOR -> State 91
      GOTO -> State 92
      CONTINUE -> State 87
      BREAK -> State 84
      RETURN -> State 96
      '(' -> State 77
      '{' -> State 49
      ';' -> State 81
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 110
      statement -> State 167
      labeled_statement -> State 115
      compound_statement -> State 105
      expression_statement -> State 111
      selection_statement -> State 122
      iteration_statement -> State 113
      jump_statement -> State 114

  State 91:
    Kernel Items:
      iteration_statement: FOR.'(' expression_statement expression_statement ')' statement
      iteration_statement: FOR.'(' expression_statement expression_statement expression ')' statement
    Reduce:
      (nil)
    Goto:
      '(' -> State 168

  State 92:
    Kernel Items:
      jump_statement: GOTO.IDENTIFIER ';'
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 169

  State 93:
    Kernel Items:
      primary_expression: IDENTIFIER., *
      labeled_statement: IDENTIFIER.':' statement
    Reduce:
      * -> [primary_expression]
    Goto:
      ':' -> State 170

  State 94:
    Kernel Items:
      selection_statement: IF.'(' expression ')' statement
      selection_statement: IF.'(' expression ')' statement ELSE statement
    Reduce:
      (nil)
    Goto:
      '(' -> State 171

  State 95:
    Kernel Items:
      unary_expression: INC_OP.unary_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 164
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 172
      unary_operator -> State 127

  State 96:
    Kernel Items:
      jump_statement: RETURN.';'
      jump_statement: RETURN.expression ';'
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      ';' -> State 173
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 174

  State 97:
    Kernel Items:
      unary_expression: SIZEOF.unary_expression
      unary_expression: SIZEOF.'(' type_name ')'
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 175
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 176
      unary_operator -> State 127

  State 98:
    Kernel Items:
      primary_expression: STRING_LITERAL., *
    Reduce:
      * -> [primary_expression]
    Goto:
      (nil)

  State 99:
    Kernel Items:
      selection_statement: SWITCH.'(' expression ')' statement
    Reduce:
      (nil)
    Goto:
      '(' -> State 177

  State 100:
    Kernel Items:
      iteration_statement: WHILE.'(' expression ')' statement
    Reduce:
      (nil)
    Goto:
      '(' -> State 178

  State 101:
    Kernel Items:
      additive_expression: additive_expression.'+' multiplicative_expression
      additive_expression: additive_expression.'-' multiplicative_expression
      shift_expression: additive_expression., *
    Reduce:
      * -> [shift_expression]
    Goto:
      '-' -> State 180
      '+' -> State 179

  State 102:
    Kernel Items:
      and_expression: and_expression.'&' equality_expression
      exclusive_or_expression: and_expression., *
    Reduce:
      * -> [exclusive_or_expression]
    Goto:
      '&' -> State 181

  State 103:
    Kernel Items:
      expression: assignment_expression., *
    Reduce:
      * -> [expression]
    Goto:
      (nil)

  State 104:
    Kernel Items:
      multiplicative_expression: cast_expression., *
    Reduce:
      * -> [multiplicative_expression]
    Goto:
      (nil)

  State 105:
    Kernel Items:
      statement: compound_statement., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 106:
    Kernel Items:
      assignment_expression: conditional_expression., *
    Reduce:
      * -> [assignment_expression]
    Goto:
      (nil)

  State 107:
    Kernel Items:
      compound_statement: '{' declaration_list.'}'
      compound_statement: '{' declaration_list.statement_list '}'
      declaration_list: declaration_list.declaration
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 93
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      TYPE_NAME -> State 21
      TYPEDEF -> State 20
      EXTERN -> State 10
      STATIC -> State 18
      AUTO -> State 5
      REGISTER -> State 15
      CHAR -> State 6
      SHORT -> State 16
      INT -> State 13
      LONG -> State 14
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 11
      DOUBLE -> State 8
      CONST -> State 7
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 9
      CASE -> State 85
      DEFAULT -> State 89
      IF -> State 94
      SWITCH -> State 99
      WHILE -> State 100
      DO -> State 90
      FOR -> State 91
      GOTO -> State 92
      CONTINUE -> State 87
      BREAK -> State 84
      RETURN -> State 96
      '(' -> State 77
      '{' -> State 49
      '}' -> State 182
      ';' -> State 81
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 110
      declaration -> State 129
      declaration_specifiers -> State 53
      storage_class_specifier -> State 34
      type_specifier -> State 38
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      enum_specifier -> State 30
      type_qualifier -> State 37
      statement -> State 124
      labeled_statement -> State 115
      compound_statement -> State 105
      statement_list -> State 183
      expression_statement -> State 111
      selection_statement -> State 122
      iteration_statement -> State 113
      jump_statement -> State 114

  State 108:
    Kernel Items:
      equality_expression: equality_expression.EQ_OP relational_expression
      equality_expression: equality_expression.NE_OP relational_expression
      and_expression: equality_expression., *
    Reduce:
      * -> [and_expression]
    Goto:
      EQ_OP -> State 184
      NE_OP -> State 185

  State 109:
    Kernel Items:
      exclusive_or_expression: exclusive_or_expression.'^' and_expression
      inclusive_or_expression: exclusive_or_expression., *
    Reduce:
      * -> [inclusive_or_expression]
    Goto:
      '^' -> State 186

  State 110:
    Kernel Items:
      expression: expression.',' assignment_expression
      expression_statement: expression.';'
    Reduce:
      (nil)
    Goto:
      ';' -> State 188
      ',' -> State 187

  State 111:
    Kernel Items:
      statement: expression_statement., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 112:
    Kernel Items:
      inclusive_or_expression: inclusive_or_expression.'|' exclusive_or_expression
      logical_and_expression: inclusive_or_expression., *
    Reduce:
      * -> [logical_and_expression]
    Goto:
      '|' -> State 189

  State 113:
    Kernel Items:
      statement: iteration_statement., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 114:
    Kernel Items:
      statement: jump_statement., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 115:
    Kernel Items:
      statement: labeled_statement., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 116:
    Kernel Items:
      logical_and_expression: logical_and_expression.AND_OP inclusive_or_expression
      logical_or_expression: logical_and_expression., *
    Reduce:
      * -> [logical_or_expression]
    Goto:
      AND_OP -> State 190

  State 117:
    Kernel Items:
      logical_or_expression: logical_or_expression.OR_OP logical_and_expression
      conditional_expression: logical_or_expression., *
      conditional_expression: logical_or_expression.'?' expression ':' conditional_expression
    Reduce:
      * -> [conditional_expression]
    Goto:
      OR_OP -> State 192
      '?' -> State 191

  State 118:
    Kernel Items:
      multiplicative_expression: multiplicative_expression.'*' cast_expression
      multiplicative_expression: multiplicative_expression.'/' cast_expression
      multiplicative_expression: multiplicative_expression.'%' cast_expression
      additive_expression: multiplicative_expression., *
    Reduce:
      * -> [additive_expression]
    Goto:
      '*' -> State 194
      '/' -> State 195
      '%' -> State 193

  State 119:
    Kernel Items:
      postfix_expression: postfix_expression.'[' expression ']'
      postfix_expression: postfix_expression.'(' ')'
      postfix_expression: postfix_expression.'(' argument_expression_list ')'
      postfix_expression: postfix_expression.'.' IDENTIFIER
      postfix_expression: postfix_expression.PTR_OP IDENTIFIER
      postfix_expression: postfix_expression.INC_OP
      postfix_expression: postfix_expression.DEC_OP
      unary_expression: postfix_expression., *
    Reduce:
      * -> [unary_expression]
    Goto:
      PTR_OP -> State 201
      INC_OP -> State 200
      DEC_OP -> State 199
      '(' -> State 196
      '[' -> State 198
      '.' -> State 197

  State 120:
    Kernel Items:
      postfix_expression: primary_expression., *
    Reduce:
      * -> [postfix_expression]
    Goto:
      (nil)

  State 121:
    Kernel Items:
      relational_expression: relational_expression.'<' shift_expression
      relational_expression: relational_expression.'>' shift_expression
      relational_expression: relational_expression.LE_OP shift_expression
      relational_expression: relational_expression.GE_OP shift_expression
      equality_expression: relational_expression., *
    Reduce:
      * -> [equality_expression]
    Goto:
      LE_OP -> State 205
      GE_OP -> State 204
      '<' -> State 202
      '>' -> State 203

  State 122:
    Kernel Items:
      statement: selection_statement., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 123:
    Kernel Items:
      shift_expression: shift_expression.LEFT_OP additive_expression
      shift_expression: shift_expression.RIGHT_OP additive_expression
      relational_expression: shift_expression., *
    Reduce:
      * -> [relational_expression]
    Goto:
      LEFT_OP -> State 206
      RIGHT_OP -> State 207

  State 124:
    Kernel Items:
      statement_list: statement., *
    Reduce:
      * -> [statement_list]
    Goto:
      (nil)

  State 125:
    Kernel Items:
      compound_statement: '{' statement_list.'}'
      statement_list: statement_list.statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 93
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      CASE -> State 85
      DEFAULT -> State 89
      IF -> State 94
      SWITCH -> State 99
      WHILE -> State 100
      DO -> State 90
      FOR -> State 91
      GOTO -> State 92
      CONTINUE -> State 87
      BREAK -> State 84
      RETURN -> State 96
      '(' -> State 77
      '{' -> State 49
      '}' -> State 208
      ';' -> State 81
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 110
      statement -> State 209
      labeled_statement -> State 115
      compound_statement -> State 105
      expression_statement -> State 111
      selection_statement -> State 122
      iteration_statement -> State 113
      jump_statement -> State 114

  State 126:
    Kernel Items:
      cast_expression: unary_expression., *
      assignment_expression: unary_expression.assignment_operator assignment_expression
    Reduce:
      * -> [cast_expression]
    Goto:
      MUL_ASSIGN -> State 216
      DIV_ASSIGN -> State 213
      MOD_ASSIGN -> State 215
      ADD_ASSIGN -> State 211
      SUB_ASSIGN -> State 219
      LEFT_ASSIGN -> State 214
      RIGHT_ASSIGN -> State 218
      AND_ASSIGN -> State 212
      XOR_ASSIGN -> State 220
      OR_ASSIGN -> State 217
      '=' -> State 210
      assignment_operator -> State 221

  State 127:
    Kernel Items:
      unary_expression: unary_operator.cast_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 142
      unary_operator -> State 127
      cast_expression -> State 222

  State 128:
    Kernel Items:
      function_definition: declarator declaration_list compound_statement., *
    Reduce:
      * -> [function_definition]
    Goto:
      (nil)

  State 129:
    Kernel Items:
      declaration_list: declaration_list declaration., *
    Reduce:
      * -> [declaration_list]
    Goto:
      (nil)

  State 130:
    Kernel Items:
      init_declarator: declarator., *
      init_declarator: declarator.'=' initializer
    Reduce:
      * -> [init_declarator]
    Goto:
      '=' -> State 70

  State 131:
    Kernel Items:
      direct_declarator: direct_declarator '(' ')'., *
    Reduce:
      * -> [direct_declarator]
    Goto:
      (nil)

  State 132:
    Kernel Items:
      identifier_list: IDENTIFIER., *
    Reduce:
      * -> [identifier_list]
    Goto:
      (nil)

  State 133:
    Kernel Items:
      parameter_declaration: declaration_specifiers.declarator
      parameter_declaration: declaration_specifiers.abstract_declarator
      parameter_declaration: declaration_specifiers., *
    Reduce:
      * -> [parameter_declaration]
    Goto:
      IDENTIFIER -> State 12
      '(' -> State 223
      '[' -> State 224
      '*' -> State 4
      declarator -> State 226
      direct_declarator -> State 29
      pointer -> State 228
      abstract_declarator -> State 225
      direct_abstract_declarator -> State 227

  State 134:
    Kernel Items:
      direct_declarator: direct_declarator '(' identifier_list.')'
      identifier_list: identifier_list.',' IDENTIFIER
    Reduce:
      (nil)
    Goto:
      ')' -> State 229
      ',' -> State 230

  State 135:
    Kernel Items:
      parameter_list: parameter_declaration., *
    Reduce:
      * -> [parameter_list]
    Goto:
      (nil)

  State 136:
    Kernel Items:
      parameter_type_list: parameter_list., ')'
      parameter_type_list: parameter_list.',' ELLIPSIS
      parameter_list: parameter_list.',' parameter_declaration
    Reduce:
      ')' -> [parameter_type_list]
    Goto:
      ',' -> State 231

  State 137:
    Kernel Items:
      direct_declarator: direct_declarator '(' parameter_type_list.')'
    Reduce:
      (nil)
    Goto:
      ')' -> State 232

  State 138:
    Kernel Items:
      direct_declarator: direct_declarator '[' ']'., *
    Reduce:
      * -> [direct_declarator]
    Goto:
      (nil)

  State 139:
    Kernel Items:
      primary_expression: IDENTIFIER., *
    Reduce:
      * -> [primary_expression]
    Goto:
      (nil)

  State 140:
    Kernel Items:
      constant_expression: conditional_expression., *
    Reduce:
      * -> [constant_expression]
    Goto:
      (nil)

  State 141:
    Kernel Items:
      direct_declarator: direct_declarator '[' constant_expression.']'
    Reduce:
      (nil)
    Goto:
      ']' -> State 233

  State 142:
    Kernel Items:
      cast_expression: unary_expression., *
    Reduce:
      * -> [cast_expression]
    Goto:
      (nil)

  State 143:
    Kernel Items:
      struct_declaration: specifier_qualifier_list.struct_declarator_list ';'
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 12
      '(' -> State 3
      ':' -> State 234
      '*' -> State 4
      struct_declarator_list -> State 237
      struct_declarator -> State 236
      declarator -> State 235
      direct_declarator -> State 29
      pointer -> State 33

  State 144:
    Kernel Items:
      struct_declaration_list: struct_declaration., *
    Reduce:
      * -> [struct_declaration_list]
    Goto:
      (nil)

  State 145:
    Kernel Items:
      struct_or_union_specifier: struct_or_union '{' struct_declaration_list.'}'
      struct_declaration_list: struct_declaration_list.struct_declaration
    Reduce:
      (nil)
    Goto:
      TYPE_NAME -> State 21
      CHAR -> State 6
      SHORT -> State 16
      INT -> State 13
      LONG -> State 14
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 11
      DOUBLE -> State 8
      CONST -> State 7
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 9
      '}' -> State 238
      type_specifier -> State 147
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      struct_declaration -> State 239
      specifier_qualifier_list -> State 143
      enum_specifier -> State 30
      type_qualifier -> State 146

  State 146:
    Kernel Items:
      specifier_qualifier_list: type_qualifier.specifier_qualifier_list
      specifier_qualifier_list: type_qualifier., *
    Reduce:
      * -> [specifier_qualifier_list]
    Goto:
      TYPE_NAME -> State 21
      CHAR -> State 6
      SHORT -> State 16
      INT -> State 13
      LONG -> State 14
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 11
      DOUBLE -> State 8
      CONST -> State 7
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 9
      type_specifier -> State 147
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      specifier_qualifier_list -> State 240
      enum_specifier -> State 30
      type_qualifier -> State 146

  State 147:
    Kernel Items:
      specifier_qualifier_list: type_specifier.specifier_qualifier_list
      specifier_qualifier_list: type_specifier., *
    Reduce:
      * -> [specifier_qualifier_list]
    Goto:
      TYPE_NAME -> State 21
      CHAR -> State 6
      SHORT -> State 16
      INT -> State 13
      LONG -> State 14
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 11
      DOUBLE -> State 8
      CONST -> State 7
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 9
      type_specifier -> State 147
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      specifier_qualifier_list -> State 241
      enum_specifier -> State 30
      type_qualifier -> State 146

  State 148:
    Kernel Items:
      struct_or_union_specifier: struct_or_union IDENTIFIER '{'.struct_declaration_list '}'
    Reduce:
      (nil)
    Goto:
      TYPE_NAME -> State 21
      CHAR -> State 6
      SHORT -> State 16
      INT -> State 13
      LONG -> State 14
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 11
      DOUBLE -> State 8
      CONST -> State 7
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 9
      type_specifier -> State 147
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      struct_declaration_list -> State 242
      struct_declaration -> State 144
      specifier_qualifier_list -> State 143
      enum_specifier -> State 30
      type_qualifier -> State 146

  State 149:
    Kernel Items:
      enumerator: IDENTIFIER '='.constant_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 142
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 140
      constant_expression -> State 243

  State 150:
    Kernel Items:
      enumerator_list: enumerator_list ','.enumerator
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 66
      enumerator -> State 244

  State 151:
    Kernel Items:
      enum_specifier: ENUM '{' enumerator_list '}'., *
    Reduce:
      * -> [enum_specifier]
    Goto:
      (nil)

  State 152:
    Kernel Items:
      enum_specifier: ENUM IDENTIFIER '{' enumerator_list.'}'
      enumerator_list: enumerator_list.',' enumerator
    Reduce:
      (nil)
    Goto:
      '}' -> State 245
      ',' -> State 150

  State 153:
    Kernel Items:
      initializer: '{'.initializer_list '}'
      initializer: '{'.initializer_list ',' '}'
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '{' -> State 153
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 154
      initializer -> State 246
      initializer_list -> State 247

  State 154:
    Kernel Items:
      initializer: assignment_expression., *
    Reduce:
      * -> [initializer]
    Goto:
      (nil)

  State 155:
    Kernel Items:
      init_declarator: declarator '=' initializer., *
    Reduce:
      * -> [init_declarator]
    Goto:
      (nil)

  State 156:
    Kernel Items:
      function_definition: declaration_specifiers declarator declaration_list compound_statement., *
    Reduce:
      * -> [function_definition]
    Goto:
      (nil)

  State 157:
    Kernel Items:
      init_declarator_list: init_declarator_list ',' init_declarator., *
    Reduce:
      * -> [init_declarator_list]
    Goto:
      (nil)

  State 158:
    Kernel Items:
      primary_expression: '(' expression.')'
      expression: expression.',' assignment_expression
    Reduce:
      (nil)
    Goto:
      ')' -> State 248
      ',' -> State 187

  State 159:
    Kernel Items:
      type_name: specifier_qualifier_list., ')'
      type_name: specifier_qualifier_list.abstract_declarator
    Reduce:
      ')' -> [type_name]
    Goto:
      '(' -> State 249
      '[' -> State 224
      '*' -> State 4
      pointer -> State 251
      abstract_declarator -> State 250
      direct_abstract_declarator -> State 227

  State 160:
    Kernel Items:
      cast_expression: '(' type_name.')' cast_expression
    Reduce:
      (nil)
    Goto:
      ')' -> State 252

  State 161:
    Kernel Items:
      jump_statement: BREAK ';'., *
    Reduce:
      * -> [jump_statement]
    Goto:
      (nil)

  State 162:
    Kernel Items:
      labeled_statement: CASE constant_expression.':' statement
    Reduce:
      (nil)
    Goto:
      ':' -> State 253

  State 163:
    Kernel Items:
      jump_statement: CONTINUE ';'., *
    Reduce:
      * -> [jump_statement]
    Goto:
      (nil)

  State 164:
    Kernel Items:
      primary_expression: '('.expression ')'
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 158

  State 165:
    Kernel Items:
      unary_expression: DEC_OP unary_expression., *
    Reduce:
      * -> [unary_expression]
    Goto:
      (nil)

  State 166:
    Kernel Items:
      labeled_statement: DEFAULT ':'.statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 93
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      CASE -> State 85
      DEFAULT -> State 89
      IF -> State 94
      SWITCH -> State 99
      WHILE -> State 100
      DO -> State 90
      FOR -> State 91
      GOTO -> State 92
      CONTINUE -> State 87
      BREAK -> State 84
      RETURN -> State 96
      '(' -> State 77
      '{' -> State 49
      ';' -> State 81
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 110
      statement -> State 254
      labeled_statement -> State 115
      compound_statement -> State 105
      expression_statement -> State 111
      selection_statement -> State 122
      iteration_statement -> State 113
      jump_statement -> State 114

  State 167:
    Kernel Items:
      iteration_statement: DO statement.WHILE '(' expression ')' ';'
    Reduce:
      (nil)
    Goto:
      WHILE -> State 255

  State 168:
    Kernel Items:
      iteration_statement: FOR '('.expression_statement expression_statement ')' statement
      iteration_statement: FOR '('.expression_statement expression_statement expression ')' statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      ';' -> State 81
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 110
      expression_statement -> State 256

  State 169:
    Kernel Items:
      jump_statement: GOTO IDENTIFIER.';'
    Reduce:
      (nil)
    Goto:
      ';' -> State 257

  State 170:
    Kernel Items:
      labeled_statement: IDENTIFIER ':'.statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 93
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      CASE -> State 85
      DEFAULT -> State 89
      IF -> State 94
      SWITCH -> State 99
      WHILE -> State 100
      DO -> State 90
      FOR -> State 91
      GOTO -> State 92
      CONTINUE -> State 87
      BREAK -> State 84
      RETURN -> State 96
      '(' -> State 77
      '{' -> State 49
      ';' -> State 81
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 110
      statement -> State 258
      labeled_statement -> State 115
      compound_statement -> State 105
      expression_statement -> State 111
      selection_statement -> State 122
      iteration_statement -> State 113
      jump_statement -> State 114

  State 171:
    Kernel Items:
      selection_statement: IF '('.expression ')' statement
      selection_statement: IF '('.expression ')' statement ELSE statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 259

  State 172:
    Kernel Items:
      unary_expression: INC_OP unary_expression., *
    Reduce:
      * -> [unary_expression]
    Goto:
      (nil)

  State 173:
    Kernel Items:
      jump_statement: RETURN ';'., *
    Reduce:
      * -> [jump_statement]
    Goto:
      (nil)

  State 174:
    Kernel Items:
      expression: expression.',' assignment_expression
      jump_statement: RETURN expression.';'
    Reduce:
      (nil)
    Goto:
      ';' -> State 260
      ',' -> State 187

  State 175:
    Kernel Items:
      primary_expression: '('.expression ')'
      unary_expression: SIZEOF '('.type_name ')'
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      TYPE_NAME -> State 21
      CHAR -> State 6
      SHORT -> State 16
      INT -> State 13
      LONG -> State 14
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 11
      DOUBLE -> State 8
      CONST -> State 7
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 9
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 158
      type_specifier -> State 147
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      specifier_qualifier_list -> State 159
      enum_specifier -> State 30
      type_qualifier -> State 146
      type_name -> State 261

  State 176:
    Kernel Items:
      unary_expression: SIZEOF unary_expression., *
    Reduce:
      * -> [unary_expression]
    Goto:
      (nil)

  State 177:
    Kernel Items:
      selection_statement: SWITCH '('.expression ')' statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 262

  State 178:
    Kernel Items:
      iteration_statement: WHILE '('.expression ')' statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 263

  State 179:
    Kernel Items:
      additive_expression: additive_expression '+'.multiplicative_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 142
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 264

  State 180:
    Kernel Items:
      additive_expression: additive_expression '-'.multiplicative_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 142
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 265

  State 181:
    Kernel Items:
      and_expression: and_expression '&'.equality_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 142
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 266

  State 182:
    Kernel Items:
      compound_statement: '{' declaration_list '}'., *
    Reduce:
      * -> [compound_statement]
    Goto:
      (nil)

  State 183:
    Kernel Items:
      compound_statement: '{' declaration_list statement_list.'}'
      statement_list: statement_list.statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 93
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      CASE -> State 85
      DEFAULT -> State 89
      IF -> State 94
      SWITCH -> State 99
      WHILE -> State 100
      DO -> State 90
      FOR -> State 91
      GOTO -> State 92
      CONTINUE -> State 87
      BREAK -> State 84
      RETURN -> State 96
      '(' -> State 77
      '{' -> State 49
      '}' -> State 267
      ';' -> State 81
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 110
      statement -> State 209
      labeled_statement -> State 115
      compound_statement -> State 105
      expression_statement -> State 111
      selection_statement -> State 122
      iteration_statement -> State 113
      jump_statement -> State 114

  State 184:
    Kernel Items:
      equality_expression: equality_expression EQ_OP.relational_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 142
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 268

  State 185:
    Kernel Items:
      equality_expression: equality_expression NE_OP.relational_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 142
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 269

  State 186:
    Kernel Items:
      exclusive_or_expression: exclusive_or_expression '^'.and_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 142
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 270

  State 187:
    Kernel Items:
      expression: expression ','.assignment_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 271

  State 188:
    Kernel Items:
      expression_statement: expression ';'., *
    Reduce:
      * -> [expression_statement]
    Goto:
      (nil)

  State 189:
    Kernel Items:
      inclusive_or_expression: inclusive_or_expression '|'.exclusive_or_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 142
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 272

  State 190:
    Kernel Items:
      logical_and_expression: logical_and_expression AND_OP.inclusive_or_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 142
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 273

  State 191:
    Kernel Items:
      conditional_expression: logical_or_expression '?'.expression ':' conditional_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 274

  State 192:
    Kernel Items:
      logical_or_expression: logical_or_expression OR_OP.logical_and_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 142
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 275

  State 193:
    Kernel Items:
      multiplicative_expression: multiplicative_expression '%'.cast_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 142
      unary_operator -> State 127
      cast_expression -> State 276

  State 194:
    Kernel Items:
      multiplicative_expression: multiplicative_expression '*'.cast_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 142
      unary_operator -> State 127
      cast_expression -> State 277

  State 195:
    Kernel Items:
      multiplicative_expression: multiplicative_expression '/'.cast_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 142
      unary_operator -> State 127
      cast_expression -> State 278

  State 196:
    Kernel Items:
      postfix_expression: postfix_expression '('.')'
      postfix_expression: postfix_expression '('.argument_expression_list ')'
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      ')' -> State 279
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      argument_expression_list -> State 280
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 281

  State 197:
    Kernel Items:
      postfix_expression: postfix_expression '.'.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 282

  State 198:
    Kernel Items:
      postfix_expression: postfix_expression '['.expression ']'
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 283

  State 199:
    Kernel Items:
      postfix_expression: postfix_expression DEC_OP., *
    Reduce:
      * -> [postfix_expression]
    Goto:
      (nil)

  State 200:
    Kernel Items:
      postfix_expression: postfix_expression INC_OP., *
    Reduce:
      * -> [postfix_expression]
    Goto:
      (nil)

  State 201:
    Kernel Items:
      postfix_expression: postfix_expression PTR_OP.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 284

  State 202:
    Kernel Items:
      relational_expression: relational_expression '<'.shift_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 142
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 285

  State 203:
    Kernel Items:
      relational_expression: relational_expression '>'.shift_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 142
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 286

  State 204:
    Kernel Items:
      relational_expression: relational_expression GE_OP.shift_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 142
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 287

  State 205:
    Kernel Items:
      relational_expression: relational_expression LE_OP.shift_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 142
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 288

  State 206:
    Kernel Items:
      shift_expression: shift_expression LEFT_OP.additive_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 142
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 289

  State 207:
    Kernel Items:
      shift_expression: shift_expression RIGHT_OP.additive_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 142
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 290

  State 208:
    Kernel Items:
      compound_statement: '{' statement_list '}'., *
    Reduce:
      * -> [compound_statement]
    Goto:
      (nil)

  State 209:
    Kernel Items:
      statement_list: statement_list statement., *
    Reduce:
      * -> [statement_list]
    Goto:
      (nil)

  State 210:
    Kernel Items:
      assignment_operator: '='., *
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 211:
    Kernel Items:
      assignment_operator: ADD_ASSIGN., *
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 212:
    Kernel Items:
      assignment_operator: AND_ASSIGN., *
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 213:
    Kernel Items:
      assignment_operator: DIV_ASSIGN., *
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 214:
    Kernel Items:
      assignment_operator: LEFT_ASSIGN., *
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 215:
    Kernel Items:
      assignment_operator: MOD_ASSIGN., *
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 216:
    Kernel Items:
      assignment_operator: MUL_ASSIGN., *
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 217:
    Kernel Items:
      assignment_operator: OR_ASSIGN., *
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 218:
    Kernel Items:
      assignment_operator: RIGHT_ASSIGN., *
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 219:
    Kernel Items:
      assignment_operator: SUB_ASSIGN., *
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 220:
    Kernel Items:
      assignment_operator: XOR_ASSIGN., *
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 221:
    Kernel Items:
      assignment_expression: unary_expression assignment_operator.assignment_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 291

  State 222:
    Kernel Items:
      unary_expression: unary_operator cast_expression., *
    Reduce:
      * -> [unary_expression]
    Goto:
      (nil)

  State 223:
    Kernel Items:
      direct_declarator: '('.declarator ')'
      direct_abstract_declarator: '('.abstract_declarator ')'
      direct_abstract_declarator: '('.')'
      direct_abstract_declarator: '('.parameter_type_list ')'
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 12
      TYPE_NAME -> State 21
      TYPEDEF -> State 20
      EXTERN -> State 10
      STATIC -> State 18
      AUTO -> State 5
      REGISTER -> State 15
      CHAR -> State 6
      SHORT -> State 16
      INT -> State 13
      LONG -> State 14
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 11
      DOUBLE -> State 8
      CONST -> State 7
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 9
      '(' -> State 223
      ')' -> State 292
      '[' -> State 224
      '*' -> State 4
      declaration_specifiers -> State 133
      storage_class_specifier -> State 34
      type_specifier -> State 38
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      enum_specifier -> State 30
      type_qualifier -> State 37
      declarator -> State 39
      direct_declarator -> State 29
      pointer -> State 228
      parameter_type_list -> State 294
      parameter_list -> State 136
      parameter_declaration -> State 135
      abstract_declarator -> State 293
      direct_abstract_declarator -> State 227

  State 224:
    Kernel Items:
      direct_abstract_declarator: '['.']'
      direct_abstract_declarator: '['.constant_expression ']'
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      ']' -> State 295
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 142
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 140
      constant_expression -> State 296

  State 225:
    Kernel Items:
      parameter_declaration: declaration_specifiers abstract_declarator., *
    Reduce:
      * -> [parameter_declaration]
    Goto:
      (nil)

  State 226:
    Kernel Items:
      parameter_declaration: declaration_specifiers declarator., *
    Reduce:
      * -> [parameter_declaration]
    Goto:
      (nil)

  State 227:
    Kernel Items:
      abstract_declarator: direct_abstract_declarator., *
      direct_abstract_declarator: direct_abstract_declarator.'[' ']'
      direct_abstract_declarator: direct_abstract_declarator.'[' constant_expression ']'
      direct_abstract_declarator: direct_abstract_declarator.'(' ')'
      direct_abstract_declarator: direct_abstract_declarator.'(' parameter_type_list ')'
    Reduce:
      * -> [abstract_declarator]
    Goto:
      '(' -> State 297
      '[' -> State 298

  State 228:
    Kernel Items:
      declarator: pointer.direct_declarator
      abstract_declarator: pointer., *
      abstract_declarator: pointer.direct_abstract_declarator
    Reduce:
      * -> [abstract_declarator]
    Goto:
      IDENTIFIER -> State 12
      '(' -> State 223
      '[' -> State 224
      direct_declarator -> State 56
      direct_abstract_declarator -> State 299

  State 229:
    Kernel Items:
      direct_declarator: direct_declarator '(' identifier_list ')'., *
    Reduce:
      * -> [direct_declarator]
    Goto:
      (nil)

  State 230:
    Kernel Items:
      identifier_list: identifier_list ','.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 300

  State 231:
    Kernel Items:
      parameter_type_list: parameter_list ','.ELLIPSIS
      parameter_list: parameter_list ','.parameter_declaration
    Reduce:
      (nil)
    Goto:
      TYPE_NAME -> State 21
      TYPEDEF -> State 20
      EXTERN -> State 10
      STATIC -> State 18
      AUTO -> State 5
      REGISTER -> State 15
      CHAR -> State 6
      SHORT -> State 16
      INT -> State 13
      LONG -> State 14
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 11
      DOUBLE -> State 8
      CONST -> State 7
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 9
      ELLIPSIS -> State 301
      declaration_specifiers -> State 133
      storage_class_specifier -> State 34
      type_specifier -> State 38
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      enum_specifier -> State 30
      type_qualifier -> State 37
      parameter_declaration -> State 302

  State 232:
    Kernel Items:
      direct_declarator: direct_declarator '(' parameter_type_list ')'., *
    Reduce:
      * -> [direct_declarator]
    Goto:
      (nil)

  State 233:
    Kernel Items:
      direct_declarator: direct_declarator '[' constant_expression ']'., *
    Reduce:
      * -> [direct_declarator]
    Goto:
      (nil)

  State 234:
    Kernel Items:
      struct_declarator: ':'.constant_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 142
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 140
      constant_expression -> State 303

  State 235:
    Kernel Items:
      struct_declarator: declarator., *
      struct_declarator: declarator.':' constant_expression
    Reduce:
      * -> [struct_declarator]
    Goto:
      ':' -> State 304

  State 236:
    Kernel Items:
      struct_declarator_list: struct_declarator., *
    Reduce:
      * -> [struct_declarator_list]
    Goto:
      (nil)

  State 237:
    Kernel Items:
      struct_declaration: specifier_qualifier_list struct_declarator_list.';'
      struct_declarator_list: struct_declarator_list.',' struct_declarator
    Reduce:
      (nil)
    Goto:
      ';' -> State 306
      ',' -> State 305

  State 238:
    Kernel Items:
      struct_or_union_specifier: struct_or_union '{' struct_declaration_list '}'., *
    Reduce:
      * -> [struct_or_union_specifier]
    Goto:
      (nil)

  State 239:
    Kernel Items:
      struct_declaration_list: struct_declaration_list struct_declaration., *
    Reduce:
      * -> [struct_declaration_list]
    Goto:
      (nil)

  State 240:
    Kernel Items:
      specifier_qualifier_list: type_qualifier specifier_qualifier_list., *
    Reduce:
      * -> [specifier_qualifier_list]
    Goto:
      (nil)

  State 241:
    Kernel Items:
      specifier_qualifier_list: type_specifier specifier_qualifier_list., *
    Reduce:
      * -> [specifier_qualifier_list]
    Goto:
      (nil)

  State 242:
    Kernel Items:
      struct_or_union_specifier: struct_or_union IDENTIFIER '{' struct_declaration_list.'}'
      struct_declaration_list: struct_declaration_list.struct_declaration
    Reduce:
      (nil)
    Goto:
      TYPE_NAME -> State 21
      CHAR -> State 6
      SHORT -> State 16
      INT -> State 13
      LONG -> State 14
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 11
      DOUBLE -> State 8
      CONST -> State 7
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 9
      '}' -> State 307
      type_specifier -> State 147
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      struct_declaration -> State 239
      specifier_qualifier_list -> State 143
      enum_specifier -> State 30
      type_qualifier -> State 146

  State 243:
    Kernel Items:
      enumerator: IDENTIFIER '=' constant_expression., *
    Reduce:
      * -> [enumerator]
    Goto:
      (nil)

  State 244:
    Kernel Items:
      enumerator_list: enumerator_list ',' enumerator., *
    Reduce:
      * -> [enumerator_list]
    Goto:
      (nil)

  State 245:
    Kernel Items:
      enum_specifier: ENUM IDENTIFIER '{' enumerator_list '}'., *
    Reduce:
      * -> [enum_specifier]
    Goto:
      (nil)

  State 246:
    Kernel Items:
      initializer_list: initializer., *
    Reduce:
      * -> [initializer_list]
    Goto:
      (nil)

  State 247:
    Kernel Items:
      initializer: '{' initializer_list.'}'
      initializer: '{' initializer_list.',' '}'
      initializer_list: initializer_list.',' initializer
    Reduce:
      (nil)
    Goto:
      '}' -> State 309
      ',' -> State 308

  State 248:
    Kernel Items:
      primary_expression: '(' expression ')'., *
    Reduce:
      * -> [primary_expression]
    Goto:
      (nil)

  State 249:
    Kernel Items:
      direct_abstract_declarator: '('.abstract_declarator ')'
      direct_abstract_declarator: '('.')'
      direct_abstract_declarator: '('.parameter_type_list ')'
    Reduce:
      (nil)
    Goto:
      TYPE_NAME -> State 21
      TYPEDEF -> State 20
      EXTERN -> State 10
      STATIC -> State 18
      AUTO -> State 5
      REGISTER -> State 15
      CHAR -> State 6
      SHORT -> State 16
      INT -> State 13
      LONG -> State 14
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 11
      DOUBLE -> State 8
      CONST -> State 7
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 9
      '(' -> State 249
      ')' -> State 292
      '[' -> State 224
      '*' -> State 4
      declaration_specifiers -> State 133
      storage_class_specifier -> State 34
      type_specifier -> State 38
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      enum_specifier -> State 30
      type_qualifier -> State 37
      pointer -> State 251
      parameter_type_list -> State 294
      parameter_list -> State 136
      parameter_declaration -> State 135
      abstract_declarator -> State 293
      direct_abstract_declarator -> State 227

  State 250:
    Kernel Items:
      type_name: specifier_qualifier_list abstract_declarator., ')'
    Reduce:
      ')' -> [type_name]
    Goto:
      (nil)

  State 251:
    Kernel Items:
      abstract_declarator: pointer., ')'
      abstract_declarator: pointer.direct_abstract_declarator
    Reduce:
      ')' -> [abstract_declarator]
    Goto:
      '(' -> State 249
      '[' -> State 224
      direct_abstract_declarator -> State 299

  State 252:
    Kernel Items:
      cast_expression: '(' type_name ')'.cast_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 142
      unary_operator -> State 127
      cast_expression -> State 310

  State 253:
    Kernel Items:
      labeled_statement: CASE constant_expression ':'.statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 93
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      CASE -> State 85
      DEFAULT -> State 89
      IF -> State 94
      SWITCH -> State 99
      WHILE -> State 100
      DO -> State 90
      FOR -> State 91
      GOTO -> State 92
      CONTINUE -> State 87
      BREAK -> State 84
      RETURN -> State 96
      '(' -> State 77
      '{' -> State 49
      ';' -> State 81
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 110
      statement -> State 311
      labeled_statement -> State 115
      compound_statement -> State 105
      expression_statement -> State 111
      selection_statement -> State 122
      iteration_statement -> State 113
      jump_statement -> State 114

  State 254:
    Kernel Items:
      labeled_statement: DEFAULT ':' statement., *
    Reduce:
      * -> [labeled_statement]
    Goto:
      (nil)

  State 255:
    Kernel Items:
      iteration_statement: DO statement WHILE.'(' expression ')' ';'
    Reduce:
      (nil)
    Goto:
      '(' -> State 312

  State 256:
    Kernel Items:
      iteration_statement: FOR '(' expression_statement.expression_statement ')' statement
      iteration_statement: FOR '(' expression_statement.expression_statement expression ')' statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      ';' -> State 81
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 110
      expression_statement -> State 313

  State 257:
    Kernel Items:
      jump_statement: GOTO IDENTIFIER ';'., *
    Reduce:
      * -> [jump_statement]
    Goto:
      (nil)

  State 258:
    Kernel Items:
      labeled_statement: IDENTIFIER ':' statement., *
    Reduce:
      * -> [labeled_statement]
    Goto:
      (nil)

  State 259:
    Kernel Items:
      expression: expression.',' assignment_expression
      selection_statement: IF '(' expression.')' statement
      selection_statement: IF '(' expression.')' statement ELSE statement
    Reduce:
      (nil)
    Goto:
      ')' -> State 314
      ',' -> State 187

  State 260:
    Kernel Items:
      jump_statement: RETURN expression ';'., *
    Reduce:
      * -> [jump_statement]
    Goto:
      (nil)

  State 261:
    Kernel Items:
      unary_expression: SIZEOF '(' type_name.')'
    Reduce:
      (nil)
    Goto:
      ')' -> State 315

  State 262:
    Kernel Items:
      expression: expression.',' assignment_expression
      selection_statement: SWITCH '(' expression.')' statement
    Reduce:
      (nil)
    Goto:
      ')' -> State 316
      ',' -> State 187

  State 263:
    Kernel Items:
      expression: expression.',' assignment_expression
      iteration_statement: WHILE '(' expression.')' statement
    Reduce:
      (nil)
    Goto:
      ')' -> State 317
      ',' -> State 187

  State 264:
    Kernel Items:
      multiplicative_expression: multiplicative_expression.'*' cast_expression
      multiplicative_expression: multiplicative_expression.'/' cast_expression
      multiplicative_expression: multiplicative_expression.'%' cast_expression
      additive_expression: additive_expression '+' multiplicative_expression., *
    Reduce:
      * -> [additive_expression]
    Goto:
      '*' -> State 194
      '/' -> State 195
      '%' -> State 193

  State 265:
    Kernel Items:
      multiplicative_expression: multiplicative_expression.'*' cast_expression
      multiplicative_expression: multiplicative_expression.'/' cast_expression
      multiplicative_expression: multiplicative_expression.'%' cast_expression
      additive_expression: additive_expression '-' multiplicative_expression., *
    Reduce:
      * -> [additive_expression]
    Goto:
      '*' -> State 194
      '/' -> State 195
      '%' -> State 193

  State 266:
    Kernel Items:
      equality_expression: equality_expression.EQ_OP relational_expression
      equality_expression: equality_expression.NE_OP relational_expression
      and_expression: and_expression '&' equality_expression., *
    Reduce:
      * -> [and_expression]
    Goto:
      EQ_OP -> State 184
      NE_OP -> State 185

  State 267:
    Kernel Items:
      compound_statement: '{' declaration_list statement_list '}'., *
    Reduce:
      * -> [compound_statement]
    Goto:
      (nil)

  State 268:
    Kernel Items:
      relational_expression: relational_expression.'<' shift_expression
      relational_expression: relational_expression.'>' shift_expression
      relational_expression: relational_expression.LE_OP shift_expression
      relational_expression: relational_expression.GE_OP shift_expression
      equality_expression: equality_expression EQ_OP relational_expression., *
    Reduce:
      * -> [equality_expression]
    Goto:
      LE_OP -> State 205
      GE_OP -> State 204
      '<' -> State 202
      '>' -> State 203

  State 269:
    Kernel Items:
      relational_expression: relational_expression.'<' shift_expression
      relational_expression: relational_expression.'>' shift_expression
      relational_expression: relational_expression.LE_OP shift_expression
      relational_expression: relational_expression.GE_OP shift_expression
      equality_expression: equality_expression NE_OP relational_expression., *
    Reduce:
      * -> [equality_expression]
    Goto:
      LE_OP -> State 205
      GE_OP -> State 204
      '<' -> State 202
      '>' -> State 203

  State 270:
    Kernel Items:
      and_expression: and_expression.'&' equality_expression
      exclusive_or_expression: exclusive_or_expression '^' and_expression., *
    Reduce:
      * -> [exclusive_or_expression]
    Goto:
      '&' -> State 181

  State 271:
    Kernel Items:
      expression: expression ',' assignment_expression., *
    Reduce:
      * -> [expression]
    Goto:
      (nil)

  State 272:
    Kernel Items:
      exclusive_or_expression: exclusive_or_expression.'^' and_expression
      inclusive_or_expression: inclusive_or_expression '|' exclusive_or_expression., *
    Reduce:
      * -> [inclusive_or_expression]
    Goto:
      '^' -> State 186

  State 273:
    Kernel Items:
      inclusive_or_expression: inclusive_or_expression.'|' exclusive_or_expression
      logical_and_expression: logical_and_expression AND_OP inclusive_or_expression., *
    Reduce:
      * -> [logical_and_expression]
    Goto:
      '|' -> State 189

  State 274:
    Kernel Items:
      conditional_expression: logical_or_expression '?' expression.':' conditional_expression
      expression: expression.',' assignment_expression
    Reduce:
      (nil)
    Goto:
      ':' -> State 318
      ',' -> State 187

  State 275:
    Kernel Items:
      logical_and_expression: logical_and_expression.AND_OP inclusive_or_expression
      logical_or_expression: logical_or_expression OR_OP logical_and_expression., *
    Reduce:
      * -> [logical_or_expression]
    Goto:
      AND_OP -> State 190

  State 276:
    Kernel Items:
      multiplicative_expression: multiplicative_expression '%' cast_expression., *
    Reduce:
      * -> [multiplicative_expression]
    Goto:
      (nil)

  State 277:
    Kernel Items:
      multiplicative_expression: multiplicative_expression '*' cast_expression., *
    Reduce:
      * -> [multiplicative_expression]
    Goto:
      (nil)

  State 278:
    Kernel Items:
      multiplicative_expression: multiplicative_expression '/' cast_expression., *
    Reduce:
      * -> [multiplicative_expression]
    Goto:
      (nil)

  State 279:
    Kernel Items:
      postfix_expression: postfix_expression '(' ')'., *
    Reduce:
      * -> [postfix_expression]
    Goto:
      (nil)

  State 280:
    Kernel Items:
      postfix_expression: postfix_expression '(' argument_expression_list.')'
      argument_expression_list: argument_expression_list.',' assignment_expression
    Reduce:
      (nil)
    Goto:
      ')' -> State 319
      ',' -> State 320

  State 281:
    Kernel Items:
      argument_expression_list: assignment_expression., *
    Reduce:
      * -> [argument_expression_list]
    Goto:
      (nil)

  State 282:
    Kernel Items:
      postfix_expression: postfix_expression '.' IDENTIFIER., *
    Reduce:
      * -> [postfix_expression]
    Goto:
      (nil)

  State 283:
    Kernel Items:
      postfix_expression: postfix_expression '[' expression.']'
      expression: expression.',' assignment_expression
    Reduce:
      (nil)
    Goto:
      ']' -> State 321
      ',' -> State 187

  State 284:
    Kernel Items:
      postfix_expression: postfix_expression PTR_OP IDENTIFIER., *
    Reduce:
      * -> [postfix_expression]
    Goto:
      (nil)

  State 285:
    Kernel Items:
      shift_expression: shift_expression.LEFT_OP additive_expression
      shift_expression: shift_expression.RIGHT_OP additive_expression
      relational_expression: relational_expression '<' shift_expression., *
    Reduce:
      * -> [relational_expression]
    Goto:
      LEFT_OP -> State 206
      RIGHT_OP -> State 207

  State 286:
    Kernel Items:
      shift_expression: shift_expression.LEFT_OP additive_expression
      shift_expression: shift_expression.RIGHT_OP additive_expression
      relational_expression: relational_expression '>' shift_expression., *
    Reduce:
      * -> [relational_expression]
    Goto:
      LEFT_OP -> State 206
      RIGHT_OP -> State 207

  State 287:
    Kernel Items:
      shift_expression: shift_expression.LEFT_OP additive_expression
      shift_expression: shift_expression.RIGHT_OP additive_expression
      relational_expression: relational_expression GE_OP shift_expression., *
    Reduce:
      * -> [relational_expression]
    Goto:
      LEFT_OP -> State 206
      RIGHT_OP -> State 207

  State 288:
    Kernel Items:
      shift_expression: shift_expression.LEFT_OP additive_expression
      shift_expression: shift_expression.RIGHT_OP additive_expression
      relational_expression: relational_expression LE_OP shift_expression., *
    Reduce:
      * -> [relational_expression]
    Goto:
      LEFT_OP -> State 206
      RIGHT_OP -> State 207

  State 289:
    Kernel Items:
      additive_expression: additive_expression.'+' multiplicative_expression
      additive_expression: additive_expression.'-' multiplicative_expression
      shift_expression: shift_expression LEFT_OP additive_expression., *
    Reduce:
      * -> [shift_expression]
    Goto:
      '-' -> State 180
      '+' -> State 179

  State 290:
    Kernel Items:
      additive_expression: additive_expression.'+' multiplicative_expression
      additive_expression: additive_expression.'-' multiplicative_expression
      shift_expression: shift_expression RIGHT_OP additive_expression., *
    Reduce:
      * -> [shift_expression]
    Goto:
      '-' -> State 180
      '+' -> State 179

  State 291:
    Kernel Items:
      assignment_expression: unary_expression assignment_operator assignment_expression., *
    Reduce:
      * -> [assignment_expression]
    Goto:
      (nil)

  State 292:
    Kernel Items:
      direct_abstract_declarator: '(' ')'., *
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 293:
    Kernel Items:
      direct_abstract_declarator: '(' abstract_declarator.')'
    Reduce:
      (nil)
    Goto:
      ')' -> State 322

  State 294:
    Kernel Items:
      direct_abstract_declarator: '(' parameter_type_list.')'
    Reduce:
      (nil)
    Goto:
      ')' -> State 323

  State 295:
    Kernel Items:
      direct_abstract_declarator: '[' ']'., *
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 296:
    Kernel Items:
      direct_abstract_declarator: '[' constant_expression.']'
    Reduce:
      (nil)
    Goto:
      ']' -> State 324

  State 297:
    Kernel Items:
      direct_abstract_declarator: direct_abstract_declarator '('.')'
      direct_abstract_declarator: direct_abstract_declarator '('.parameter_type_list ')'
    Reduce:
      (nil)
    Goto:
      TYPE_NAME -> State 21
      TYPEDEF -> State 20
      EXTERN -> State 10
      STATIC -> State 18
      AUTO -> State 5
      REGISTER -> State 15
      CHAR -> State 6
      SHORT -> State 16
      INT -> State 13
      LONG -> State 14
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 11
      DOUBLE -> State 8
      CONST -> State 7
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 9
      ')' -> State 325
      declaration_specifiers -> State 133
      storage_class_specifier -> State 34
      type_specifier -> State 38
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      enum_specifier -> State 30
      type_qualifier -> State 37
      parameter_type_list -> State 326
      parameter_list -> State 136
      parameter_declaration -> State 135

  State 298:
    Kernel Items:
      direct_abstract_declarator: direct_abstract_declarator '['.']'
      direct_abstract_declarator: direct_abstract_declarator '['.constant_expression ']'
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      ']' -> State 327
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 142
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 140
      constant_expression -> State 328

  State 299:
    Kernel Items:
      abstract_declarator: pointer direct_abstract_declarator., *
      direct_abstract_declarator: direct_abstract_declarator.'[' ']'
      direct_abstract_declarator: direct_abstract_declarator.'[' constant_expression ']'
      direct_abstract_declarator: direct_abstract_declarator.'(' ')'
      direct_abstract_declarator: direct_abstract_declarator.'(' parameter_type_list ')'
    Reduce:
      * -> [abstract_declarator]
    Goto:
      '(' -> State 297
      '[' -> State 298

  State 300:
    Kernel Items:
      identifier_list: identifier_list ',' IDENTIFIER., *
    Reduce:
      * -> [identifier_list]
    Goto:
      (nil)

  State 301:
    Kernel Items:
      parameter_type_list: parameter_list ',' ELLIPSIS., ')'
    Reduce:
      ')' -> [parameter_type_list]
    Goto:
      (nil)

  State 302:
    Kernel Items:
      parameter_list: parameter_list ',' parameter_declaration., *
    Reduce:
      * -> [parameter_list]
    Goto:
      (nil)

  State 303:
    Kernel Items:
      struct_declarator: ':' constant_expression., *
    Reduce:
      * -> [struct_declarator]
    Goto:
      (nil)

  State 304:
    Kernel Items:
      struct_declarator: declarator ':'.constant_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 142
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 140
      constant_expression -> State 329

  State 305:
    Kernel Items:
      struct_declarator_list: struct_declarator_list ','.struct_declarator
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 12
      '(' -> State 3
      ':' -> State 234
      '*' -> State 4
      struct_declarator -> State 330
      declarator -> State 235
      direct_declarator -> State 29
      pointer -> State 33

  State 306:
    Kernel Items:
      struct_declaration: specifier_qualifier_list struct_declarator_list ';'., *
    Reduce:
      * -> [struct_declaration]
    Goto:
      (nil)

  State 307:
    Kernel Items:
      struct_or_union_specifier: struct_or_union IDENTIFIER '{' struct_declaration_list '}'., *
    Reduce:
      * -> [struct_or_union_specifier]
    Goto:
      (nil)

  State 308:
    Kernel Items:
      initializer: '{' initializer_list ','.'}'
      initializer_list: initializer_list ','.initializer
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '{' -> State 153
      '}' -> State 331
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 154
      initializer -> State 332

  State 309:
    Kernel Items:
      initializer: '{' initializer_list '}'., *
    Reduce:
      * -> [initializer]
    Goto:
      (nil)

  State 310:
    Kernel Items:
      cast_expression: '(' type_name ')' cast_expression., *
    Reduce:
      * -> [cast_expression]
    Goto:
      (nil)

  State 311:
    Kernel Items:
      labeled_statement: CASE constant_expression ':' statement., *
    Reduce:
      * -> [labeled_statement]
    Goto:
      (nil)

  State 312:
    Kernel Items:
      iteration_statement: DO statement WHILE '('.expression ')' ';'
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 333

  State 313:
    Kernel Items:
      iteration_statement: FOR '(' expression_statement expression_statement.')' statement
      iteration_statement: FOR '(' expression_statement expression_statement.expression ')' statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      ')' -> State 334
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 335

  State 314:
    Kernel Items:
      selection_statement: IF '(' expression ')'.statement
      selection_statement: IF '(' expression ')'.statement ELSE statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 93
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      CASE -> State 85
      DEFAULT -> State 89
      IF -> State 94
      SWITCH -> State 99
      WHILE -> State 100
      DO -> State 90
      FOR -> State 91
      GOTO -> State 92
      CONTINUE -> State 87
      BREAK -> State 84
      RETURN -> State 96
      '(' -> State 77
      '{' -> State 49
      ';' -> State 81
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 110
      statement -> State 338
      labeled_statement -> State 115
      compound_statement -> State 105
      expression_statement -> State 111
      selection_statement -> State 122
      iteration_statement -> State 113
      jump_statement -> State 114

  State 315:
    Kernel Items:
      unary_expression: SIZEOF '(' type_name ')'., *
    Reduce:
      * -> [unary_expression]
    Goto:
      (nil)

  State 316:
    Kernel Items:
      selection_statement: SWITCH '(' expression ')'.statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 93
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      CASE -> State 85
      DEFAULT -> State 89
      IF -> State 94
      SWITCH -> State 99
      WHILE -> State 100
      DO -> State 90
      FOR -> State 91
      GOTO -> State 92
      CONTINUE -> State 87
      BREAK -> State 84
      RETURN -> State 96
      '(' -> State 77
      '{' -> State 49
      ';' -> State 81
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 110
      statement -> State 339
      labeled_statement -> State 115
      compound_statement -> State 105
      expression_statement -> State 111
      selection_statement -> State 122
      iteration_statement -> State 113
      jump_statement -> State 114

  State 317:
    Kernel Items:
      iteration_statement: WHILE '(' expression ')'.statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 93
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      CASE -> State 85
      DEFAULT -> State 89
      IF -> State 94
      SWITCH -> State 99
      WHILE -> State 100
      DO -> State 90
      FOR -> State 91
      GOTO -> State 92
      CONTINUE -> State 87
      BREAK -> State 84
      RETURN -> State 96
      '(' -> State 77
      '{' -> State 49
      ';' -> State 81
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 110
      statement -> State 340
      labeled_statement -> State 115
      compound_statement -> State 105
      expression_statement -> State 111
      selection_statement -> State 122
      iteration_statement -> State 113
      jump_statement -> State 114

  State 318:
    Kernel Items:
      conditional_expression: logical_or_expression '?' expression ':'.conditional_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 142
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 341

  State 319:
    Kernel Items:
      postfix_expression: postfix_expression '(' argument_expression_list ')'., *
    Reduce:
      * -> [postfix_expression]
    Goto:
      (nil)

  State 320:
    Kernel Items:
      argument_expression_list: argument_expression_list ','.assignment_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      '(' -> State 77
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 342

  State 321:
    Kernel Items:
      postfix_expression: postfix_expression '[' expression ']'., *
    Reduce:
      * -> [postfix_expression]
    Goto:
      (nil)

  State 322:
    Kernel Items:
      direct_abstract_declarator: '(' abstract_declarator ')'., *
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 323:
    Kernel Items:
      direct_abstract_declarator: '(' parameter_type_list ')'., *
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 324:
    Kernel Items:
      direct_abstract_declarator: '[' constant_expression ']'., *
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 325:
    Kernel Items:
      direct_abstract_declarator: direct_abstract_declarator '(' ')'., *
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 326:
    Kernel Items:
      direct_abstract_declarator: direct_abstract_declarator '(' parameter_type_list.')'
    Reduce:
      (nil)
    Goto:
      ')' -> State 343

  State 327:
    Kernel Items:
      direct_abstract_declarator: direct_abstract_declarator '[' ']'., *
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 328:
    Kernel Items:
      direct_abstract_declarator: direct_abstract_declarator '[' constant_expression.']'
    Reduce:
      (nil)
    Goto:
      ']' -> State 344

  State 329:
    Kernel Items:
      struct_declarator: declarator ':' constant_expression., *
    Reduce:
      * -> [struct_declarator]
    Goto:
      (nil)

  State 330:
    Kernel Items:
      struct_declarator_list: struct_declarator_list ',' struct_declarator., *
    Reduce:
      * -> [struct_declarator_list]
    Goto:
      (nil)

  State 331:
    Kernel Items:
      initializer: '{' initializer_list ',' '}'., *
    Reduce:
      * -> [initializer]
    Goto:
      (nil)

  State 332:
    Kernel Items:
      initializer_list: initializer_list ',' initializer., *
    Reduce:
      * -> [initializer_list]
    Goto:
      (nil)

  State 333:
    Kernel Items:
      expression: expression.',' assignment_expression
      iteration_statement: DO statement WHILE '(' expression.')' ';'
    Reduce:
      (nil)
    Goto:
      ')' -> State 345
      ',' -> State 187

  State 334:
    Kernel Items:
      iteration_statement: FOR '(' expression_statement expression_statement ')'.statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 93
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      CASE -> State 85
      DEFAULT -> State 89
      IF -> State 94
      SWITCH -> State 99
      WHILE -> State 100
      DO -> State 90
      FOR -> State 91
      GOTO -> State 92
      CONTINUE -> State 87
      BREAK -> State 84
      RETURN -> State 96
      '(' -> State 77
      '{' -> State 49
      ';' -> State 81
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 110
      statement -> State 346
      labeled_statement -> State 115
      compound_statement -> State 105
      expression_statement -> State 111
      selection_statement -> State 122
      iteration_statement -> State 113
      jump_statement -> State 114

  State 335:
    Kernel Items:
      expression: expression.',' assignment_expression
      iteration_statement: FOR '(' expression_statement expression_statement expression.')' statement
    Reduce:
      (nil)
    Goto:
      ')' -> State 347
      ',' -> State 187

  State 336:
    Kernel Items:
      selection_statement: IF '(' expression ')' statement., *
      selection_statement: IF '(' expression ')' statement.ELSE statement
    Reduce:
      * -> [selection_statement]
    Goto:
      ELSE -> State 348

  State 337:
    Kernel Items:
      selection_statement: IF '(' expression ')' statement., '!'
      selection_statement: IF '(' expression ')' statement., '&'
      selection_statement: IF '(' expression ')' statement., '('
      selection_statement: IF '(' expression ')' statement., '*'
      selection_statement: IF '(' expression ')' statement., '+'
      selection_statement: IF '(' expression ')' statement., '-'
      selection_statement: IF '(' expression ')' statement., ';'
      selection_statement: IF '(' expression ')' statement., '{'
      selection_statement: IF '(' expression ')' statement., '}'
      selection_statement: IF '(' expression ')' statement., '~'
      selection_statement: IF '(' expression ')' statement., BREAK
      selection_statement: IF '(' expression ')' statement., CASE
      selection_statement: IF '(' expression ')' statement., CONSTANT
      selection_statement: IF '(' expression ')' statement., CONTINUE
      selection_statement: IF '(' expression ')' statement., DEC_OP
      selection_statement: IF '(' expression ')' statement., DEFAULT
      selection_statement: IF '(' expression ')' statement., DO
      selection_statement: IF '(' expression ')' statement., ELSE
      selection_statement: IF '(' expression ')' statement., FOR
      selection_statement: IF '(' expression ')' statement., GOTO
      selection_statement: IF '(' expression ')' statement., IDENTIFIER
      selection_statement: IF '(' expression ')' statement., IF
      selection_statement: IF '(' expression ')' statement., INC_OP
      selection_statement: IF '(' expression ')' statement., RETURN
      selection_statement: IF '(' expression ')' statement., SIZEOF
      selection_statement: IF '(' expression ')' statement., STRING_LITERAL
      selection_statement: IF '(' expression ')' statement., SWITCH
      selection_statement: IF '(' expression ')' statement., WHILE
      selection_statement: IF '(' expression ')' statement.ELSE statement , '!'
      selection_statement: IF '(' expression ')' statement.ELSE statement , '&'
      selection_statement: IF '(' expression ')' statement.ELSE statement , '('
      selection_statement: IF '(' expression ')' statement.ELSE statement , '*'
      selection_statement: IF '(' expression ')' statement.ELSE statement , '+'
      selection_statement: IF '(' expression ')' statement.ELSE statement , '-'
      selection_statement: IF '(' expression ')' statement.ELSE statement , ';'
      selection_statement: IF '(' expression ')' statement.ELSE statement , '{'
      selection_statement: IF '(' expression ')' statement.ELSE statement , '}'
      selection_statement: IF '(' expression ')' statement.ELSE statement , '~'
      selection_statement: IF '(' expression ')' statement.ELSE statement , BREAK
      selection_statement: IF '(' expression ')' statement.ELSE statement , CASE
      selection_statement: IF '(' expression ')' statement.ELSE statement , CONSTANT
      selection_statement: IF '(' expression ')' statement.ELSE statement , CONTINUE
      selection_statement: IF '(' expression ')' statement.ELSE statement , DEC_OP
      selection_statement: IF '(' expression ')' statement.ELSE statement , DEFAULT
      selection_statement: IF '(' expression ')' statement.ELSE statement , DO
      selection_statement: IF '(' expression ')' statement.ELSE statement , ELSE
      selection_statement: IF '(' expression ')' statement.ELSE statement , FOR
      selection_statement: IF '(' expression ')' statement.ELSE statement , GOTO
      selection_statement: IF '(' expression ')' statement.ELSE statement , IDENTIFIER
      selection_statement: IF '(' expression ')' statement.ELSE statement , IF
      selection_statement: IF '(' expression ')' statement.ELSE statement , INC_OP
      selection_statement: IF '(' expression ')' statement.ELSE statement , RETURN
      selection_statement: IF '(' expression ')' statement.ELSE statement , SIZEOF
      selection_statement: IF '(' expression ')' statement.ELSE statement , STRING_LITERAL
      selection_statement: IF '(' expression ')' statement.ELSE statement , SWITCH
      selection_statement: IF '(' expression ')' statement.ELSE statement , WHILE
    Reduce:
      IDENTIFIER -> [selection_statement]
      CONSTANT -> [selection_statement]
      STRING_LITERAL -> [selection_statement]
      SIZEOF -> [selection_statement]
      INC_OP -> [selection_statement]
      DEC_OP -> [selection_statement]
      CASE -> [selection_statement]
      DEFAULT -> [selection_statement]
      IF -> [selection_statement]
      ELSE -> [selection_statement]
      SWITCH -> [selection_statement]
      WHILE -> [selection_statement]
      DO -> [selection_statement]
      FOR -> [selection_statement]
      GOTO -> [selection_statement]
      CONTINUE -> [selection_statement]
      BREAK -> [selection_statement]
      RETURN -> [selection_statement]
      '(' -> [selection_statement]
      '{' -> [selection_statement]
      '}' -> [selection_statement]
      ';' -> [selection_statement]
      '*' -> [selection_statement]
      '-' -> [selection_statement]
      '+' -> [selection_statement]
      '&' -> [selection_statement]
      '!' -> [selection_statement]
      '~' -> [selection_statement]
    Goto:
      ELSE -> State 348
    Shift/reduce conflict symbols:
      [ELSE]

  State 338:
    Kernel Items:
      selection_statement: IF '(' expression ')' statement., ELSE
      selection_statement: IF '(' expression ')' statement., WHILE
      selection_statement: IF '(' expression ')' statement.ELSE statement , ELSE
      selection_statement: IF '(' expression ')' statement.ELSE statement , WHILE
    Reduce:
      ELSE -> [selection_statement]
      WHILE -> [selection_statement]
    Goto:
      ELSE -> State 348
    Shift/reduce conflict symbols:
      [ELSE]

  State 339:
    Kernel Items:
      selection_statement: SWITCH '(' expression ')' statement., *
    Reduce:
      * -> [selection_statement]
    Goto:
      (nil)

  State 340:
    Kernel Items:
      iteration_statement: WHILE '(' expression ')' statement., *
    Reduce:
      * -> [iteration_statement]
    Goto:
      (nil)

  State 341:
    Kernel Items:
      conditional_expression: logical_or_expression '?' expression ':' conditional_expression., *
    Reduce:
      * -> [conditional_expression]
    Goto:
      (nil)

  State 342:
    Kernel Items:
      argument_expression_list: argument_expression_list ',' assignment_expression., *
    Reduce:
      * -> [argument_expression_list]
    Goto:
      (nil)

  State 343:
    Kernel Items:
      direct_abstract_declarator: direct_abstract_declarator '(' parameter_type_list ')'., *
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 344:
    Kernel Items:
      direct_abstract_declarator: direct_abstract_declarator '[' constant_expression ']'., *
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 345:
    Kernel Items:
      iteration_statement: DO statement WHILE '(' expression ')'.';'
    Reduce:
      (nil)
    Goto:
      ';' -> State 349

  State 346:
    Kernel Items:
      iteration_statement: FOR '(' expression_statement expression_statement ')' statement., *
    Reduce:
      * -> [iteration_statement]
    Goto:
      (nil)

  State 347:
    Kernel Items:
      iteration_statement: FOR '(' expression_statement expression_statement expression ')'.statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 93
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      CASE -> State 85
      DEFAULT -> State 89
      IF -> State 94
      SWITCH -> State 99
      WHILE -> State 100
      DO -> State 90
      FOR -> State 91
      GOTO -> State 92
      CONTINUE -> State 87
      BREAK -> State 84
      RETURN -> State 96
      '(' -> State 77
      '{' -> State 49
      ';' -> State 81
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 110
      statement -> State 350
      labeled_statement -> State 115
      compound_statement -> State 105
      expression_statement -> State 111
      selection_statement -> State 122
      iteration_statement -> State 113
      jump_statement -> State 114

  State 348:
    Kernel Items:
      selection_statement: IF '(' expression ')' statement ELSE.statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 93
      CONSTANT -> State 86
      STRING_LITERAL -> State 98
      SIZEOF -> State 97
      INC_OP -> State 95
      DEC_OP -> State 88
      CASE -> State 85
      DEFAULT -> State 89
      IF -> State 94
      SWITCH -> State 99
      WHILE -> State 100
      DO -> State 90
      FOR -> State 91
      GOTO -> State 92
      CONTINUE -> State 87
      BREAK -> State 84
      RETURN -> State 96
      '(' -> State 77
      '{' -> State 49
      ';' -> State 81
      '*' -> State 78
      '-' -> State 80
      '+' -> State 79
      '&' -> State 76
      '!' -> State 75
      '~' -> State 83
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 110
      statement -> State 351
      labeled_statement -> State 115
      compound_statement -> State 105
      expression_statement -> State 111
      selection_statement -> State 122
      iteration_statement -> State 113
      jump_statement -> State 114

  State 349:
    Kernel Items:
      iteration_statement: DO statement WHILE '(' expression ')' ';'., *
    Reduce:
      * -> [iteration_statement]
    Goto:
      (nil)

  State 350:
    Kernel Items:
      iteration_statement: FOR '(' expression_statement expression_statement expression ')' statement., *
    Reduce:
      * -> [iteration_statement]
    Goto:
      (nil)

  State 351:
    Kernel Items:
      selection_statement: IF '(' expression ')' statement ELSE statement., *
    Reduce:
      * -> [selection_statement]
    Goto:
      (nil)

Number of states: 351
Number of shift actions: 2989
Number of reduce actions: 246
Number of shift/reduce conflicts: 2
Number of reduce/reduce conflicts: 0
*/
