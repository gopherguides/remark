package remark

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gomarkdown/markdown/ast"
)

type List struct {
	*ast.List
	Elements Elements
}

func (l List) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"elements": l.Elements,
		"literal":  l.String(),
	}
	return json.MarshalIndent(map[string]interface{}{
		fmt.Sprintf("%T", l): m,
	}, "", "  ")
}

func (l List) String() string {
	var lines []string
	for _, e := range l.Elements {
		lines = append(lines, e.String())
	}
	return strings.Join(lines, "\n")
}

func NewList(node *ast.List) (*List, error) {
	l := &List{
		List: node,
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
