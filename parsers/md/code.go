package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewCode(node *ast.Code) (*remark.Code, error) {
	c := &remark.Code{
		Text: string(node.Literal),
	}
	return c, nil
}
