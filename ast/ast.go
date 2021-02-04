package ast

import (
	"fmt"
	"strings"
)

// QM is
type QM struct {
	// Pos        lexer.Position
	Package string   `parser:"'package' @Ident"`
	Entries []*Entry `parser:"@@*"`
}

// Entry is
type Entry struct {
	Component *Component `parser:"@@"`
}

// Component is
type Component struct {
	// Pos  lexer.Position
	Name     string           `parser:"'component' @Ident"`
	Contents []*ComponentBody `parser:"'{' @@* '}'"`
}

func (c *Component) String() string {

	var params string
	var destructParams string
	propsKey := c.getPropsKey()

	if len(propsKey) > 0 {
		defaultProps := c.defaultProps()
		params = fmt.Sprintf(`props = %s`, defaultProps)

		if len(defaultProps) > 0 {
			destructParams = fmt.Sprintf(`const { %s } = props;`, strings.Join(propsKey, ", "))
		}
	}

	return fmt.Sprintf(
		`export function %s(%s) {
			%s
			return <div>This component's name is {name} : %s</div>;
		}`,
		c.Name,
		params,
		destructParams,
		c.Name,
	)
}

func (c *Component) defaultProps() string {
	var props []string

	for _, content := range c.Contents {
		if content.Properties.Value != nil {
			propString := fmt.Sprintf(`%s: %s`, content.Properties.Key, *content.Properties.Value.String)
			props = append(props, propString)
		}
	}

	propsJoined := strings.Join(props, ", ")
	return fmt.Sprintf(`{ % s }`, propsJoined)
}

func (c *Component) getPropsKey() []string {
	var props []string

	for _, content := range c.Contents {
		props = append(props, content.Properties.Key)
	}

	return props
}

// ComponentBody is
type ComponentBody struct {
	Fragments  *Fragment `parser:"@@"`
	Properties *Property `parser:"| @@"`
	// Templates  *Template `parser:"| @@"`
}

// Property is
type Property struct {
	Key     string `parser:"'prop' @Ident"`
	TypeRef string `parser:"@Ident"`
	Value   *Value `parser:"('=' @@)?"`
}

func (prop *Property) String() (value string, ok bool) {

	if prop.Value == nil {
		return "", false
	}

	return fmt.Sprintf(`%s: %s`, prop.Key, *prop.Value.String), true
}

// Fragment is
type Fragment struct {
	Key     string `parser:"'fragment' @Ident"`
	TypeRef string `parser:"@Ident"`
}

// Template is
type Template struct {
	Tags HtmlTag `parser:"'template' '{' @@ '}'"`
}

type HtmlTag struct {
	Tag      string `parser:"'<'@Ident'>'"`
	children string `parser:"@String '</'Ident'>'"`
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
