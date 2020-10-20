package remark

import (
	"encoding/json"
	"fmt"

	"github.com/gomarkdown/markdown/ast"
)

// SuperScript is strikethrough
type SuperScript struct {
	*ast.Superscript
}

func (script SuperScript) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"literal": script.String(),
	}

	return json.MarshalIndent(map[string]interface{}{
		fmt.Sprintf("%T", script): m,
	}, "", "  ")
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
