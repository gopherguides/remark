package remark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gomarkdown/markdown/ast"
)

type Section struct {
	*Heading
	Elements Elements
	Sections Sections
	Nodes    []ast.Node
}

func (s *Section) MarshalJSON() ([]byte, error) {
	bits := map[string]interface{}{
		"elements": s.Elements,
		"heading":  s.Heading,
		"literal":  s.String(),
		"sections": s.Sections,
	}
	return json.MarshalIndent(map[string]interface{}{
		fmt.Sprintf("%T", s): bits,
	}, "", "  ")
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

func NewSection(head *Heading, children []ast.Node) (*Section, error) {
	s := &Section{
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

func (sections Sections) MarshalJSON() ([]byte, error) {
	bits := make([]json.Marshaler, 0, len(sections))
	for _, s := range sections {
		bits = append(bits, s)
	}
	return json.MarshalIndent(map[string]interface{}{
		fmt.Sprintf("%T", sections): bits,
	}, "", "  ")
}
