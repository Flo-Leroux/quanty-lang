package ast

// File is
type File struct {
	Component *Component `parser:"@@" json:"component"`
}

// Component is
type Component struct {
	Name string `parser:"'component' @Ident '{}'" json:"name"`
}

// ComponentConnect is
type ComponentConnect struct {
	Name string `parser:"@Ident" json:"name"`
}
