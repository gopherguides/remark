package remark

import (
	"bytes"
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

func (s *Section) Parent() Element {
	return s.parent
}

func (s Section) String() string {
	bb := &bytes.Buffer{}
	// fmt.Fprintf(bb, "%s\n\n", s.Heading)
	fmt.Fprintf(bb, s.Elements.String())
	return bb.String()
}

func NewSection(h *Heading) (*Section, error) {
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
