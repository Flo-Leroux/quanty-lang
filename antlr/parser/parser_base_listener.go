// Code generated from ./Parser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Parser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseParserListener is a complete listener for a parse tree produced by Parser.
type BaseParserListener struct{}

var _ ParserListener = &BaseParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDocument is called when production document is entered.
func (s *BaseParserListener) EnterDocument(ctx *DocumentContext) {}

// ExitDocument is called when production document is exited.
func (s *BaseParserListener) ExitDocument(ctx *DocumentContext) {}

// EnterComponentDef is called when production componentDef is entered.
func (s *BaseParserListener) EnterComponentDef(ctx *ComponentDefContext) {}

// ExitComponentDef is called when production componentDef is exited.
func (s *BaseParserListener) ExitComponentDef(ctx *ComponentDefContext) {}

// EnterVariableDefList is called when production variableDefList is entered.
func (s *BaseParserListener) EnterVariableDefList(ctx *VariableDefListContext) {}

// ExitVariableDefList is called when production variableDefList is exited.
func (s *BaseParserListener) ExitVariableDefList(ctx *VariableDefListContext) {}

// EnterVariableDef is called when production variableDef is entered.
func (s *BaseParserListener) EnterVariableDef(ctx *VariableDefContext) {}

// ExitVariableDef is called when production variableDef is exited.
func (s *BaseParserListener) ExitVariableDef(ctx *VariableDefContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseParserListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseParserListener) ExitVariable(ctx *VariableContext) {}

// EnterSelectionSet is called when production selectionSet is entered.
func (s *BaseParserListener) EnterSelectionSet(ctx *SelectionSetContext) {}

// ExitSelectionSet is called when production selectionSet is exited.
func (s *BaseParserListener) ExitSelectionSet(ctx *SelectionSetContext) {}

// EnterTagDef is called when production tagDef is entered.
func (s *BaseParserListener) EnterTagDef(ctx *TagDefContext) {}

// ExitTagDef is called when production tagDef is exited.
func (s *BaseParserListener) ExitTagDef(ctx *TagDefContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *BaseParserListener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BaseParserListener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaseParserListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseParserListener) ExitArgument(ctx *ArgumentContext) {}
