package token

// Type -
type Type string

// String -
func (t Type) String() string {
	return string(t)
}

// Token -
type Token struct {
	Type    Type
	Literal string
}

// New -
func New(typ Type, ch byte) Token {
	return Token{Type: typ, Literal: string(ch)}
}

// keywords -
var keywords = map[string]Type{
	"component": COMPONENT,
}

const (

	//
	// Special Types
	//

	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	STRING = "STRING"

	//
	// Identifiers & Literals
	//

	IDENT = "IDENT"

	//
	// Delimiters
	//

	COMMA   = "COMMA"
	COMMENT = "COMMENT"

	LBRACE = "LBRACE"
	RBRACE = "RBRACE"

	LPAREN = "LPAREN"
	RPAREN = "RPAREN"

	DOT = "DOT"

	COLON = "COLON"

	NEWLINE = "NEWLINE"

	//
	// Keywords
	//

	COMPONENT = "COMPONENT"

	//
	// Prefix
	//

	//
	// Logical
	//
)

// LookupKeywords -
func LookupKeywords(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
