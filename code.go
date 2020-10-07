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
	return c.text
}

func (c Code) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"text": c.Text,
	}
	return json.Marshal(map[string]interface{}{
		fmt.Sprintf("%T", h): m,
	})
}

func NewCode(node *ast.Code) (*Code, error) {
	c := &Code{
		Code: node,
		text: fmt.Sprintf("`%s`", string(node.Literal)),
	}
	return c, nil
}
