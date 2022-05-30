// Code generated from ./Parser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Parser

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseParserVisitor) VisitFile(ctx *FileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseParserVisitor) VisitModuleDef(ctx *ModuleDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseParserVisitor) VisitImportDef(ctx *ImportDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseParserVisitor) VisitComponentDef(ctx *ComponentDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseParserVisitor) VisitVariableDefList(ctx *VariableDefListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseParserVisitor) VisitVariableDef(ctx *VariableDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseParserVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseParserVisitor) VisitSelectionSet(ctx *SelectionSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseParserVisitor) VisitSelectString(ctx *SelectStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseParserVisitor) VisitSelectTag(ctx *SelectTagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseParserVisitor) VisitTagDef(ctx *TagDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseParserVisitor) VisitArgumentList(ctx *ArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseParserVisitor) VisitArgument(ctx *ArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}
