package parser_test

import (
	"testing"

	"github.com/Flo-Leroux/quanty-lang/internal/language/ast"
	"github.com/Flo-Leroux/quanty-lang/internal/language/parser"
	"github.com/Flo-Leroux/quanty-lang/internal/language/token"
	. "github.com/smartystreets/goconvey/convey"
)

func BenchmarkParser(b *testing.B) {
	component := `
component Main {
	div {
		span
		p 
		div {
			span
			p 
			div {
				span
				p
			}
		}
		div {
			span
			p
			div {
				span
				p
				div {
					span
					p
				}
				div {
					span
					p
				}
				div {
					span
					p
				}
			}
		}
	}
	img
}
`

	for n := 0; n < b.N; n++ {
		parser.NewParser(component).Parse()
	}
}

// TestParser -
func TestParser(t *testing.T) {
	Convey("Subject: Parser", t, func() {
		Convey("When component", func() {
			Convey("With empty selection", func() {
				component := `
component Main {
}
`

				schemaParser := parser.NewParser(component)

				schema := schemaParser.Parse()

				result := &ast.Schema{
					Statements: []ast.Statement{
						&ast.ComponentStatement{
							Token: token.Token{
								Type:    token.COMPONENT,
								Literal: "component",
							},
							Name: token.Token{
								Type:    token.IDENT,
								Literal: "Main",
							},
							Fields: []*ast.Field{},
						},
					},
				}

				So(schema, ShouldResemble, result)
				So(schemaParser.Errors(), ShouldBeEmpty)
			})

			Convey("With 1 level of selection", func() {

				component := `
component Main {
	div
	img
}
`

				schemaParser := parser.NewParser(component)

				schema := schemaParser.Parse()

				result := &ast.Schema{
					Statements: []ast.Statement{
						&ast.ComponentStatement{
							Token: token.Token{
								Type:    token.COMPONENT,
								Literal: "component",
							},
							Name: token.Token{
								Type:    token.IDENT,
								Literal: "Main",
							},
							Fields: []*ast.Field{
								{
									Name: token.Token{
										Type:    token.IDENT,
										Literal: "div",
									},
									Selections: []*ast.Field{},
								},
								{
									Name: token.Token{
										Type:    token.IDENT,
										Literal: "img",
									},
									Selections: []*ast.Field{},
								},
							},
						},
					},
				}

				So(schema, ShouldResemble, result)
				So(schemaParser.Errors(), ShouldBeEmpty)
			})

			Convey("With 2 levels of selection", func() {

				Convey("Should be parsed", func() {
					component := `
component Main {
	div {
		span
		p
	}
	img
}
`

					schemaParser := parser.NewParser(component)

					schema := schemaParser.Parse()

					result := &ast.Schema{
						Statements: []ast.Statement{
							&ast.ComponentStatement{
								Token: token.Token{
									Type:    token.COMPONENT,
									Literal: "component",
								},
								Name: token.Token{
									Type:    token.IDENT,
									Literal: "Main",
								},
								Fields: []*ast.Field{
									{
										Name: token.Token{
											Type:    token.IDENT,
											Literal: "div",
										},
										Selections: []*ast.Field{
											{
												Name: token.Token{
													Type:    token.IDENT,
													Literal: "span",
												},
												Selections: []*ast.Field{},
											},
											{
												Name: token.Token{
													Type:    token.IDENT,
													Literal: "p",
												},
												Selections: []*ast.Field{},
											},
										},
									},
									{
										Name: token.Token{
											Type:    token.IDENT,
											Literal: "img",
										},
										Selections: []*ast.Field{},
									},
								},
							},
						},
					}

					So(schema, ShouldResemble, result)
					So(schemaParser.Errors(), ShouldBeEmpty)
				})

				Convey("When missing braces at level 1", func() {
					Convey("Case left", func() {
						component := `
component Main {
	div
		span
		p
	}
	img
}
`

						p := parser.NewParser(component)

						p.Parse()

						err := p.Errors()

						So(err, ShouldNotBeEmpty)
					})

					Convey("Case right", func() {
						component := `
component Main {
	div {
		span
		p

	img
}
`

						p := parser.NewParser(component)

						p.Parse()

						err := p.Errors()

						So(err, ShouldNotBeEmpty)
					})
				})
			})

			Convey("When missing braces", func() {
				Convey("Case left", func() {
					component := "component Main }"

					p := parser.NewParser(component)

					p.Parse()

					err := p.Errors()

					So(err, ShouldNotBeEmpty)
				})

				Convey("Case right", func() {
					component := "component Main {"

					p := parser.NewParser(component)

					p.Parse()

					err := p.Errors()

					So(err, ShouldNotBeEmpty)
				})

				Convey("Case left and right", func() {
					component := "component Main"

					p := parser.NewParser(component)

					p.Parse()

					err := p.Errors()

					So(err, ShouldNotBeEmpty)
				})
			})
		})
	})
}
