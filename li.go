package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type ListItem struct {
	Attributes
	Elements Elements
}

func (li ListItem) Name() string {
	return "li"
}

func (li *ListItem) HTML() template.HTML {
	t := tags.New(li.Name(), li.Options())
	t.Append(li.Elements.HTML())
	return t.HTML()
}
