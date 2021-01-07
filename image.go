package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type Image struct {
	Attributes
	Elems Elements
}

func (img Image) Elements() Elements {
	return img.Elems
}

func (img Image) Name() string {
	return "img"
}

func (img *Image) HTML() template.HTML {
	t := tags.New(img.Name(), img.Options())
	t.Append(img.Elements().HTML())
	return t.HTML()
}
