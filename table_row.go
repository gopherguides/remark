package remark

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark/internal/md"
	"github.com/olekukonko/tablewriter"
)

// TableRow is strikethrough
type TableRow struct {
	*ast.TableRow
	Elements Elements
}

func (row TableRow) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"content": row.Elements,
		"literal": row.String(),
	}

	return json.MarshalIndent(map[string]interface{}{
		fmt.Sprintf("%T", row): m,
	}, "", "  ")
}

func (row TableRow) String() string {
	bb := &bytes.Buffer{}
	w := md.NewTableWriter(bb)
	row.WriteTable(w)
	return bb.String()
}

func (row TableRow) WriteTable(w *tablewriter.Table) {
	if len(row.Elements) == 0 {
		return
	}

	var rows []string
	for _, e := range row.Elements {
		c, ok := e.(*TableCell)
		if !ok {
			panic(e)
		}
		rows = append(rows, c.String())
	}
	if _, ok := row.Parent.(*ast.TableHeader); ok {
		w.SetHeader(rows)
		return
	}
	w.Append(rows)
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
