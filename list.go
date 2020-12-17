package remark

import (
	"html/template"
	"strings"

	"github.com/gobuffalo/tags/v3"
)

type ListType int

const (
	ListTypeOrdered ListType = iota
	ListTypeUnordered
)

type List struct {
	Type     ListType
	Elements Elements
}

func (l List) HTML() template.HTML {
	var t *tags.Tag
	switch l.Type {
	case ListTypeOrdered:
		t = tags.New("ol", nil)
	default:
		t = tags.New("ul", nil)
	}
	t.Append(l.Elements.HTML())
	return t.HTML()
}

func (l List) String() string {
	var lines []string
	// for _, e := range l.Elements {
	// 	s := e.String()
	// 	if _, ok := l.Parent.(*ast.ListItem); ok {
	// 		s = "  " + s
	// 	}
	// 	lines = append(lines, s)
	// }
	return strings.Join(lines, "\n")
}
