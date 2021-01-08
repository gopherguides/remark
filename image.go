package remark

type Image struct {
	Attributes
	Elems Elements
}

func (img Image) Elements() Elements {
	return img.Elems
}

func (img Image) Name() string {
	return "img"
}
