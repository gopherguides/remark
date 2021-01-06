package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type Heading struct {
	Attributes
	Elements Elements
}

func (h Heading) Name() string {
	return "h"
}

func (h *Heading) HTML() template.HTML {
	t := tags.New(h.Name(), h.Options())
	t.Append(h.Elements.HTML())
	return t.HTML()
}
