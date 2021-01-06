package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewCodeSpan(node *ast.HTMLSpan) (*remark.CodeSpan, error) {
	c := &remark.CodeSpan{}
	err := set(c, node)
	return c, err
}
