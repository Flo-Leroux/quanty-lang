package html_test

import (
	"bytes"
	"testing"

	"github.com/Flo-Leroux/quanty-lang/internal/language/ast"
	"github.com/Flo-Leroux/quanty-lang/internal/transpilers/html"
	. "github.com/smartystreets/goconvey/convey"
)

func TestFmt(t *testing.T) {
	Convey("Subject: Html", t, func() {

		var h *html.HTML

		type test struct {
			name string
			args ast.Visitable
			want string
		}

		Convey("New Html", func() {
			h = html.New()

			So(h, ShouldResemble, &html.HTML{})
		})

		Convey("Use io.Write", func() {
			h = html.New()

			tt := test{
				name: "With 1 Component",
				args: ast.NewSchema().
					WithStatements(
						ast.NewComponentStatement("Main"),
					),
				want: `<Main></Main>`,
			}

			var b bytes.Buffer
			_, err := h.Write(&b, tt.args)

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
					want: `<Main></Main>`,
				},
				{
					name: "With 2 Components",
					args: ast.NewSchema().
						WithStatements(
							ast.NewComponentStatement("Main"),
							ast.NewComponentStatement("Nav"),
						),
					want: `<Main></Main>
<Nav></Nav>`,
				},
			}

			for _, tt := range tests {
				Convey(tt.name, func() {
					h = html.New()
					rendered := h.String(tt.args)

					So(rendered, ShouldEqual, tt.want)
				})
			}
		})

		Convey("Component", func() {
			tests := []test{
				{
					name: "When Empty",
					args: ast.NewComponentStatement("Main"),
					want: `<Main></Main>`,
				},
				{
					name: "With 1 Field",
					args: ast.NewComponentStatement("Main").
						WithSelections(
							ast.NewField("div"),
						),
					want: `<Main>
	<div></div>
</Main>`,
				},
				{
					name: "With 1 StringValue",
					args: ast.NewComponentStatement("Main").
						WithSelections(
							ast.NewStringValue("Hello World!"),
						),
					want: `<Main>
	Hello World!
</Main>`,
				},
				{
					name: "With 1 StringValue AND 1 Field",
					args: ast.NewComponentStatement("Main").
						WithSelections(
							ast.NewField("div"),
							ast.NewStringValue("Hello World!"),
						),
					want: `<Main>
	<div></div>
	Hello World!
</Main>`,
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
					want: `<Main>
	<div>
		<p>
			<span>
				Hello
			</span>
			<span>
				World!
			</span>
		</p>
	</div>
	Hello World!
</Main>`,
				},
			}

			for _, tt := range tests {
				Convey(tt.name, func() {
					h = html.New()
					rendered := h.String(tt.args)

					So(rendered, ShouldEqual, tt.want)
				})
			}
		})

		Convey("Field", func() {
			tests := []test{
				{
					name: "When Empty",
					args: ast.NewField("div"),
					want: "<div></div>",
				},
				{
					name: "With 1 selection level",
					args: ast.NewField("div").
						WithSelections(
							ast.NewField("p"),
						),
					want: `<div>
	<p></p>
</div>`,
				},
			}

			for _, tt := range tests {
				Convey(tt.name, func() {
					h = html.New()
					rendered := h.String(tt.args)

					So(rendered, ShouldEqual, tt.want)
				})
			}
		})

		Convey("StringValue", func() {
			tests := []test{
				{
					name: "With string value",
					args: ast.NewStringValue("Hello World!"),
					want: `Hello World!`,
				},
			}

			for _, tt := range tests {
				Convey(tt.name, func() {
					h = html.New()
					rendered := h.String(tt.args)

					So(rendered, ShouldEqual, tt.want)
				})
			}
		})

	})

}
