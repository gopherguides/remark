package remark

import (
	"github.com/gomarkdown/markdown/ast"
)

type Paragraph struct {
	*ast.Paragraph
	Elements Elements
}

func (p Paragraph) String() string {
	return p.Elements.String()
}

func NewParagraph(node *ast.Paragraph) (Paragraph, error) {
	p := Paragraph{
		Paragraph: node,
	}

	elems, err := NewElements(node.Children)
	if err != nil {
		return p, err
	}

	p.Elements = elems
	return p, nil
}
