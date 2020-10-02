package remark

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark/internal/md"
)

type Paragraph struct {
	*ast.Paragraph
}

func (p Paragraph) String() string {
	return md.String(p.Paragraph)
}
