package remark

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
