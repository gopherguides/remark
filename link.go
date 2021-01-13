package remark

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
