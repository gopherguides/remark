package remark

import (
	"fmt"

	"github.com/gomarkdown/markdown/ast"
)

// Footnotes is strikethrough
type Footnotes struct {
	*ast.Footnotes
	Elements Elements
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
