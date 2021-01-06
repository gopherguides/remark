package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type Paragraph struct {
	Attributes
	Elements Elements
}

func (p Paragraph) Name() string {
	return "p"
}

func (p *Paragraph) HTML() template.HTML {
	t := tags.New(p.Name(), p.Options())
	t.Append(p.Elements.HTML())
	return t.HTML()
}
