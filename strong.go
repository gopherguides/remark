package remark

import (
	"fmt"
	"html/template"

	"github.com/gobuffalo/tags/v3"
	"github.com/gomarkdown/markdown/ast"
)

type Strong struct {
	*ast.Strong
	Elements Elements
}

func (s Strong) HTML() template.HTML {
	t := tags.New("strong", nil)
	t.Append(s.Elements.HTML())
	return t.HTML()
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
