package remark

import (
	"encoding/json"
	"fmt"

	"github.com/gomarkdown/markdown/ast"
)

type Text struct {
	*ast.Text
}

func (t Text) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(map[string]interface{}{
		fmt.Sprintf("%T", t): t.String(),
	}, "", "  ")
}

func NewText(node *ast.Text) (*Text, error) {
	t := &Text{
		Text: node,
	}
	return t, nil
}

func (t Text) String() string {
	return string(t.Literal)
}
