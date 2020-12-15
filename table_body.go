package remark

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/gobuffalo/tags/v3"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark/internal/md"
	"github.com/olekukonko/tablewriter"
)

var _ TableWriter = &TableBody{}

type TableBody struct {
	*ast.TableBody
	Elements Elements
}

func (body TableBody) HTML() template.HTML {
	t := tags.New("tbody", nil)
	t.Append(body.Elements.HTML())
	return t.HTML()
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
