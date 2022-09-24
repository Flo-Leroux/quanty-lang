package ast

type Selection interface {
	Statement
	selectionNode()
}

type SelectionList = []Selection
