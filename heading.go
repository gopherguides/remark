package remark

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/gomarkdown/markdown/ast"
)

type Heading struct {
	*ast.Heading
	Elements Elements
}

func (h Heading) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"content":  h.Elements.String(),
		"elements": h.Elements,
		"literal":  h.String(),
	}
	return json.MarshalIndent(map[string]interface{}{
		fmt.Sprintf("%T", h): m,
	}, "", "  ")
}

func (h Heading) String() string {
	if h.Heading == nil {
		return ""
	}
	if h.Level == 0 {
		h.Level = 1
	}

	bb := &bytes.Buffer{}
	for i := 0; i < h.Level; i++ {
		fmt.Fprint(bb, "#")
	}
	fmt.Fprintf(bb, " %s", h.Elements)
	return bb.String()
}

type Headings []Heading

func NewHeading(node *ast.Heading) (*Heading, error) {
	h := &Heading{
		Heading: node,
	}

	for _, n := range node.Children {
		el, err := NewElement(n)
		if err != nil {
			return nil, err
		}
		h.Elements = append(h.Elements, el)
	}
	return h, nil
}
