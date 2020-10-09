package remark

import (
	"encoding/json"
	"fmt"

	"github.com/gomarkdown/markdown/ast"
)

type ListItem struct {
	*ast.ListItem
	Elements Elements
}

func (li ListItem) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"elements": li.Elements,
		"literal":  li.String(),
	}
	return json.MarshalIndent(map[string]interface{}{
		fmt.Sprintf("%T", li): m,
	}, "", "  ")
}

func (li ListItem) String() string {
	return fmt.Sprintf("%s %s", string(li.BulletChar), li.Elements)
}

func NewListItem(node *ast.ListItem) (*ListItem, error) {
	li := &ListItem{
		ListItem: node,
	}

	for _, n := range node.Children {
		el, err := NewElement(n)
		if err != nil {
			return nil, err
		}
		li.Elements = append(li.Elements, el)
	}
	return li, nil
}
