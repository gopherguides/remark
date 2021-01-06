package remark

import (
	"fmt"
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type Heading struct {
	Attributes
	Level    int
	Elements Elements
}

func (h Heading) Name() string {
	lvl := h.Level
	if lvl == 0 {
		lvl = 1
	}
	return fmt.Sprintf("h%d", lvl)
}

func (h *Heading) HTML() template.HTML {
	t := tags.New(h.Name(), h.Options())
	t.Append(h.Elements.HTML())
	return t.HTML()
}
