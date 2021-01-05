package md

import (
	"html/template"

	"github.com/gomarkdown/markdown/ast"
)

type Footnotes struct {
	*ast.Footnotes
	Elements Elements
}

func (feet Footnotes) HTML() template.HTML {
	panic(feet)
	// return ""
}

func NewFootnotes(node *ast.Footnotes) (*Footnotes, error) {
	elems, err := NewElements(node.Children)
	if err != nil {
		return nil, err
	}

	d := &Footnotes{
		Footnotes: node,
		Elements:  elems,
	}

	return d, nil
}
