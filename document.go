package remark

type Document struct {
	Elems Elements
}

func (d Document) Name() string {
	return "document"
}

func (d Document) Elements() Elements {
	return d.Elems
}
