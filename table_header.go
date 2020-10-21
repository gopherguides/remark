package remark

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark/internal/md"
	"github.com/olekukonko/tablewriter"
)

var _ TableWriter = &TableHeader{}

// TableHeader is strikethrough
type TableHeader struct {
	*ast.TableHeader
	Elements Elements
}

func (th TableHeader) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"content": th.Elements,
		"literal": th.String(),
	}

	return json.MarshalIndent(map[string]interface{}{
		fmt.Sprintf("%T", th): m,
	}, "", "  ")
}

func (th TableHeader) String() string {
	bb := &bytes.Buffer{}
	table := md.NewTableWriter(bb)
	th.WriteTable(table)
	return bb.String()
}

func (row TableHeader) WriteTable(w *tablewriter.Table) error {
	if len(row.Elements) == 0 {
		return nil
	}

	var cells []*TableCell
	for _, e := range row.Elements {
		if row, ok := e.(*TableRow); ok {
			cells = row.Cells()
		}
		tw, ok := e.(TableWriter)
		if !ok {
			return fmt.Errorf("expected TableWriter, got %T", e)
		}
		tw.WriteTable(w)
	}

	var seps []string
	for _, c := range cells {
		var s string
		if c.Align == ast.TableAlignmentCenter {
			s += ":"
		}
		for i := 0; i < len(c.String()); i++ {
			s += "-"
		}
		if c.Align == ast.TableAlignmentCenter || c.Align == ast.TableAlignmentRight {
			s += ":"
		}
		seps = append(seps, s)
	}
	w.Append(seps)
	return nil
}

func NewTableHeader(node *ast.TableHeader) (*TableHeader, error) {
	elems, err := NewElements(node.Children)
	if err != nil {
		return nil, err
	}

	th := &TableHeader{
		TableHeader: node,
		Elements:    elems,
	}
	return th, nil
}
