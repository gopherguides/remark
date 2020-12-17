package remark

import (
	"fmt"
	"html/template"

	"github.com/gobuffalo/tags/v3"
	"github.com/gomarkdown/markdown/ast"
)

type Emph struct {
	*ast.Emph
	Elements Elements
}

func (s Emph) HTML() template.HTML {
	t := tags.New("em", nil)
	t.Append(s.Elements.HTML())
	return t.HTML()
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
