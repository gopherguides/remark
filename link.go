package remark

import (
	"encoding/json"
	"fmt"

	"github.com/gomarkdown/markdown/ast"
)

type Link struct {
	*ast.Link
	Elements Elements
}

func (l Link) URL() string {
	return string(l.Destination)
}

func (l Link) Caption() string {
	return string(l.Title)
}

func (l Link) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"caption":  l.Caption(),
		"content":  l.Elements.String(),
		"elements": l.Elements,
		"literal":  l.String(),
		"url":      l.URL(),
	}
	return json.MarshalIndent(map[string]interface{}{
		fmt.Sprintf("%T", l): m,
	}, "", "  ")
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
