package ifaces

import (
	"fmt"
	"html/template"
	"strings"
)

type Element interface {
	fmt.Stringer
	HTML() template.HTML
}

type Elements []Element

func (elems Elements) String() string {
	lines := make([]string, 0, len(elems))
	for _, e := range elems {
		lines = append(lines, e.String())
	}
	return strings.Join(lines, "")
}

func (elems Elements) HTML() template.HTML {
	lines := make([]string, 0, len(elems))
	for _, e := range elems {
		lines = append(lines, string(e.HTML()))
	}

	s := strings.Join(lines, "")
	return template.HTML(s)
}
