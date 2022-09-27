package format_test

import (
	"bytes"
	"testing"

	"github.com/Flo-Leroux/quanty-lang/internal/format"
	"github.com/Flo-Leroux/quanty-lang/internal/language/ast"
	. "github.com/smartystreets/goconvey/convey"
)

func TestFmt(t *testing.T) {
	Convey("Subject: Fmt", t, func() {

		var fmt *format.Fmt

		type test struct {
			name string
			args ast.Visitable
			want string
		}

		Convey("New Fmt", func() {
			fmt = format.New()

			So(fmt, ShouldResemble, &format.Fmt{})
		})

		Convey("Use io.Write", func() {
			fmt = format.New()

			tt := test{
				name: "With 1 Component",
				args: ast.NewSchema().
					WithStatements(
						ast.NewComponentStatement("Main"),
					),
				want: `component Main {}
`,
			}

			var b bytes.Buffer
			_, err := fmt.Write(&b, tt.args)

			So(err, ShouldBeNil)
			So(b.String(), ShouldEqual, tt.want)
		})

		Convey("Schema", func() {
			tests := []test{
				{
					name: "When Empty",
					args: ast.NewSchema(),
					want: "",
				},
				{
					name: "With 1 Component",
					args: ast.NewSchema().
						WithStatements(
							ast.NewComponentStatement("Main"),
						),
					want: `component Main {}
`,
				},
				{
					name: "With 2 Components",
					args: ast.NewSchema().
						WithStatements(
							ast.NewComponentStatement("Main"),
							ast.NewComponentStatement("Nav"),
						),
					want: `component Main {}

component Nav {}
`,
				},
			}

			for _, tt := range tests {
				Convey(tt.name, func() {
					fmt = format.New()
					rendered := fmt.String(tt.args)

					So(rendered, ShouldEqual, tt.want)
				})
			}
		})

		Convey("Component", func() {
			tests := []test{
				{
					name: "When Empty",
					args: ast.NewComponentStatement("Main"),
					want: `component Main {}`,
				},
				{
					name: "With 1 Field",
					args: ast.NewComponentStatement("Main").
						WithSelections(
							ast.NewField("div"),
						),
					want: `component Main {
	div
}`,
				},
				{
					name: "With 1 StringValue",
					args: ast.NewComponentStatement("Main").
						WithSelections(
							ast.NewStringValue("Hello World!"),
						),
					want: `component Main {
	"Hello World!"
}`,
				},
				{
					name: "With 1 StringValue AND 1 Field",
					args: ast.NewComponentStatement("Main").
						WithSelections(
							ast.NewField("div"),
							ast.NewStringValue("Hello World!"),
						),
					want: `component Main {
	div
	"Hello World!"
}`,
				},
				{
					name: "With multiple selections",
					args: ast.NewComponentStatement("Main").
						WithSelections(
							ast.NewField("div").
								WithSelections(
									ast.NewField("p").
										WithSelections(
											ast.NewField("span").
												WithSelections(
													ast.NewStringValue("Hello"),
												),
											ast.NewField("span").
												WithSelections(
													ast.NewStringValue("World!"),
												),
										),
								),
							ast.NewStringValue("Hello World!"),
						),
					want: `component Main {
	div {
		p {
			span {
				"Hello"
			}
			span {
				"World!"
			}
		}
	}
	"Hello World!"
}`,
				},
			}

			for _, tt := range tests {
				Convey(tt.name, func() {
					fmt = format.New()
					rendered := fmt.String(tt.args)

					So(rendered, ShouldEqual, tt.want)
				})
			}
		})

		Convey("Field", func() {
			tests := []test{
				{
					name: "When Empty",
					args: ast.NewField("div"),
					want: "div",
				},
				{
					name: "With 1 selection level",
					args: ast.NewField("div").
						WithSelections(
							ast.NewField("p"),
						),
					want: `div {
	p
}`,
				},
			}

			for _, tt := range tests {
				Convey(tt.name, func() {
					fmt = format.New()
					rendered := fmt.String(tt.args)

					So(rendered, ShouldEqual, tt.want)
				})
			}
		})

		Convey("StringValue", func() {
			tests := []test{
				{
					name: "With string value",
					args: ast.NewStringValue("Hello World!"),
					want: `"Hello World!"`,
				},
			}

			for _, tt := range tests {
				Convey(tt.name, func() {
					fmt = format.New()
					rendered := fmt.String(tt.args)

					So(rendered, ShouldEqual, tt.want)
				})
			}
		})

	})

}
