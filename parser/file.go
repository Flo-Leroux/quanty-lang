package parser

import (
	. "quanty/ast"
	quantyerror "quanty/error"
	"quanty/lexer"
)

func ParseFile(source *Source) (*File, *quantyerror.Error) {
	p := parser{
		lexer: lexer.New(source),
	}
	return p.parseFile(), p.err
}

func (p *parser) parseFile() *File {
	var file File
	for p.peek().Kind != lexer.EOF {
		if p.err != nil {
			return &file
		}
		file.Position = p.peekPos()
		switch p.peek().Kind {
		case lexer.Name:
			switch p.peek().Value {
			case "module":
				file.Module = p.parseModule()
			case "import":
				file.Imports = p.parseImports()
			case "component":
				file.Component = p.parseComponent()
			// case "fragment":
			// 	doc.Fragments = append(doc.Fragments, p.parseFragmentDefinition())
			default:
				p.unexpectedError()
			}
		default:
			p.unexpectedError()
		}
	}

	return &file
}
