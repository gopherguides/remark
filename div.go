package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type Div struct {
	Attributes
	Elements Elements
}

func (div Div) Name() string {
	return "div"
}

func (div *Div) HTML() template.HTML {
	t := tags.New(div.Name(), div.Options())
	t.Append(div.Elements.HTML())
	return t.HTML()
}
