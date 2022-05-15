package ast

type Document struct {
	Kind        string `json:"kind"`
	FileName    string `json:"fileName"`
	Definitions []Node `json:"definitions"`
}

func NewDocument(fileName string) *Document {
	return &Document{
		Kind:        document.String(),
		FileName:    fileName,
		Definitions: []Node{},
	}
}

func (d *Document) String() string {
	return ""
}

func (d *Document) AddDefinition(node Node) *Document {
	d.Definitions = append(d.Definitions, node)
	return d
}
