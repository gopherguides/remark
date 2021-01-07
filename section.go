package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type Section struct {
	Attributes
	Elems Elements
}

func (s Section) Elements() Elements {
	return s.Elems
}

func (sec Section) Name() string {
	return "section"
}

func (sec *Section) HTML() template.HTML {
	t := tags.New(sec.Name(), sec.Options())
	t.Append(sec.Elements().HTML())
	return t.HTML()
}
