package ast_test

import (
	"testing"

	"github.com/Flo-Leroux/quanty-lang/internal/language/ast"
	"github.com/Flo-Leroux/quanty-lang/internal/language/token"
	. "github.com/smartystreets/goconvey/convey"
)

// TestField -
func TestField(t *testing.T) {
	type test struct {
		name   string
		input  *ast.Field
		result string
	}

	tests := []test{
		{
			name: "When selection is empty",
			input: &ast.Field{
				Name:       token.Token{Type: token.IDENT, Literal: "div"},
				Selections: []*ast.Field{},
			},
			result: "div",
		},
		{
			name: "When selection level 1",
			input: &ast.Field{
				Name: token.Token{
					Type:    token.IDENT,
					Literal: "div",
				},
				Selections: []*ast.Field{
					{
						Name:       token.Token{Type: token.IDENT, Literal: "p"},
						Selections: []*ast.Field{},
					},
					{
						Name:       token.Token{Type: token.IDENT, Literal: "img"},
						Selections: []*ast.Field{},
					},
				},
			},
			result: "div { p img }",
		},
		{
			name: "When selection level 2",
			input: &ast.Field{
				Name: token.Token{
					Type:    token.IDENT,
					Literal: "div",
				},
				Selections: []*ast.Field{
					{
						Name: token.Token{
							Type:    token.IDENT,
							Literal: "p",
						},
						Selections: []*ast.Field{
							{
								Name:       token.Token{Type: token.IDENT, Literal: "span"},
								Selections: []*ast.Field{},
							},
						},
					},
					{
						Name:       token.Token{Type: token.IDENT, Literal: "img"},
						Selections: []*ast.Field{},
					},
				},
			},
			result: "div { p { span } img }",
		},
	}

	Convey("Subject: Stringify", t, func() {

		for _, test := range tests {
			Convey(test.name, func() {
				So(test.input.String(), ShouldEqual, test.result)
			})
		}
	})
}
