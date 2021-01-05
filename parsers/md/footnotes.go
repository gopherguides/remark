package md

import (
	"fmt"
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

func (feet Footnotes) String() string {
	return fmt.Sprintf("~~%s~~", feet.Elements)
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
