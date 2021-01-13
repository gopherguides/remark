package remark

type Image struct {
	Attributes
	Elems    Elements
	metadata Metadata
}

func (img *Image) Metadata() Metadata {
	if img.metadata == nil {
		img.metadata = Metadata{}
	}
	return img.metadata
}

func (img Image) Elements() Elements {
	return img.Elems
}

func (img Image) Name() string {
	return "img"
}
