package remark

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/gomarkdown/markdown/ast"
)

type Section struct {
	Heading
	Elements    Elements
	SubSections Sections
	Nodes       []ast.Node
}

func (s Section) IsZero() bool {
	if s.Heading.Heading != nil {
		return false
	}

	if len(s.Elements) > 0 || len(s.SubSections) > 0 || len(s.Nodes) > 0 {
		return false
	}

	return true
}

func (s Section) String() string {
	bb := &bytes.Buffer{}
	fmt.Fprintf(bb, "%s\n\n", s.Heading)

	if len(s.Elements) > 0 {
		fmt.Fprintf(bb, "%s\n\n", s.Elements)
	}

	if len(s.SubSections) > 0 {
		fmt.Fprintf(bb, "%s\n\n", s.SubSections)
	}
	return bb.String()
}

// func (s Section) String() string {
// 	lines := make([]string, 0, len(s.Children))
//
// 	for _, c := range s.Children {
// 		lines = append(lines, md.String(c))
// 	}
// 	return strings.Join(lines, "\n\n")
// }

type Sections []Section

func (ss Sections) String() string {
	lines := make([]string, 0, len(ss))

	for _, s := range ss {
		lines = append(lines, s.String())
	}

	return strings.Join(lines, "\n\n")
}
