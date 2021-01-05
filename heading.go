package remark

import (
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

type Headings []Heading
