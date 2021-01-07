package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewSuperScript(node *ast.Superscript) (remark.Tag, error) {
	sup := &remark.Generic{
		TagName: "sup",
	}
	err := set(sup, node)
	return sup, err
}
