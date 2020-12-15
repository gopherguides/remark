package remark

import (
	"fmt"
	"html/template"

	"github.com/gobuffalo/tags/v3"
	"github.com/gomarkdown/markdown/ast"
)

type Code struct {
	*ast.Code
	Text string
}

func (c Code) String() string {
	return fmt.Sprintf("`%s`", c.Text)
}

func (c Code) HTML() template.HTML {
	pre := tags.New("pre", nil)
	code := tags.New("code", nil)
	code.Append(c.Text)
	pre.Append(code)

	return pre.HTML()
}

func NewCode(node *ast.Code) (*Code, error) {
	c := &Code{
		Code: node,
		Text: string(node.Literal),
	}
	return c, nil
}
