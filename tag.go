package remark

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/tags/v3"
)

type Tag interface {
	Attrs() Attributes
	GetChildren() Tags
	Options() tags.Options
	TagName() string
	fmt.Stringer
}

type Tags []Tag

func (tags Tags) Overview() string {
	for _, t := range tags {
		m, ok := t.(Metadatable)
		if !ok {
			continue
		}

		if _, ok := m.Metadata()["overview"]; !ok {
			continue
		}

		var lines []string
		for _, c := range m.GetChildren() {
			switch t := c.(type) {
			case *Heading:
				continue
			case String:
				lines = append(lines, t.String())
			default:
				lines = append(lines, t.GetChildren().String())
			}
		}
		return strings.Join(lines, " ")
	}
	return ""
}

func (tags Tags) String() string {
	lines := make([]string, 0, len(tags))
	for _, t := range tags {
		lines = append(lines, t.String())
	}
	return strings.Join(lines, " ")
}

func (tags Tags) Body() (Tag, bool) {
	return tags.FindFirst("body")
}

func (tags Tags) FindFirst(name string) (Tag, bool) {
	for _, e := range tags {
		if e.TagName() == name {
			return e, true
		}
		if el, ok := e.GetChildren().FindFirst(name); ok {
			return el, true
		}
	}
	return nil, false
}

func (tags Tags) FindAll(name string) Tags {
	var res Tags
	for _, e := range tags {
		if e.TagName() == name {
			res = append(res, e)
			continue
		}
		r := e.GetChildren().FindAll(name)
		res = append(res, r...)
	}

	return res
}
