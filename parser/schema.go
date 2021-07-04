package parser

import (
	. "quanty/ast"
	"quanty/lexer"
	"quanty/quantyerror"
)

func ParseFile(source *Source) (*SchemaDocument, *quantyerror.Error) {
	p := parser{
		lexer: lexer.New(source),
	}
	ast, err := p.parseSchemaDocument(), p.err
	if err != nil {
		return nil, err
	}

	for _, def := range ast.Definitions {
		def.BuiltIn = source.BuiltIn
	}
	for _, def := range ast.Extensions {
		def.BuiltIn = source.BuiltIn
	}

	return ast, nil
}

func ParseFiles(inputs ...*Source) (*SchemaDocument, *quantyerror.Error) {
	ast := &SchemaDocument{}
	for _, input := range inputs {
		inputAst, err := ParseFile(input)
		if err != nil {
			return nil, err
		}
		ast.Merge(inputAst)
	}
	return ast, nil
}

func (p *parser) parseSchemaDocument() *SchemaDocument {
	var doc SchemaDocument
	doc.Position = p.peekPos()
	for p.peek().Kind != lexer.EOF {
		if p.err != nil {
			return nil
		}

		var description string
		if p.peek().Kind == lexer.BlockString || p.peek().Kind == lexer.String {
			description = p.parseDescription()
		}

		if p.peek().Kind != lexer.Name {
			p.unexpectedError()
			break
		}

		switch p.peek().Value {
		case "scalar", "type", "interface", "union", "enum", "input":
			doc.Definitions = append(doc.Definitions, p.parseTypeSystemDefinition(description))
		case "schema":
			doc.Schema = append(doc.Schema, p.parseSchemaDefinition(description))
		case "directive":
			doc.Directives = append(doc.Directives, p.parseDirectiveDefinition(description))
		case "extend":
			if description != "" {
				p.unexpectedToken(p.prev)
			}
			p.parseTypeSystemExtension(&doc)
		default:
			p.unexpectedError()
			return nil
		}
	}

	return &doc
}
