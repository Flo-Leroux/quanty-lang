package ast

// File is
type File struct {
	// Pos        lexer.Position
	Component *Component `parser:"@@" json:"component"`
}

// Component is
type Component struct {
	// Pos  lexer.Position
	Name    string            `parser:"'component' @Ident" json:"name"`
	Connect *ComponentConnect `parser:"'{' @@* '}'" json:"connect"`
}

// ComponentConnect is
type ComponentConnect struct {
	Name string `parser:"'connect' @Ident" json:"name`
}

// Boolean is
type Boolean bool

// Capture is
func (b *Boolean) Capture(values []string) error {
	*b = values[0] == "true"
	return nil
}

// Value is
type Value struct {
	Float  *float64 `parser:"  @Float"`
	Int    *int     `parser:"| @Int"`
	String *string  `parser:"| @String"`
	Bool   *Boolean `parser:"| @('true' | 'false')"`
}
