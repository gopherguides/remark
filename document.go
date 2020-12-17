package remark

import (
	"html/template"
	"strings"

	"github.com/gomarkdown/markdown/ast"
)

type Document struct {
	Document *ast.Document `json:"-"`
	Sections Sections
}

func (m Document) Parent() Element {
	return nil
}

func (m Document) HTML() template.HTML {
	return m.Sections.HTML()
}

func (m Document) String() string {
	s := m.Sections.String()

	return strings.TrimSpace(s)
}

func NewDocument(node *ast.Document) (*Document, error) {
	m := &Document{
		Document: node,
	}

	var kids []ast.Node
	var head *Heading
	var err error
	for _, n := range node.Children {
		switch t := n.(type) {
		case *ast.Heading:
			if t.Level > 1 {
				kids = append(kids, t)
				continue
			}

			if len(kids) > 0 {
				sec, err := NewSection(head, kids)
				if err != nil {
					return nil, err
				}
				m.Sections = append(m.Sections, sec)
				kids = []ast.Node{}
			}

			head, err = NewHeading(t)
			if err != nil {
				return nil, err
			}

		default:
			kids = append(kids, t)
		}
	}

	if len(kids) > 0 {
		sec, err := NewSection(head, kids)
		if err != nil {
			return nil, err
		}
		m.Sections = append(m.Sections, sec)
	}

	return m, nil
}
