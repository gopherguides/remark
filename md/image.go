package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewImage(node *ast.Image) (*remark.Image, error) {
	elems, err := NewElements(node.Children)
	if err != nil {
		return nil, err
	}

	d := &remark.Image{
		Attributes: attrs(node.Attribute),
		Elements:   elems,
	}

	return d, nil
}
