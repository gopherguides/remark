package remark

import "github.com/gomarkdown/markdown/ast"

type Strong struct {
	*ast.Strong
}

func (s Strong) String() string {
	return ""
}

func NewStrong(node *ast.Strong) (Strong, error) {
	s := Strong{
		Strong: node,
	}
	return s, nil
}
