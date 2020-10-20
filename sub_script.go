package remark

import (
	"encoding/json"
	"fmt"

	"github.com/gomarkdown/markdown/ast"
)

// SubScript is strikethrough
type SubScript struct {
	*ast.Subscript
}

func (script SubScript) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"literal": script.String(),
	}

	return json.MarshalIndent(map[string]interface{}{
		fmt.Sprintf("%T", script): m,
	}, "", "  ")
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
