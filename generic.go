package remark

import (
	"html"

	"github.com/gobuffalo/tags/v3"
)

type Generic struct {
	TagName string
	Attributes
	Elems    Elements
	metadata Metadata
}

func (gen *Generic) Metadata() Metadata {
	if gen.metadata == nil {
		gen.metadata = Metadata{}
	}
	return gen.metadata
}

func (g Generic) Elements() Elements {
	return g.Elems
}

func (g Generic) Name() string {
	return g.TagName
}

func (g Generic) String() string {
	t := tags.New(g.Name(), g.Options())
	t.Append(g.Elems)
	return html.UnescapeString(t.String())
}
