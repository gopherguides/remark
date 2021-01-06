package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type Image struct {
	Attributes
	Elements Elements
}

func (img Image) Name() string {
	return "img"
}

func (img *Image) HTML() template.HTML {
	t := tags.New(img.Name(), img.Options())
	t.Append(img.Elements.HTML())
	return t.HTML()
}
