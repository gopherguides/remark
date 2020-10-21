package remark

import (
	"github.com/gomarkdown/markdown/ast"
)

// TableCell is strikethrough
type TableCell struct {
	*ast.TableCell
	Elements Elements
}

func (cell TableCell) String() string {
	s := cell.Elements.String()
	return s
}

func NewTableCell(node *ast.TableCell) (*TableCell, error) {
	elems, err := NewElements(node.Children)
	if err != nil {
		return nil, err
	}

	cell := &TableCell{
		TableCell: node,
		Elements:  elems,
	}
	return cell, nil
}
