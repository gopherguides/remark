package remark

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type Heading struct {
	Level    int
	Elements Elements
}

func (h Heading) HTML() template.HTML {
	if h.Level == 0 {
		h.Level = 1
	}

	ht := tags.New(fmt.Sprintf("h%d", h.Level), nil)
	ht.Append(h.Elements.HTML())

	return ht.HTML()
}

func (h Heading) String() string {
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
