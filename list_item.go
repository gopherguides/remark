package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type ListItem struct {
	Attributes Attributes
	Bullet     string
	Delimiter  string
	Elements   Elements
}

func (li ListItem) HTML() template.HTML {
	t := tags.New("li", nil)
	t.Append(li.Elements.HTML())
	return t.HTML()
}
