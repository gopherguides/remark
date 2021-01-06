package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewCode(node *ast.Code) (*remark.Code, error) {
	c := &remark.Code{}
	if err := set(c, node); err != nil {
		return nil, err
	}
	return c, nil
}
