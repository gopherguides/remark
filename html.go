package remark

import "github.com/gomarkdown/markdown/ast"

type HTML struct {
	*ast.HTMLSpan
}

func (h HTML) String() string {
	return ""
}

func NewHTML(node *ast.HTMLSpan) (HTML, error) {
	h := HTML{
		HTMLSpan: node,
	}
	return h, nil
}
