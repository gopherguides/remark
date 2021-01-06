package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type Code struct {
	Attributes
	Language string
	Elements Elements
}

func (code Code) Name() string {
	return "code"
}

func (code *Code) HTML() template.HTML {
	t := tags.New(code.Name(), code.Options())
	t.Append(code.Elements.HTML())
	return t.HTML()
}
