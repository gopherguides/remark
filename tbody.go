package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type TableBody struct {
	Attributes
	Elements Elements
}

func (tbody TableBody) Name() string {
	return "tbody"
}

func (tbody *TableBody) HTML() template.HTML {
	t := tags.New(tbody.Name(), tbody.Options())
	t.Append(tbody.Elements.HTML())
	return t.HTML()
}
