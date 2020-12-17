package md

import (
	"html/template"
	"strings"

	"github.com/gomarkdown/markdown/ast"
)

func NewSection(head *remark.Heading, children []ast.Node) (*remark.Section, error) {
	s := &remark.Section{
		Heading: head,
	}

	var kids []ast.Node
	// var err error
	for _, n := range children {
		switch t := n.(type) {
		case *ast.Heading:
			h, err := NewHeading(t)
			if err != nil {
				return nil, err
			}
			sec, err := NewSection(h, kids)
			if err != nil {
				return nil, err
			}
			s.Sections = append(s.Sections, sec)
			kids = []ast.Node{}

		default:
			kids = append(kids, t)
		}
	}

	for _, k := range kids {
		el, err := NewElement(k)
		if err != nil {
			return nil, err
		}
		s.Elements = append(s.Elements, el)
	}
	return s, nil
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
