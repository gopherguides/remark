package md

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
	"github.com/gomarkdown/markdown/ast"
)

type TableCell struct {
	*ast.TableCell
	Elements Elements
}

func (cell TableCell) HTML() template.HTML {
	t := tags.New("td", nil)
	t.Append(cell.Elements.HTML())
	return t.HTML()
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
