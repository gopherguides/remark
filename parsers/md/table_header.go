package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewTableHeader(node *ast.TableHeader) (remark.Tag, error) {
	th := &remark.Generic{
		TagName: "th",
	}
	err := set(th, node)
	return th, err
}
