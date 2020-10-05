package remark

import (
	"github.com/gomarkdown/markdown/ast"
)

type Module struct {
	*ast.Document
	Sections Sections
}

type Modules []Module

func NewModule(node *ast.Document) (Module, error) {
	m := Module{
		Document: node,
	}
	return m, nil
}
