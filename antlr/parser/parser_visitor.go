// Code generated from ./Parser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Parser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by Parser.
type ParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by Parser#file.
	VisitFile(ctx *FileContext) interface{}

	// Visit a parse tree produced by Parser#moduleDef.
	VisitModuleDef(ctx *ModuleDefContext) interface{}

	// Visit a parse tree produced by Parser#importDef.
	VisitImportDef(ctx *ImportDefContext) interface{}

	// Visit a parse tree produced by Parser#componentDef.
	VisitComponentDef(ctx *ComponentDefContext) interface{}

	// Visit a parse tree produced by Parser#variableDefList.
	VisitVariableDefList(ctx *VariableDefListContext) interface{}

	// Visit a parse tree produced by Parser#variableDef.
	VisitVariableDef(ctx *VariableDefContext) interface{}

	// Visit a parse tree produced by Parser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by Parser#selectionSet.
	VisitSelectionSet(ctx *SelectionSetContext) interface{}

	// Visit a parse tree produced by Parser#selectString.
	VisitSelectString(ctx *SelectStringContext) interface{}

	// Visit a parse tree produced by Parser#selectTag.
	VisitSelectTag(ctx *SelectTagContext) interface{}

	// Visit a parse tree produced by Parser#tagDef.
	VisitTagDef(ctx *TagDefContext) interface{}

	// Visit a parse tree produced by Parser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}

	// Visit a parse tree produced by Parser#argument.
	VisitArgument(ctx *ArgumentContext) interface{}
}
