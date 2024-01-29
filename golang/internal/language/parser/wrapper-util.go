package parser

import "github.com/Flo-Leroux/quanty-lang/internal/language/token"

// Wrap -
type Wrap struct {
	Open  token.Type
	Close token.Type
}

var BraceWrapper = Wrap{
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

// func (w Wrapper) findCloseWithOpen(t token.Type) token.Type {
// 	for _, wrap := range w {
// 		if wrap.Open == t {
// 			return wrap.Close
// 		}
// 	}

// 	return token.ILLEGAL
// }

func (w Wrapper) findOpenWithClose(t token.Type) token.Type {
	for _, wrap := range w {
		if wrap.Close == t {
			return wrap.Open
		}
	}

	return token.ILLEGAL
}

var wrappers = &Wrapper{
	BraceWrapper,
}

// wrapWith -
func (p *Parser) wrapWith(wrapper Wrap, fn func()) {
	if !p.expectAndNext(wrapper.Open) {
		return
	}

	for !p.currentTokenIs(wrapper.Close) {
		fn()
		p.next()

		if p.currentTokenIs(token.EOF) {
			p.peekError(wrapper.Close)
			return
		}
	}
}
