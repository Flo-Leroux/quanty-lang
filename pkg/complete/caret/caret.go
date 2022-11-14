package caret

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"golang.org/x/exp/constraints"
)

type TokenPosition struct {
	Index   int
	Context antlr.ParseTree
	Text    string
}

type CaretPosition struct {
	Line   int
	Column int
}

func ComputeTokenPosition(
	parseTree antlr.ParseTree,
	tokens antlr.TokenStream,
	caretPosition CaretPosition,
	identifierTokenTypes []int,
) *TokenPosition {
	fmt.Printf("Underlying Type: %T\n", parseTree)

	if tree, ok := interface{}(parseTree).(antlr.TerminalNode); ok {
		fmt.Println("OK")
		return computeTokenPositionOfTerminal(tree, tokens, caretPosition, identifierTokenTypes)
	} else {
		// return nil
		return computeTokenPositionOfChildNode(parseTree.(antlr.ParserRuleContext), tokens, caretPosition, identifierTokenTypes)
	}

	// switch tree := parseTree.(type) {
	// case *antlr.TerminalNodeImpl:
	// case *antlr.ErrorNodeImpl:
	// 	return nil
	// default:
	// }
}

func positionOfToken(
	token antlr.Token,
	text string,
	caretPosition CaretPosition,
	identifierTokenTypes []int,
	parseTree antlr.ParseTree,
) *TokenPosition {
	start := token.GetTokenSource().GetCharPositionInLine()
	stop := start + len(token.GetText())

	fmt.Println(token, "=>", start, stop, caretPosition.Line, caretPosition.Column)
	fmt.Println()

	if token.GetLine() == caretPosition.Line && start <= caretPosition.Column && caretPosition.Column <= stop {
		index := token.GetTokenIndex()
		if Contains(identifierTokenTypes, token.GetTokenType()) {
			index--
		}
		return &TokenPosition{
			Index:   index,
			Context: parseTree,
			Text:    text[0 : caretPosition.Column-start],
		}
	} else {
		return nil
	}
}

func computeTokenPositionOfTerminal(
	parseTree antlr.TerminalNode,
	tokens antlr.TokenStream,
	caretPosition CaretPosition,
	identifierTokenTypes []int,
) *TokenPosition {
	token := parseTree.GetSymbol()
	text := parseTree.GetText()
	return positionOfToken(token, text, caretPosition, identifierTokenTypes, parseTree)
}

func computeTokenPositionOfChildNode(
	parseTree antlr.ParserRuleContext,
	tokens antlr.TokenStream,
	caretPosition CaretPosition,
	identifierTokenTypes []int,
) *TokenPosition {

	fmt.Println(parseTree.GetStart(), parseTree.GetStop())

	if parseTree.GetStart() != nil && parseTree.GetStart().GetLine() > caretPosition.Line ||
		parseTree.GetStop() != nil && parseTree.GetStop().GetLine() < caretPosition.Line {
		return nil
	}

	for i := 0; i < parseTree.GetChildCount(); i++ {
		position := ComputeTokenPosition(parseTree.GetChild(i).(antlr.ParseTree), tokens, caretPosition, identifierTokenTypes)
		if position != nil {
			return position
		}
	}

	if parseTree.GetStart() != nil && parseTree.GetStop() != nil {
		for i := parseTree.GetStart().GetTokenIndex(); i <= parseTree.GetStop().GetTokenIndex(); i++ {
			pos := positionOfToken(tokens.Get(i), tokens.Get(i).GetText(), caretPosition, identifierTokenTypes, parseTree)
			if pos != nil {
				return pos
			}
		}
	}

	return nil
}

func Contains[T constraints.Ordered](arr []T, toCheck T) bool {
	for _, item := range arr {
		if item == toCheck {
			return true
		}
	}

	return false
}

func FindCursorTokenIndex(tokenStream antlr.TokenStream, cursor CaretPosition) int {
	fmt.Println("Size", tokenStream.Size())
	// NOTE: cursor position is 1-based, while token's charPositionInLine is 0-based
	cursorCol := cursor.Column - 1
	for i := 0; i < tokenStream.Size(); i++ {
		t := tokenStream.Get(i)

		tokenStartCol := t.GetColumn()
		tokenEndCol := tokenStartCol + len(t.GetText())

		// tokenStartLine := t.GetLine()
		// tokenEndLine := tokenStartLine + 0 // len(t.GetText())

		if cursor.Line == t.GetLine() {

			if tokenStartCol <= cursorCol && cursorCol <= tokenEndCol {
				fmt.Println("Col", tokenStartCol, tokenEndCol, cursorCol)
				return t.GetTokenIndex()
			}
			// fmt.Println("Line", tokenStartLine, tokenEndLine, cursor.Line)
		}

		// fmt.Println(tokenEndLine > cursor.Line || tokenStartLine == cursor.Line && tokenEndCol > cursorCol)

		// NOTE: tokenEndCol makes sense only of tokenStartLine === tokenEndLine
		// if tokenEndLine > cursor.Line || tokenStartLine == cursor.Line && tokenEndCol > cursorCol {
		// 	if i > 0 && tokenStartLine == cursor.Line && tokenStartCol == cursorCol {
		// 		// 	// 	//   possibleIdentifierPrefix.test(tokenStream.get(i - 1).text as string)
		// 		return i - 1
		// 	} else if t.GetChannel() == antlr.TokenHiddenChannel {
		// 		return i + 1
		// 	} else {
		// 		return i
		// 	}
		// }
	}
	return -1
}
