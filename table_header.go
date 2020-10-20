package remark

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark/internal/md"
	"github.com/olekukonko/tablewriter"
)

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

func (row TableHeader) WriteTable(w *tablewriter.Table) {
	if len(row.Elements) == 0 {
		return
	}

	for _, e := range row.Elements {
		tw, ok := e.(TableWriter)
		if !ok {
			panic(e)
		}
		tw.WriteTable(w)
	}

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
