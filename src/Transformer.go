package main

type Transformer struct {
	ast *Root
}

// NewTransformer returns a nex instance of Parser.
func NewTransformer(ast *Root) *Transformer {
	return &Transformer{
		ast,
	}
}

func (t *Transformer) transformComponent(c *Component) string {

	props := ""

	if len(c.Properties) > 0 {
		values := ""
		props += "{"

		for _, property := range c.Properties {
			if props != "{" {
				props += ", "
			}

			props += property.Name

			if property.Value != "" {
				if values == "" {
					values += "{ "
				} else {
					values += ", "
				}
				values += property.Name + ":" + property.Value
			}

		}

		props += "}"

		if values != "" {
			values += " }"
			props += " = "
			props += values
		}
	}

	template := "<div>My name is {name}</div>"

	return "export function " + c.Name + "(" + props + ") { return " + template + "; }"
}
