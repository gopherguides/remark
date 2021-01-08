package remark

type Section struct {
	Attributes
	Elems Elements
}

func (s Section) Elements() Elements {
	return s.Elems
}

func (sec Section) Name() string {
	return "section"
}
