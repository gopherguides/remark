package remark

import (
	"encoding/json"
	"fmt"

	"github.com/gomarkdown/markdown/ast"
)

type HR struct {
	*ast.HorizontalRule
}

func (hr HR) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(map[string]interface{}{
		fmt.Sprintf("%T", hr): hr.String(),
	}, "", "  ")
}

func (hr HR) String() string {
	return string(hr.Literal)
}

func NewHR(node *ast.HorizontalRule) (*HR, error) {
	h := &HR{
		HorizontalRule: node,
	}
	return h, nil
}
