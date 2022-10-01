package parser

import "github.com/Flo-Leroux/quanty-lang/pkg/dsl/token"

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

var Wrappers = &Wrapper{
	BraceWrapper,
}

// wrapWith -
func (p *Parser) WrapWith(wrapper Wrap, fn func()) {
	if !p.ExpectAndNext(wrapper.Open) {
		return
	}

	for !p.CurrentTokenIs(wrapper.Close) {
		fn()
		p.Next()

		if p.CurrentTokenIs(token.EOF) {
			p.PeekError(wrapper.Close)
			return
		}
	}
}
