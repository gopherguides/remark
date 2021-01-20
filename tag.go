package remark

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/tags/v3"
)

type Tag interface {
	Attrs() *Attributes
	GetChildren() Tags
	Options() tags.Options
	TagName() string
	fmt.Stringer
}

type Tags []Tag

func (tags Tags) Metadata() Metadata {
	md := Metadata{}
	for _, t := range tags {
		if m, ok := t.(Metadatable); ok {
			for k, v := range m.Metadata() {
				md[k] = v
			}
		}
		for k, v := range t.GetChildren().Metadata() {
			md[k] = v
		}
	}
	return md
}

func (tags Tags) String() string {
	lines := make([]string, 0, len(tags))
	for _, t := range tags {
		lines = append(lines, t.String())
	}
	return strings.Join(lines, "")
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
