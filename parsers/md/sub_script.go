package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewSubScript(node *ast.Subscript) (remark.Tag, error) {
	sub := &remark.Generic{
		TagName: "sub",
	}
	err := set(sub, node)
	return sub, err
}
