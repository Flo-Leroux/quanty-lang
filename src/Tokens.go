package main

type Token int

const (
	// Special tokens

	ILLEGAL Token = iota
	EOF
	WS

	// Literals

	IDENT // fields, table_name

	// Misc characters

	OPEN_BRACKET  // {
	CLOSE_BRACKET // }
	ASTERISK      // *
	COMMA         // ,
	EQUAL         // =
	END           // ;

	// Keywords

	COMPONENT // component
	PROP      // property
)

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}

var eof = rune(0)
