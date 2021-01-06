package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewTable(node *ast.Table) (*remark.Table, error) {
	table := &remark.Table{}
	err := set(table, node)
	return table, err
}
