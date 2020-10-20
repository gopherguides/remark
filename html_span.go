package remark

import (
	"encoding/json"
	"fmt"

	"github.com/gomarkdown/markdown/ast"
)

type HTMLSpan struct {
	*ast.HTMLSpan
}

func (h HTMLSpan) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"literal": h.String(),
	}
	return json.MarshalIndent(map[string]interface{}{
		fmt.Sprintf("%T", h): m,
	}, "", "  ")
}

func (h HTMLSpan) String() string {
	return string(h.Literal)
}

// NewHTMLSpan returns an element, because if it is a
// <code> span, for example, it'll return a CodeSpan type instead.
func NewHTMLSpan(node *ast.HTMLSpan) (Element, error) {
	h := &HTMLSpan{
		HTMLSpan: node,
	}

	// TODO: Parse HTML to get the attributes etc...
	return h, nil
}
