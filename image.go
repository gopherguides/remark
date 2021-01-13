package remark

import "github.com/gobuffalo/tags/v3"

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

func (g Image) String() string {
	t := tags.New(g.Name(), g.Options())
	t.Append(g.Elems)
	return t.String()
}
