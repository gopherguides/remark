package remark

import (
	"fmt"

	"github.com/gomarkdown/markdown/ast"
)

type Code struct {
	*ast.Code
	parent Element
}

func (c *Code) Parent() Element {
	return c.parent
}

func (c Code) String() string {
	return fmt.Sprintf("`%s`", string(c.Literal))
}

func NewCode(node *ast.Code) (*Code, error) {
	c := &Code{
		Code: node,
	}
	return c, nil
}
