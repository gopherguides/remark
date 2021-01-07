package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type Link struct {
	Attributes
	Elems Elements
}

func (a Link) Elements() Elements {
	return a.Elems
}

func (a Link) Name() string {
	return "a"
}

func (a *Link) HTML() template.HTML {
	t := tags.New(a.Name(), a.Options())
	t.Append(a.Elements().HTML())
	return t.HTML()
}
