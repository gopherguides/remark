package remark

import (
	"html/template"
	"strings"
)

type Document struct {
	Sections Sections
}

func (m Document) HTML() template.HTML {
	return m.Sections.HTML()
}

func (m Document) String() string {
	s := m.Sections.String()

	return strings.TrimSpace(s)
}
