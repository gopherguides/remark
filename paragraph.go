package remark

import (
	"encoding/json"
	"fmt"

	"github.com/gomarkdown/markdown/ast"
)

type Paragraph struct {
	*ast.Paragraph
	Elements Elements
}

func (p Paragraph) MarshalJSON() ([]byte, error) {
	bits := map[string]interface{}{
		"elements": p.Elements,
		"literal":  p.String(),
	}
	return json.MarshalIndent(map[string]interface{}{
		fmt.Sprintf("%T", p): bits,
	}, "", "  ")
}

func (p Paragraph) String() string {
	return p.Elements.String()
}

func NewParagraph(node *ast.Paragraph) (*Paragraph, error) {
	elems, err := NewElements(node.Children)
	if err != nil {
		return nil, err
	}

	p := &Paragraph{
		Paragraph: node,
		Elements:  elems,
	}

	return p, nil
}
