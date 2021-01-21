package remark

type Metadatable interface {
	Tag
	Metadata() Metadata
}

type Appendable interface {
	Tag
	Append(tag Tag)
}
