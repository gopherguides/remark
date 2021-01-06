package remark

import (
	"html/template"
	"strings"
)

type Element interface {
	Name() string
}

type Elements []Element

func (elems Elements) Name() string {
	if len(elems) == 1 {
		return elems[0].Name()
	}
	return ""
}

func (elems Elements) HTML() template.HTML {
	var lines []string
	for _, e := range elems {
		if t, ok := e.(Tag); ok {
			lines = append(lines, string(t.HTML()))
		}
	}

	return template.HTML(strings.Join(lines, " "))
}
