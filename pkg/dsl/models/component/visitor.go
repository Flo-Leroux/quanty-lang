package component

type Visitor interface {
	VisitComponent(*Component)
}

func (c *Component) Accept(v Visitor) {
	v.VisitComponent(c)
}
