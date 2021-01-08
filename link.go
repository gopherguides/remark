package remark

type Link struct {
	Attributes
	Elems Elements
}

func (a Link) Elements() Elements {
	return a.Elems
}

func (a Link) Name() string {
	return "a"
}
