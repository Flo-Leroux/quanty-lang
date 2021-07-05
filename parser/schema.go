package parser

import (
	. "quanty/ast"
	quantyerror "quanty/error"
	"quanty/lexer"
)

func ParseFile(source *Source) (*SchemaDocument, *quantyerror.Error) {
	p := parser{
		lexer: lexer.New(source),
	}
	ast, err := p.parseSchemaDocument(), p.err
	if err != nil {
		return nil, err
	}

	// for _, def := range ast.Definitions {
	// 	def.BuiltIn = source.BuiltIn
	// }
	// for _, def := range ast.Extensions {
	// 	def.BuiltIn = source.BuiltIn
	// }

	return ast, nil
}

func MustParseFile(source *Source) *SchemaDocument {
	p := parser{
		lexer: lexer.New(source),
	}
	ast, err := p.parseSchemaDocument(), p.err
	if err != nil {
		panic(err)
	}

	return ast

	// for _, def := range ast.Definitions {
	// 	def.BuiltIn = source.BuiltIn
	// }
	// for _, def := range ast.Extensions {
	// 	def.BuiltIn = source.BuiltIn
	// }
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

		// var description string
		// if p.peek().Kind == lexer.BlockString || p.peek().Kind == lexer.String {
		// 	description = p.parseDescription()
		// }

		if p.peek().Kind != lexer.Name {
			p.unexpectedError()
			break
		}

		switch p.peek().Value {
		case "component":
			doc.Definitions = append(doc.Definitions, p.parseTypeSystemDefinition())
		// case "scalar", "type", "interface", "union", "enum", "input":
		// 	doc.Definitions = append(doc.Definitions, p.parseTypeSystemDefinition(description))
		// case "schema":
		// 	doc.Schema = append(doc.Schema, p.parseSchemaDefinition(description))
		// case "directive":
		// 	doc.Directives = append(doc.Directives, p.parseDirectiveDefinition(description))
		// case "extend":
		// 	if description != "" {
		// 		p.unexpectedToken(p.prev)
		// 	}
		// 	p.parseTypeSystemExtension(&doc)
		default:
			p.unexpectedError()
			return nil
		}
	}

	return &doc
}

func (p *parser) parseTypeSystemDefinition() *Definition {
	tok := p.peek()
	if tok.Kind != lexer.Name {
		p.unexpectedError()
		return nil
	}

	switch tok.Value {
	case "component":
		return p.parseComponentTypeDefinition()
	// case "type":
	// 	return p.parseObjectTypeDefinition(description)
	// case "interface":
	// 	return p.parseInterfaceTypeDefinition(description)
	// case "union":
	// 	return p.parseUnionTypeDefinition(description)
	// case "enum":
	// 	return p.parseEnumTypeDefinition(description)
	// case "input":
	// 	return p.parseInputObjectTypeDefinition(description)
	default:
		p.unexpectedError()
		return nil
	}
}

func (p *parser) parseComponentTypeDefinition() *Definition {
	p.expectKeyword("component")

	var def Definition
	def.Position = p.peekPos()
	def.Kind = Component
	def.Name = p.parseName()
	// def.Interfaces = p.parseImplementsInterfaces()
	// def.Directives = p.parseDirectives(true)
	def.Fields = p.parseFieldsDefinition()
	return &def
}

func (p *parser) parseFieldsDefinition() FieldList {
	var defs FieldList
	p.some(lexer.BraceL, lexer.BraceR, func() {
		defs = append(defs, p.parseFieldDefinition())
	})
	return defs
}

func (p *parser) parseFieldDefinition() *FieldDefinition {
	var def FieldDefinition
	def.Position = p.peekPos()
	// def.Description = p.parseDescription()
	def.Name = p.parseName()
	// def.Arguments = p.parseArgumentDefs()
	// p.expect(lexer.Colon)
	// def.Type = p.parseTypeReference()
	// def.Directives = p.parseDirectives(true)

	return &def
}
