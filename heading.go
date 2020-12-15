package remark

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/gobuffalo/tags/v3"
	"github.com/gomarkdown/markdown/ast"
)

type Heading struct {
	*ast.Heading
	Elements Elements
}

func (h Heading) HTML() template.HTML {
	if h.Heading == nil {
		return ""
	}
	if h.Level == 0 {
		h.Level = 1
	}

	ht := tags.New(fmt.Sprintf("h%d", h.Level), nil)
	ht.Append(h.Elements.HTML())

	return ht.HTML()
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

func NewHeading(node *ast.Heading) (*Heading, error) {
	h := &Heading{
		Heading: node,
	}

	for _, n := range node.Children {
		el, err := NewElement(n)
		if err != nil {
			return nil, err
		}
		h.Elements = append(h.Elements, el)
	}
	return h, nil
}
