package parser

import (
	"github.com/Flo-Leroux/quanty-lang/pkg/dsl/models"
	"github.com/Flo-Leroux/quanty-lang/pkg/dsl/models/node"
)

type RegisterFnNode = func(*Parser) node.Node

func (p *Parser) RegisterNode(name models.Name, fn RegisterFnNode) {
	if _, ok := p.nodeParsers[name]; ok {
		p.nodeParsers[name] = fn
	}
}

func (p *Parser) ParseNode(name models.Name) node.Node {
	if fn, ok := p.nodeParsers[name]; ok {
		return fn(p)
	}

	return nil
}

type RegisterFnStatement = func(*Parser) node.Statement

func (p *Parser) RegisterStatement(name models.Name, fn RegisterFnStatement) {
	if _, ok := p.statementParsers[name]; ok {
		p.statementParsers[name] = fn
	}
}

func (p *Parser) ParseStatement(name models.Name) node.Statement {
	if fn, ok := p.statementParsers[name]; ok {
		return fn(p)
	}

	return nil
}
