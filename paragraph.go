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

func NewParagraph(node *ast.Paragraph) (*Paragraph, error) {
	elems, err := NewElements(node.Children)
	if err != nil {
		return nil, err
	}

	p := &Paragraph{
		Paragraph: node,
		Elements:  elems,
	}

	return p, nil
}
