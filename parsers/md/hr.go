package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewHR(node *ast.HorizontalRule) (remark.Tag, error) {
	h := &remark.Generic{
		TagName: "hr",
	}
	return h, nil
}
