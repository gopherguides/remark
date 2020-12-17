package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewHeading(node *ast.Heading) (*remark.Heading, error) {
	h := &remark.Heading{
		Heading: node,
	}

	for _, n := range node.Children {
		el, err := NewElement(n)
		if err != nil {
			return nil, err
		}
		h.Elements = append(h.Elements, el)
	}
	return h, nil
}
