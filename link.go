package remark

import (
	"encoding/json"
	"fmt"

	"github.com/gomarkdown/markdown/ast"
)

type Link struct {
	*ast.Link
	Caption  string
	Elements Elements
	URL      string
}

func (l Link) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"caption":  l.Caption,
		"elements": l.Elements,
		"literal":  l.String(),
		"url":      l.URL,
	}
	return json.MarshalIndent(map[string]interface{}{
		fmt.Sprintf("%T", l): m,
	}, "", "  ")
}

func (l Link) String() string {
	return fmt.Sprintf("[%s](%s)", l.Elements, l.URL)
}

func NewLink(node *ast.Link) (*Link, error) {
	l := &Link{
		Caption: string(node.Title),
		Link:    node,
		URL:     string(node.Destination),
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
