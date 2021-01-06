package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewSubScript(node *ast.Subscript) (*remark.Sub, error) {
	sub := &remark.Sub{}
	err := set(sub, node)
	return sub, err
}
