package remark

import (
	"encoding/json"
	"fmt"

	"github.com/gomarkdown/markdown/ast"
)

type List struct {
	*ast.List
}

func (l List) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(map[string]interface{}{
		fmt.Sprintf("%T", l): nil,
	}, "", "  ")
}

func (l List) String() string {
	panic("remark.List")
}

func NewList(node *ast.List) (*List, error) {
	l := &List{
		List: node,
	}
	return l, nil
}
