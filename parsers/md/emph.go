package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewEmph(node *ast.Emph) (*remark.Em, error) {
	s := &remark.Em{}
	err := set(s, node)
	return s, err
}
