package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type HTML struct {
	Attributes
	Elements Elements
}

func (html HTML) Name() string {
	return "html"
}

func (html *HTML) HTML() template.HTML {
	t := tags.New(html.Name(), html.Options())
	t.Append(html.Elements.HTML())
	return t.HTML()
}
