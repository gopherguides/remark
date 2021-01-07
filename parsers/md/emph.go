package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewEmph(node *ast.Emph) (remark.Tag, error) {
	s := &remark.Generic{
		TagName: "em",
	}
	err := set(s, node)
	return s, err
}
