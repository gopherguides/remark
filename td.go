package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type TD struct {
	Attributes
	Elements Elements
}

func (td TD) Name() string {
	return "td"
}

func (td *TD) HTML() template.HTML {
	t := tags.New(td.Name(), td.Options())
	t.Append(td.Elements.HTML())
	return t.HTML()
}
