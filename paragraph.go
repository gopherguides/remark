package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type Paragraph struct {
	Attributes Attributes
	Elements   Elements
}

func (p *Paragraph) HTML() template.HTML {
	t := tags.New("p", nil)
	t.Append(p.Elements.HTML())
	return t.HTML()
}
