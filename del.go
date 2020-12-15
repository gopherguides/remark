package remark

import (
	"fmt"
	"html/template"

	"github.com/gobuffalo/tags/v3"
	"github.com/gomarkdown/markdown/ast"
)

// Del is strikethrough
type Del struct {
	*ast.Del
	Elements Elements
}

func (del Del) String() string {
	return fmt.Sprintf("~~%s~~", del.Elements)
}

func (del Del) HTML() template.HTML {
	strike := tags.New("strike", nil)
	strike.Append(del.Elements)
	return strike.HTML()
}

func NewDel(node *ast.Del) (*Del, error) {
	elems, err := NewElements(node.Children)
	if err != nil {
		return nil, err
	}

	d := &Del{
		Del:      node,
		Elements: elems,
	}

	return d, nil
}
