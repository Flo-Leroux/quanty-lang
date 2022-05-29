package ssr

import (
	"fmt"
	"io/ioutil"
	"log"
	"quanty/antlr/parser"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var logger = log.New(ioutil.Discard, "", log.LstdFlags)

type Visitor struct {
	parser.BaseParserVisitor
}

func NewVisitor() *Visitor {
	v := new(Visitor)
	return v
}

// Visit - Returns a CalcReturn.
func (q *Visitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(q)
}

// VisitChildren - This visit children implementation is only used for the prog node so it prints the statement.
func (q *Visitor) VisitChildren(node antlr.RuleNode) interface{} {
	for _, n := range node.GetChildren() {
		q.Visit(n.(antlr.ParseTree))
	}
	return nil
}

func (v *Visitor) VisitFile(ctx *parser.FileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitModuleDef(ctx *parser.ModuleDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitComponentDef(ctx *parser.ComponentDefContext) string {
	return v.VisitSelectionSet(ctx.SelectionSet().(*parser.SelectionSetContext))
}

func (v *Visitor) VisitVariableDefList(ctx *parser.VariableDefListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitVariableDef(ctx *parser.VariableDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitVariable(ctx *parser.VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitSelectionSet(ctx *parser.SelectionSetContext) string {

	res := make([]string, 0)

	for _, child := range ctx.AllSelection() {
		switch t := child.(type) {
		case *parser.SelectTagContext:
			res = append(res, v.VisitSelectTag(t))
		case *parser.SelectStringContext:
			res = append(res, v.VisitSelectString(t))
		}
		// ssr += v.VisitSelection(selection.(*parser.SelectionContext))
	}

	return strings.Join(res, "")
}

func (v *Visitor) VisitSelection(ctx *parser.SelectionContext) string {
	res := make([]string, 0)
	for _, child := range ctx.GetChildren() {
		switch t := child.(type) {
		case *parser.SelectTagContext:
			res = append(res, v.VisitSelectTag(t))
		case *parser.SelectStringContext:
			res = append(res, v.VisitSelectString(t))
		}
	}

	return strings.Join(res, "")
}

func (v *Visitor) VisitSelectString(ctx *parser.SelectStringContext) string {
	str := ctx.GetText()

	return strings.TrimPrefix(
		strings.TrimSuffix(
			str,
			"\"",
		),
		"\"",
	)
}

func (v *Visitor) VisitSelectTag(ctx *parser.SelectTagContext) string {
	return v.VisitTagDef(ctx.TagDef().(*parser.TagDefContext))
}

func (v *Visitor) VisitTagDef(ctx *parser.TagDefContext) string {

	var arguments string

	if ctx.ArgumentList() != nil {
		arguments = v.VisitArgumentList(ctx.ArgumentList().(*parser.ArgumentListContext))
	}

	openTag := fmt.Sprintf(
		"<%s %s>",
		ctx.IDEN().GetText(),
		arguments,
	)
	closeTag := fmt.Sprintf("</%s>", ctx.IDEN().GetText())

	children := v.VisitSelectionSet(ctx.SelectionSet().(*parser.SelectionSetContext))

	return fmt.Sprintf("%s %s %s", openTag, children, closeTag)
}

func (v *Visitor) VisitArgumentList(ctx *parser.ArgumentListContext) string {

	args := make(map[string][]string)
	for _, arg := range ctx.AllArgument() {
		argObj := v.VisitArgument(arg.(*parser.ArgumentContext))

		if args[argObj.Key] == nil {
			args[argObj.Key] = make([]string, 0)
		}

		args[argObj.Key] = append(args[argObj.Key], argObj.Value)
	}

	argsStrings := make([]string, 0)
	for key, values := range args {
		argsStrings = append(
			argsStrings,
			fmt.Sprintf("%s=\"%s\"", key, strings.Join(values, " ")),
		)
	}

	return strings.Join(argsStrings, " ")
}

func (v *Visitor) VisitArgument(ctx *parser.ArgumentContext) struct {
	Key   string
	Value string
} {
	return struct {
		Key   string
		Value string
	}{
		Key: ctx.GetKey().GetText(),
		Value: strings.TrimPrefix(
			strings.TrimSuffix(
				ctx.GetValue().GetText(),
				"\"",
			),
			"\"",
		),
	}
}
