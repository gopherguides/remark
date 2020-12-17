package remark

import (
	"bufio"
	"fmt"
	"html/template"
	"strings"

	"github.com/gobuffalo/tags/v3"
)

type ListItem struct {
	Attributes Attributes
	Bullet     string
	Delimiter  string
	Elements   Elements
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
	s := fmt.Sprintf("%s%s %s", li.Bullet, li.Delimiter, strings.Join(lines, ""))
	return s
}
