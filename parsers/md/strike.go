package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewDel(node *ast.Del) (remark.Tag, error) {
	d := &remark.Generic{
		TagName: "strike",
	}

	err := set(d, node)
	return d, err
}
