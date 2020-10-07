package remark

import "github.com/gomarkdown/markdown/ast"

type Strong struct {
	*ast.Strong
	parent Element
}

func (s *Strong) Parent() Element {
	return s.parent
}

func (s Strong) String() string {
	return ""
}

func NewStrong(node *ast.Strong) (*Strong, error) {
	s := &Strong{
		Strong: node,
	}
	return s, nil
}
