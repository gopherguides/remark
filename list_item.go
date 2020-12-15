package remark

import (
	"bufio"
	"fmt"
	"html/template"
	"strings"

	"github.com/gobuffalo/tags/v3"
	"github.com/gomarkdown/markdown/ast"
)

type ListItem struct {
	*ast.ListItem
	Elements Elements
}

func (li ListItem) HTML() template.HTML {
	t := tags.New("li", nil)
	t.Append(li.Elements.HTML())
	return t.HTML()
}

func (li ListItem) String() string {
	lines := make([]string, 0, len(li.Elements))
	for _, e := range li.Elements {
		switch t := e.(type) {
		case *List:
			scanner := bufio.NewScanner(strings.NewReader(t.String()))
			var sub []string
			for scanner.Scan() {
				text := scanner.Text()
				sub = append(sub, text)
			}
			lines = append(lines, "\n"+strings.Join(sub, "\n"))
		default:
			s := e.String()
			lines = append(lines, s)
		}
	}
	s := fmt.Sprintf("%s%s %s", string(li.BulletChar), string(li.Delimiter), strings.Join(lines, ""))
	return s
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
