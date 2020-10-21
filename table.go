package remark

import (
	"bytes"
	"fmt"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark/internal/md"
	"github.com/olekukonko/tablewriter"
)

var _ TableWriter = &Table{}

// Table is strikethrough
type Table struct {
	*ast.Table
	Elements Elements
}

func (table Table) WriteTable(w *tablewriter.Table) error {
	for _, e := range table.Elements {
		sw, ok := e.(TableWriter)
		if !ok {
			return fmt.Errorf("expected TableWriter, got %T", e)
		}
		sw.WriteTable(w)
	}
	w.Render()
	return nil
}

func (tb Table) String() string {
	bb := &bytes.Buffer{}
	table := md.NewTableWriter(bb)
	tb.WriteTable(table)
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
