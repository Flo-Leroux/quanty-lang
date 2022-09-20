package ast_test

import (
	"testing"

	"github.com/Flo-Leroux/quanty-lang/internal/language/ast"
	"github.com/Flo-Leroux/quanty-lang/internal/language/token"
	. "github.com/smartystreets/goconvey/convey"
)

// TestComponent -
func TestComponent(t *testing.T) {
	Convey("Subject: Stringify", t, func() {
		component := &ast.ComponentStatement{
			Token: token.Token{
				Type:    token.COMPONENT,
				Literal: "component",
			},
			Name: token.Token{
				Type:    token.IDENT,
				Literal: "Main",
			},
		}

		Convey("When selection is empty", func() {
			strResult := `component Main {
}`

			So(component.String(), ShouldEqual, strResult)
		})
	})
}
