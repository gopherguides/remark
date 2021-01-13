package remark

import (
	"fmt"

	"github.com/gobuffalo/tags/v3"
)

type Heading struct {
	Attributes
	Elems    Elements
	Level    int
	metadata Metadata
}

func (head *Heading) Metadata() Metadata {
	if head.metadata == nil {
		head.metadata = Metadata{}
	}
	return head.metadata
}

func (h Heading) Elements() Elements {
	return h.Elems
}

func (h Heading) Name() string {
	lvl := h.Level
	if lvl == 0 {
		lvl = 1
	}
	return fmt.Sprintf("h%d", lvl)
}

func (g Heading) String() string {
	t := tags.New(g.Name(), g.Options())
	t.Append(g.Elems)
	return t.String()
}
