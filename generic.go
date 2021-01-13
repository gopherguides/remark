package remark

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
