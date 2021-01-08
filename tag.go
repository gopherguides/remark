package remark

import "github.com/gobuffalo/tags/v3"

type Tag interface {
	Element
	Options() tags.Options
}
