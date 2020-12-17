package md

import (
	"fmt"
	"html/template"

	"github.com/gobuffalo/tags/v3"
	"github.com/gomarkdown/markdown/ast"
)

type SubScript struct {
	*ast.Subscript
}

func (s SubScript) HTML() template.HTML {
	t := tags.New("sub", tags.Options{
		"body": string(s.Literal),
	})
	return t.HTML()
}

func (script SubScript) String() string {
	return fmt.Sprintf("~%s~", script.Literal)
}

func NewSubScript(node *ast.Subscript) (*SubScript, error) {
	d := &SubScript{
		Subscript: node,
	}

	return d, nil
}
