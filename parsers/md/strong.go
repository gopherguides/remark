package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewStrong(node *ast.Strong) (*remark.Strong, error) {
	s := &remark.Strong{}
	err := set(s, node)
	return s, err
}
