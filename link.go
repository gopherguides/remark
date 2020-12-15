package remark

import (
	"fmt"
	"html/template"

	"github.com/gobuffalo/tags/v3"
	"github.com/gomarkdown/markdown/ast"
)

type Link struct {
	*ast.Link
	Elements Elements
}

func (l Link) HTML() template.HTML {
	a := tags.New("a", tags.Options{
		"href": l.URL(),
		"alt":  l.Caption(),
	})
	a.Append(l.Elements.HTML())
	return a.HTML()
}

func (l Link) URL() string {
	return string(l.Destination)
}

func (l Link) Caption() string {
	return string(l.Title)
}

func (link Link) String() string {
	if len(link.Title) == 0 {
		return fmt.Sprintf("[%s](%s)", link.Elements, link.URL())
	}
	return fmt.Sprintf("[%s](%s %q)", link.Elements, link.URL(), link.Caption())
}

func NewLink(node *ast.Link) (*Link, error) {
	l := &Link{
		Link: node,
	}

	for _, n := range node.Children {
		el, err := NewElement(n)
		if err != nil {
			return nil, err
		}
		l.Elements = append(l.Elements, el)
	}
	return l, nil
}
