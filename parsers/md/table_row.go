package md

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/gobuffalo/tags/v3"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark/internal/md"
	"github.com/olekukonko/tablewriter"
)

var _ TableWriter = &TableRow{}

// TableRow is strikethrough
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

func (row TableRow) String() string {
	bb := &bytes.Buffer{}
	w := md.NewTableWriter(bb)
	row.WriteTable(w)
	return bb.String()
}

func (row TableRow) WriteTable(w *tablewriter.Table) error {
	if len(row.Elements) == 0 {
		return nil
	}

	var rows []string
	for _, e := range row.Elements {
		c, ok := e.(*TableCell)
		if !ok {
			return fmt.Errorf("expected *TableCell, got %T", e)
		}
		rows = append(rows, c.String())
	}

	if _, ok := row.Parent.(*ast.TableHeader); ok {
		w.SetHeader(rows)
		return nil
	}

	w.Append(rows)
	return nil
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
