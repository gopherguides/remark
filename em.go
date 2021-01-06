package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type Em struct {
	Attributes
	Elements Elements
}

func (em Em) Name() string {
	return "em"
}

func (em *Em) HTML() template.HTML {
	t := tags.New(em.Name(), em.Options())
	t.Append(em.Elements.HTML())
	return t.HTML()
}
