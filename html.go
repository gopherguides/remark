package remark

import (
	"encoding/json"
	"fmt"

	"github.com/gomarkdown/markdown/ast"
)

type HTML struct {
	*ast.HTMLSpan
}

func (h HTML) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"text": "",
	}
	return json.Marshal(map[string]interface{}{
		fmt.Sprintf("%T", h): m,
	})
}

func (h HTML) String() string {
	return ""
}

func NewHTML(node *ast.HTMLSpan) (*HTML, error) {
	h := &HTML{
		HTMLSpan: node,
	}
	return h, nil
}
