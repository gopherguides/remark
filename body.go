package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type Body struct {
	Attributes
	Elements Elements
}

func (body Body) Name() string {
	return "body"
}

func (body *Body) HTML() template.HTML {
	t := tags.New(body.Name(), body.Options())
	t.Append(body.Elements.HTML())
	return t.HTML()
}
