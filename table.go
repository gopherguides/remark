package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type Table struct {
	Attributes
	Elems Elements
}

func (table Table) Elements() Elements {
	return table.Elems
}

func (table Table) Name() string {
	return "table"
}

func (table *Table) HTML() template.HTML {
	t := tags.New(table.Name(), table.Options())
	t.Append(table.Elements().HTML())
	return t.HTML()
}
