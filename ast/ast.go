package ast

import (
	"bytes"
	"quanty/token"
	"strings"
)

type Definition interface {
	TokenLiteral() string
	String() string
}

type Document struct {
	// Kind        token.Token // Token.DOCUMENT
	Definitions []Definition
}

func (d *Document) TokenLiteral() string {
	if len(d.Definitions) > 0 {
		return d.Definitions[0].TokenLiteral()
	}
	return ""
}

func (d *Document) String() string {
	var out bytes.Buffer

	for _, def := range d.Definitions {
		out.WriteString(def.String())
	}

	return standardizeSpace(out.String())
}

type ComponentDefinition struct {
	Kind         token.Token // Token.COMPONENT
	Name         *Identifier
	SelectionSet *SelectionSet
}

func (c *ComponentDefinition) TokenLiteral() string {
	return c.Kind.Literal
}

func (c *ComponentDefinition) String() string {
	var out bytes.Buffer

	out.WriteString(c.TokenLiteral() + " ")
	out.WriteString(c.Name.String() + " ")

	out.WriteString(c.SelectionSet.String())

	return out.String()
}

type Identifier struct {
	Kind  token.Token // Token.IDENT
	Value string
}

func (i *Identifier) TokenLiteral() string {
	return i.Kind.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

type SelectionSet struct {
	Kind      token.Token // Token.SELECTION_SET
	Selection []Definition
}

func (s *SelectionSet) TokenLiteral() string {
	return s.Kind.Literal
}

func (s *SelectionSet) String() string {
	var out bytes.Buffer
	if len(s.Selection) == 0 {
		return out.String()
	}

	addSpace(&out)

	out.WriteString("{")
	addSpace(&out)

	for _, t := range s.Selection {
		addSpace(&out)
		out.WriteString(t.String())
		addSpace(&out)
	}

	addSpace(&out)
	out.WriteString("}")

	addSpace(&out)
	return out.String()
}

type TagDefinition struct {
	Kind         token.Token // Token.TAG
	Name         *Identifier
	SelectionSet *SelectionSet
}

func (t *TagDefinition) TokenLiteral() string {
	return t.Kind.Literal
}

func (t *TagDefinition) String() string {
	var out bytes.Buffer

	out.WriteString(t.Name.String())
	addSpace(&out)

	if t.SelectionSet != nil {
		out.WriteString(t.SelectionSet.String())
	}

	addSpace(&out)
	return out.String()
}

type StringLiteral struct {
	Kind  token.Token // token.STRING
	Value string
}

func (sl *StringLiteral) TokenLiteral() string {
	return sl.Kind.Literal
}

func (sl *StringLiteral) String() string {
	var out bytes.Buffer

	out.WriteString("\"")
	out.WriteString(sl.Value)
	out.WriteString("\"")

	return out.String()
}

func standardizeSpace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func addSpace(b *bytes.Buffer) {
	b.WriteString(" ")
}
