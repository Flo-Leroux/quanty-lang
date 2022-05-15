package ast

type VariableDefinition struct {
	Kind string `json:"kind"`
	Name Name   `json:"name"`
}

func NewVariableDefinition(name Name) *VariableDefinition {
	return &VariableDefinition{
		Kind: variableDefinition.String(),
		Name: name,
	}
}
