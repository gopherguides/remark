package remark

import (
	"fmt"

	"github.com/gomarkdown/markdown/ast"
)

type Code struct {
	*ast.Code
}

func (c Code) String() string {
	return fmt.Sprintf("`%s`", string(c.Literal))
}

func NewCode(node *ast.Code) (Code, error) {
	c := Code{
		Code: node,
	}
	return c, nil
}
