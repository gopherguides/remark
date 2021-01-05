package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type Link struct {
	Attributes Attributes
	Elements   Elements
}

func (l Link) HTML() template.HTML {
	opts := tags.Options{}
	for k, v := range l.Attributes {
		opts[k] = v
	}
	a := tags.New("a", opts)
	a.Append(l.Elements.HTML())
	return a.HTML()
}

func (l Link) URL() string {
	return l.Attributes["href"]
}

func (l Link) Alt() string {
	return l.Attributes["alt"]
}
