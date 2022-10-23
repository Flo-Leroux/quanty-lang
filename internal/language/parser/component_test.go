package parser_test

import (
	"testing"

	"github.com/Flo-Leroux/quanty-lang/internal/language/ast"
	"github.com/Flo-Leroux/quanty-lang/internal/language/parser"
	. "github.com/smartystreets/goconvey/convey"
)

// // BenchmarkParserComponent -
// func BenchmarkParserComponent(b *testing.B) {
// 	component := `
// component Main {
// 	div {
// 		span
// 		p
// 		div {
// 			span
// 			p
// 			div {
// 				span
// 				p
// 			}
// 		}
// 		div {
// 			span
// 			p
// 			div {
// 				span
// 				p
// 				div {
// 					span
// 					p
// 				}
// 				div {
// 					span
// 					p
// 				}
// 				div {
// 					span
// 					p
// 				}
// 			}
// 		}
// 	}
// 	img
// 	"Hello World!"
// }
// `

// 	// for n := 0; n < b.N; n++ {
// 	// 	input := antlr.NewInputStream(component)
// 	// 	lexer := antlr4.NewquantyLexer(input)
// 	// 	stream := antlr.NewCommonTokenStream(lexer, 0)
// 	// 	p := antlr4.NewquantyParser(stream)
// 	// 	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
// 	// 	p.BuildParseTrees = true
// 	// 	p.Schema()

// 	// 	// antlr.ParseTreeWalkerDefault.Walk(antlr4.NewQuantyListener(), tree)
// 	// }
// }

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

				result := ast.NewSchema().
					WithStatements(
						ast.NewComponentStatement("Main"),
					)

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

				result := ast.NewSchema().
					WithStatements(
						ast.NewComponentStatement("Main").
							WithSelections(
								ast.NewField("div"),
								ast.NewField("img"),
							),
					)

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

					result := ast.NewSchema().
						WithStatements(
							ast.NewComponentStatement("Main").
								WithSelections(
									ast.NewField("div").
										WithSelections(
											ast.NewField("span"),
											ast.NewField("p"),
										),
									ast.NewField("img"),
								),
						)

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

				result := ast.NewSchema().
					WithStatements(
						ast.NewComponentStatement("Main").
							WithSelections(
								ast.NewStringValue("Hello World!"),
							),
					)

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

				result := ast.NewSchema().
					WithStatements(
						ast.NewComponentStatement("Main").
							WithSelections(
								ast.NewField("p").
									WithSelections(
										ast.NewStringValue("Hello World - Level 2-1!"),
									),
								ast.NewStringValue("Hello World - Level 1!"),
								ast.NewField("p").
									WithSelections(
										ast.NewStringValue("Hello World - Level 2-2!"),
									),
							),
					)

				So(schema, ShouldResemble, result)
				So(schemaParser.Errors(), ShouldBeEmpty)
			})

		})
	})
}
