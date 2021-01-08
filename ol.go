package remark

type OrderedList struct {
	Attributes
	Elems Elements
}

func (ol OrderedList) Elements() Elements {
	return ol.Elems
}

func (ol OrderedList) Name() string {
	return "ol"
}
