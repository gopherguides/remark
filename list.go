package remark

import (
	"html/template"
	"strings"

	"github.com/gobuffalo/tags/v3"
	"github.com/gomarkdown/markdown/ast"
)

type List struct {
	*ast.List
	Elements Elements
}

func (l List) HTML() template.HTML {
	var t *tags.Tag
	switch l.ListFlags {
	case ast.ListTypeOrdered:
		t = tags.New("ol", nil)
	default:
		t = tags.New("ul", nil)
	}
	t.Append(l.Elements.HTML())
	return t.HTML()
}

func (l List) String() string {
	var lines []string
	for _, e := range l.Elements {
		s := e.String()
		if _, ok := l.Parent.(*ast.ListItem); ok {
			s = "  " + s
		}
		lines = append(lines, s)
	}
	return strings.Join(lines, "\n")
}

func NewList(node *ast.List) (*List, error) {
	l := &List{
		List: node,
	}

	for _, n := range node.Children {
		// always li
		el, err := NewElement(n)
		if err != nil {
			return nil, err
		}
		l.Elements = append(l.Elements, el)
	}
	return l, nil
}
