package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewListItem(node *ast.ListItem) (remark.Tag, error) {
	li := &remark.Generic{
		TagName: "li",
	}
	err := set(li, node)
	return li, err
}
