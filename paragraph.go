package remark

import (
	"encoding/json"

	"github.com/gomarkdown/markdown/ast"
)

type Paragraph struct {
	*ast.Paragraph `json:"-"`
	Elements       Elements `json:"elements"`
	parent         Element  `json:"parent"`
}

func (p Paragraph) MarshalJSON() ([]byte, error) {
	bits := map[string]interface{}{
		"elements": p.Elements,
	}
	return json.Marshal(bits)
}

func (p *Paragraph) Parent() Element {
	return p.parent
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
