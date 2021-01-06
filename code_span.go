package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type CodeSpan struct {
	Attributes
	Elements Elements
}

func (code CodeSpan) Name() string {
	return "code"
}

func (code *CodeSpan) HTML() template.HTML {
	t := tags.New(code.Name(), code.Options())
	t.Append(code.Elements.HTML())
	return t.HTML()
}
