package remark

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark/internal/md"
)

type Heading struct {
	*ast.Heading
}

func (h Heading) String() string {
	bb := &bytes.Buffer{}
	for i := 0; i < h.Level; i++ {
		bb.WriteString("#")
	}
	var lines []string
	for _, c := range h.Children {
		lines = append(lines, md.String(c))
	}
	fmt.Fprintf(bb, " %s", strings.Join(lines, ""))
	return bb.String()
}

type Headings []Heading
