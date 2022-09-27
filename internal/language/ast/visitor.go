package ast

// Visitable -
type Visitable interface {
	Accept(Visitor)
}

// Visitor -
type Visitor interface {
	VisitSchema(*Schema)
	VisitComponentStatement(*ComponentStatement)
	VisitField(*Field)
	VisitStringValue(*StringValue)
}
