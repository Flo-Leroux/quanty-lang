package language

import (
	"path/filepath"
	"strings"

	"github.com/Flo-Leroux/quanty-lang/internal/language/ast"
	"github.com/Flo-Leroux/quanty-lang/internal/language/parser"
)

type File struct {
	name   string
	path   string
	raw    string
	parser *parser.Parser
	schema *ast.Schema
}

func (f *File) Schema() *ast.Schema {
	return f.schema
}

func toFile(path, directory, raw string) (*File, error) {
	p := parser.NewParser(raw)
	s := p.Parse()

	if err := p.Error(); err != nil {
		return nil, err
	}

	// Get the relative file path
	// ./views/html/index.tmpl -> index.tmpl
	rel, err := filepath.Rel(directory, path)
	if err != nil {
		return nil, err
	}

	// Reverse slashes '\' -> '/' and
	// partials\footer.tmpl -> partials/footer.tmpl
	name := filepath.ToSlash(rel)

	// Remove ext from name 'index.tmpl' -> 'index'
	name = strings.TrimSuffix(name, EXTENSION)
	// name = strings.Replace(name, e.extension, "", -1)

	file := &File{
		name:   name,
		path:   path,
		raw:    raw,
		parser: p,
		schema: s,
	}

	return file, nil
}
