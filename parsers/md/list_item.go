package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewListItem(node *ast.ListItem) (*remark.ListItem, error) {
	li := &remark.ListItem{
		Attributes: attrs(node.Attribute),
	}

	for _, n := range node.Children {
		el, err := NewElement(n)
		if err != nil {
			return nil, err
		}
		li.Elements = append(li.Elements, el)

	}
	return li, nil
}
