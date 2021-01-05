package remark

import (
	"fmt"
	"html/template"
	"strings"
)

type Section struct {
	*Heading
	Elements Elements
	Sections Sections
}

func (s Section) HTML() template.HTML {
	return s.Elements.HTML()
}

type Sections []*Section

func (s Sections) HTML() template.HTML {
	lines := make([]string, 0, len(s))
	for _, e := range s {
		fmt.Printf(">>>TODO section.go:46: e (%T): %v \n", e, e)
		fmt.Printf(">>>TODO section.go:21: e.HTML() (%T): %v \n", e.HTML(), e.HTML())
		lines = append(lines, string(e.HTML()))
	}
	sl := strings.Join(lines, "")
	return template.HTML(sl)
}
