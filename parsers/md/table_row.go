package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewTableRow(node *ast.TableRow) (*remark.TR, error) {
	row := &remark.TR{}
	err := set(row, node)
	return row, err
}
