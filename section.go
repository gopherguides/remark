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
	Elements    Elements
	SubSections Sections
	Nodes       []ast.Node
}

func (s *Section) MarshalJSON() ([]byte, error) {
	bits := map[string]interface{}{
		"heading":      s.Heading,
		"elements":     s.Elements,
		"sub_sections": s.SubSections,
	}
	return json.MarshalIndent(map[string]interface{}{
		fmt.Sprintf("%T", s): bits,
	}, "", "  ")
}

func (s Section) String() string {
	bb := &bytes.Buffer{}
	// fmt.Fprintf(bb, "%s\n\n", s.Heading)
	fmt.Fprintf(bb, s.Elements.String())
	return bb.String()
}

func NewSection(h *Heading, children []ast.Node) (*Section, error) {
	s := &Section{
		Heading: h,
	}
	fmt.Printf(">>>TODO section.go:40: len(children) (%T): %v \n", len(children), len(children))
	// for _, node := range children {
	// 	switch t := node.(type) {
	// 	case *ast.Heading:
	// 		if t.Level == 1 {
	// 			return s, nil
	// 		}
	// 		// h, err := NewHeading(t)
	// 		// if err != nil {
	// 		// 	return nil, err
	// 		// }
	// 		// if t.Level == 1 {
	// 		// 	return nil, fmt.Errorf("can not have H1 under another H1 %v", t)
	// 		// }
	// 		// sec, err := NewSection(h, children[i:])
	// 		// if err != nil {
	// 		// 	return nil, err
	// 		// }
	// 		// fmt.Printf(">>>TODO section.go:47: sec (%T): %v \n", sec, sec)
	// 	default:
	// 		fmt.Printf(">>>TODO section.go:48: t (%T): %v \n", t, t)
	// 	}
	// }
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
