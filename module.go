package remark

import (
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
	b, _ := m.MarshalJSON()
	return string(b)
	// return m.Sections.String()
}

// func (m Module) Format(st fmt.State, verb rune) {
// 	switch verb {
// 	case 'v':
// 		b, err := json.MarshalIndent(m, "", "  ")
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		st.Write(b)
// 		return
// 	case 'q':
// 		fmt.Fprintf(st, "%q", m.String())
// 		return
// 	}
// 	fmt.Fprint(st, m.String())
// }

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

	// section := &Section{}
	// for i, n := range node.Children {
	// 	switch t := n.(type) {
	// 	case *ast.Heading:
	// 		m.Sections = append(m.Sections, section)
	// 		h, err := NewHeading(t)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		section, err = NewSection(h, node.Children[i:])
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		fmt.Printf(">>>TODO module.go:60: section (%T): %v \n", section, section)
	// 	default:
	// 		if section.Heading != nil {
	// 			continue
	// 		}
	// 		el, err := NewElement(t)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	//
	// 		section.Elements = append(section.Elements, el)
	// 	}
	// }
	// m.Sections = append(m.Sections, section)

	return m, nil
}
