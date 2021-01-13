package remark

import (
	"fmt"
)

type Heading struct {
	Attributes
	metadata Metadata
	Elems    Elements
	Level    int
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
