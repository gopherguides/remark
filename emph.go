package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type Emph struct {
	Elements Elements
}

func (s Emph) HTML() template.HTML {
	t := tags.New("em", nil)
	t.Append(s.Elements.HTML())
	return t.HTML()
}
