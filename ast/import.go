package ast

type ImportFileList []*ImportFileDefinition

type ImportFileDefinition struct {
	Name     string
	Path     string
	Position *Position `dump:"-" json:"-"`
}
