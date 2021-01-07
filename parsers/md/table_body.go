package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewTableBody(node *ast.TableBody) (remark.Tag, error) {
	body := &remark.Generic{
		TagName: "tbody",
	}
	err := set(body, node)
	return body, err
}
