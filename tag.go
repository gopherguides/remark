package remark

import "github.com/gobuffalo/tags/v3"

type Tag interface {
	Element
	Attrs() *Attributes
	Options() tags.Options
}
