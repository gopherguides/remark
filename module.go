package remark

import (
	"encoding/json"
	"fmt"
	"log"

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
	return m.Sections.String()
}

func (m Module) Format(st fmt.State, verb rune) {
	switch verb {
	case 'v':
		b, err := json.Marshal(m)
		if err != nil {
			log.Fatal(err)
		}
		st.Write(b)
		return
	case 'q':
		fmt.Fprintf(st, "%q", m.String())
		return
	}
	fmt.Fprint(st, m.String())
}

func NewModule(node *ast.Document) (*Module, error) {
	m := &Module{
		Document: node,
	}

	section := &Section{}
	for i, n := range node.Children {
		switch t := n.(type) {
		case *ast.Heading:
			m.Sections = append(m.Sections, section)
			h, err := NewHeading(t)
			if err != nil {
				return nil, err
			}
			section, err = NewSection(h, node.Children[i:])
			if err != nil {
				return nil, err
			}
		default:
			el, err := NewElement(t)
			if err != nil {
				return nil, err
			}

			fmt.Printf(">>>TODO module.go:47: el (%T): %v \n", el, el)
			section.Elements = append(section.Elements, el)
		}
	}
	m.Sections = append(m.Sections, section)

	return m, nil
}
