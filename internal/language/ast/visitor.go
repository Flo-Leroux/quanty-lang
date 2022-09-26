package ast

// Visitable T -
type Visitable[T any] interface {
	Accept(Visitor[T])
}

// Visitor T -
type Visitor[T any] interface {
	VisitSchema(*Schema) T
	VisitComponentStatement(*ComponentStatement) T
	VisitStringValue(*StringValue) T
	VisitField(*Field) T
}
