package remark

import (
	"fmt"

	"github.com/gobuffalo/tags/v3"
)

type Metadata map[string]string

var _ Tag = Metadata{}

func (md Metadata) TagName() string {
	return "metadata"
}

func (md Metadata) String() string {
	tag := tags.New(md.TagName(), md.Options())
	for k, v := range md {
		tag.Append(fmt.Sprintf("%s:%s\n", k, v))
	}
	return tag.String()
}

func (md Metadata) GetChildren() Tags {
	return Tags{}
}

func (md Metadata) Options() tags.Options {
	return tags.Options{}
}

func (md Metadata) Attrs() *Attributes {
	return &Attributes{}
}
