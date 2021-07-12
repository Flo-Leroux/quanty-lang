package parser

import (
	. "quanty/ast"
	"quanty/lexer"
)

func (p *parser) parseValueLiteral(isConst bool) *Value {
	token := p.peek()

	var kind ValueKind
	switch token.Kind {
	case lexer.BracketL:
		return p.parseList(isConst)
	case lexer.BraceL:
		return p.parseObject(isConst)
	case lexer.Dollar:
		if isConst {
			p.unexpectedError()
			return nil
		}
		return &Value{Position: &token.Pos, Raw: p.parseVariable(), Kind: Variable}
	case lexer.Int:
		kind = IntValue
	case lexer.Float:
		kind = FloatValue
	case lexer.String:
		kind = StringValue
	case lexer.BlockString:
		kind = BlockValue
	case lexer.Name:
		switch token.Value {
		case "true", "false":
			kind = BooleanValue
		case "null":
			kind = NullValue
		default:
			kind = EnumValue
		}
	default:
		p.unexpectedError()
		return nil
	}

	p.next()

	return &Value{Position: &token.Pos, Raw: token.Value, Kind: kind}
}

func (p *parser) parseList(isConst bool) *Value {
	var values ChildValueList
	pos := p.peekPos()
	p.many(lexer.BracketL, lexer.BracketR, func() {
		values = append(values, &ChildValue{Value: p.parseValueLiteral(isConst)})
	})

	return &Value{Children: values, Kind: ListValue, Position: pos}
}

func (p *parser) parseObject(isConst bool) *Value {
	var fields ChildValueList
	pos := p.peekPos()
	p.many(lexer.BraceL, lexer.BraceR, func() {
		fields = append(fields, p.parseObjectField(isConst))
	})

	return &Value{Children: fields, Kind: ObjectValue, Position: pos}
}

func (p *parser) parseObjectField(isConst bool) *ChildValue {
	field := ChildValue{}
	field.Position = p.peekPos()
	field.Name = p.parseName()

	p.expect(lexer.Colon)

	field.Value = p.parseValueLiteral(isConst)
	return &field
}

func (p *parser) parseTypeReference() *Type {
	var typ Type

	if p.skip(lexer.BracketL) {
		typ.Position = p.peekPos()
		typ.Elem = p.parseTypeReference()
		p.expect(lexer.BracketR)
	} else {
		typ.Position = p.peekPos()
		typ.NamedType = p.parseName()
	}

	if p.skip(lexer.Bang) {
		typ.Position = p.peekPos()
		typ.NonNull = true
	}
	return &typ
}
