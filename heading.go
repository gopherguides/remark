package remark

import (
	"bytes"
	"fmt"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark/internal/md"
)

type Heading struct {
	*ast.Heading
	Sections Sections
}

func (h Heading) String() string {
	bb := &bytes.Buffer{}
	for i := 0; i < h.Level; i++ {
		bb.WriteString("#")
	}
	fmt.Fprintf(bb, " %s", md.String(h.Heading))
	return bb.String()
}

type Headings []Heading
