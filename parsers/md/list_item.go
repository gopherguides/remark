package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewListItem(node *ast.ListItem) (*remark.ListItem, error) {
	li := &remark.ListItem{}
	err := set(li, node)
	return li, err
}
