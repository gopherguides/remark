package remark

type Metadatable interface {
	Tag
	Metadata() Metadata
}

type Appendable interface {
	Tag
	Append(tag Tag)
}

type Overviewable interface {
	Tag
	Overview() bool
}

type Summaryable interface {
	Tag
	Summary() bool
}
