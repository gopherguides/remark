package remark

import (
	"fmt"
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type Link struct {
	Attributes Attributes
	Elements   Elements
}

func (l Link) HTML() template.HTML {
	a := tags.New("a", tags.Options{
		"href": l.URL(),
		"alt":  l.Alt(),
	})
	a.Append(l.Elements.HTML())
	return a.HTML()
}

func (l Link) URL() string {
	return l.Attributes["href"]
}

func (l Link) Alt() string {
	return l.Attributes["alt"]
}

func (link Link) String() string {
	alt := link.Alt()
	if len(alt) == 0 {
		return fmt.Sprintf("[%s](%s)", link.Elements, link.URL())
	}
	return fmt.Sprintf("[%s](%s %q)", link.Elements, link.URL(), alt)
}
