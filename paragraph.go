package remark

import (
	"bytes"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark/internal/md"
)

type Paragraph struct {
	*ast.Paragraph
}

func (p Paragraph) String() string {
	bb := &bytes.Buffer{}

	for _, c := range p.Children {
		bb.WriteString(md.String(c))
	}

	return bb.String()
}
