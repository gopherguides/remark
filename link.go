package remark

import (
	"encoding/json"
	"fmt"

	"github.com/gomarkdown/markdown/ast"
)

type Link struct {
	*ast.Link
}

func (l Link) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(map[string]interface{}{
		fmt.Sprintf("%T", l): nil,
	}, "", "  ")
}

func (l Link) String() string {
	panic("remark.Link")
}

func NewLink(node *ast.Link) (*Link, error) {
	l := &Link{
		Link: node,
	}
	return l, nil
}
