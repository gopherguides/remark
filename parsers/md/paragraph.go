package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewParagraph(node *ast.Paragraph) (*remark.Paragraph, error) {
	p := &remark.Paragraph{}
	err := set(p, node)
	return p, err
}
