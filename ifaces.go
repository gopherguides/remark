package remark

type Appendable interface {
	Tag
	Append(tag Tag)
}
