package remark

import (
	"bytes"
	"fmt"

	"github.com/gomarkdown/markdown/ast"
)

type Heading struct {
	*ast.Heading
	Elements Elements
	parent   Element
}

func (h *Heading) Parent() Element {
	return h.parent
}

func (h Heading) String() string {
	if h.Heading == nil {
		return ""
	}
	if h.Level == 0 {
		h.Level = 1
	}

	bb := &bytes.Buffer{}
	for i := 0; i < h.Level; i++ {
		fmt.Fprint(bb, "#")
	}
	fmt.Fprintf(bb, " %s", h.Elements)
	return bb.String()
}

type Headings []Heading

// func (s Headings) String() string {
// 	lines := make([]string, 0, len(s))
// 	for _, e := range s {
// 		lines = append(lines, e.String())
// 	}
// 	return strings.Join(lines, "\n\n")
// }

func NewHeading(node *ast.Heading) (*Heading, error) {
	h := &Heading{
		Heading: node,
	}
	return h, nil
}
