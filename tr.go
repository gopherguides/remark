package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type TR struct {
	Attributes
	Elements Elements
}

func (tr TR) Name() string {
	return "tr"
}

func (tr *TR) HTML() template.HTML {
	t := tags.New(tr.Name(), tr.Options())
	t.Append(tr.Elements.HTML())
	return t.HTML()
}
