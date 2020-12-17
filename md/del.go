package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewDel(node *ast.Del) (*remark.Del, error) {
	elems, err := NewElements(node.Children)
	if err != nil {
		return nil, err
	}

	d := &remark.Del{
		Elements: elems,
	}

	return d, nil
}
