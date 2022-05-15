package ast

type Component struct {
	Kind string `json:"kind"`
	Name Name   `json:"name"`

	VariableDefinitions []*VariableDefinition `json:"variableDefinitions"`
}

func NewComponent() *Component {
	return &Component{
		Kind: componentDefinition.String(),

		VariableDefinitions: []*VariableDefinition{},
	}
}

func (c *Component) String() string {
	return c.Name.String()
}

func (c *Component) SetName(name string) *Component {
	c.Name = NewName(name)
	return c
}

func (c *Component) AddVariables(variable *VariableDefinition) *Component {
	c.VariableDefinitions = append(c.VariableDefinitions, variable)
	return c
}
