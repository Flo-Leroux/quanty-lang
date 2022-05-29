// Code generated from ./Parser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Parser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ParserListener is a complete listener for a parse tree produced by Parser.
type ParserListener interface {
	antlr.ParseTreeListener

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// EnterModuleDef is called when entering the moduleDef production.
	EnterModuleDef(c *ModuleDefContext)

	// EnterComponentDef is called when entering the componentDef production.
	EnterComponentDef(c *ComponentDefContext)

	// EnterVariableDefList is called when entering the variableDefList production.
	EnterVariableDefList(c *VariableDefListContext)

	// EnterVariableDef is called when entering the variableDef production.
	EnterVariableDef(c *VariableDefContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterSelectionSet is called when entering the selectionSet production.
	EnterSelectionSet(c *SelectionSetContext)

	// EnterSelectString is called when entering the selectString production.
	EnterSelectString(c *SelectStringContext)

	// EnterSelectTag is called when entering the selectTag production.
	EnterSelectTag(c *SelectTagContext)

	// EnterTagDef is called when entering the tagDef production.
	EnterTagDef(c *TagDefContext)

	// EnterArgumentList is called when entering the argumentList production.
	EnterArgumentList(c *ArgumentListContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitModuleDef is called when exiting the moduleDef production.
	ExitModuleDef(c *ModuleDefContext)

	// ExitComponentDef is called when exiting the componentDef production.
	ExitComponentDef(c *ComponentDefContext)

	// ExitVariableDefList is called when exiting the variableDefList production.
	ExitVariableDefList(c *VariableDefListContext)

	// ExitVariableDef is called when exiting the variableDef production.
	ExitVariableDef(c *VariableDefContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitSelectionSet is called when exiting the selectionSet production.
	ExitSelectionSet(c *SelectionSetContext)

	// ExitSelectString is called when exiting the selectString production.
	ExitSelectString(c *SelectStringContext)

	// ExitSelectTag is called when exiting the selectTag production.
	ExitSelectTag(c *SelectTagContext)

	// ExitTagDef is called when exiting the tagDef production.
	ExitTagDef(c *TagDefContext)

	// ExitArgumentList is called when exiting the argumentList production.
	ExitArgumentList(c *ArgumentListContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)
}
