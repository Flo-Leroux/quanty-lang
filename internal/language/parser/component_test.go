package parser_test

import (
	"testing"

	"github.com/Flo-Leroux/quanty-lang/internal/language/ast"
	"github.com/Flo-Leroux/quanty-lang/internal/language/parser"
	"github.com/Flo-Leroux/quanty-lang/internal/language/token"
	. "github.com/smartystreets/goconvey/convey"
)

// BenchmarkParserComponent -
func BenchmarkParserComponent(b *testing.B) {
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
	"Hello World!"
}
`

	for n := 0; n < b.N; n++ {
		parser.NewParser(component).Parse()
	}
}

// TestParserComponent -
func TestParserComponent(t *testing.T) {
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
							Selections: ast.SelectionList{},
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
					Statements: ast.StatementList{
						&ast.ComponentStatement{
							Token: token.Token{
								Type:    token.COMPONENT,
								Literal: "component",
							},
							Name: token.Token{
								Type:    token.IDENT,
								Literal: "Main",
							},
							Selections: ast.SelectionList{
								&ast.Field{
									Name: token.Token{
										Type:    token.IDENT,
										Literal: "div",
									},
									Selections: ast.SelectionList{},
								},
								&ast.Field{
									Name: token.Token{
										Type:    token.IDENT,
										Literal: "img",
									},
									Selections: ast.SelectionList{},
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
								Selections: ast.SelectionList{
									&ast.Field{
										Name: token.Token{
											Type:    token.IDENT,
											Literal: "div",
										},
										Selections: ast.SelectionList{
											&ast.Field{
												Name: token.Token{
													Type:    token.IDENT,
													Literal: "span",
												},
												Selections: ast.SelectionList{},
											},
											&ast.Field{
												Name: token.Token{
													Type:    token.IDENT,
													Literal: "p",
												},
												Selections: ast.SelectionList{},
											},
										},
									},
									&ast.Field{
										Name: token.Token{
											Type:    token.IDENT,
											Literal: "img",
										},
										Selections: ast.SelectionList{},
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

			Convey("With string content at level 1", func() {
				component := `
				component Main {
					"Hello World!"
				}`

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
							Selections: ast.SelectionList{
								&ast.StringValue{
									Token: token.Token{
										Type:    token.STRING,
										Literal: "Hello World!",
									},
									Value: "Hello World!",
								},
							},
						},
					},
				}

				So(schema, ShouldResemble, result)
				So(schemaParser.Errors(), ShouldBeEmpty)
			})

			Convey("With string content at level 2", func() {
				component := `
				component Main {
					p {
						"Hello World - Level 2-1!"
					}
					"Hello World - Level 1!"
					p {
						"Hello World - Level 2-2!"
					}
				}`

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
							Selections: ast.SelectionList{
								&ast.Field{
									Name: token.Token{
										Type:    token.IDENT,
										Literal: "p",
									},
									Selections: []ast.Selection{
										&ast.StringValue{
											Token: token.Token{
												Type:    token.STRING,
												Literal: "Hello World - Level 2-1!",
											},
											Value: "Hello World - Level 2-1!",
										},
									},
								},
								&ast.StringValue{
									Token: token.Token{
										Type:    token.STRING,
										Literal: "Hello World - Level 1!",
									},
									Value: "Hello World - Level 1!",
								},
								&ast.Field{
									Name: token.Token{
										Type:    token.IDENT,
										Literal: "p",
									},
									Selections: []ast.Selection{
										&ast.StringValue{
											Token: token.Token{
												Type:    token.STRING,
												Literal: "Hello World - Level 2-2!",
											},
											Value: "Hello World - Level 2-2!",
										},
									},
								},
							},
						},
					},
				}

				So(schema, ShouldResemble, result)
				So(schemaParser.Errors(), ShouldBeEmpty)
			})

		})
	})
}
