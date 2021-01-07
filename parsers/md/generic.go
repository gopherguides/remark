package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func (p *Parser) NewGeneric(name string, node ast.Node) (*remark.Generic, error) {
	g := &remark.Generic{
		TagName: name,
	}
	err := p.set(g, node)
	return g, err
}
