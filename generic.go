package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type Generic struct {
	TagName string
	Attributes
	Elems Elements
}

func (g Generic) Elements() Elements {
	return g.Elems
}

func (g Generic) Name() string {
	return g.TagName
}

func (g *Generic) HTML() template.HTML {
	t := tags.New(g.Name(), g.Options())
	t.Append(g.Elements().HTML())
	return t.HTML()
}
