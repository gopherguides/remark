package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type Head struct {
	Attributes
	Elements Elements
}

func (head Head) Name() string {
	return "head"
}

func (head *Head) HTML() template.HTML {
	t := tags.New(head.Name(), head.Options())
	t.Append(head.Elements.HTML())
	return t.HTML()
}
