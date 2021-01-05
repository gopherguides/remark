package ifaces

import (
	"html/template"
	"strings"
)

type Element interface {
	HTML() template.HTML
}

type Elements []Element

func (elems Elements) HTML() template.HTML {
	lines := make([]string, 0, len(elems))
	for _, e := range elems {
		lines = append(lines, string(e.HTML()))
	}

	s := strings.Join(lines, "")
	return template.HTML(s)
}
