package remark

import (
	"fmt"

	"github.com/gomarkdown/markdown/ast"
)

type ListItem struct {
	*ast.ListItem
	Elements Elements
}

func (li ListItem) String() string {
	return fmt.Sprintf("%s%s %s", string(li.BulletChar), string(li.Delimiter), li.Elements)
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
