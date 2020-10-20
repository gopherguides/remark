package remark

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark/internal/md"
	"github.com/olekukonko/tablewriter"
)

// Table is strikethrough
type Table struct {
	*ast.Table
	Elements Elements
}

func (table Table) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"content": table.Elements,
		"literal": table.String(),
	}

	return json.MarshalIndent(map[string]interface{}{
		fmt.Sprintf("%T", table): m,
	}, "", "  ")
}

type TableWriter interface {
	WriteTable(w *tablewriter.Table)
}

func (tb Table) String() string {
	bb := &bytes.Buffer{}

	table := md.NewTableWriter(bb)
	// table := tablewriter.NewWriter(bb)
	// table.SetCenterSeparator("|")
	// table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	// table.SetHeader([]string{"Date", "Description", "CV2", "Amount"})

	// table.AppendBulk(data) // Add Bulk Data
	for _, e := range tb.Elements {
		sw, ok := e.(TableWriter)
		if !ok {
			panic(e)
		}
		sw.WriteTable(table)
	}

	table.Render()
	return bb.String()
}

func NewTable(node *ast.Table) (*Table, error) {
	elems, err := NewElements(node.Children)
	if err != nil {
		return nil, err
	}

	table := &Table{
		Table:    node,
		Elements: elems,
	}
	return table, nil
}
