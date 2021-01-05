package md

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
	"github.com/gomarkdown/markdown/ast"
)

type TableHeader struct {
	*ast.TableHeader
	Elements Elements
}

func (th TableHeader) HTML() template.HTML {
	t := tags.New("th", nil)
	t.Append(th.Elements.HTML())
	return t.HTML()
}

func NewTableHeader(node *ast.TableHeader) (*TableHeader, error) {
	elems, err := NewElements(node.Children)
	if err != nil {
		return nil, err
	}

	th := &TableHeader{
		TableHeader: node,
		Elements:    elems,
	}
	return th, nil
}
