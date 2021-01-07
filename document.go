package remark

import "html/template"

type Document struct {
	Elements Elements
}

func (d Document) Name() string {
	return "document"
}

func (d Document) HTML() template.HTML {
	return d.Elements.HTML()
}
