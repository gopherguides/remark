package md

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
	"github.com/gomarkdown/markdown/ast"
)

type Paragraph struct {
	*ast.Paragraph
	Elements Elements
}

func (p Paragraph) HTML() template.HTML {
	pt := tags.New("p", nil)
	pt.Append(p.Elements.HTML())

	return pt.HTML()
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
