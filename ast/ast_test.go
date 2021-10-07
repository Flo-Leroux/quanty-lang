package ast

import (
	"quanty/token"
	"testing"
)

func TestComponentString(t *testing.T) {
	doc := &Document{
		Definitions: []Definition{
			&ComponentDefinition{
				Kind: token.Token{Type: token.COMPONENT, Literal: "component"},
				Name: &Identifier{
					Kind:  token.Token{Type: token.IDENT, Literal: "Main"},
					Value: "Main",
				},
				SelectionSet: &SelectionSet{
					Kind: token.Token{Type: token.SELECTION_SET, Literal: ""},
					Selection: []Definition{
						&TagDefinition{
							Kind: token.Token{Type: token.TAG, Literal: "tag"},
							Name: &Identifier{
								Kind:  token.Token{Type: token.IDENT, Literal: "p"},
								Value: "p",
							},
							SelectionSet: &SelectionSet{
								Kind: token.Token{Type: token.TAG, Literal: "tag"},
								Selection: []Definition{
									&TagDefinition{
										Kind: token.Token{Type: token.TAG, Literal: "tag"},
										Name: &Identifier{
											Kind:  token.Token{Type: token.IDENT, Literal: "span"},
											Value: "span",
										},
										SelectionSet: &SelectionSet{},
									},
								},
							},
						},
						&TagDefinition{
							Kind: token.Token{Type: token.TAG, Literal: "tag"},
							Name: &Identifier{
								Kind:  token.Token{Type: token.IDENT, Literal: "p"},
								Value: "p",
							},
							SelectionSet: &SelectionSet{
								Kind: token.Token{Type: token.TAG, Literal: "tag"},
								Selection: []Definition{
									&TagDefinition{
										Kind: token.Token{Type: token.TAG, Literal: "tag"},
										Name: &Identifier{
											Kind:  token.Token{Type: token.IDENT, Literal: "bold"},
											Value: "bold",
										},
										SelectionSet: &SelectionSet{
											Kind: token.Token{Type: token.SELECTION_SET, Literal: ""},
											Selection: []Definition{
												&StringLiteral{
													Kind:  token.Token{Type: token.STRING, Literal: "Hello World!"},
													Value: "Hello World!",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	output := `component Main { p { span } p { bold { "Hello World!" } } }`

	if doc.String() != output {
		t.Errorf("document.String() wong, got=%q", doc.String())
	}
}
