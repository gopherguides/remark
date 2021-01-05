package remark

import (
	"html/template"
)

type Document struct {
	Sections Sections
}

func (m Document) HTML() template.HTML {
	return m.Sections.HTML()
}
