package ast

type Name struct {
	Kind  string `json:"kind"`
	Value string `json:"value"`
}

func NewName(value string) Name {
	return Name{
		Kind:  name.String(),
		Value: value,
	}
}

func (n *Name) String() string {
	return n.Value
}
