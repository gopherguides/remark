package remark

import (
	"github.com/gomarkdown/markdown/ast"
)

type Heading struct {
	*ast.Heading
	Elements Elements
}

type Headings []Heading

func NewHeading(node *ast.Heading) (Heading, error) {
	h := Heading{
		Heading: node,
	}
	return h, nil
}
