package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func (p *Parser) NewTable(node *ast.Table) (*remark.Table, error) {
	table := &remark.Table{}
	err := p.set(table, node)
	return table, err
}
