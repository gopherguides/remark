package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewBlockQuote(node *ast.BlockQuote) (remark.Tag, error) {
	g := &remark.Generic{
		TagName: "blockquote",
	}

	err := set(g, node)
	return g, err
}
