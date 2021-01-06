package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type Sub struct {
	Attributes
	Elements Elements
}

func (sub Sub) Name() string {
	return "sub"
}

func (sub *Sub) HTML() template.HTML {
	t := tags.New(sub.Name(), sub.Options())
	t.Append(sub.Elements.HTML())
	return t.HTML()
}
