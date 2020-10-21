package remark

import (
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

func NewCode(node *ast.Code) (*Code, error) {
	c := &Code{
		Code: node,
		Text: string(node.Literal),
	}
	return c, nil
}
