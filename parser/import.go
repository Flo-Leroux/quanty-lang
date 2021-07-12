package parser

import (
	. "quanty/ast"
	"quanty/lexer"
)

func (p *parser) parseImports() *ImportFileList {
	p.expectKeyword("import")

	var defs ImportFileList = make(ImportFileList, 0)

	// p.expect(lexer.ParenL)
	// p.expect(lexer.ParenR)

	// p.next()
	// tok := p.next()
	// if tok.Kind != lexer.Import {
	// 	p.unexpectedToken(tok)
	// }

	p.many(lexer.ParenL, lexer.ParenR, func() {
		defs = append(defs, p.parseImportFile())
	})

	return &defs
}

func (p *parser) parseImportFile() *ImportFileDefinition {
	var def ImportFileDefinition
	def.Position = p.peekPos()

	def.Name = p.parseName()
	def.Path = p.expect(lexer.String).Value

	return &def
}
