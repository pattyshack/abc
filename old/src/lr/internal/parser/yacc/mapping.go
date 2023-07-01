package yacc

import (
	"fmt"

	"github.com/pattyshack/abc/src/lr/internal/parser"
)

func LRSymbolTypeToYaccTokenNum(tt parser.LRSymbolId) int {
	switch tt {
	case ':', '<', '>', '|', ';':
		return int(tt)
	case parser.LRTokenToken:
		return TOKEN
	case parser.LRTypeToken:
		return TYPE
	case parser.LRStartToken:
		return START
	case parser.Arrow:
		return ARROW
	case parser.LRRuleDefToken:
		return RULE_DEF
	case parser.LRLabelToken:
		return LABEL
	case parser.LRIdentifierToken:
		return IDENTIFIER
	case parser.LRCharacterToken:
		return CHARACTER
	case parser.LRSectionMarkerToken:
		return SECTION_MARKER
	case parser.LRSectionContentToken:
		return SECTION_CONTENT
	}

	panic(fmt.Sprintf("Unexpected LRSymbolId: %v", tt))
}
