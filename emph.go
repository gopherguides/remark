package remark

import (
	"fmt"

	"github.com/gomarkdown/markdown/ast"
)

type Emph struct {
	*ast.Emph
	Elements Elements
}

func (s Emph) String() string {
	return fmt.Sprintf("_%s_", s.Elements)
}

func NewEmph(node *ast.Emph) (*Emph, error) {
	s := &Emph{
		Emph: node,
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
