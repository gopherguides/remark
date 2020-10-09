package remark

import (
	"encoding/json"
	"fmt"

	"github.com/gomarkdown/markdown/ast"
)

type Strong struct {
	*ast.Strong
}

func (s Strong) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(map[string]interface{}{
		fmt.Sprintf("%T", s): nil,
	}, "", "  ")
}

func (s Strong) String() string {
	panic("remark.Strong")
}

func NewStrong(node *ast.Strong) (*Strong, error) {
	s := &Strong{
		Strong: node,
	}
	return s, nil
}
