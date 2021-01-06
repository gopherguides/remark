package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewDel(node *ast.Del) (*remark.Strike, error) {
	d := &remark.Strike{}

	if err := set(d, node); err != nil {
		return nil, err
	}

	return d, nil
}
