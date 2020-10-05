package remark

import (
	"github.com/gomarkdown/markdown/ast"
)

type Section struct {
	Heading
	Elements    Elements
	SubSections Sections
	Nodes       []ast.Node
}

func NewSection(h Heading) (Section, error) {
	s := Section{
		Heading: h,
	}

	return s, nil
}

type Sections []Section
