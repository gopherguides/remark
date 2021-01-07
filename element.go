package remark

import (
	"html/template"
	"strings"
)

type Element interface {
	Name() string
}

type Elemental interface {
	Elements() Elements
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
			h := string(t.HTML())
			lines = append(lines, h)
		}
	}

	return template.HTML(strings.Join(lines, ""))
}
