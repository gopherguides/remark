package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewTableCell(node *ast.TableCell) (*remark.TD, error) {
	cell := &remark.TD{}
	err := set(cell, node)
	return cell, err
}
