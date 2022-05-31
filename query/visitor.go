package query

import (
	"io/ioutil"
	"log"
	"quanty/language/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var logger = log.New(ioutil.Discard, "", log.LstdFlags)

// Visit - Returns a CalcReturn.
func (q *Program) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(q)
}

// VisitChildren - This visit children implementation is only used for the prog node so it prints the statement.
func (q *Program) VisitChildren(node antlr.RuleNode) interface{} {
	for _, n := range node.GetChildren() {
		q.Visit(n.(antlr.ParseTree))
	}
	return nil
}

func (v *Program) VisitFile(ctx *parser.FileContext) interface{} {

	module := v.VisitModuleDef(ctx.ModuleDef().(*parser.ModuleDefContext))

	for _, i := range ctx.AllImportDef() {
		importDef := v.VisitImportDef(i.(*parser.ImportDefContext))
		module.AppendImport(importDef)
	}

	for _, c := range ctx.AllComponentDef() {
		component := v.VisitComponentDef(c.(*parser.ComponentDefContext))
		module.AppendComponent(component)
	}

	v.AppendModule(module)

	return v.VisitChildren(ctx)
}

func (v *Program) VisitImportDef(ctx *parser.ImportDefContext) *Import {
	return NewImport(ctx)
}

func (v *Program) VisitModuleDef(ctx *parser.ModuleDefContext) *Module {
	return NewModule(ctx)
}

func (v *Program) VisitComponentDef(ctx *parser.ComponentDefContext) *Component {
	return NewComponent(ctx)
}

func (v *Program) VisitVariableDefList(ctx *parser.VariableDefListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Program) VisitVariableDef(ctx *parser.VariableDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Program) VisitVariable(ctx *parser.VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Program) VisitSelectionSet(ctx *parser.SelectionSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Program) VisitSelection(ctx *parser.SelectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Program) VisitSelectString(ctx *parser.SelectStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Program) VisitSelectTag(ctx *parser.SelectTagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Program) VisitTagDef(ctx *parser.TagDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Program) VisitArgumentList(ctx *parser.ArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Program) VisitArgument(ctx *parser.ArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}
