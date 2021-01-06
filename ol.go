package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type OrderedList struct {
	Attributes
	Elements Elements
}

func (ol OrderedList) Name() string {
	return "ol"
}

func (ol *OrderedList) HTML() template.HTML {
	t := tags.New(ol.Name(), ol.Options())
	t.Append(ol.Elements.HTML())
	return t.HTML()
}
