package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type TH struct {
	Attributes
	Elements Elements
}

func (th TH) Name() string {
	return "th"
}

func (th *TH) HTML() template.HTML {
	t := tags.New(th.Name(), th.Options())
	t.Append(th.Elements.HTML())
	return t.HTML()
}
