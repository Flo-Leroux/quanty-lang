package values

type Visitor interface {
	VisitString(*String)
}

func (s *String) Accept(v Visitor) {
	v.VisitString(s)
}
