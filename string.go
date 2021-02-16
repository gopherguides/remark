package remark

import "github.com/gobuffalo/tags/v3"

type String string

var _ Tag = String("")

func (s String) Options() tags.Options {
	return tags.Options{}
}

func (s String) Attrs() Attributes {
	return Attributes{}
}

func (s String) String() string {
	return string(s)
}

func (s String) TagName() string {
	return "string"
}

func (s String) GetChildren() Tags {
	return Tags{}
}
