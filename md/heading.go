package md

import (
	"fmt"

	"github.com/gopherguides/remark"

	"github.com/gomarkdown/markdown/ast"
)

func (p *Parser) NewHeading(node *ast.Heading) (*remark.Heading, error) {
	h := &remark.Heading{
		Generic: remark.NewGeneric(fmt.Sprintf("h%d", node.Level)),
		Level:   node.Level,
	}
	return h, nil
}
