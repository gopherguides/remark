package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type Sup struct {
	Attributes
	Elements Elements
}

func (sup Sup) Name() string {
	return "sup"
}

func (sup *Sup) HTML() template.HTML {
	t := tags.New(sup.Name(), sup.Options())
	t.Append(sup.Elements.HTML())
	return t.HTML()
}
