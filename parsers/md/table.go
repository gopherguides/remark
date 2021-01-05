package md

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
	"github.com/gomarkdown/markdown/ast"
)

type Table struct {
	*ast.Table
	Elements Elements
}

func (tb Table) HTML() template.HTML {
	table := tags.New("table", nil)
	table.Append(tb.Elements.HTML())
	return table.HTML()
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
