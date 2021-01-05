package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewHR(node *ast.HorizontalRule) (*remark.HR, error) {
	h := &remark.HR{}
	return h, nil
}
