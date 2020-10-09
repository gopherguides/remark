package remark

import (
	"encoding/json"
	"fmt"

	"github.com/gomarkdown/markdown/ast"
)

type ListItem struct {
	*ast.ListItem
}

func (li ListItem) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(map[string]interface{}{
		fmt.Sprintf("%T", li): nil,
	}, "", "  ")
}

func (li ListItem) String() string {
	panic("remark.ListItem")
}

func NewListItem(node *ast.ListItem) (*ListItem, error) {
	li := &ListItem{
		ListItem: node,
	}
	return li, nil
}
