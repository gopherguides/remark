package ast

import (
	"fmt"
	"strings"
)

type Header struct {
	Text  *Paragraph
	Level int
	Nodes Nodes
}

func (h Header) String() string {
	var b strings.Builder
	for i := 0; i < h.Level; i++ {
		b.WriteString("#")
	}
	fmt.Fprintf(&b, " %s", h.Text)

	if len(h.Nodes) == 0 {
		return b.String()
	}

	fmt.Fprintf(&b, "\n%s", h.Nodes)
	return b.String()
}
