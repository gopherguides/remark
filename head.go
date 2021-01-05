package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type Head struct {
	Elements Elements
}

func (m Head) HTML() template.HTML {
	t := tags.New("head", nil)
	t.Append(m.Elements.HTML())
	return t.HTML()
}
