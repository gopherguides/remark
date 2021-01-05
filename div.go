package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type Div struct {
	Attributes Attributes
	Elements   Elements
}

func (d Div) HTML() template.HTML {
	t := tags.New("div", nil)
	t.Append(d.Elements.HTML())
	return t.HTML()
}
