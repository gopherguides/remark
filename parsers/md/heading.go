package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func (p *Parser) NewHeading(node *ast.Heading) (*remark.Heading, error) {
	h := &remark.Heading{
		Level: node.Level,
	}
	err := p.set(h, node)
	return h, err
}
