package md

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
	"github.com/gomarkdown/markdown/ast"
)

type TableRow struct {
	*ast.TableRow
	Elements Elements
}

func (row TableRow) HTML() template.HTML {
	tr := tags.New("tr", nil)
	tr.Append(row.Elements.HTML())
	return tr.HTML()
}

func (row TableRow) Cells() []*TableCell {
	cells := make([]*TableCell, 0, len(row.Elements))
	for _, e := range row.Elements {
		if c, ok := e.(*TableCell); ok {
			cells = append(cells, c)
		}
	}
	return cells
}

func NewTableRow(node *ast.TableRow) (*TableRow, error) {
	elems, err := NewElements(node.Children)
	if err != nil {
		return nil, err
	}

	row := &TableRow{
		TableRow: node,
		Elements: elems,
	}
	return row, nil
}
