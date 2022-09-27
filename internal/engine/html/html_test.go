package html_test

import (
	"bytes"
	"testing"

	"github.com/Flo-Leroux/quanty-lang/internal/engine/html"
	"github.com/Flo-Leroux/quanty-lang/internal/language/ast"
	. "github.com/smartystreets/goconvey/convey"
)

func TestFmt(t *testing.T) {
	Convey("Subject: Html", t, func() {

		var engine *html.Engine

		type test struct {
			name string
			args ast.Visitable
			want string
		}

		Convey("New Html", func() {
			engine = html.New()

			So(engine, ShouldResemble, &html.Engine{})
		})

		Convey("Use io.Write", func() {
			engine = html.New()

			tt := test{
				name: "Witengine 1 Component",
				args: ast.NewSchema().
					WithStatements(
						ast.NewComponentStatement("Main"),
					),
				want: `<Main></Main>`,
			}

			var b bytes.Buffer
			_, err := engine.Write(&b, tt.args)

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
					name: "Witengine 1 Component",
					args: ast.NewSchema().
						WithStatements(
							ast.NewComponentStatement("Main"),
						),
					want: `<Main></Main>`,
				},
				{
					name: "Witengine 2 Components",
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
					engine = html.New()
					rendered := engine.String(tt.args)

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
					name: "Witengine 1 Field",
					args: ast.NewComponentStatement("Main").
						WithSelections(
							ast.NewField("div"),
						),
					want: `<Main>
	<div></div>
</Main>`,
				},
				{
					name: "Witengine 1 StringValue",
					args: ast.NewComponentStatement("Main").
						WithSelections(
							ast.NewStringValue("Hello World!"),
						),
					want: `<Main>
	Hello World!
</Main>`,
				},
				{
					name: "Witengine 1 StringValue AND 1 Field",
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
					name: "Witengine multiple selections",
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
					engine = html.New()
					rendered := engine.String(tt.args)

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
					name: "Witengine 1 selection level",
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
					engine = html.New()
					rendered := engine.String(tt.args)

					So(rendered, ShouldEqual, tt.want)
				})
			}
		})

		Convey("StringValue", func() {
			tests := []test{
				{
					name: "Witengine string value",
					args: ast.NewStringValue("Hello World!"),
					want: `Hello World!`,
				},
			}

			for _, tt := range tests {
				Convey(tt.name, func() {
					engine = html.New()
					rendered := engine.String(tt.args)

					So(rendered, ShouldEqual, tt.want)
				})
			}
		})

	})

}
