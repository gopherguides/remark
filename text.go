package remark

import (
	"encoding/json"
	"fmt"

	"github.com/gomarkdown/markdown/ast"
)

type Text struct {
	*ast.Text
	Body string
}

func (t Text) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(map[string]interface{}{
		fmt.Sprintf("%T", t): t.Body,
	}, "", "  ")
}

func NewText(node *ast.Text) (*Text, error) {
	t := &Text{
		Text: node,
		Body: string(node.Literal),
	}
	return t, nil
}

func (t Text) String() string {
	return t.Body
}
