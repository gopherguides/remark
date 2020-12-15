package remark

import (
	"html/template"

	"github.com/gomarkdown/markdown/ast"
)

type HR struct {
	*ast.HorizontalRule
}

func (hr HR) String() string {
	return string(hr.Literal)
}

func NewHR(node *ast.HorizontalRule) (*HR, error) {
	h := &HR{
		HorizontalRule: node,
	}
	return h, nil
}

func (hr HR) HTML() template.HTML {
	return template.HTML(`<hr>`)
}
