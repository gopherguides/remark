package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewCode(node *ast.Code) (remark.Tag, error) {
	c := &remark.Generic{
		TagName: "code",
	}
	err := set(c, node)
	return c, err
}
