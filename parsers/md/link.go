package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewLink(node *ast.Link) (*remark.Link, error) {
	l := &remark.Link{}
	err := set(l, node)
	return l, err
}
