package ast

import (
	"strings"

	"github.com/Flo-Leroux/quanty-lang/internal/language/token"
)

// Field -
type Field struct {
	Name       token.Token // token.IDENT
	Selections SelectionList
}

func (f *Field) statementNode() {}

func (f *Field) selectionNode() {}

// TokenLiteral -
func (ls *Field) TokenLiteral() string {
	return ls.Name.Literal
}

// String -
func (f *Field) String() string {
	var sb strings.Builder
	sb.WriteString(f.Name.Literal)

	if len(f.Selections) != 0 {
		sb.WriteString(" { ")

		selectionsStr := make([]string, 0)
		for _, sel := range f.Selections {
			selectionsStr = append(
				selectionsStr,
				strings.Trim(
					sel.String(),
					" ",
				),
			)
		}
		sb.WriteString(
			strings.Join(selectionsStr, " "),
		)

		sb.WriteString(" }")
	}

	return sb.String()
}

// StringValue -
type StringValue struct {
	Token token.Token // token.STRING
	Value string
}

func (s *StringValue) statementNode() {}

func (s *StringValue) selectionNode() {}

// TokenLiteral -
func (s *StringValue) TokenLiteral() string {
	return s.Token.Literal
}

// String -
func (s *StringValue) String() string {
	return s.Value
}
