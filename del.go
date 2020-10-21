package remark

import (
	"fmt"

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
