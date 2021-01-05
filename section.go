package remark

import (
	"html/template"
	"strings"
)

type Section struct {
	Elements Elements
}

func (s Section) HTML() template.HTML {
	return s.Elements.HTML()
}

type Sections []*Section

func (s Sections) HTML() template.HTML {
	lines := make([]string, 0, len(s))
	for _, e := range s {
		lines = append(lines, string(e.HTML()))
	}
	sl := strings.Join(lines, "")
	return template.HTML(sl)
}
