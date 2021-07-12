package parser

func (p *parser) parseModule() string {
	p.expectKeyword("module")
	return p.parseName()
}
