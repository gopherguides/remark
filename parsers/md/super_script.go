package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewSuperScript(node *ast.Superscript) (*remark.Sup, error) {
	sup := &remark.Sup{}
	err := set(sup, node)
	return sup, err
}
