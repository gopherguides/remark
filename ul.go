package remark

import "github.com/gobuffalo/tags/v3"

type UnorderedList struct {
	Attributes
	Elems    Elements
	metadata Metadata
}

func (ul *UnorderedList) Metadata() Metadata {
	if ul.metadata == nil {
		ul.metadata = Metadata{}
	}
	return ul.metadata
}

func (ul UnorderedList) Elements() Elements {
	return ul.Elems
}

func (ul UnorderedList) Name() string {
	return "ul"
}

func (g UnorderedList) String() string {
	t := tags.New(g.Name(), g.Options())
	t.Append(g.Elems)
	return t.String()
}
