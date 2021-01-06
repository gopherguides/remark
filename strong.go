package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type Strong struct {
	Attributes
	Elements Elements
}

func (strong Strong) Name() string {
	return "strong"
}

func (strong *Strong) HTML() template.HTML {
	t := tags.New(strong.Name(), strong.Options())
	t.Append(strong.Elements.HTML())
	return t.HTML()
}
