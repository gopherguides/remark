package remark

type Generic struct {
	TagName string
	Attributes
	Elems Elements
}

func (g Generic) Elements() Elements {
	return g.Elems
}

func (g Generic) Name() string {
	return g.TagName
}
