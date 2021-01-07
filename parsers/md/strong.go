package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewStrong(node *ast.Strong) (remark.Tag, error) {
	s := &remark.Generic{
		TagName: "strong",
	}
	err := set(s, node)
	return s, err
}
