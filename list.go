package remark

import (
	"html/template"

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
