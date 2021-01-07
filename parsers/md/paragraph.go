package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewParagraph(node *ast.Paragraph) (remark.Tag, error) {
	g := &remark.Generic{
		TagName: "p",
	}

	err := set(g, node)
	return g, err
}
