package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func (p *Parser) NewLink(node *ast.Link) (*remark.Link, error) {
	l := &remark.Link{}
	if err := p.set(l, node); err != nil {
		return nil, err
	}
	l.Set("href", string(node.Destination))

	t := string(node.Title)
	if len(t) > 0 {
		l.Set("title", t)
	}
	return l, nil
}
