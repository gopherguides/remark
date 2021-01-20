package remark

import (
	"strings"

	"github.com/gobuffalo/tags/v3"
)

type Generic struct {
	Attributes
	Children Tags
	Name     string
	Data     Metadata
}

func (t *Generic) Append(tag Tag) {
	t.Children = append(t.Children, tag)
}

func (t *Generic) GetChildren() Tags {
	return t.Children
}

func (t *Generic) TagName() string {
	return t.Name
}

func (t *Generic) Metadata() Metadata {
	if t.Data == nil {
		t.Data = Metadata{}
	}
	return t.Data
}

func (tag Generic) String() string {
	if len(tag.Name) == 0 {
		return tag.GetChildren().String()
	}

	t := tags.New(tag.Name, tag.Options())

	if len(tag.Children) > 0 {
		t.Append(tag.Children.String())
	}

	return t.String()
}

func NewGeneric(name string) *Generic {
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)

	return &Generic{
		Name:       name,
		Data:       Metadata{},
		Attributes: Attributes{},
	}
}
