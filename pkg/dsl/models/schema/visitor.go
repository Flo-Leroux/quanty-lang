package schema

type Visitor interface {
	VisitSchema(*Schema)
}

func (s *Schema) Accept(v Visitor) {
	v.VisitSchema(s)
}
