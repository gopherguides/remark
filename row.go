package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type Row struct {
	Attributes
	Elements Elements
}

func (row Row) Name() string {
	return "row"
}

func (row *Row) HTML() template.HTML {
	t := tags.New(row.Name(), row.Options())
	t.Append(row.Elements.HTML())
	return t.HTML()
}
