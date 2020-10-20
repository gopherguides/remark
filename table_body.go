package remark

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gomarkdown/markdown/ast"
	"github.com/olekukonko/tablewriter"
)

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
	var lines []string
	for _, e := range body.Elements {
		lines = append(lines, e.String())
	}
	return strings.Join(lines, "\n")
}

func (body TableBody) WriteTable(w *tablewriter.Table) {
	for _, e := range body.Elements {
		sw, ok := e.(TableWriter)
		if !ok {
			panic(e)
		}
		sw.WriteTable(w)
	}
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
