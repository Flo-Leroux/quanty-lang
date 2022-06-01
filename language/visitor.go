package language

import (
	"io/ioutil"
	"log"
	"quanty/language/ast"
	"quanty/language/parser"
)

var logger = log.New(ioutil.Discard, "", log.LstdFlags)

type visitor struct {
	*parser.BaseParserVisitor
}

func (v *visitor) VisitDocument(ctx *parser.DocumentContext) *ast.DocumentNode {

	def := &ast.DocumentNode{
		Kind:        ast.DOCUMENT,
		Imports:     make([]string, 0),
		Definitions: make([]ast.DefinitionNode, 0),
	}

	def.Module = v.VisitModuleDef(ctx.ModuleDef().(*parser.ModuleDefContext))

	for _, child := range ctx.AllImportDef() {
		importDef := v.VisitImportDef(child.(*parser.ImportDefContext))
		def.Imports = append(def.Imports, importDef)
	}

	for _, child := range ctx.AllOperationDef() {
		var operationDef ast.DefinitionNode
		switch t := child.(type) {
		case *parser.ComponentDefContext:
			operationDef = v.VisitComponentDef(t)
		}

		if operationDef != nil {
			def.Definitions = append(def.Definitions, operationDef)
		}
	}

	return def
}

func (v *visitor) VisitModuleDef(ctx *parser.ModuleDefContext) string {
	return ctx.GetName().GetText()
}

func (v *visitor) VisitImportDef(ctx *parser.ImportDefContext) string {
	return ctx.GetModule().GetText()
}

func (v *visitor) VisitComponentDef(ctx *parser.ComponentDefContext) *ast.OperationDefinitionNode {

	def := &ast.OperationDefinitionNode{
		Kind:      ast.OPERATION_DEFINITION,
		Operation: ast.COMPONENT,
		Name: &ast.NameNode{
			Kind:  ast.NAME,
			Value: ctx.GetName().GetText(),
		},
		VariableDefinitions: make([]*ast.VariableDefinitionNode, 0),

		SelectionSet: v.VisitSelectionSet(ctx.SelectionSet().(*parser.SelectionSetContext)),
	}

	return def
}

func (v *visitor) VisitVariableDefList(ctx *parser.VariableDefListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitVariableDef(ctx *parser.VariableDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitNamedType(ctx *parser.NamedTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitNonNullType(ctx *parser.NonNullTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitListType(ctx *parser.ListTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *visitor) VisitSelectionSet(ctx *parser.SelectionSetContext) *ast.SelectionSetNode {
	def := &ast.SelectionSetNode{
		Kind:       ast.SELECTION_SET,
		Selections: make([]ast.SelectionNode, 0),
	}

	for _, child := range ctx.AllSelection() {
		var sel ast.SelectionNode
		switch t := child.(type) {
		case *parser.SelectStringContext:
			sel = v.VisitSelectString(t)
		case *parser.SelectFieldContext:
			sel = v.VisitSelectField(t)
		}

		if sel != nil {
			def.Selections = append(def.Selections, sel)
		}
	}

	return def
}

func (v *visitor) VisitSelectString(ctx *parser.SelectStringContext) *ast.StringValueNode {

	str := ctx.GetText()

	return &ast.StringValueNode{
		Kind:  ast.STRING,
		Value: str,
	}
}

func (v *visitor) VisitSelectField(ctx *parser.SelectFieldContext) *ast.FieldNode {
	if field := ctx.Field(); field != nil {
		return v.VisitField(field.(*parser.FieldContext))
	}

	return nil
}

func (v *visitor) VisitField(ctx *parser.FieldContext) *ast.FieldNode {
	def := &ast.FieldNode{
		Kind: ast.FIELD,
		Name: &ast.NameNode{
			Kind:  ast.NAME,
			Value: ctx.GetName().GetText(),
		},
		Arguments: make([]*ast.ArgumentNode, 0),
	}

	if alias := ctx.GetAlias(); alias != nil {
		def.Alias = &ast.NameNode{
			Kind:  ast.NAME,
			Value: alias.GetText(),
		}
	}

	if ctx.ArgumentList() != nil {
		def.Arguments = v.VisitArgumentList(ctx.ArgumentList().(*parser.ArgumentListContext))
	}

	if ctx.SelectionSet() != nil {
		def.SelectionSet = v.VisitSelectionSet(ctx.SelectionSet().(*parser.SelectionSetContext))
	}

	return def
}

func (v *visitor) VisitArgumentList(ctx *parser.ArgumentListContext) []*ast.ArgumentNode {
	list := make([]*ast.ArgumentNode, 0)

	for _, child := range ctx.AllArgument() {
		var arg *ast.ArgumentNode
		switch t := child.(type) {
		case *parser.ArgumentStringContext:
			arg = v.VisitArgumentString(t)
		case *parser.ArgumentVariableContext:
			arg = v.VisitArgumentVariable(t)
		}

		if arg != nil {
			list = append(list, arg)
		}
	}

	return list
}

func (v *visitor) VisitArgumentVariable(ctx *parser.ArgumentVariableContext) *ast.ArgumentNode {
	return &ast.ArgumentNode{
		Kind: ast.ARGUMENT,
		Value: &ast.VariableNode{
			Kind: ast.VARIABLE,
			Name: &ast.NameNode{
				Kind:  ast.NAME,
				Value: ctx.GetValue().GetText(),
			},
		},
	}
}

func (v *visitor) VisitArgumentString(ctx *parser.ArgumentStringContext) *ast.ArgumentNode {
	return &ast.ArgumentNode{
		Kind: ast.ARGUMENT,
		Value: &ast.StringValueNode{
			Kind:  ast.VARIABLE,
			Value: ctx.GetValue().GetText(),
		},
	}
}
