package remark

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
)

type Section struct {
	*Heading
	Elements Elements
	Sections Sections
}

func (s Section) String() string {
	bb := &bytes.Buffer{}

	if s.Heading != nil {
		fmt.Fprintf(bb, "%s", s.Heading)
	}

	for _, e := range s.Elements {
		fmt.Fprintf(bb, "\n\n%s", e)
	}

	for _, e := range s.Sections {
		fmt.Fprintf(bb, "\n\n%s", e)
	}
	return bb.String()
}

type Sections []*Section

func (s Sections) String() string {
	lines := make([]string, 0, len(s))
	for _, e := range s {
		lines = append(lines, e.String())
	}
	return strings.Join(lines, "")
}

func (s Sections) HTML() template.HTML {
	lines := make([]string, 0, len(s))
	for _, e := range s {
		lines = append(lines, string(e.HTML()))
	}
	sl := strings.Join(lines, "")
	return template.HTML(sl)
}
