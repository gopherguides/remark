package md

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
	"github.com/gomarkdown/markdown/ast"
)

type TableBody struct {
	*ast.TableBody
	Elements Elements
}

func (body TableBody) HTML() template.HTML {
	t := tags.New("tbody", nil)
	t.Append(body.Elements.HTML())
	return t.HTML()
}

func NewTableBody(node *ast.TableBody) (*TableBody, error) {
	elems, err := NewElements(node.Children)
	if err != nil {
		return nil, err
	}

	body := &TableBody{
		TableBody: node,
		Elements:  elems,
	}
	return body, nil
}
