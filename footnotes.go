package remark

import (
	"encoding/json"
	"fmt"

	"github.com/gomarkdown/markdown/ast"
)

// Footnotes is strikethrough
type Footnotes struct {
	*ast.Footnotes
	Elements Elements
}

func (feet Footnotes) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"content": feet.Elements,
		"literal": feet.String(),
	}

	return json.MarshalIndent(map[string]interface{}{
		fmt.Sprintf("%T", feet): m,
	}, "", "  ")
}

func (feet Footnotes) String() string {
	return fmt.Sprintf("~~%s~~", feet.Elements)
}

func NewFootnotes(node *ast.Footnotes) (*Footnotes, error) {
	elems, err := NewElements(node.Children)
	if err != nil {
		return nil, err
	}

	d := &Footnotes{
		Footnotes: node,
		Elements:  elems,
	}

	return d, nil
}
