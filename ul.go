package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type UnorderedList struct {
	Attributes
	Elems Elements
}

func (ul UnorderedList) Elements() Elements {
	return ul.Elems
}

func (ul UnorderedList) Name() string {
	return "ul"
}

func (ul *UnorderedList) HTML() template.HTML {
	t := tags.New(ul.Name(), ul.Options())
	t.Append(ul.Elements().HTML())
	return t.HTML()
}
