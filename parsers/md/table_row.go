package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewTableRow(node *ast.TableRow) (remark.Tag, error) {
	row := &remark.Generic{
		TagName: "tr",
	}
	err := set(row, node)
	return row, err
}
