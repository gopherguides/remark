package remark

import (
	"encoding/json"
	"fmt"

	"github.com/gomarkdown/markdown/ast"
)

type HR struct {
	*ast.HorizontalRule
	Text string
}

func (hr HR) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"text": hr.Text,
	}
	return json.MarshalIndent(map[string]interface{}{
		fmt.Sprintf("%T", hr): m,
	}, "", "  ")
}

func (hr HR) String() string {
	return hr.Text
}

func NewHR(node *ast.HorizontalRule) (*HR, error) {
	h := &HR{
		HorizontalRule: node,
		Text:           string(node.Literal),
	}
	return h, nil
}
