package remark

import "github.com/gobuffalo/tags/v3"

type Link struct {
	Attributes
	Elems    Elements
	metadata Metadata
}

func (link *Link) Metadata() Metadata {
	if link.metadata == nil {
		link.metadata = Metadata{}
	}
	return link.metadata
}

func (a Link) Elements() Elements {
	return a.Elems
}

func (a Link) Name() string {
	return "a"
}

func (g Link) String() string {
	t := tags.New(g.Name(), g.Options())
	t.Append(g.Elems)
	return t.String()
}
