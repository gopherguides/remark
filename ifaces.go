package remark

type Textable interface {
	Tag
	Text() string
}

type Metadatable interface {
	Tag
	Metadata() Metadata
}

type Appendable interface {
	Tag
	Append(tag Tag)
}
