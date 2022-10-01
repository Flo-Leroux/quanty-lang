package dsl

import (
	"github.com/Flo-Leroux/quanty-lang/pkg/dsl/models"
	"github.com/Flo-Leroux/quanty-lang/pkg/dsl/models/component"
	"github.com/Flo-Leroux/quanty-lang/pkg/dsl/parser"
)

type Program struct {
}

func New() *Program {
	p := parser.New("")

	p.RegisterStatement(models.Schema, component.Parser)

	p.ParseStatement(models.Schema)

	return &Program{}
}
