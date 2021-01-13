package remark

import "github.com/gobuffalo/tags/v3"

type OrderedList struct {
	Attributes
	metadata Metadata
	Elems    Elements
}

func (ol *OrderedList) Metadata() Metadata {
	if ol.metadata == nil {
		ol.metadata = Metadata{}
	}
	return ol.metadata
}

func (ol OrderedList) Elements() Elements {
	return ol.Elems
}

func (ol OrderedList) Name() string {
	return "ol"
}

func (g OrderedList) String() string {
	t := tags.New(g.Name(), g.Options())
	t.Append(g.Elems)
	return t.String()
}
