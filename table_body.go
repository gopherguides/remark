package remark

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark/internal/md"
	"github.com/olekukonko/tablewriter"
)

var _ TableWriter = &TableBody{}

// TableBody is strikethrough
type TableBody struct {
	*ast.TableBody
	Elements Elements
}

func (body TableBody) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"content": body.Elements,
		"literal": body.String(),
	}

	return json.MarshalIndent(map[string]interface{}{
		fmt.Sprintf("%T", body): m,
	}, "", "  ")
}

func (body TableBody) String() string {
	bb := &bytes.Buffer{}
	table := md.NewTableWriter(bb)
	body.WriteTable(table)
	return bb.String()
}

func (body TableBody) WriteTable(w *tablewriter.Table) error {
	for _, e := range body.Elements {
		sw, ok := e.(TableWriter)
		if !ok {
			return fmt.Errorf("expected TableWriter, got %T", e)
		}
		sw.WriteTable(w)
	}
	return nil
}

func NewTableBody(node *ast.TableBody) (*TableBody, error) {
	elems, err := NewElements(node.Children)
	if err != nil {
		return nil, err
	}

	body := &TableBody{
		TableBody: node,
		Elements:  elems,
	}
	return body, nil
}
