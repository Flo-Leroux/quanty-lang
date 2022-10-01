package field

type Visitor interface {
	VisitField(*Field)
}

func (f *Field) Accept(v Visitor) {
	v.VisitField(f)
}
