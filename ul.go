package remark

type UnorderedList struct {
	Attributes
	Elems Elements
}

func (ul UnorderedList) Elements() Elements {
	return ul.Elems
}

func (ul UnorderedList) Name() string {
	return "ul"
}
