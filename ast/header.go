package ast

import (
	"fmt"
	"strings"
)

type Header struct {
	Text  string
	Level int
}

func (h Header) String() string {
	var b strings.Builder
	for i := 0; i < h.Level; i++ {
		b.WriteString("#")
	}
	fmt.Fprintf(&b, " %s", h.Text)
	return b.String()
}
