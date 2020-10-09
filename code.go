package remark

import (
	"encoding/json"
	"fmt"

	"github.com/gomarkdown/markdown/ast"
)

type Code struct {
	*ast.Code
	Text string
}

func (c Code) String() string {
	return fmt.Sprintf("`%s`", c.Text)
}

func (c Code) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"content": c.Text,
		"literal": c.String(),
	}
	return json.MarshalIndent(map[string]interface{}{
		fmt.Sprintf("%T", c): m,
	}, "", " ")
}

func NewCode(node *ast.Code) (*Code, error) {
	c := &Code{
		Code: node,
		Text: string(node.Literal),
	}
	return c, nil
}
