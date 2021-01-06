package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type Strike struct {
	Attributes
	Elements Elements
}

func (strike Strike) Name() string {
	return "strike"
}

func (strike *Strike) HTML() template.HTML {
	t := tags.New(strike.Name(), strike.Options())
	t.Append(strike.Elements.HTML())
	return t.HTML()
}
