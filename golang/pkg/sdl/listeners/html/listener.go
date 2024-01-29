package html

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/Flo-Leroux/quanty-lang/pkg/sdl/parser"
)

type HtmlQuantyListener struct {
	filename  string
	templates *template.Template

	components   map[string]*strings.Builder
	crtComponent string

	*parser.BaseQuantyParserListener
}

func New(filename string) *HtmlQuantyListener {
	return &HtmlQuantyListener{
		filename:   filename,
		templates:  template.New(filename),
		components: make(map[string]*strings.Builder),
	}
}

func (x *HtmlQuantyListener) EnterComponent(ctx *parser.ComponentContext) {
	name := ctx.GetName().GetText()
	x.crtComponent = name

	x.current().WriteString(
		fmt.Sprintf(`{{ define "%s" }}`, name),
	)
}

func (x *HtmlQuantyListener) EnterSelection(ctx *parser.SelectionContext) {
	isString := ctx.STRING()
	if isString != nil {
		x.current().WriteString(isString.GetText())
	}
}

func (x *HtmlQuantyListener) EnterField(ctx *parser.FieldContext) {
	name := ctx.GetName().GetText()

	x.current().WriteString(
		fmt.Sprintf(`<%s>`, name),
	)
}

func (x *HtmlQuantyListener) ExitField(ctx *parser.FieldContext) {
	name := ctx.GetName().GetText()

	x.current().WriteString(
		fmt.Sprintf(`</%s>`, name),
	)
}

func (x *HtmlQuantyListener) ExitComponent(ctx *parser.ComponentContext) {
	x.current().WriteString(`{{ end }}`)

	x.templates.Parse(x.current().String())

	x.crtComponent = ""
}

func (x *HtmlQuantyListener) current() *strings.Builder {
	if crt := x.components[x.crtComponent]; crt != nil {
		return crt
	}
	x.components[x.crtComponent] = &strings.Builder{}
	return x.components[x.crtComponent]
}

func (x *HtmlQuantyListener) Print() {
	for _, str := range x.components {
		fmt.Println(str.String())
	}
}

func (x *HtmlQuantyListener) Lookup(name string) *template.Template {
	return x.templates.Lookup(name)
}
