package parser

import (
	"quanty/lexer"
	"strings"
	"testing"
)

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestComponents(t *testing.T) {
	input := standardizeSpace(
		`component Main {
			p {
				bold
			}
			User
		}

		component User {
			src
			p {
				span {
					"Hello World!"
				}
			}
		}
		`,
	)

	l := lexer.New(input)
	p := New(l)

	output := p.ParseDocument()
	checkParserErrors(t, p)

	if input != output.String() {
		t.Fatalf("ParseDocument() failed. expected='%s', got='%s'", input, output.String())
	}
}

func standardizeSpace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
