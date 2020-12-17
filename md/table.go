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

var _ TableWriter = &Table{}

type Table struct {
	*ast.Table
	Elements Elements
}

func (tb Table) HTML() template.HTML {
	table := tags.New("table", nil)
	table.Append(tb.Elements.HTML())
	return table.HTML()
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
