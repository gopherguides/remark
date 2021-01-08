package remark

import (
	"fmt"
)

type Heading struct {
	Attributes
	Elems Elements
	Level int
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
