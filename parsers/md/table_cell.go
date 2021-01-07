package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewTableCell(node *ast.TableCell) (remark.Tag, error) {
	cell := &remark.Generic{
		TagName: "td",
	}
	err := set(cell, node)
	return cell, err
}
