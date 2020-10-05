package remark

import (
	"github.com/gomarkdown/markdown/ast"
)

type Paragraph struct {
	*ast.Paragraph
}

func NewParagraph(node *ast.Paragraph) (Paragraph, error) {
	p := Paragraph{
		Paragraph: node,
	}
	return p, nil
}
