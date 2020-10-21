package remark

import (
	"fmt"

	"github.com/gomarkdown/markdown/ast"
)

type Strong struct {
	*ast.Strong
	Elements Elements
}

func (s Strong) String() string {
	return fmt.Sprintf("**%s**", s.Elements)
}

func NewStrong(node *ast.Strong) (*Strong, error) {
	s := &Strong{
		Strong: node,
	}

	for _, n := range node.Children {
		el, err := NewElement(n)
		if err != nil {
			return nil, err
		}
		s.Elements = append(s.Elements, el)
	}
	return s, nil
}
