package remark

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
)

type Metadata map[string]string

var _ Tag = Metadata{}

func (md Metadata) TagName() string {
	return "metadata"
}

func (md Metadata) HTML() template.HTML {
	return ""
}

func (md Metadata) String() string {
	return ""
}

func (md Metadata) GetChildren() Tags {
	return Tags{}
}

func (md Metadata) Options() tags.Options {
	return tags.Options{}
}

func (md Metadata) Attrs() Attributes {
	return Attributes{}
}
