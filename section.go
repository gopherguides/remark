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
	parent      Element
}

func (s *Section) MarshalJSON() ([]byte, error) {
	bits := map[string]interface{}{
		"elements":     s.Elements,
		"sub_sections": s.SubSections,
	}
	return json.MarshalIndent(map[string]interface{}{
		fmt.Sprintf("%T", s): bits,
	}, "", "  ")
}

func (s *Section) Parent() Element {
	return s.parent
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
