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

func NewHTMLSpan(node *ast.HTMLSpan) (*HTMLSpan, error) {
	h := &HTMLSpan{
		HTMLSpan: node,
	}

	return h, nil
}
