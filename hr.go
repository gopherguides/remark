package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type HR struct {
	Attributes
	Elements Elements
}

func (hr HR) Name() string {
	return "hr"
}

func (hr *HR) HTML() template.HTML {
	t := tags.New(hr.Name(), hr.Options())
	t.Append(hr.Elements.HTML())
	return t.HTML()
}
