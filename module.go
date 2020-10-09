package remark

import (
	"strings"

	"github.com/gomarkdown/markdown/ast"
)

type Module struct {
	Document *ast.Document `json:"-"`
	Sections Sections
}

func (m Module) Parent() Element {
	return nil
}

func (m Module) MarshalJSON() ([]byte, error) {
	return m.Sections.MarshalJSON()
}

func (m Module) String() string {
	s := m.Sections.String()

	return strings.TrimSpace(s)
}

func NewModule(node *ast.Document) (*Module, error) {
	m := &Module{
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
