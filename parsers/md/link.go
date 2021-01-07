package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func (p *Parser) NewLink(node *ast.Link) (*remark.Link, error) {
	l := &remark.Link{}
	err := p.set(l, node)
	return l, err
}
