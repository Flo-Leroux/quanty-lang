package parser

import "github.com/Flo-Leroux/quanty-lang/internal/language/token"

// Wrap -
type Wrap struct {
	Open  token.Type
	Close token.Type
}

var BRACE_WRAPPER = Wrap{
	Open:  token.LBRACE,
	Close: token.RBRACE,
}

type Wrapper []Wrap

func (w Wrapper) isOpen(t token.Token) bool {
	for _, wrap := range w {
		if wrap.Open == t.Type {
			return true
		}
	}

	return false
}

func (w Wrapper) isClose(t token.Token) bool {
	for _, wrap := range w {
		if wrap.Close == t.Type {
			return true
		}
	}

	return false
}

func (w Wrapper) findCloseWithOpen(t token.Type) token.Type {
	for _, wrap := range w {
		if wrap.Open == t {
			return wrap.Close
		}
	}

	return token.ILLEGAL
}

func (w Wrapper) findOpenWithClose(t token.Type) token.Type {
	for _, wrap := range w {
		if wrap.Close == t {
			return wrap.Open
		}
	}

	return token.ILLEGAL
}

var wrappers = &Wrapper{
	BRACE_WRAPPER,
}
