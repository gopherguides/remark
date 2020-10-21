package remark

import (
	"fmt"

	"github.com/gomarkdown/markdown/ast"
)

// SuperScript is strikethrough
type SuperScript struct {
	*ast.Superscript
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
