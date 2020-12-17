package remark

import (
	"fmt"
	"html/template"

	"github.com/gobuffalo/tags/v3"
	"github.com/gomarkdown/markdown/ast"
)

type SuperScript struct {
	*ast.Superscript
}

func (s SuperScript) HTML() template.HTML {
	t := tags.New("sup", tags.Options{
		"body": string(s.Literal),
	})

	return t.HTML()
}

func (script SuperScript) String() string {
	return fmt.Sprintf("^%s^", script.Literal)
}

func NewSuperScript(node *ast.Superscript) (*SuperScript, error) {
	d := &SuperScript{
		Superscript: node,
	}

	return d, nil
}
