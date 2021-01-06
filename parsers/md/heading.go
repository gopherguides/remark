package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewHeading(node *ast.Heading) (*remark.Heading, error) {
	h := &remark.Heading{
		Level: node.Level,
	}
	err := set(h, node)
	return h, err
}
