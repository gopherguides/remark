package remark

import (
	"fmt"

	"github.com/gomarkdown/markdown/ast"
)

// SubScript is strikethrough
type SubScript struct {
	*ast.Subscript
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
